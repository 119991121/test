// JavaScript Document
function changeColor1(x) {//光标移动高亮事件
    x.style.background = '#3071a9';
}
function changeColorBack1(x) {
    x.style.background = '#428bca';
}
function changeColor2(x) {
    x.style.background = '#449d44';
}
function changeColorBack2(x) {
    x.style.background = '#5cb85c';
}
function changeColor3(x) {
    x.style.background = '#c9302c';
}
function changeColorBack3(x) {
    x.style.background = '#d9534f';
}
function changeColor4(x) {
    x.style.background = '#449d44';
}
function changeColorBack4(x) {
    x.style.background = '#5cb85c';
}
function changeColor5(x) {
    x.style.background = '#7e7c71';
}
function changeColorBack5(x) {
    x.style.background = '#e4dbc0';
}
var zoom=function(){//点击按钮改变frame1高度
	var t=document.getElementById("frame1");
	t.style.height="500px";
	}
var del=function(){//点击按钮消除frame4
	var d=document.getElementById("frame4");
	var a=d.lastChild;
	d.remove(a);
	}
	
  var user = document.getElementById('user'),//账号记忆功能（未测试是否实现）
      localUser = localStorage.getItem('user') || ''; 
      user.value = localUser;
      if (localUser !== '') {
        check.setAttribute('checked', '');
      }

    function fn() {
      if (check.checked) {
        localStorage.setItem('user', user.value);
      } else {
        localStorage.setItem('user', '');
      }
    }
