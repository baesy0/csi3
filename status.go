package main

import "errors"

// Status 자료구조는 상태 자료구조이다.
type Status struct {
	ID          string  `json:"id"`          // ID, 상태코드
	TextColor   string  `json:"textcolor"`   // TEXT 색상
	BGColor     string  `json:"bgcolor"`     // BG 상태 색상
	BorderColor string  `json:"bordercolor"` // Border 색상
	Description string  `json:"description"` // 설명
	Order       float64 `json:"order"`       // Status 우선순위
	DefaultOn   bool    `json:"defaulton"`   // 기본선택 여부
	InitStatus  bool    `json:"initstatus"`  // 아이템 생성시 최초 설정되는 Status 설정값
}

// CheckError 메소드는 Status 자료구조의 에러를 체크한다.
func (s *Status) CheckError() error {
	if s.ID == "" {
		return errors.New("ID가 빈 문자열 입니다")
	}
	if !regexWebColor.MatchString(s.TextColor) {
		return errors.New("웹컬러 문자열이 아닙니다")
	}
	if !regexWebColor.MatchString(s.BGColor) {
		return errors.New("웹컬러 문자열이 아닙니다")
	}
	if !regexWebColor.MatchString(s.BorderColor) {
		return errors.New("웹컬러 문자열이 아닙니다")
	}
	return nil
}
