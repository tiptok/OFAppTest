package api

import (
	"bytes"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/common"
	"github.com/tiptok/OFAppTest/src/tool/gencode/constant"
	"github.com/tiptok/OFAppTest/src/tool/gencode/model"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// 通过api dsl描述语言 生成对应的api服务
func RunApiSever(ctx *cli.Context) {
	var (
		o        = apiSvrOptions{}
		results  = make(chan *GenResult, 10)
		serveGen = serveGenFactory()
	)
	o.ProjectPath = ctx.String("p") //项目文件根目录
	o.SaveTo = ctx.String("st")
	o.Language = ctx.String("lang")
	o.Lib = ctx.String("lib")

	if _, ok := o.Valid(); !ok {
		return
	}
	controllers, err := ReadApiModels(o.ProjectPath)
	if err != nil {
		fmt.Println("read api models err:", err)
		return
	}
	for i := 0; i < len(controllers); i++ {
		c := controllers[i]
		if err := serveGen.GenController(c, o, results); err != nil {
			fmt.Println("gen controller error:", err)
			return
		}
		if err := serveGen.GenRouter(c, o, results); err != nil {
			fmt.Println("gen router error:", err)
			return
		}
	}
	close(results)
	var done sync.WaitGroup
	done.Add(1)
	go func() {
		for result := range results {
			err := common.SaveTo(result.Root, result.SaveTo, result.FileName, result.FileData)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		done.Done()
	}()
	done.Wait()
}

func ReadApiModels(p string) (controllers []Controller, err error) {
	var f os.FileInfo
	if f, err = os.Stat(p); err != nil {
		return
	}
	readOne := func(file string) error {
		c := Controller{}
		if err := common.ReadModelFromJsonFile(file, &c); err != nil {
			return err
		}
		controllers = append(controllers, c)
		return nil
	}
	if !f.IsDir() {
		err = readOne(p)
		return
	}
	root := filepath.Join(p, constant.Api)
	files, e := ioutil.ReadDir(root)
	if e != nil {
		err = e
		return
	}
	for i := range files {
		fileItem := files[i]
		if fileItem.IsDir() {
			continue
		}
		if err = readOne(filepath.Join(root, fileItem.Name())); err != nil {
			return
		}
	}
	return
}

// serve生成器
func serveGenFactory() ServeGen {
	return GoBeeApiServeGen{}
}

type ServeGen interface {
	GenController(c Controller, options apiSvrOptions, result chan<- *GenResult) error
	GenRouter(c Controller, options apiSvrOptions, result chan<- *GenResult) error
	GenApplication(o Operation, options apiSvrOptions, result chan<- *GenResult) error
	GenProtocol(o Operation, options apiSvrOptions, result chan<- *GenResult) error
}

// golang beego 框架 serve生成器
type GoBeeApiServeGen struct{}

func (g GoBeeApiServeGen) GenController(c Controller, options apiSvrOptions, result chan<- *GenResult) error {
	//fmt.Println("gen controller:",c.Controller)
	buf := bytes.NewBuffer(nil)
	if err := common.ExecuteTmpl(buf, beegonController, map[string]interface{}{
		"Module":          "",
		"ControllerLower": common.LowFirstCase(c.Controller),
		"Controller":      c.Controller,
	}); err != nil {
		return err
	}

	for i := 0; i < len(c.Paths); i++ {
		buf.WriteString("\n")
		p := c.Paths[i]
		pName, req, rsp := p.ParsePath()
		//log.Println(pName,req,rsp)
		if err := common.ExecuteTmpl(buf, beegoControllerMethod, map[string]interface{}{
			"Method":         common.UpperFirstCase(pName),
			"Desc":           p.Summary,
			"ControllerName": c.Controller,
			"RequestModel":   req,
			"ResponseModel":  rsp,
			"Application":    pName,
		}); err != nil {
			return err
		}
	}

	result <- &GenResult{
		Root:     options.SaveTo,
		SaveTo:   constant.WithController(options.Lib),
		FileName: common.LowCasePaddingUnderline(c.Controller) + ".go",
		FileData: buf.Bytes(),
	}
	return nil
}
func (g GoBeeApiServeGen) GenRouter(c Controller, options apiSvrOptions, result chan<- *GenResult) error {
	return nil
}
func (g GoBeeApiServeGen) GenApplication(o Operation, options apiSvrOptions, result chan<- *GenResult) error {
	return nil
}
func (g GoBeeApiServeGen) GenProtocol(o Operation, options apiSvrOptions, result chan<- *GenResult) error {
	return nil
}

type GenResult struct {
	Root     string
	SaveTo   string
	FileName string
	FileData []byte
}
type Operation struct {
	Url      path
	Request  model.CustomerModel
	Response model.CustomerModel
}

// 服务参数
type apiSvrOptions struct {
	// 项目路径
	ProjectPath string
	// 保存路径
	SaveTo string
	// 服务语言
	Language string
	// 框架
	Lib string
}

func (o apiSvrOptions) Valid() (error, bool) {
	return nil, true
}
