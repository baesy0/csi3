package main

import "gopkg.in/mgo.v2/bson"

// Review 는 리뷰데이터 자료구조 이다.
type Review struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"` // ID
	Project       string        `json:"project"`                 // 프로젝트
	Name          string        `json:"name"`                    // 샷네임, 에셋네임
	Task          string        `json:"task"`                    // 태스크
	Createtime    string        `json:"createtime"`              // 생성시간
	Updatetime    string        `json:"updatetime"`              // 업데이트 시간
	Author        string        `json:"author"`                  // 작성자
	AuthorNameKor string        `json:"authornamekor"`           // 작성자 한글 이름
	Path          string        `json:"path"`                    // 리뷰경로
	Status        string        `json:"status"`                  // 상태 approve, comment, waiting
	ProcessStatus string        `json:"processstatus"`           // 연산상태. wait, error, done
	Sketches      []Sketch      `json:"sketches"`                // 스케치 프레임
	Playlist      []string      `json:"playlist"`                // 플레이리스트 목록
	Comments      []Comment     `json:"comments"`                // 댓글
	Description   string        `json:"description"`             // 설명
	Progress      int           `json:"progress"`                // 진행률
	CameraInfo    string        `json:"camerainfo"`              // 카메라정보
	CreatedMp4    bool          `json:"createmp4"`               // Mp4 생성여부
	Fps           float64       `json:"fps"`                     // fps
	Log           string        `json:"log"`                     // Log 예로 Errlog가 있다.
	MainVersion   int           `json:"mainversion"`             // Main Version
	SubVersion    int           `json:"subversion"`              // Sub Version
}

// Sketch 는 스케치 자료구조이다.
type Sketch struct {
	Frame      int    `json:"frame"`      // 프레임수
	Duration   int    `json:"duration"`   // 스케치의 길이
	Author     string `json:"author"`     // 스케치를 그린사람
	SketchPath string `json:"sketchpath"` // 스케치 경로
	Createtime string `json:"createtime"` // 생성시간
	Updatetime string `json:"updatetime"` // 업데이트 시간
}
