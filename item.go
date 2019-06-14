package main

import (
	"sort"
)

// TASKS 는 작업에 사용되는 테스크 리스트이다.
var TASKS = []string{"Model", "Mm", "Ani", "Fx", "Mg", "Fur", "Sim", "Light", "Comp", "Matte", "Crowd", "Layout", "Env", "Temp1", "Concept", "Previz"}

const (
	// CLIENT 클라이언트 컨펌상태
	CLIENT = "9"
	// OMIT 작업취소 상태
	OMIT = "8"
	// CONFIRM 내부 컨펌상태
	CONFIRM = "7"
	// WIP 작업중 상태
	WIP = "6"
	// READY 작업준비중 상태
	READY = "5"
	// ASSIGN 작업자 배정을 기다리는 상태
	ASSIGN = "4"
	// OUT 외주상태
	OUT = "3"
	// DONE 작업완료 상태
	DONE = "2"
	// HOLD 작업중단 상태
	HOLD = "1"
	// NONE 상태없음. 예) 소스
	NONE = "0"
)

// Infobarnum 검색이후 나온 각 상태별 수를 담기위한 자료구조이다.
type Infobarnum struct {
	Assign  int
	Ready   int
	Wip     int
	Confirm int
	Done    int
	Omit    int
	Hold    int
	Out     int
	None    int
	Total   int
	Search  int
	Shot    int
	Shot2d  int
	Shot3d  int
	Assets  int
}

// Media 자료구조는 미디어 형식을 담기위한 자료구조이다.
type Media struct {
	ID       int       // 미디어 아이디. 추후 미디어 아이디를 발급받게 되면, 사용할 예정이다.
	Path     string    // 미디어 경로
	Versions Version   // 버전정보
	Comments []Comment // 코맨트
}

// Comment 자료구조는 답글을 남길 때 사용하는 자료구조이다.
type Comment struct {
	Date   string // 코맨트 등록시간 RFC3339
	Author string // 작성자
	Title  string // 제목
	Body   string // 내용
}

// Version 자료구조는 버전정보를 담을 때 사용하는 자료구조이다.
type Version struct {
	Main int // 메인버전 v01 또는 v01_w2 형태에서 앞부분 "1" 이다.
	Sub  int // 서브버전 v02_w02 형태에서 뒷부분 "2" 이다.
}

