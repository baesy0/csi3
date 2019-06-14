package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

// handleCmd 함수는 /cmd URI를 통해서 입력받는 값을 이용해서 필요한 명령을 웹으로 수행하는 함수이다.
func handleCmd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"DB에 접속할 수 없습니다.\"}\n")
		return
	}
	defer session.Close()
	q := r.URL.Query()
	funcname := q.Get("funcname")
	project := q.Get("project")
	switch funcname {
	case "rmoverlaponsetnote":
		err = RmOverlapOnsetnote(session, project)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprintln(w, project+" 프로젝트의 작업내용 중복값이 제거되었습니다.")
		return
	default:
		fmt.Fprintf(w, "%s 이름으로 선언된 웹 함수가 없습니다.", funcname)
		return
	}
}
