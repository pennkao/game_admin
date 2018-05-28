//数据结构
//@author 黄承武

package datastruct

import(
	"time"
)



//用户基础信息数据结构（redis数据库中用户基础信息储存结构）
type UserBaseinfo struct
{
	OpenId string //平台账号
	Id int //唯一索引
	Name string //用户昵称
	Maxpoint int //最大关卡数
	Createtime time.Time //用户创建时间
	Logintime time.Time //最近登陆时间
	NewPlayer bool //是否是新用户（true：是；false：不是）
}



//消息错误数据结构
type MsgError struct
{
	Code int //错误码
	Errmsg string //错误消息
}



//发送给请求者的消息结构
type Msg struct
{
	Code int //消息码
	Data interface{} //数据
}



//日登陆数据结构体
type DauInfo struct
{
	Info []interface{} //日登陆用户数
}