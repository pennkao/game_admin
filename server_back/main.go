//手指画线数据后台服务
//@author 黄承武

package main

import(
	"runtime"
	"server_back/web"
	// "server_back/tool"
)


//主函数
func main(){

	//启动Go语言最佳性能，创建的线程数量使用CPU的数量.这样的话Go创建的协程
	//程序的线程不会超过CPU，在线程切换时CPU不会有太多的消耗。
	runtime.GOMAXPROCS(runtime.NumCPU());
	
	//写入测试数据
	// tool.WriteTestDatas();

	//启动http服务
	web.StartHttpserver();
}



