package gitBeegoT

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/astaxie/beego/config"
	//如果你使用xml 或者 yaml 驱动就需要手工安装引入包
	_ "github.com/astaxie/beego/config/xml"

	"github.com/astaxie/beego/utils"
)

/*
读取xml配置信息

*/
func TestXmlConfig(t *testing.T) {
	var (
		//xml parse should incluce in <config></config> tags
		xmlcontext = `<?xml version="1.0" encoding="UTF-8"?>
<config>
<appname>beeapi</appname>
<httpport>8080</httpport>
<mysqlport>3600</mysqlport>
<PI>3.1415976</PI>
<runmode>dev</runmode>
<autorender>false</autorender>
<copyrequestbody>true</copyrequestbody>
<path1>${GOPATH}</path1>
<path2>${GOPATH||/home/go}</path2>
<mysection>
<id>1</id>
<name>MySection</name>
</mysection>
</config>
`
	)
	f, err := os.Create("param.xml")
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	_, err = f.WriteString(xmlcontext)
	f.Close()
	defer os.Remove("param.xml")

	xmlconfig, err := config.NewConfig("xml", "param.xml")
	if err != nil {
		t.Fatal(err)
	}
	httpport, err := xmlconfig.Int64("httpport")
	//Configer 读取参数
	log.Println(xmlconfig)
	log.Println("httpport:", httpport)

	//读取系统参数
	env := utils.NewBeeMap()
	for _, e := range os.Environ() {
		splits := strings.Split(e, "=")
		env.Set(splits[0], os.Getenv(splits[0]))
		log.Println(splits[0], ":", os.Getenv(splits[0]))
	}
	//log.Println(env.Items())
}

/*
	读取json配置
*/
func TestJsonConfig(t *testing.T) {
	var (
		jsonstr = `{
		"Name":"tiptok",
		"Adult":true,
		"age":22,
		"Course":"C#;Go"
	}`
	)

	cf, _ := os.Create("param.json")
	cf.WriteString(jsonstr)

	cf.Close()

	jsonConfig, _ := config.NewConfig("json", "param.json")
	defer os.Remove("param.json")
	Name := jsonConfig.String("Name")
	IsAdult, _ := jsonConfig.Bool("Adult")
	Age, _ := jsonConfig.Int("age")
	//数组
	course := jsonConfig.Strings("Course")
	log.Println(Name, IsAdult, Age, course)
	t.Fatal("xx")
}

/*
	读取ini配置
*/
func TestIniConfig(t *testing.T) {
	inicontext := `
	;comment one
	#comment two
	Name = tiptok
	Adult = true
	age = 22
	Course = C#;Go
	`
	cf, _ := os.Create("param.conf")
	cf.WriteString(inicontext)

	cf.Close()

	jsonConfig, _ := config.NewConfig("ini", "param.conf")
	defer os.Remove("param.conf")
	Name := jsonConfig.String("Name")
	IsAdult, _ := jsonConfig.Bool("Adult")
	Age, _ := jsonConfig.Int("age")
	//数组
	course := jsonConfig.Strings("Course")
	log.Println(Name, IsAdult, Age, course)
	t.Fatal("xx")
}
