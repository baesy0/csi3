package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/digital-idea/dilog"
	"github.com/digital-idea/dipath"
	"gopkg.in/mgo.v2"
)

// handleAPISetTags 함수는 아이템에 태그를 교체합니다.
func handleAPISetTags(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}
	defer session.Close()
	tokenID, _, err := TokenHandler(r, session)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}
	r.ParseForm() // 받은 문자를 파싱합니다. 파싱되면 map이 됩니다.
	var project string
	var name string
	var tags string
	args := r.PostForm
	for key, values := range args {
		switch key {
		case "project":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
				return
			}
			project = v
		case "name":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
				return
			}
			name = v
		case "tag", "tags":
			if len(values) == 0 {
				tags = ""
			} else {
				tags = values[0]
			}
		}
	}
	err = SetTags(session, project, name, Str2List(tags))
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}\n", err)
		return
	}
	// 로그처리
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dilog.Add(*flagDBIP, host, fmt.Sprintf("Edit tag: %s", tags), project, name, "csi3", tokenID, 180)
	fmt.Fprintf(w, "{\"error\":\"\"}\n")
}

// handleAPISetTaskMov 함수는 Task에 mov를 설정한다.
func handleAPISetTaskMov(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}
	type Recipe struct {
		Project string `json:"project"`
		Name    string `json:"name"`
		Task    string `json:"task"`
		Mov     string `json:"mov"`
		UserID  string `json:"userid"`
		Error   string `json:"error"`
	}
	rcp := Recipe{}
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()
	rcp.UserID, _, err = TokenHandler(r, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	r.ParseForm()
	for key, values := range r.PostForm {
		switch key {
		case "project":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rcp.Project = v
		case "name", "shot", "asset":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rcp.Name = v
		case "task":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rcp.Task = v
		case "mov": // 앞뒤샷 포함 여러개의 mov를 등록할 수 있다.
			rcp.Mov = strings.Join(values, ";")
		case "userid":
			v, err := PostFormValueInList(key, values)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if rcp.UserID == "unknown" && v != "" {
				rcp.UserID = v
			}
		default:
			http.Error(w, key+"키는 사용할 수 없습니다.(project, shot, asset, task, mov 키값만 사용가능합니다.)", http.StatusBadRequest)
			return
		}
	}
	rcp.Mov = dipath.Win2lin(rcp.Mov) // 내부적으로 모든 경로는 unix 경로를 사용한다.
	_, err = setTaskMov(session, rcp.Project, rcp.Name, rcp.Task, rcp.Mov)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// log
	err = dilog.Add(*flagDBIP, host, fmt.Sprintf("Setmov: %s %s", rcp.Task, rcp.Mov), rcp.Project, rcp.Name, "csi3", rcp.UserID, 180)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// slack log
	err = slacklog(session, rcp.Project, fmt.Sprintf("Setmov: %s %s\nProject: %s, Name: %s, Author: %s", rcp.Task, rcp.Mov, rcp.Project, rcp.Name, rcp.UserID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// json 으로 결과 전송
	data, _ := json.Marshal(rcp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
