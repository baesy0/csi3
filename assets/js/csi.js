var OSName="Linux";
if (navigator.appVersion.indexOf("Win") != -1) OSName="Win";
if (navigator.appVersion.indexOf("Mac") != -1) OSName="Mac";
if (navigator.appVersion.indexOf("X11") != -1) OSName="Linux";
if (navigator.appVersion.indexOf("Linux") != -1) OSName="Linux";

// onload 함수는 브라우저에서 테스트 할 때 에러가 많이 일어난다. 리뷰해본적이 없는 코드같다. 점검해보기.
function onload(){
	// projectId 변수 설정
	let project = "";
	if (document.getElementById("projectId")) {
		project = document.getElementById("projectId").innerHTML; //프로젝트명
	}

	// 검색 결과가 없을때 <title>내용에 페이지 정보를 출력한다.
	// items 페이지, 프로젝트명 출력.
	if (project !== "") {
		document.title = "CSI : " + projectId.innerHTML;
	// editItem 페이지, 프로젝트명과 샷네임(slug)출력.
	} else if (document.getElementById("projectTitle")) {
		document.title = document.getElementById("slugTitle").innerHTML  + " | " + document.getElementById("projectTitle").innerHTML;
	// editProject 페이지, 프로젝트명 출력.
	} else if (document.getElementById("editProjectTitle")) {
		document.title = "CSI : edit " + document.getElementById("editProjectTitle").innerHTML;
	// CSI 하위 메뉴 출력 (help,ProjectInfo)
	} else {
		var url = document.location.href;
		var path = new URL(url).pathname;
		var pageTitle = path.replace("/","");
		if (pageTitle != "") {
			document.title = "CSI : " + pageTitle;
		} else {
			document.title = "CSI";
		}
	}

	// 강조할 단어 리스트
	var wordList = ["2d_","3d_","concept_","model_","mm_","matchmove_","layout_","ani_","animation_","fx_","mg_","motion_","previz_","fursim_","sim_","crowd_","light_","comp_","matte_","env_","rig_","just_","dir_","sup_"];
	var wordListRe = new RegExp(wordList.join("|"), "gi");

	// 현장내용/히스토리 영역
	var onsets = document.getElementsByClassName('historyOnsetnote');

	for (i = 0; i < onsets.length; i++) {
		// 현장내용 Text에 \n문자가 있다면 줄바꿈을 한다.
		onsets[i].innerHTML = onsets[i].innerHTML.replace(/^\n/,"").replace(/\n/g,"<br>");
		// 현장내용의 Task와 팀정보를 강조한다.
		onsets[i].innerHTML = onsets[i].innerHTML.replace(wordListRe, function(match) {
			return "<highlight>" + match + "</highlight>";
		});
	}
}

//달력 .DatePicker Class에 적용된다.
$(function() {
	$( ".DatePicker" ).datepicker({
		dateFormat: 'yy-mm-dd'
	});
});

//checkbox all
function selectmode(){
	var onnum = 0
	// 체크가 되어있는 갯수를 구한다.
	for(var i=0; i<document.getElementsByClassName('StatusCheckBox').length;i++) {
		if (document.getElementsByClassName('StatusCheckBox')[i].checked == true) {
			onnum = onnum + 1
		}
	}

	if (onnum == 9) {
		for(var i=0; i<document.getElementsByClassName('StatusCheckBox').length;i++) {
			document.getElementsByClassName('StatusCheckBox')[i].checked=false
		}
	} else if (onnum == 0) {
		// 이 모드는 자주 사용하는 사용자 선택패턴이다.
		document.getElementsByClassName('StatusCheckBox')[0].checked=true // assign
		document.getElementsByClassName('StatusCheckBox')[1].checked=true // ready
		document.getElementsByClassName('StatusCheckBox')[2].checked=true // wip
		document.getElementsByClassName('StatusCheckBox')[3].checked=true // confirm
		document.getElementsByClassName('StatusCheckBox')[4].checked=false // done
		document.getElementsByClassName('StatusCheckBox')[5].checked=false // omit
		document.getElementsByClassName('StatusCheckBox')[6].checked=false // hold
		document.getElementsByClassName('StatusCheckBox')[7].checked=false // out
		document.getElementsByClassName('StatusCheckBox')[8].checked=false // none
	} else {
		for(var i=0; i<document.getElementsByClassName('StatusCheckBox').length;i++) {
			document.getElementsByClassName('StatusCheckBox')[i].checked=true
		}
	}
}