// Item 자료구조는 하나의 항목에 대한 자료구조이다.
type Item struct {
	Project string `json:"project"` // 프로젝트명
	ID      string `json:"id"`      // id. 추후 slug를 대체한다.

	// 현장정보
	// 현장에서 사용하는 카메라 데이터 이름. 슈퍼바이저 툴과 연동하기 위해서 Key로 사용된다.
	// 일반적으로 스캔이름과 같지만 항상 동일하지 않다.
	// 카메라 데이터 A037C012_160708_R717.[11694428-1172226].ari 형태에서 A037C012_160708_R717 부분이 데이터 이름이다.
	Dataname   string `json:"dataname"`   // 영화카메라(Red,Alexa등)이 자동 생성시키는 이미지 파일명이다.
	OnsetScene string `json:"onsetscene"` // 현장툴 TakeD와 맞출 Scene 문자열
	OnsetCut   string `json:"onsetcut"`   // 현장툴 TakeD와 맞출 Cut 문자열

	// 작업이 필요한 정보
	Scanname    string   `json:"scanname"`    // 스캔이름
	Platesize   string   `json:"platesize"`   // 플레이트 이미지사이즈
	Name        string   `json:"name"`        // 샷이름 SS_0010
	Seq         string   `json:"seq"`         // 시퀀스이름 SS_0010 에서 SS문자에 해당하는값. 에셋이면 "" 문자열이 들어간다.
	Type        string   `json:"type"`        // org, org1, src, asset..
	UseType     string   `json:"usetype"`     // 재스캔상황시 실제로 사용해야하는 타입표기
	Scantime    string   `json:"scantime"`    // 스캔 등록시간 RFC3339
	Thumpath    string   `json:"thumpath"`    // 썸네일경로
	Thummov     string   `json:"thummov"`     // 썸네일 mov 경로
	Beforemov   string   `json:"beforemov"`   // 전에 들어갈 mov. 만약 2개 이상이라면 space로 구분한다.
	Aftermov    string   `json:"aftermov"`    // 후에 들어갈 mov. 만약 2개 이상이라면 space로 구분한다.
	Retimeplate string   `json:"retimeplate"` // 리타임 플레이트 경로
	Platepath   string   `json:"platepath"`   // 플레이트 경로
	Slug        string   `json:"slug"`        // Name + Type이며 DB내부 컬렉션에서 고유ID로 활용한다.
	Shottype    string   `json:"shottype"`    // 2D, 3D
	Onsetnote   []string `json:"onsetnote"`   // 현장내용, 작업내용
	Ddline3d    string   `json:"ddline3d"`    // 3D 데드라인 RFC3339
	Ddline2d    string   `json:"ddline2d"`    // 2D 데드라인 RFC3339
	Rnum        string   `json:"rnum"`        // 롤넘버, 영화를 권으로 나누었을 때 이 샷의 권 번호. A는 1권을 H는 8권을 의미한다.
	Tag         []string `json:"tag"`         // 태그리스트
	Assettags   []string `json:"assettags"`   // 에셋그룹 태그
	Pmnote      []string `json:"pmnote"`      // PM 수정사항
	Finver      string   `json:"finver"`      // 파이널된 버젼
	Clientver   string   `json:"clientver"`   // 클라이언트에게 보낸 버전
	Findate     string   `json:"findate"`     // 파이널 데이터가 나간 날짜
	Outinfo     string   `json:"outinfo"`     // 삭제예정 : 외주가 나갔다면 외주회사정보.
	Rstate      string   `json:"rstate"`      // 삭제예정 : 0,1 재스캔상태
	Link        []string `json:"link"`        // 링크된 자료구조
	Linkslug    []string `json:"linkslug"`    // 링크된 slug리스트
	Dsize       string   `json:"dsize"`       // 디스토션 사이즈
	Rendersize  string   `json:"rendersize"`  // 특수상황시 렌더사이즈. 예) 5k플레이트를 3D에서 2k영역만 잡아서 최종 아웃풋까지 이어질 때
	Status      string   `json:"status"`      // 샷 상태.
	Assettype   string   `json:"assettype"`   // char, env, global, prop, comp, plant, vehicle, group 추후 Assettags로 사용된다.
	Updatetime  string   `json:"updatetime"`  // 업데이트 시간 RFC3339
	Finname     string   `json:"finname"`     // 파이널 파일이름
	Focal       string   `json:"focal"`       // 렌즈 미리수
	Stereotype  string   `json:"stereotype"`  // parallel(*), conversions
	Stereoeye   string   `json:"stereoeye"`   // left(*), right
	Outputname  string   `json:"outputname"`  // 프로젝트중 클라이언트가 제시하는 아웃풋 이름
	OCIOcc      string   `json:"ociocc"`      // Neutural Grading Pipeline에 사용하는 .cc 파일의 경로.

	//시간에 관련된 데이터이다.
	ScanKeycodeIn   string `json:"scankeycodein"`   // 제거예정. 스캔플레이트 키코드 In
	ScanKeycodeOut  string `json:"scankeycodeout"`  // 제거예정. 스캔플레이트 키코드 Out
	JustKeycodeIn   string `json:"justkeycodein"`   // 제거예정. 저스트 키코드 In
	JustKeycodeOut  string `json:"justkeycodeout"`  // 제거에정. 저스트 키코드 Out
	ScanFrame       int    `json:"scanframe"`       // 스캔 프레임수
	ScanTimecodeIn  string `json:"scantimecodein"`  // 스캔플레이트 타임코드 In
	ScanTimecodeOut string `json:"scantimecodeout"` // 스캔플레이트 타임코드 Out
	ScanIn          int    `json:"scanin"`          // 스캔 Frame In
	ScanOut         int    `json:"scanout"`         // 스캔 Frame Out
	HandleIn        int    `json:"handlein"`        // 핸들 Frame In
	HandleOut       int    `json:"handleout"`       // 핸들 Frame Out
	JustIn          int    `json:"justin"`          // 저스트 Frame In
	JustOut         int    `json:"justout"`         // 저스트 Frame Out
	JustTimecodeIn  string `json:"justtimecodein"`  // 저스트 타임코드 In
	JustTimecodeOut string `json:"justtimecodeout"` // 저스트 타임코드 Out
	PlateIn         int    `json:"platein"`         // 플레이트 Frame In
	PlateOut        int    `json:"plateout"`        // 플레이트 Frame Out

	//아래는 자주사용하지 않지만 역사가 만들어낸 자료구조이다.
	Soundfile string `json:"soundfile"` // 사운드파일 필요시 사운드파일 경로

	//task
	Model         Task                   `json:"model"`   // 모델링팀 정보.
	Fur           Task                   `json:"fur"`     // 털
	Mm            Task                   `json:"mm"`      // 매치무브
	Ani           Task                   `json:"ani"`     // 에니메이션
	Fx            Task                   `json:"fx"`      // FX
	Mg            Task                   `json:"mg"`      // 모션그래픽
	Light         Task                   `json:"light"`   // 라이팅
	Texture       Task                   `json:"texture"` // 텍스쳐
	Lookdev       Task                   `json:"lookdev"` // 룩뎁
	Comp          Task                   `json:"comp"`    // 합성
	Roto          Task                   `json:"roto"`    // 로토
	Prep          Task                   `json:"prep"`    // 입체작업전 프렙작업
	Stereo        Task                   `json:"stereo"`  // 입체작업
	Matte         Task                   `json:"matte"`   // 매트
	Env           Task                   `json:"env"`     // 환경
	Sim           Task                   `json:"sim"`     // 시뮬레이션
	Layout        Task                   `json:"layout"`  // 레이아웃
	Crowd         Task                   `json:"crowd"`   // 군중
	Temp1         Task                   `json:"temp1"`   // 기타1
	Temp2         Task                   `json:"temp2"`   // 기타2
	Concept       Task                   `json:"concept"` // 컨셉
	Previz        Task                   `json:"previz"`  // 프리비즈
	OnsetCam      `json:"onsetcam"`      // 현장 카메라 정보
	ProductionCam `json:"productioncam"` // 포스트 프로덕션 카메라 정보
	Rollmedia     string                 `json:"rollmedia"`   // 현장데이터의 Rollmedia 문자. 수동으로 현장데이터와 연결할 때 사용한다.
	ObjectidIn    int                    `json:"objectidin"`  // ObjectID 시작번호. Deep이미지의 DeepID를 만들기 위해서 파이프라인상 필요하다.
	ObjectidOut   int                    `json:"objectidout"` // ObjectID 끝번호. Deep이미지의 DeepID를 만들기 위해서 파인라인상 필요하다.
}

