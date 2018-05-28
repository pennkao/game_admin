//工具包
//@author 黄承武

package tool


import(
	"server_back/sql"
	"server_back/datastruct"
	"strconv"
	"time"
	"encoding/json"
	"fmt"
	"strings"
)


//写入测试数据（共写入3天数据。默认第1天为开服日）
//注：调用此函数时需要清理后台工具链接的数据库数据.
func WriteTestDatas(){
	c := sql.SqlOperate().Get();
	defer c.Close();
	
	//写入开服日数据
	c.Do("select", 0);
	var id=0;
	for i:=0; i<5; i++{
		id++
		uinfo := datastruct.UserBaseinfo{};
		uinfo.Id = id;
		uinfo.OpenId = "robot_" + strconv.Itoa(uinfo.Id);
		uinfo.Name = uinfo.OpenId;
		uinfo.Maxpoint = 0;
		uinfo.NewPlayer = true;
		uinfo.Createtime = time.Now();
		uinfo.Logintime = time.Now();

		data, _ := json.Marshal(uinfo);
		_, err2 := c.Do("rpush", uinfo.OpenId, data);
		if err2!=nil{
			fmt.Println("[error] : insert fail!");
		}
	}
	c.Do("select", 10);
	d1 := datastruct.DauInfo{};
	d1ary := make([]interface{}, 0);
	d1ary = append(d1ary, 0);
	d1ary = append(d1ary, 5);
	d1ary = append(d1ary, 5);
	d1ary = append(d1ary, make([]string, 0));
	d1ary = append(d1ary, []string{"robot_1","robot_2","robot_3","robot_4","robot_5"});
	d1.Info = d1ary;
	d1byte, _ := json.Marshal(d1);
	ttstr := time.Now().Format("2006-01-02 15:04:05");
	c.Do("rpush", strings.Split(ttstr, " ")[0], d1byte);

	//写入开服第2天数据
	c.Do("select", 0);
	for j:=0; j<3; j++{
		id ++;
		uinfo2 := datastruct.UserBaseinfo{};
		uinfo2.Id = id;
		uinfo2.OpenId = "robot_" + strconv.Itoa(uinfo2.Id);
		uinfo2.Name = uinfo2.OpenId;
		uinfo2.Maxpoint = 0;
		uinfo2.NewPlayer = true;
		t1 := time.Now();
		t2 := t1.AddDate(0, 0, 1);
		uinfo2.Createtime = t2;
		uinfo2.Logintime = t2;

		data2, _ := json.Marshal(uinfo2);
		_, err2 := c.Do("rpush", uinfo2.OpenId, data2);
		if err2!=nil{
			fmt.Println("[error] : insert fail!");
		}
	}
	for k:=0; k<2; k++{
		val, _ := c.Do("lindex", "robot_"+strconv.Itoa(k+1), 0);
		ino := datastruct.UserBaseinfo{};
		json.Unmarshal(val.([]byte), &ino);
		ino.NewPlayer = false;
		ino.Logintime = time.Now().AddDate(0,0,1);
		inodata, _ := json.Marshal(ino);
		c.Do("lset", "robot_"+strconv.Itoa(k+1), 0, inodata);
	}

	c.Do("select", 10);
	d2 := datastruct.DauInfo{};
	d2ary := make([]interface{}, 0);
	d2ary = append(d2ary, 2);
	d2ary = append(d2ary, 3);
	d2ary = append(d2ary, 5);
	d2ary = append(d2ary, []string{"robot_1","robot_2"});
	d2ary = append(d2ary, []string{"robot_6","robot_7","robot_8"});
	d2.Info = d2ary;
	d2byte, _ := json.Marshal(d2);
	ttstr2 := time.Now().AddDate(0,0,1).Format("2006-01-02 15:04:05");
	c.Do("rpush", strings.Split(ttstr2, " ")[0], d2byte);


	//写入开服3日数据
	c.Do("select", 0);
	for l:=0; l<2; l++{
		id ++;
		uinfo3 := datastruct.UserBaseinfo{};
		uinfo3.Id = id;
		uinfo3.OpenId = "robot_" + strconv.Itoa(uinfo3.Id);
		uinfo3.Name = uinfo3.OpenId;
		uinfo3.Maxpoint = 0;
		uinfo3.NewPlayer = true;
		t1 := time.Now();
		t2 := t1.AddDate(0, 0, 2);
		uinfo3.Createtime = t2;
		uinfo3.Logintime = t2;

		data2, _ := json.Marshal(uinfo3);
		_, err2 := c.Do("rpush", uinfo3.OpenId, data2);
		if err2!=nil{
			fmt.Println("[error] : insert fail!");
		}
	}
	val, _ := c.Do("lindex", "robot_1", 0);
	ino := datastruct.UserBaseinfo{};
	json.Unmarshal(val.([]byte), &ino);
	ino.NewPlayer = false;
	ino.Logintime = time.Now().AddDate(0,0,2);
	inodata, _ := json.Marshal(ino);
	c.Do("lset", "robot_1", 0, inodata);

	val2, _ := c.Do("lindex", "robot_6", 0);
	ino2 := datastruct.UserBaseinfo{};
	json.Unmarshal(val2.([]byte), &ino2);
	ino2.NewPlayer = false;
	ino2.Logintime = time.Now().AddDate(0,0,2);
	inodata2, _ := json.Marshal(ino2);
	c.Do("lset", "robot_6", 0, inodata2);
}