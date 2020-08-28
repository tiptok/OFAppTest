package api

const protocol = `{
	"name":"{{.Protocol}}",
	"value_type":"dto",
	"fields":[
		{
			"name":"Id",
			"type":"int64",
            "desc":"唯一标识",
            "required":false
		}
	]
}
`

// ControllerLower auth
// Module   github.com/tiptok/gocommon
// Controller Auth
const beegonController = `package controllers

import (
	"{{.Module}}/pkg/application/{{.ControllerLower}}"
	"github.com/tiptok/gocomm/pkg/log"
	"{{.Module}}/pkg/protocol"
	protocolx "{{.Module}}/pkg/protocol/{{.ControllerLower}}"
)

type {{.Controller}}Controller struct {
	BaseController
}
`

// ControllerName Auth
// Description    Desc
// Method         Login
// RequestModel   LoginRequest
// ResponseModel  ResponseModel
// Application    auth
const beegoControllerMethod = `// {{.Method}} 
// {{.Desc}}
func (this *{{.ControllerName}}Controller) {{.Method}}() {
	var msg *protocol.ResponseMessage
	defer func() {
		this.Resp(msg)
	}()
	var request *protocolx.{{.RequestModel}}
	if err := this.JsonUnmarshal(&request); err != nil {
		msg = protocol.BadRequestParam(1)
		return
	}
	if b, m := this.Valid(request); !b {
		msg = m
		return
	}
	header := this.GetRequestHeader(this.Ctx)
	data, err := {{.Application}}.{{.Method}}(header, request)
	if err != nil {
		log.Error(err)
	}
	msg = protocol.NewReturnResponse(data, err)
}
`

// Module mod名
// Routers
const beegoRouters = `package routers

import (
	"github.com/astaxie/beego"
	"{{.Module}}/pkg/port/beego/controllers"
)

func init() {
{{.Routers}}
}
`

const beegoRouter = `	beego.Router("{{.Url}}", &controllers.{{.Controller}}{}, "{{.HttpMethod}}:{{.Method}}")`

const beegoRouterInit = `package beego

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "{{.Module}}/pkg/port/beego/routers"
)

func init(){
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}`

const application = `package {{.Package}}

import (
	"github.com/tiptok/gocomm/pkg/log"
	"{{.Module}}/pkg/protocol"
	"{{.Module}}/pkg/application/factory"
	protocolx "{{.Module}}/pkg/protocol/{{.Package}}"
)

{{.Methods}}
`

//Method Login
//
const applicationMethod = `func {{.Method}}(header *protocol.RequestHeader, request *protocolx.{{.Method}}Request) (rsp *protocolx.{{.Method}}Response, err error) {
	var (
		transactionContext, _          = factory.CreateTransactionContext(nil)
	)
	rsp = &protocolx.{{.Method}}Response{}
	if err = transactionContext.StartTransaction(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		transactionContext.RollbackTransaction()
	}()
	
	err = transactionContext.CommitTransaction()
	return
}`

const protocolModel = `package {{.Package}}

type {{.Model}} struct {
{{.Fields}}
}
`

const protocolField = `	// {{.Desc}}
	{{.Column}} {{.Type}} {{.Tags}}
`

const beegoBaseController = `package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"{{.Module}}/pkg/protocol"
	"strconv"
)

type BaseController struct {
	beego.Controller
}

func (controller BaseController) JsonUnmarshal(v interface{}) error {
	body := controller.Ctx.Input.RequestBody
	if len(body)==0{
		body = []byte("{}")
	}
	return json.Unmarshal(body, v)
}

func (controller BaseController) GetLimitInfo() (offset int, limit int) {
	offset, _ = controller.GetInt("offset")
	limit, _ = controller.GetInt("limit")
	return
}

//Valid  valid struct
func (controller *BaseController) Valid(obj interface{}) (result bool, msg *protocol.ResponseMessage) {
	/*校验*/
	var err error
	valid := validation.Validation{}
	result, err = valid.Valid(obj)
	if err != nil {
	}
	if !result {
		msg = protocol.BadRequestParam(2)
		return
	}

	return
}

func (this *BaseController)  Resp(msg *protocol.ResponseMessage) {
	this.Data["json"] = msg
	this.Ctx.Input.SetData("outputData", msg)
	this.ServeJSON()
}

func (this *BaseController) RespH5(msg *protocol.ResponseMessage) {
	if msg.Errno != 0 {
		msg.Errno = -1
	}
	this.Data["json"] = msg
	this.Ctx.Input.SetData("outputData", msg)
	this.ServeJSON()
}

//获取请求头信息
func (this *BaseController) GetRequestHeader(ctx *context.Context) *protocol.RequestHeader {
	h := &protocol.RequestHeader{}
	h.UserId, _ = strconv.ParseInt(ctx.Input.Header("x-mmm-id"), 10, 64)
	return h
}

`
const pgFactoryTransaction = `package factory

import (
	"{{.Module}}/pkg/infrastructure/pg"
	"{{.Module}}/pkg/infrastructure/pg/transaction"
)

func CreateTransactionContext(options map[string]interface{}) (*transaction.TransactionContext, error) {
	return &transaction.TransactionContext{
		PgDd: pg.DB,
	}, nil
}
`

const beegoMain = `package main

import (
	"github.com/astaxie/beego"
	_ "{{.Module}}/pkg/constant"
	_ "{{.Module}}/pkg/infrastructure/pg"
	 _ "{{.Module}}/pkg/port/beego"
)

func main() {
	defer func() {

	}()
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}`