// Task 자료구조는 태크스 정보를 담는 자료구조이다.
type Task struct {
	User       string   `json:"user"`       // 아티스트명
	Status     string   `json:"status"`     // 상태
	Startdate  string   `json:"startdate"`  // 작업시작일 RFC3339
	Predate    string   `json:"predate"`    // 1차 마감일 RFC3339
	Date       string   `json:"date"`       // 마감일 RFC3339
	Mov        string   `json:"mov"`        // mov 경로
	Mdate      string   `json:"mdate"`      // mov 업데이트된 날짜 RFC3339
	Note       []string `json:"note"`       // 작업노트
	Movhistory []string `json:"movhistory"` // mov 히스토리
	Pubfile    string   `json:"pubfile"`    // Pubfile
	Due        int      `json:"due"`        // 예측 멘데이
	Promday    int      `json:"promday"`    // 실제멘데이
	Level      string   `json:"level"`      // 샷 레벨
	Title      string   `json:"title"`      // 테스크 네임. Temp1, Temp2의 표기 네임을 바꾸기 위해서 사용함.
	UserNote   string   `json:"usernote"`   // 아티스트와 관련된 엘리먼트등의 정보를 입력하기 위해 사용.
}

// updateStatus는 각 팀의 상태를 조합해서 샷 상태를 업데이트하는 함수이다.
func (item *Item) updateStatus() {
	tasks := []Task{
		item.Model,
		item.Fur,
		item.Mm,
		item.Ani,
		item.Fx,
		item.Mg,
		item.Light,
		item.Texture,
		item.Lookdev,
		item.Comp,
		item.Matte,
		item.Env,
		item.Sim,
		item.Layout,
		item.Crowd,
		item.Temp1,
	}
	maxstatus := "0"
	for _, t := range tasks {
		if t.Status > maxstatus {
			maxstatus = t.Status
		}
	}
	item.Status = maxstatus
}

