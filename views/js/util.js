//访问地址，请勿改动  zhouyong
var url = "{request_url}";
var jsonpurl = "http://linjie.ngrok.cc/pws";
// var jsonpurl = "http://127.0.0.1:8080/pws";

function loadHtml(url,target){
    var args = '';
    if(url.split("?")[1]){
        args = '<div id=\'args\' style=\'display: none;\' data-args=\''+url.split("?")[1]+'\'></div>';
    }
    $.ajax({
        type: "get",
        url: url,
        dataType: "html",
        async: false,
        success: function (data) {
            $(target).html(args+data);
        }
    });
}
var browser={
    versions:function(){
        var u = navigator.userAgent, app = navigator.appVersion;
        return {//移动终端浏览器版本信息
            trident: u.indexOf('Trident') > -1, //IE内核
            presto: u.indexOf('Presto') > -1, //opera内核
            webKit: u.indexOf('AppleWebKit') > -1, //苹果、谷歌内核
            gecko: u.indexOf('Gecko') > -1 && u.indexOf('KHTML') == -1, //火狐内核
            mobile: !!u.match(/AppleWebKit.*Mobile.*/)||!!u.match(/AppleWebKit/), //是否为移动终端
            ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), //ios终端
            android: u.indexOf('Android') > -1 || u.indexOf('Linux') > -1, //android终端或者uc浏览器
            iPhone: u.indexOf('iPhone') > -1 || u.indexOf('Mac') > -1, //是否为iPhone或者QQHD浏览器
            iPad: u.indexOf('iPad') > -1, //是否iPad
            webApp: u.indexOf('Safari') == -1 //是否web应该程序，没有头部与底部
        };
    }(),
    language:(navigator.browserLanguage || navigator.language).toLowerCase()
};

/* 
* 微信图片浏览
* curSrc:  当前图片url
* srcList: 图片list
 */
function imagePreview(curSrc,srcList) {
    if(!curSrc || !srcList || srcList.length == 0) {
        return;
    }
    WeixinJSBridge.invoke('imagePreview', {
        'current' : curSrc,
        'urls' : srcList
    });
}
/*是否微信浏览器*/
function isWechatBrower(){
    var ua = navigator.userAgent.toLowerCase();
    if(ua.match(/MicroMessenger/i)=="micromessenger") {
        return true;
    } else {
        return false;
    }
}

/*从地址栏中获取参数*/
function getQueryString(name) { 
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i"); 
    var r = window.location.search.substr(1).match(reg); 
    if (r != null) return unescape(r[2]); return null; 
}
/*
* 从url获取参数
* var foo = getArgs()["foo"]
*/
function getArgs(){
    var args = {};
    var match = null;
    var search = decodeURIComponent(location.search.substring(1));
    var reg = /(?:([^&]+)=([^&]+))/g;
    while((match = reg.exec(search))!==null){
        args[match[1]] = match[2];
    }
    return args;
}

/*
* AJAX POST请求
* url:      接口地址
* post:     接口名(接口名前需要加"/")
* data:     参数
* callback: 回调函数
*/
function postAjax(port, data, callback){
    var data = {};
    $.ajax({
        type        : 'POST',
        url         : url + port,
        data        : data,
        cache       : false,
        timeout     : 20000,
        contentType : "application/json; charset=utf-8",
        success     : function (r) {
            callback(r);
        },
        error       : function(err){
            console.log(err);
        }
    })
}

/*
* AJAX GET跨域请求
* jsonpurl: 接口地址
* port:     接口名(接口名前需要加"/")
* data:     参数
* callback: 回调函数
*/
function jsonpAjax(port, data, callback){
	Ajax.init({
		type        : 'GET',
		url         : jsonpurl+port,
		dataType    : 'jsonp',
		jsonp       : 'jsoncallback',
		data        : data,
		time        : 20000,
		contentType : 'application/json; charset=utf-8',
		success:function(r){
		    callback(r)
		},
		error:function(e){
		    console.error(e);
		}
	});
}

/*
* AJAX POST请求
* url:      接口地址
* port:     接口名(接口名前需要加"/")
* data:     参数
* callback: 回调函数
*/
function jsonAjax(port, data, callback){
	Ajax.init({
		type        : 'POST',
		url         : url+port,
		dataType    : 'json',
		data        : data,
		time        : 20000,
		contentType : 'application/json; charset=utf-8',
		success:function(r){
		    callback(r)
		},
		error:function(e){
			// 错误回调
		    console.error(e);
		    !!vm?vm.isLoading=false:null;
		    alert("当前网络不稳定，请求失败。");
		}
	});
}

