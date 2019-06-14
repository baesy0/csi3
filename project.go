package main

// Projectinfo DB에 들어가는 자료구조이다.

// ProjectStatus 는 숫자이다.
type ProjectStatus int

// 프로젝트 상태
const (
	UnknownProjectStatus = ProjectStatus(iota) // 0
	PreProjectStatus                           // 1
	PostProjectStatus                          // 2
	LayoverProjectStatus                       // 3 중단상태 예) 중간입금이 되지않아서 내부판단하 장기중단. ITEM자료구조에 HOLD 라는 값이 존재하여 LAYOVER 표현을 사용함.
	BackupProjectStatus                        // 4
	ArchiveProjectStatus                       // 5
	LawsuitProjectStatus                       // 6 소송상태
)

// Milestone 자료구조
type Milestone struct {
	Name string `json:"name"` // 마일스톤 이름
	Date string `json:"date"` // 날짜 RFC3339
}

// Mov 정보를 담는 자료구조
type Mov struct {
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	Codec         string  `json:"codec"`
	Fps           float64 `json:"fps"`
	InColorspace  string  `json:"incolorspace"`
	OutColorspace string  `json:"outcolorspace"`
}

// Project 정보를 담는 자료구조
type Project struct {
	ID                       string        `json:"id"`                       // 프로젝트 코드
	Name                     string        `json:"name"`                     // 프로젝트 한글이름
	MailHead                 string        `json:"mailhead"`                 // 이메일헤드 "[부산행]"
	Style                    string        `json:"style"`                    // 영화, 에니메이션, 광고, VR
	Stereo                   bool          `json:"stereo"`                   // 입체 프로젝트
	Screenx                  bool          `json:"screenx"`                  // 스크린X 프로젝트
	Director                 string        `json:"director"`                 // 감독
	Super                    string        `json:"super"`                    // 슈퍼바이져
	CgSuper                  string        `json:"cgsuper"`                  // CG 슈퍼바이져
	Pd                       string        `json:"pd"`                       // PD
	Pm                       string        `json:"pm"`                       // PM
	Pa                       string        `json:"pa"`                       // PA
	Message                  string        `json:"message"`                  // CSI 상단에 표시되는 공지사항
	Wiki                     string        `json:"wiki"`                     // 위키 URL
	Edit                     string        `json:"edit"`                     // 편집본 경로. 추후 URL
	AspectRatio              float64       `json:"aspectratio"`              // 픽셀 AspectRatio. 아나모픽 렌즈를 사용한 프로젝트는 AspectRatio가 다르다.
	Issue                    string        `json:"issue"`                    // 주요 CG내용. Preproduction 단계시 PostProject단계인 사람들에게 중요하게 표시되도록. - 부분장님 요청사항
	Camera                   string        `json:"camera"`                   // 촬영에 사용된 카메라
	PlateWidth               int           `json:"platewidth"`               // 아웃풋 플레이트 Width
	PlateHeight              int           `json:"plateheight"`              // 아웃풋 플레이트 Height
	ResizeType               string        `json:"resizetype"`               // 아웃풋 플레이트 레터박스 리사이즈타입.(fill:가로,세로자동판단, width:가로기준)
	PlateExt                 string        `json:"plateext"`                 // 아웃풋 플레이트 확장자. 간혹 mov -> exr 로 나가는 프로젝트가 있다.
	PlateInColorspace        string        `json:"plateincolorspace"`        // 아웃풋 플레이트 IN  컬러스페이스. 넘벳 프록시 이미지 렌더링시 사용된다.
	PlateOutColorspace       string        `json:"plateoutcolorspace"`       // 아웃풋 플레이트 OUT 컬러스페이스
	ProxyOutColorspace       string        `json:"proxyoutcolorspace"`       // 프록시 이미지 OUT 컬러스페이스(플레이트를 이용해서 프록시 이미지 생성시 사용할 Out 컬러스페이스)
	OutputMov                Mov           `json:"outputmov"`                // 아웃풋 Mov 포멧
	EditMov                  Mov           `json:"editmov"`                  // 편집실 Mov 포멧 - 레터박스 이슈로 아웃풋 Mov포멧과 다를때가 있다.
	Milestones               []Milestone   `json:"milestones"`               // CrankIn, CrankUp, 심의일, 시작일, 마감일, 기술시사, 칸, 예고편, 촬영종료, 촬영시작등에 해당하는 일정리스트
	Status                   ProjectStatus `json:"status"`                   // 프로젝트 상태코드가 들어간다. Preproduction -> Postproduction -> (Layover:중단) -> Backup -> Archive -> (Lawsuit:소송상태)
	Lut                      string        `json:"lut"`                      // 프로젝트 메인 LUT파일
	LutInColorspace          string        `json:"lutincolorspace"`          // 프로젝트 LUT IN  컬러스페이스
	LutOutColorspace         string        `json:"lutoutcolorspace"`         // 프로젝트 LUT OUT 컬러스페이스
	Description              string        `json:"description"`              // 필요한 자세한 설명
	Updatetime               string        `json:"updatetime"`               // 업데이트 시간
	StartFrame               int           `json:"startframe"`               // 시작프레임 회사는 1001로 시작함.
	VersionNum               int           `json:"versionnum"`               // 버전의 자릿수. 회사 기본 자릿수는 2자리. 외부 협력사와 작업시 3자리, 4자리도 간혹 보인다.
	SeqNum                   int           `json:"seqnum"`                   // 시퀀스 자릿수. 보통 4~8자리까지 다양하게 사용된다.
	Aeskey                   string        `json:"aeskey"`                   // 프로젝트 정보중 암호화가 필요한 부분에 사용할 AES키
	NukeGizmo                string        `json:"nukegizmo"`                // 슬레이트기즈모 경로
	CropAspectRatio          float64       `json:"cropaspectratio"`          // CropMask의 AspectRatio를 입력.
	PostProductionProxyCodec string        `json:"postproductionproxycodec"` // 이미지의 퀄리티가 상관없는 테스크에서 사용할 가벼운 코덱
	FxElement                string        `json:"fxelement"`                // 프로젝트에 사용하는 FX elemets 이다. 이 정보는 houdini pluto 에서 사용된다.
	MayaCropMaskSize         string        `json:"mayacropmasksize"`         // Maya CropMask에 사용되는 size 정보이다.
	HoudiniImportScale       float64       `json:"houdiniimportscale"`       // Houdini에서 사용하는 Import Scale 값입니다. 기본값은 0.1입니다.
	ScreenxOverlay           float64       `json:"screenxoverlay"`           // ScreenX에 사용되는 카메라 Overlay 값입니다. 기본값은 1.0입니다.
	Netapp                   bool          `json:"netapp"`                   // Netapp을 사용하는지 체크
	NoteHighlight            string        `json:"notehighlight"`            // onsetnote, pmnote에 입력되는 특정 단어를 강조하기 위해 사용된다.
}

// NewProject 함수는 기본 설정된 프로젝트 자료구조를 반환한다.
func NewProject(name string) *Project {
	return &Project{
		ID:                       name,
		Updatetime:               Now(),
		AspectRatio:              1.0,
		CropAspectRatio:          1.0,
		StartFrame:               1001,
		VersionNum:               2,
		SeqNum:                   4,
		ResizeType:               "fill",
		PostProductionProxyCodec: "Apple_Prores_422LT",
		FxElement:                "explosion,destruction,smoke,dust",
		HoudiniImportScale:       0.1,
		ScreenxOverlay:           1.0,
		Netapp:                   true,
	}
}
