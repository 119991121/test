// JavaScript
var btn = document.getElementsById('button1');
function run(){ 
  var time = 60;
  timer = setInterval(
    function(){ 
      btn.disabled = true;
      btn.innerText = time+"秒后重新发送";
      time--;
      if(time ==0){ 
      btn.innerText='重新发送验证码';
      btn.disabled = false;
      clearInterval(timer);
      } 
    }
	,6000) 
} 