// setRumTag는 특정 항목이 입력이 되었을때 알맞은 태그를 자동으로 넣거나 삭제할 때 사용한다.
// 예를 들어 "A0001" 이라는 롤넘버가 셋팅되면 태그리스트에 "1권" 이라는 단어를 넣어준다.
func (item *Item) setRnumTag() {
	if item.Rnum == "" {
		return
	}
	var rnumTag string
	switch item.Rnum[0] {
	case 'a', 'A':
		rnumTag = "1권"
	case 'b', 'B':
		rnumTag = "2권"
	case 'c', 'C':
		rnumTag = "3권"
	case 'd', 'D':
		rnumTag = "4권"
	case 'e', 'E':
		rnumTag = "5권"
	case 'f', 'F':
		rnumTag = "6권"
	case 'g', 'G':
		rnumTag = "7권"
	case 'h', 'H':
		rnumTag = "8권"
	}
	var newTags []string
	for _, t := range item.Tag {
		if !(t == "1권" || t == "2권" || t == "3권" || t == "4권" || t == "5권" || t == "6권" || t == "7권" || t == "8권") {
			newTags = append(newTags, t)
		}
	}
	newTags = append(newTags, rnumTag)
	sort.Strings(newTags)
	item.Tag = newTags
}

// 팀의 mov가 업데이트 되었다면 업데이트된 시간을 DB에 저장한다.
func (item *Item) updateMdate(olditem *Item) {
	if item.Model.Mov != olditem.Model.Mov {
		item.Model.Mdate = Now()
	}
	if item.Mm.Mov != olditem.Mm.Mov {
		item.Mm.Mdate = Now()
	}
	if item.Layout.Mov != olditem.Layout.Mov {
		item.Layout.Mdate = Now()
	}
	if item.Ani.Mov != olditem.Ani.Mov {
		item.Ani.Mdate = Now()
	}
	if item.Fx.Mov != olditem.Fx.Mov {
		item.Fx.Mdate = Now()
	}
	if item.Mg.Mov != olditem.Mg.Mov {
		item.Mg.Mdate = Now()
	}
	if item.Temp1.Mov != olditem.Temp1.Mov {
		item.Temp1.Mdate = Now()
	}
	if item.Fur.Mov != olditem.Fur.Mov {
		item.Fur.Mdate = Now()
	}
	if item.Sim.Mov != olditem.Sim.Mov {
		item.Sim.Mdate = Now()
	}
	if item.Crowd.Mov != olditem.Crowd.Mov {
		item.Crowd.Mdate = Now()
	}
	if item.Light.Mov != olditem.Light.Mov {
		item.Light.Mdate = Now()
	}
	if item.Comp.Mov != olditem.Comp.Mov {
		item.Comp.Mdate = Now()
	}
	if item.Matte.Mov != olditem.Matte.Mov {
		item.Matte.Mdate = Now()
	}
	if item.Env.Mov != olditem.Env.Mov {
		item.Env.Mdate = Now()
	}
}
