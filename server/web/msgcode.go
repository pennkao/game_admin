//消息码
//@author 黄承武


package web


//处理请求 - 注册用户数量
const MSG_CODE_1001 int = 1001;

//处理请求 - 每个关卡的人员分布
const MSG_CODE_1002 int = 1002;

//处理请求 - 拉取每日用户登陆数据
const MSG_CODE_1003 int = 1003;


//检测消息码合法性
//@param c 消息码
//@return true：合法；false：不合法
func CheckMsgcode(c int) (bool){

	if c!=MSG_CODE_1001 && c!=MSG_CODE_1002 && c!=MSG_CODE_1003 {
		return false;
	}

	return true;
}


//-----------------------------------------
//--错误码


//消息操作码不合法
const ERR_MSG_CODE_1001 int = -1;