/**
   服务器配置文件
   配置http服务的地址、端口号，数据库相关信息以及缓存服务的相关信息

   @author 黄承武

   @create time 2018.4.23

  */
package config

//http服务地址
//const HTTP_URL string = "192.168.2.122:1201";
 const HTTP_URL string = "172.21.0.47:1201";

//数据库地址、端口号
//const SQL_URL_PORT string = "192.168.2.254:5000";
 const SQL_URL_PORT string = "172.21.0.14:6379";

//数据库登陆密码
const SQL_PASSWORD string = "@doinggame2018!";

//最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
const SQL_MAXACTIVE int = 20;

//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
const SQL_MAXIDLE int = 10;

//数据库连接超时时间（单位：秒）
const SQL_TIMEOUT int = 300;


//关卡总数（配置总数 +1. 这个第1位是从0开始）
const CARD_NUMBER int = 506 +1;
