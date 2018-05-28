//http消息处理
//@author 黄承武

package web

import (
	"strings"
	// "strings"
	"time"
	"strconv"
	"net/http"
	"server_back/sql"
	"fmt"
	"encoding/json"
	"server_back/datastruct"
)



//处理各种http消息
//@param w 回执消息操作对象
//@param data 数据
func DoHttphanlder(w http.ResponseWriter, data []string){
	val, _ := strconv.Atoi(data[0])
	switch val{
	case MSG_CODE_1001: //拉取注册人数
		doGetRegs(w, val);

	case MSG_CODE_1002: //拉取每个关卡的用户分布
		doGetCardDistribution(w, val);

	case MSG_CODE_1003: //拉取每日用户登陆数据
		doGetDayLoginNumber(w, data);
	}
}




//拉取注册人数
func doGetRegs(w http.ResponseWriter, c int){
	data := datastruct.Msg{};
	data.Code = c;
	data.Data = sql.GetRegs();
	// fmt.Println(data.Data);
	msg, _ := json.Marshal(data);
	fmt.Fprintf(w, string(msg));
}



//拉取每个关卡的用户分布
func doGetCardDistribution(w http.ResponseWriter, c int){
	data := datastruct.Msg{};
	data.Code = c;
	data.Data = sql.GetCardDistribution();
	msg, _ := json.Marshal(data);
	fmt.Fprintf(w, string(msg));
}



//拉取每日用户登陆数据
func doGetDayLoginNumber(w http.ResponseWriter, d []string){
	//取出当前日期
	crtdate := time.Now().Format("2006-01-02 15:04:05");
	crtdate_ary := strings.Split(crtdate, " ");

	//参数时间
	rdate := strings.Split(d[1], "T");//选择日期
	rdate2 := strings.Split(d[2], "T");//对比日期

	//创建消息体
	data := datastruct.Msg{};
	cod , _ := strconv.Atoi(d[0]);
	data.Code = cod;

	if crtdate_ary[0] == rdate[0] { //查询日期等于当前日期，计算当前查询时登陆的用户数
		// fmt.Println("crtdata");
		v1 := sql.GetDAUCrtDate(crtdate_ary[0]);
		v1 = append(v1, sql.GetDAUForeDate(rdate2[0]));
		data.Data = v1;

	}else{ //查询日期不等于当前日期，则从数据库中查询数据
		// fmt.Println("other data");
		v2 := sql.GetDAUForeDate(rdate[0]);
		v2 = append(v2, sql.GetDAUForeDate(rdate2[0]));
		data.Data = v2;
	}

	msg, _ := json.Marshal(data);
	fmt.Fprintf(w, string(msg));
}