//샷 체크박스 F5 눌렀을때 없애는 기능.
function uncheck() {
	var checkboxes = document.getElementsByName('select_slug');
	for (var i=0; i<checkboxes.length; i++) {
		console.log(checkboxes[i].type)
		if (checkboxes[i].type == 'checkbox') {
			checkboxes[i].checked = false;
		}
	}
}

function changeProject(objS)
{
	document.links["assign"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/assign";
	document.links["ready"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/ready";
	document.links["wip"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/wip";
	document.links["confirm"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/confirm";
	document.links["done"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/done";
	document.links["omit"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/omit";
	document.links["hold"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/hold";
	document.links["out"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/out";
	document.links["none"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/none";
	document.links["2d"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/2d";
	document.links["3d"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/3d";
	document.links["asset"].href = "/tag/" + objS.options[objS.selectedIndex].value + "/asset";
	var csititle = "CSI: "
	document.title = csititle.concat(objS.options[objS.selectedIndex].value);
}


function playmov(address)
{
	myWindow=window.open(address,"PlayWindows","width=1280, height=720, toolbar=no, menubar=no, location=no");
	myWindow.focus();
}

function keypressed(){
	if(event.keyCode==122) self.close();
	else return false;
}
document.omkeydown=keypressed;


function onlyNumber(event) {
	event = event || window.event;
	var keyID = (event.which) ? event.which : event.keyCode;
	if ( (keyID >=48 && keyID <= 57) || (keyID >= 96 && keyID <= 105) || keyID == 8 || keyID == 46 || keyID == 37 || keyID == 39)
		return;
	else
		return false;
}

function removeChar(event) {
	event = event || window.event;
	var keyID = (event.which) ? event.which : event.keyCode;
	if (keyID == 8 || keyID == 46 || keyID == 37 || keyID == 39)
		return;
	else
		event.target.value = event.target.value.replace(/[^0-9]/g,"");
}


function removeWhiteSpace(event) {
	event.value = event.value.replace(/ /g, '');
}

// *문자를 x문자로 바꾼다.
// X를 x문자로 바꾼다.
// 공백을 제거한다.
// 렌즈디스토션값을 입력시 2048*1280 -> 2048x1280 형태로 바꾸기 위함이다.
// 숫자와 x를 제외한 영문입력시 삭제됩니다.
function widthxHeight(event) {
	event = event || window.event;
	event.target.value = event.target.value.replace("*","x");
	event.target.value = event.target.value.replace("X","x");
	event.target.value = event.target.value.replace(/[^\d\x]/gi,"");
}

//drop된 file의 "file://" 제거
function rmFileProtocol(event) {	
	event = event || window.event;
	event.preventDefault();
	
	var data= event.dataTransfer.getData('text/plain'); //드래그한 데이터 자료를 얻는다.
	event.target.value = "";
	event.target.value = data.replace("file://","");
}

//버튼을 클릭 하면 editItem 언디스토션사이즈 form에 placesize(scansize) 값을 입력한다.
//Checkbox를 체크하면 undistort value에 platesize가 입력된다.
//Checkbox의 체크를 해제하면 undistort value가 ""이 된다.
function inputNone(checkbox) {
	if (checkbox.checked) {
		document.getElementById("undistort").value = document.getElementById("platesize").value;
	} else {
		document.getElementById("undistort").value = "";
	}
}

// ScreenX 버튼이 클릭될때 체크 여부에 따라 이벤트가 발생한다.
// ScreenX가 체크되면 ScreenxOverlay가 활성화 된다.
// ScreenX가체크가 해제되면 ScreenxOverlay가 비활성화되고 1.0의 값이 들어간다.
function checkScreenx(event) {
	event = event || window.event;
	if (event.target.checked) {
		document.getElementById("screenxoverlay").readOnly = false;
	} else {
		document.getElementById("screenxoverlay").readOnly = true;
		document.getElementById("screenxoverlay").value = 1.0;
	}
}