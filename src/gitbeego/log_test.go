package gitbeego

import (
	"testing"

	"log"

	"github.com/astaxie/beego/logs"
)

func TestNewLog(t *testing.T) {

	/*******************通用方式********************/
	logger := logs.GetLogger("AppTest")
	logger.Println("20180307CCC")
	//2018/03/07 17:49:49.650 [APPTEST] 20180307CCC

	logs.SetLogger(logs.AdapterFile, `{"filename":"App.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//设置是否显示调用函数
	logs.EnableFuncCallDepth(true)
	//函数名称深度
	logs.SetLogFuncCallDepth(3)
	logs.Debug("20180307CCC")
	logs.Info("20180307CCC")
	logs.Warn("20180307CCC")
	logs.Error("20180307CCC")
	logs.Critical("20180307CCC")
	t.Log("end")

	log.Println("20180307CCC")

	/*******************单独声明使用********************/
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Info("Instance.")
}
