window.onload = function(){
	var Remind = function(){
		this.init();
	};
	Remind.prototype = {
		init: function(){
			if(!document.querySelector("#reminding")){
				var remind = document.createElement("div");
				remind.id = "reminding";
				remind.style.display = "none";
				remind.innerHTML = this.addHtml();
				document.querySelector("body").appendChild(remind);
				console.warn("add id name 'reminding' element");
				this.addCss();
			}
		},
		addCss: function(){
			var style = document.createElement("style");
			style.innerHTML = '#reminding{position:fixed;top:0;left:0;width:100%;height:100%;background:rgba(0,0,0,.6);z-index:4}#reminding>div{top:50%;position:absolute;padding:0 20px;width:100%;box-sizing:border-box;text-align:center;-webkit-transform:translateY(-50%);transform:translateY(-50%)}.pop-head{line-height:1.21875rem;font-size:.4375rem;color:#5babff;background:#fff}.pop-cont{background:#f5f5f9;font-size:.40625rem;padding:1.09375rem 0;color:#333;border-top:1px solid #eee;border-bottom:1px solid #eee}.pop-btns>button{float:left;width:100%;line-height:1.21875rem;color:#5babff;border:0;background:#fff;font-size:.4375rem;box-sizing:border-box}.pop-btns>button:first-child{border-right:1px solid #eee}';
			document.querySelector("head").appendChild(style);
		},
		addHtml: function(){
			return '<div><div class="pop-head">提示</div><div class="pop-cont"></div><div class="pop-btns"><button id="closeBtn">知道了</button></div></div>';
		},
		alert: function(msg, callback){
			if(!msg) return "msg is null";
			document.querySelector("#reminding .pop-cont").innerHTML = msg;
			document.querySelector("#reminding").style.display = "block";
			document.querySelector("#closeBtn").onclick = function(){
				document.querySelector("#reminding").style.display = "none";
				callback&&callback()
			}
		}
	}
	window["Remind"] = new Remind();
};