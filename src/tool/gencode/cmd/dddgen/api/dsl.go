package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/common"
	"github.com/tiptok/OFAppTest/src/tool/gencode/constant"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

// 描述语言生成  go run gencode.go api-dsl -c User --url /user/userInfo
// 生成api dsl描述语言
func RunApiDSL(ctx *cli.Context) {
	var (
		o          = dslOptions{}
		controller = Controller{}
	)
	o.ProjectPath = ctx.String("p") //项目文件根目录
	o.Controller = ctx.String("c")
	o.Path = ctx.String("url")
	o.Method = ctx.String("m")
	if strings.Index(o.Path, "/") < 0 {
		o.Path = "/" + o.Path
	}
	if err := o.Valid(); err != nil {
		fmt.Println(err)
		return
	}

	filename := fmt.Sprintf("api_%v.json", common.LowCasePaddingUnderline(o.Controller))
	f := filepath.Join(o.ProjectPath, "api", filename)
	if common.FileExists(f) {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = json.Unmarshal(data, &controller); err != nil {
			fmt.Println(err)
			return
		}
	}
	if len(controller.Paths) == 0 {
		controller.Controller = o.Controller
	}
	splits := strings.Split(o.Path, "/")
	protcolName := splits[len(splits)-1]
	RequestProtocol := common.UpperFirstCase(protcolName) + "Request"
	ResponseProtocol := common.UpperFirstCase(protcolName) + "Response"
	path := path{
		Path:     o.Path,
		Method:   o.Method,
		Summary:  "测试",
		Content:  "json",
		Request:  refObject{RefPath: strings.Join([]string{constant.ApiProtocol, common.LowFirstCase(controller.Controller), RequestProtocol}, "/")},
		Response: refObject{RefPath: strings.Join([]string{constant.ApiProtocol, common.LowFirstCase(controller.Controller), ResponseProtocol}, "/")},
	}
	fmt.Println(filepath.Join(constant.ApiProtocol, controller.Controller, RequestProtocol), constant.ApiProtocol, controller.Controller, RequestProtocol)
	fmt.Println(path.Request.RefPath)
	for i := range controller.Paths {
		p := controller.Paths[i]
		if p.Path == o.Path {
			//路由已存在
			controller.Paths[i] = path
			fmt.Println("exists:", o.ProjectPath, o.Controller, o.Method)
			return
		}
	}
	controller.Paths = append(controller.Paths, path)
	//save file
	if err := common.SaveJsonTo(o.ProjectPath, "api", filename, controller, true); err != nil {
		fmt.Println(err)
		return
	}
	protocolPath := filepath.Join(constant.ApiProtocol, common.LowCasePaddingUnderline(controller.Controller))
	if err := saveProtocol(o.ProjectPath, protocolPath, RequestProtocol, false); err != nil {
		fmt.Println(err)
		return
	}
	if err := saveProtocol(o.ProjectPath, protocolPath, ResponseProtocol, false); err != nil {
		fmt.Println(err)
		return
	}
}
func saveProtocol(root string, st string, filename string, recover bool) error {
	f := filepath.Join(root, st, filename)
	if common.FileExists(f) && !recover {
		return nil
	}
	tP, err := template.New("controller").Parse(protocol)
	if err != nil {
		log.Fatal(err)
	}
	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["Protocol"] = filename
	tP.Execute(bufTmpl, m)

	return common.SaveTo(root, st, fmt.Sprintf("%v.json", filename), bufTmpl.Bytes())
}

type dslOptions struct {
	ProjectPath string
	Controller  string
	Path        string
	Method      string
	Summary     string
}

func (o dslOptions) Valid() error {
	if len(o.Controller) == 0 {
		return fmt.Errorf("controller is require")
	}
	if len(o.Path) == 0 {
		return fmt.Errorf("path is require")
	}
	return nil
}

type Controller struct {
	Controller string `json:"controller"`
	Paths      []path `json:"paths"`
}
type path struct {
	Path     string    `json:"path"`
	Method   string    `json:"method"`
	Summary  string    `json:"summary"`
	Content  string    `json:"content"`
	Request  refObject `json:"request"`
	Response refObject `json:"response"`
}

func (p path) ParsePath() (string, string, string) {
	splits := strings.Split(p.Path, "/")
	protcolName := splits[len(splits)-1]
	RequestProtocol := common.UpperFirstCase(protcolName) + "Request"
	ResponseProtocol := common.UpperFirstCase(protcolName) + "Response"
	return protcolName, RequestProtocol, ResponseProtocol
}

type refObject struct {
	RefPath string `json:"ref"`
}
