/**
 * 文档代码
 * @author 黄承武
 */


/**
 * 文档初始化完毕
 */
function initcomplete(){
	var t1 = document.getElementById("dt");
	var t2 = document.getElementById("dt2");
	var date = new Date();
	
	var m = date.getMonth()+1;
	var d = date.getDate();
	var h = date.getHours();
	var min = date.getMinutes();
	
	var tim = date.getFullYear()+"-"+ (m<10?"0"+m:m ) +"-"+ (d<10?"0"+d:d ) +"T"+(h<10?"0"+h:h )+":"+(min<10?"0"+min:min );
//	console.log(tim);
	t1.value = tim;
	t2.value = tim;
}



/**
 * 发送请求
 * @param {Object} url 地址
 * @param {Object} data 数据（多个数据之间用逗号分隔）
 * @param {Object} callback（请求的回调函数）
 */
function requstData(url, data, callback){
	
	
	var xmlhttp;
    if (window.ActiveXObject) {
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    } else {
        xmlhttp = new XMLHttpRequest();
    }


//	console.log(data);
    xmlhttp.open("post", url , true);
    xmlhttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
    xmlhttp.send(data);
    
    xmlhttp.onreadystatechange = function(){
        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            var obj = JSON.parse(xmlhttp.responseText);
            if(callback!=null){
                callback(obj);
            }
        }
    }
    
}


/**
 * 按钮点击方法
 * @param {Object} data 数据
 */
function btnClick(data, fn){
	var infotext = document.getElementById("info");
	infotext.style.color = "#333333";
	infotext.innerText = "正在查询，请稍后...\n(注：由于库存储量很大，部分查询耗时比较长.)";
	
	
	if(data == "1003"){//查询DAU及留存指令
		var dtime = document.getElementById("dt");
		var dtime2 = document.getElementById("dt2");
		if(dtime.value == "" || dtime2.value == ""){
			infotext.innerText = "请求失败! 请设置查询及要比对的日期和时间.";
			infotext.style.color = "#FF0000";
			return;
		}
		var d1 = dtime.value.split("T")[0];
		var d2 = dtime2.value.split("T")[0];
		if (d1 == d2){
			infotext.innerText = "请求失败! 查询日期不能与对比日期相同.";
			infotext.style.color = "#FF0000";
			return;
		}
		var curVal = (Date.parse(new Date(d2)) - Date.parse(new Date(d1))) / 1000 / 60 / 60 / 24;
		if(curVal>0){
			infotext.innerText = "请求失败! 查询日期不能小于对比日期.";
			infotext.style.color = "#FF0000";
			return;
		}
		
		data += ","+dtime.value+","+dtime2.value;
	}
	
	
	//发送消息
	var d = "data=" + data;
	requstData(HOST_PORT + FORM, d, fn);
}




/**
 * 显示总注册人数
 * @param {Object} val
 */
function showRegs(val){
	var infotext = document.getElementById("info");
	infotext.innerText = "当前注册总人数：" + val.Data;
}




/**
 * 显示关卡分布人员分布情况
 * @param {Object} val
 */
function showCards(val){
	var infotext = document.getElementById("info");
	var data = val.Data;
	var len = data.length;
	
	var ary = [];
	for(var i=0; i<len; i++){
		ary.push("关卡："+i + " / 包含用户数：" + data[i]);
	}
	
	var s = ary.join("\n\n");
	infotext.innerText = s;
}



/**
 * 显示日登陆数据
 * @param {Object} val
 */
function showDAU(val){
//	console.log(val.Data);
	var infotext = document.getElementById("info");
	
	var s = "当前查询的日期登陆的用户数（DAU）：\n\n";
	if (val.Data.length==1){
		s += "查询日期无任何数据.";
		infotext.innerText = s;
		return;
	}
	
	
	s += "老用户登陆数：" + (val.Data[0] || "0")+ "\n";
	s += "新用户登陆数：" + (val.Data[1] || "0") + "\n";
	s += "总登陆数：" + (val.Data[2] || "0");
	s += "\n\n-----------------------------------\n\n";
	s += "与对比日期的留存：（查询日期登录的老用户包含对比日期的新增用户数 / 对比日期的新增用户数）\n\n";
	
	
	if(val.Data[5].length==0){
		s += "查询的对比日期无任何记录."
	}else{
		s += "对比日期的新增用户数：" + val.Data[5][1] + "\n";
		var num = 0;
		var ind = -1;
		for(var i=0; i<val.Data[0]; i++){
			ind = val.Data[5][4].indexOf(val.Data[3][i]);
			if(ind!=-1){ num++; };
		}
		
		s += "查询日期的老用户中包含对比日期的新增用户数量：" +num + "\n";
		s += "查询日期相对对比日期的留存为：" + ((num/val.Data[5][1]) * 100).toFixed(2) + "%";
	}
	infotext.innerText = s;
	
}



/**清理*/
function btnClickClear(){
	var infotext = document.getElementById("info");
	infotext.innerText = "暂无任何查询";
}