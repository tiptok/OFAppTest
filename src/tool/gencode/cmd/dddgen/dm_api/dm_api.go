package dm_api

import (
	"bytes"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/cmd/dddgen/api"
	"github.com/tiptok/OFAppTest/src/tool/gencode/cmd/dddgen/dm"
	"github.com/tiptok/OFAppTest/src/tool/gencode/common"
	"github.com/tiptok/OFAppTest/src/tool/gencode/constant"
	"github.com/tiptok/OFAppTest/src/tool/gencode/model"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func DmApiRun(ctx *cli.Context) {
	var (
		o       dm.DMOptions = dm.DMOptions{}
		results              = make(chan *api.GenResult, 100)
	)
	o.ProjectPath = ctx.String("p")
	o.SaveTo = ctx.String("st")
	o.Lib = ctx.String("lib")
	o.Language = ctx.String("lang")
	o.ModulePath = common.GoModuleName(o.SaveTo)

	dms := dm.ReadDomainModels(filepath.Join(o.ProjectPath, "domain-model"))
	if len(dms) == 0 {
		log.Println("domain-model not found")
		return
	}
	controllers, err := dmsToApiControllers(dms)
	if err != nil {
		log.Println(err)
	}
	var genServer api.ServeGen = GoBeeDomainApiServeGen{}

	api.GenApiServer(genServer, o.SvrOptions, controllers, results)
}

func dmsToApiControllers(dms []dm.DomainModel) (controllers []api.Controller, err error) {
	for _, dm := range dms {
		if !dm.NeedRestful() {
			continue
		}
		newCtr := api.Controller{
			Controller: dm.Name,
			Paths:      make([]api.ApiPath, 0),
		}
		newCtr.Paths = append(newCtr.Paths, []api.ApiPath{
			getApi(dm, "Create", http.MethodPost),
			getApi(dm, "Update", http.MethodPut),
			getApi(dm, "Get", http.MethodGet),
			getApi(dm, "Delete", http.MethodDelete),
			getApi(dm, "List", http.MethodGet),
		}...)
		controllers = append(controllers, newCtr)
	}
	return
}

func getApi(dm dm.DomainModel, prefix, method string) api.ApiPath {
	function := prefix + dm.Name
	lowercase := common.LowCasePaddingUnderline(dm.Name)
	lowerFirst := common.LowFirstCase(dm.Name)
	apiPath := api.ApiPath{
		Path:    fmt.Sprintf("/%v/:%vId", lowercase, lowerFirst),
		Method:  method,
		Content: "json",
	}
	apiPath.ServiceName = function
	apiPath.Operator = string(constant.COMMAND)
	switch strings.ToUpper(method) {
	case http.MethodPost:
		apiPath.Path = fmt.Sprintf("/%v/", lowercase)
	case http.MethodPut:
		apiPath.Path = fmt.Sprintf("/%v/:%vId", lowercase, lowerFirst)
	case http.MethodDelete:
		apiPath.Path = fmt.Sprintf("/%v/:%vId", lowercase, lowerFirst)
	case http.MethodGet:
		apiPath.Operator = string(constant.QUERY)
		apiPath.Path = fmt.Sprintf("/%v/:%vId", lowercase, lowerFirst)
	}
	if prefix == "List" {
		apiPath.Path = fmt.Sprintf("/%v/", lowercase)
	}
	apiPath.Summary = fmt.Sprintf("%v execute %v  %v  %v", function, apiPath.Operator, strings.ToLower(prefix), dm.Name)
	apiPath.Request = api.RefObject{RefPath: function + "Request"}
	apiPath.Response = api.RefObject{RefPath: function + "Response"}
	return apiPath
}

type GoBeeDomainApiServeGen struct {
	api.GoBeeApiServeGen
}

func (g GoBeeDomainApiServeGen) GenApplication(c api.Controller, options model.SvrOptions, result chan<- *api.GenResult) error {
	buf := bytes.NewBuffer(nil)
	bufMethods := bytes.NewBuffer(nil)
	for i := 0; i < len(c.Paths); i++ {
		bufMethods.WriteString("\n\n")
		p := c.Paths[i]
		pName, _, _ := p.ParsePath()
		//log.Println(pName,req,rsp)
		if err := common.ExecuteTmpl(bufMethods, api.ApplicationMethod, map[string]interface{}{
			"Method":  common.UpperFirstCase(pName),
			"Service": c.Controller,
			"Logic":   g.getApplicationLogic(c, p, options),
		}); err != nil {
			return err
		}
	}

	if err := common.ExecuteTmpl(buf, api.Application, map[string]interface{}{
		"Package": common.LowCasePaddingUnderline(c.Controller),
		"Module":  options.ModulePath,
		"Service": c.Controller,
		"Methods": bufMethods.String(),
	}); err != nil {
		return err
	}

	result <- &api.GenResult{
		Root:     options.SaveTo,
		SaveTo:   constant.WithApplication(common.LowCasePaddingUnderline(c.Controller)),
		FileName: common.LowCasePaddingUnderline(c.Controller) + ".go",
		FileData: buf.Bytes(),
	}
	return nil
}

func (g GoBeeDomainApiServeGen) getApplicationLogic(c api.Controller, path api.ApiPath, options model.SvrOptions) string {
	buf := bytes.NewBuffer(nil)
	domainPath := filepath.Join(options.ProjectPath, "domain-model", common.LowCasePaddingUnderline(c.Controller)+".json")
	models := dm.ReadDomainModels(domainPath)
	if len(models) == 0 {
		return ""
	}
	domainModel := models[0]
	if strings.HasPrefix(path.ServiceName, "Create") {
		buf.WriteString(fmt.Sprintf("	new%v:=&domain.%v{\n", c.Controller, c.Controller))
		for _, field := range domainModel.Fields {
			if containAnyInArray(field.Name, "Id") {
				continue
			}
			if containAnyInArray(field.Name, "At", "Time") {
				buf.WriteString(fmt.Sprintf("		%v: time.Now(),\n", field.Name))
				continue
			}
			buf.WriteString(fmt.Sprintf("		%v: request.%v,\n", field.Name, field.Name))
		}
		buf.WriteString("	}\n")
		buf.WriteString(fmt.Sprintf(`	
    var %vRepository,_ = factory.Create%vRepository(transactionContext)
	if m,err:=%vRepository.Save(new%v);err!=nil{
		return nil,err
	}else{
		rsp = m
	}`, c.Controller, c.Controller, c.Controller, c.Controller))
		return buf.String()
	}
	if strings.HasPrefix(path.ServiceName, "Update") {
		buf.WriteString(fmt.Sprintf("new%v:=&domain.%v{", c.Controller, c.Controller))
		for _, field := range domainModel.Fields {
			if containAnyInArray(field.Name, "Id") {
				continue
			}
			if containAnyInArray(field.Name, "At", "Time") {
				buf.WriteString(fmt.Sprintf("%v: time.Now()", field.Name))
				continue
			}
			buf.WriteString(fmt.Sprintf("%v: request.%v", field.Name, field.Name))
		}
		buf.WriteString("}")
		return buf.String()
	}
	return buf.String()
}

func containAnyInArray(c string, array ...string) bool {
	for i := range array {
		if strings.Contains(c, array[i]) {
			return true
		}
	}
	return false
}
