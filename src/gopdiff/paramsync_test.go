package gopdiff

import (
	"log"
	"testing"

	"io/ioutil"
	"strings"

	"os"

	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/xml"
)

/*
	同步配置
	1.param.xml
	2.paramN.xml
	3.节点改名 data -> config
*/
func TestParamDiff(t *testing.T) {
	//更新配置项到临时文件 paramt.xml paramNt.xml
	dataO, _ := ioutil.ReadFile("param.xml")
	sdataO := strings.Replace(string(dataO), "<Data>", "<config>", 1)
	sdataO = strings.Replace(sdataO, "</Data>", "</config>", 1)
	ioutil.WriteFile("paramt.xml", []byte(sdataO), 666)
	defer os.Remove("paramt.xml")

	dataN, _ := ioutil.ReadFile("paramN.xml")
	sdataN := strings.Replace(string(dataN), "<Data>", "<config>", 1)
	sdataN = strings.Replace(sdataN, "</Data>", "</config>", 1)
	ioutil.WriteFile("paramNt.xml", []byte(sdataN), 666)
	defer os.Remove("paramNt.xml")

	//读取配置
	pOldConfig, err := config.NewConfig("xml", "paramt.xml")
	if err != nil {
		log.Println("Read param.xml Error.", err)
		return
	}
	pOldSection, err := pOldConfig.GetSection("param")
	if err != nil {
		log.Println("GetSection param.xml Error.", err)
		return
	}
	pNewConfig, err := config.NewConfig("xml", "paramNt.xml")
	if err != nil {
		log.Println("Read paramN.xml Error.", err)
	}
	pNewSection, err := pNewConfig.GetSection("param")

	data, _ := ioutil.ReadFile("paramN.xml")
	sData := string(data)
	if err != nil {
		log.Println("Read ParamN.xml Error:", err)
		return
	}
	for okey, ov := range pOldSection {
		if nv, isexists := pNewSection[okey]; isexists {
			if ov != nv {
				log.Printf("同步配置 %v->%v  <%s>%v</%s>", nv, ov, okey, ov, okey)
				//pNewSection[key] = v
				oldFormat := "<" + okey + ">" + ov + "</" + okey + ">"
				newFormat := "<" + okey + ">" + nv + "</" + okey + ">"
				sData = strings.Replace(sData, newFormat, oldFormat, 1)
			}
		}
	}
	log.Println("**********************更新配置参数****************")
	log.Println(sData)
	sData = strings.Replace(sData, "<config>", "<Data>", 1)
	sData = strings.Replace(sData, "</config>", "</Data>", 1)
	ioutil.WriteFile("paramN.xml", []byte(sData), 666)
}
