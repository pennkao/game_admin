//http服务(主服务，主线程)
//@author 黄承武


package web

import(
	"net/http"
	"server_back/config"
	"fmt"
	"strconv"
	"server_back/datastruct"
	"encoding/json"
	"strings"
)


type CustomData struct{
	Data interface{}
}
//处理消息
//@private
func msgHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type");
	w.Header().Set("content-type", "application/json");

	//表单形式读取数据
	data := r.FormValue("data");
	datas := strings.Split(data, ",");

	code, _ := strconv.Atoi( datas[0] );
	bol := CheckMsgcode( code );
	if !bol {
		errmsg := datastruct.MsgError{};
		errmsg.Code = ERR_MSG_CODE_1001;
		errmsg.Errmsg = "message code error! code = "+strconv.Itoa(code);
		fmt.Println(errmsg.Errmsg);
		result, _ := json.Marshal(errmsg);
		fmt.Fprintf(w, string(result));
		return;
	}

	fmt.Println("received new message. code = " + strconv.Itoa(code));
	DoHttphanlder(w, datas);
}




//启动http服务器
func StartHttpserver(){

	//注册http请求处理函数
	http.HandleFunc("/backMsgHandler", msgHandler);

	fmt.Println("-> start http server (main)...");

	//启动每日登陆记录服务（子线程）
	StartRecordLoginServer();

	//监听http
	err := http.ListenAndServe(config.HTTP_URL, nil);
	if err != nil {
		fmt.Println("-> [Error] : start http server fail! err = ", err.Error());
	}
}