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
	"github.com/tiptok/gocommon/pkg/log"
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

const beegoRouter = `
`
