//数据库redis相关操作
//注 ： redis 10库 ： 负责存储日登陆用户的总数（DAU）

//@author 黄承武

package sql

import(
	"github.com/garyburd/redigo/redis"
	"server_back/config"
	"time"
	"fmt"
	"encoding/json"
	"server_back/datastruct"
	"strings"
)


//sql数据链接池
func SqlOperate() *redis.Pool{

	return &redis.Pool{
		MaxActive : config.SQL_MAXACTIVE,
		MaxIdle : config.SQL_MAXIDLE,
		IdleTimeout : time.Duration(config.SQL_TIMEOUT) * time.Second,

		//链接
		Dial : func()(redis.Conn, error){
			c, err := redis.Dial("tcp", config.SQL_URL_PORT);
			if err!=nil {
				return nil, err;
			}
			if _, err := c.Do("AUTH", config.SQL_PASSWORD); err != nil{
				c.Close();
				return nil, err;
			}
			return c, err;
		},

		//ping测试（在网页上）
		TestOnBorrow : func(c redis.Conn, t time.Time) error{
			if time.Since(t) < time.Minute {
				return nil;
			}
			_, err := c.Do("PING");
			return err;
		},
	};
}



//----------------------------------------------------
//--后台数据操作



//查询注册用户数量
//@return 返回注册人数 (-1:说明查询失败)
func GetRegs() int64{
	c := SqlOperate().Get();
	defer c.Close();

	//选择库
	_, err := c.Do("select", 0);
	if err!=nil{
		return -1;
	}

	val, err := c.Do("dbsize");
	if err!=nil {
		return -1;
	}

	return val.(int64);
}



//获取各个关卡的分布数据
func GetCardDistribution() ([config.CARD_NUMBER]int){
	c := SqlOperate().Get();
	defer c.Close();

	//创建一个切片
	ary := [config.CARD_NUMBER]int{};


	//选择库
	_, err := c.Do("select", 0);
	if err!=nil {
		return ary;
	}

	infos, err := redis.Values( c.Do("keys", "*") );
	if err != nil {
		return ary;
	}

	//累计查询
	for _, v := range infos{
		
		val,  _ := c.Do("lindex", string(v.([]byte)), 0);
		
		baseinfo := datastruct.UserBaseinfo{};
		json.Unmarshal(val.([]byte), &baseinfo);

		ary[baseinfo.Maxpoint] = ary[baseinfo.Maxpoint] +1;
		
	}

	return ary;
}




//获取日登陆数据（当前日期）
//@return 返回日登陆数据
func GetDAUCrtDate(val string)([]interface{}){
	c := SqlOperate().Get();
	defer c.Close();

	_, err := c.Do("select", 0);
	if err!=nil {
		return make([]interface{}, 0);
	}

	infos, err2 := redis.Values( c.Do("keys", "*") );
	if err2 != nil {
		return make([]interface{}, 0);
	}


	//累计查询
	var oldnum int=0; //老用户数
	var olds []string; //老用户OPENID
	var newnum int=0; //新用户数
	var news []string; //新用户OPENID
	var num int = 0 //总数
	var ts []string;
	for _, v := range infos{
		
		dt, _ := c.Do("lindex", string(v.([]byte)), 0);
		
		baseinfo := datastruct.UserBaseinfo{};
		json.Unmarshal(dt.([]byte), &baseinfo);

		t := baseinfo.Logintime.Format("2006-01-02 15:04:05");
		
		ts = strings.Split(t, " ");
		if(ts[0] == val){//取出当前日登陆的用户
			if baseinfo.NewPlayer {
				news = append(news, baseinfo.OpenId);
				newnum ++;
			}else{
				olds = append(olds, baseinfo.OpenId);
				oldnum ++;
			}
		}
	}
	num = oldnum + newnum;
 
	//数据
	numbrs := make([]interface{}, 0);
	numbrs = append(numbrs, oldnum);
	numbrs = append(numbrs, newnum);
	numbrs = append(numbrs, num);
	numbrs = append(numbrs, olds);
	numbrs = append(numbrs, news);

	return numbrs;
}



//设置日登陆数据（当前日期）
func SetDAUCrtDate(val string){
	ns := GetDAUCrtDate(val);

	c := SqlOperate().Get();
	defer c.Close();

	_, err := c.Do("select", 10);
	if err != nil { fmt.Println("[error]: set current date login(DAU) error. sql select 10 lib fail!"); return; };

	dau := datastruct.DauInfo{};
	dau.Info = ns;
	infodata, _ := json.Marshal(dau);
	_, err2 := c.Do("rpush", val, infodata);
	
	if err2 != nil { fmt.Println("[error]: set current date login(DAU) error. insert failed!"); };

	fmt.Println("record current date DAU success! Info = ", ns);
}



//获取日登陆数据（当前日之前的）
func GetDAUForeDate(val string) ([]interface{}){
	c := SqlOperate().Get();
	defer c.Close();

	_, err := c.Do("select", 10);
	if err != nil {return make([]interface{}, 0); };

	n, err2 := c.Do("lindex", val, 0);
	if err2 != nil{ return make([]interface{}, 0);};

	if n==nil { return make([]interface{}, 0);};

	dau := datastruct.DauInfo{};
	err3 := json.Unmarshal(n.([]byte), &dau);
	if err3 != nil { return make([]interface{}, 0); };

	return dau.Info;
}