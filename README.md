# CSI3

![travisCI](https://secure.travis-ci.org/digital-idea/csi3.png)

![screenshot](figures/screenshot.png)

프로젝트 매니징을 도와주는 웹 어플리케이션 입니다.

속도, 간결함(검색어 Based), 교육의 최소화, 단일파일 배포를 중점으로 개발되고 있습니다.

### 다운로드

- [Linux 64bit](https://github.com/digital-idea/csi3/releases/download/v3.0.1/csi3_linux_x86-64.tgz)
- [Linux 64bit for Digitalidea](https://github.com/digital-idea/csi3/releases/download/v3.0.1/csi3_linux_di_x86-64.tgz): 회사가 필요한 인수가 자동으로 설정되어 있습니다.
- [macOS 64bit](https://github.com/digital-idea/csi3/releases/download/v3.0.1/csi3_darwin_x86-64.tgz)
- [macOS 64bit for Digitalidea](https://github.com/digital-idea/csi3/releases/download/v3.0.1/csi3_darwin_di_x86-64.tgz): 회사가 필요한 인수가 자동으로 설정되어 있습니다.

> 윈도우즈 경우 Server가 아니면 동시접속자수가 제한되어 있습니다. 서버로서 가치가 적어서 빌드하지 않습니다.

### Roadmap
- [x] 빌드환경 구축 ![build](http://progressed.io/bar/100)
- [ ] 유저자료구조 추가 ![user](http://progressed.io/bar/30)
- [ ] CSI2 restAPI 추가
- [ ] multi Task 기능추가
- [ ] 3D 파이프라인과 필요한 자료구조 및 API 생성(샷트레킹, 버전)

### mongoDB 설치, 실행

CentOS
```bash
$ sudo yum install mongodb mongodb-server
$ sudo service mongod start
```

macOS
```bash
$ brew install mongodb
$ brew services start mongodb
```

### CSI 실행

```bash
$ sudo csi3 -http :80
```

> macOS는 아파치 서버가 Built-in 되어있습니다. 80포트를 사용하기 위해서는 `$ sudo apachectl stop` 명령어를 입력해주세요.

csi는 [wfs](https://github.com/digital-idea/wfs), [dilog](https://github.com/digital-idea/dilog), [dilink](https://github.com/digital-idea/dilink)과 연동됩니다. 아래 서비스도 같이 실행해주세요.

```bash
$ dilog -http :8080
$ wfs -http :8081
```

### 명령어 사용법
- 프로젝트 생성
```bash
$ csi3 -add project -name [projectname]
```

- 프로젝트 삭제: Root 계정에서만 작동됩니다.
```
# csi3 -rm project -name [projectname]
```

- 샷등록
DB값만 생성되며, 샷 폴더가 생성되지는 않습니다.

```bash
$ csi3 -add item -project [projectname] -name [SS_0010] -type [org]
```

- 샷,에셋 삭제
```bash
$ csi3 -rm item -project [projectname] -name [SS_0010] -type [org]
```

- 에셋등록 예: 에셋이 prop 타입이고 component 형태일 때
```
$ csi3 -add item -type asset -project [projectname] -name [Assetname] -assettype prop -assettags prop,component
```

- 2016-12-05 에 아티스트가 업로드한 mov를 rvplayer로 한번에 모아보기.
```bash
$ csi3 -date 2016-12-05 -play &
```

- 특정 프로젝트의 데일리 mov rvplayer로 모아보기.
```bash
$ csi3 -date 2016-12-05 -play -project [projectname] &
$ csi3 -date 2016-12-05 -play -project [projectname] -task model & // 해당 프로젝트의 model 테스크만 보기
```

- 사용자 패스워드 초기화
사용자의 패스워드를 `Welcome2csi!`로 초기화 하고 싶다면 아래 명령어를 사용합니다.
관리자만 처리할 수 있습니다.

```bash
$ sudo csi3 -initpass Welcom2csi! -id [userid]
```

- 사용자 제거

```bash
$ sudo csi3 -rm user -id [userid]
```

### 썸네일 경로
- 썸네일주소 : `/thumbnail/{projectname}/{slug}.jpg`
- 사용자이미지 : `/thumbnail/user/{id}.jpg`

### RestAPI
- [Project](documents/rest_project.md)
- [Item](documents/rest_item.md)
- [User](documents/rest_user.md)

### Process
- [디자인 프로세스](documents/process_designer.md)
- [개발 프로세스](documents/process_developer.md)
- [Onset Setellite](documents/setellite.md)
- [DB관리](documents/dbbackup.md)

### Developer
- Maintainer: HanwoongKim(hello@lazypic.org)
- Committer: 
- Contributors: 

### Infomation
- [History](documents/history.md): csi의 역사
- License: BSD 3-Clause License
- 회사 전용 빌드문의: hello@lazypic.org
- 참고: [CentOS7에서 mongoDB 설치](https://github.com/cgiseminar/curriculum/blob/master/docs/install_mongodb.md)