/*
* 封装AJAX
*/
;(function(){
	var Ajax=function(params){
		this.config={
			url:"",
			type:"get",
			async:true,
			dataType:"json",
			contentType:"application/x-www-form-urlencoded; charset=UTF-8",
			data:{}
		};
		this.start(params);
	};
	var xhr = null;
	Ajax.init=function(params){
		new Ajax(params);
	};
	Ajax.prototype={
		constructor: Ajax,
		createXHR:function(){
			if(typeof XMLHttpRequest!='undefined'){
				return new XMLHttpRequest();
			}else if(typeof ActiveXObject!='undefined'){
				if(typeof arguments.callee.activeXString!='string'){
					var versions=["MSXML2.XMLHttp.6.0","MSXML2.XMLHttp.3.0","MSXML2.XMLHttp"],i,len;
					for(i=0,len=versions.length;i<len;i++){
						try{
							new ActiveXObject(versions[i]);
							arguments.callee.activeXString=versions[i];
							break;
						}catch(ex){

						}
					}
				}
				return new ActiveXObject(arguments.callee.activeXString);
			}else{
				throw new Error("No XHR object available.");
			}
		},
		start:function(params){
			xhr=new this.createXHR();
			if(params.url){
				this.config.url=params.url;
			}else{
				throw new Error("url cannot be null!");
			}
			if(params.type){
				this.config.type=params.type;
			}
			if(params.async){
				this.config.async=params.async;
			}
			if(params.dataType){
				this.config.dataType=params.dataType;
			}
			if(params.data){
				this.config.data=params.data;
			}
			if(params.jsonp){
				this.config.jsonp=params.jsonp;
			}
			if(params.success){
				this.config.success=params.success;
			}
			if(params.error){
				this.config.error=params.error;
			}
			if(params.beforeSend){
				params.beforeSend();
			}

			var complete=function(){
				if(xhr.readyState==4){
						if((xhr.status>=200&&xhr.status<300)||xhr.status==304){
							if(params.success){
								params.success(JSON.parse(xhr.responseText));
							}
						}else{
							if(params.error){
								params.error();
							}else{
								throw new Error("Request was unsucessful:"+xhr.status);
							}
						}
				}
			}

			if(this.config.dataType.toLowerCase()=="json"){//非跨域
				if(this.config.type.toLowerCase()=="get"){
					for(var item in this.config.data){
						this.config.url=addURLParam(this.config.url,item,this.config.data[item]);
					}
					xhr.onreadystatechange=complete;
					xhr.open(this.config.type,this.config.url,this.config.async);
					xhr.send(null);
				}
				if(this.config.type.toLowerCase()=="post"){
					xhr.addEventListener('readystatechange',complete);
					xhr.open(this.config.type,this.config.url,this.config.async);
					if(params.contentType){
						this.config.contentType=params.contentType;
					}
					xhr.setRequestHeader("Content-Type",this.config.contentType);
					xhr.send(serialize(this.config.data));
				}
			}else if(this.config.dataType.toLowerCase()=="jsonp"){//跨域
				if(this.config.type.toLowerCase()=="get"){//jsonp只能进行get请求跨域
					if(!params.url){
						throw new Error("url is null!");
					}
					//创建script标签
				    if (window["count"]) {
				        window["count"]++;
				    } else {
				        window["count"] = 1;
				    }
				    // cbName 唯一jsonp请求
				    var cbName = "jsonp"+new Date().getTime()+window['count'];

					var head=document.getElementsByTagName('head')[0];
					this.config[this.config.callback]=cbName;
					var scriptTag=document.createElement('script');
					head.appendChild(scriptTag);

					//创建jsonp的回调函数
					window[cbName]=function(json){
						head.removeChild(scriptTag);
						clearTimeout(scriptTag.timer);
						window[cbName]=null;
						params.success&&params.success(json);
					};
					//超时处理
					if(params.timeout){
						scriptTag.timer=setTimeout(function(){
							head.removeChild(scriptTag);
							params.error&&params.error({message:"timeout"});
							window[cbName]=null;
						},params.timeout);
					}
					this.config.url=this.config.url+"?"+(this.config.jsonp?this.config.jsonp:'callback')+"="+cbName;
					for(var item in this.config.data){
						this.config.url=addURLParam(this.config.url,item,this.config.data[item]);
					}
                    scriptTag.src=this.config.url;
				}
			}else{
				throw new Error("dataType is error!");
			}
		}
	}
	function addURLParam(url,name,value){
		url+=(url.indexOf("?")==-1 ? "?" : "&");
		url+=encodeURIComponent(name)+"="+encodeURIComponent(value);
		return url;
	}
	//序列化函数
	function serialize(data){
		var val="";
		var str="";
		for(var item in data){
			str=item+"="+data[item];
			val+=str+'&';
		}
		return val.slice(0,val.length-1);
	}
	window["Ajax"]=Ajax;
})();