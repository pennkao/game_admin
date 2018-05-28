//负责记录每日登陆用户数据服务（子服务，子线程）
//@author 黄承武


package web

import(
	"time"
	"fmt"
	"strings"
	"server_back/sql"
)


//启动
func StartRecordLoginServer(){
	fmt.Println("-> start record login server (child)...");

	go func(){
		for{
			now := time.Now();
			next := now.Add(time.Hour * 24);
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location());
			t := time.NewTimer(next.Sub(now));
			<-t.C;

			//执行存储
			execute();
		}
	}();
}



func execute(){
	now := time.Now();
	prev := now.AddDate(0,0,-1);
	// fmt.Println(prev);
	timestr := prev.Format("2006-01-02 15:04:05");
	// fmt.Println(timestr);
	times := strings.Split(timestr, " ");
	fmt.Println("record current date DAU: ",times[0]);

	sql.SetDAUCrtDate(times[0]);
}