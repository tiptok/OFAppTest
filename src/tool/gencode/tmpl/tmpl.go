package tmpl

//param  -c Auth -m Login
//Controller Auth
//ControllerLowcase auth
//Method Login
//MethodLowcase login
var ControllerMethod =`
//{{.Method}}
func(this *{{.Controller}}Controller){{.Method}}(){
	var msg *mybeego.Message
	defer func(){
		this.Resp(msg)
	}()
	var request *protocol.{{.Method}}Request
	if err:=json.Unmarshal(this.ByteBody,&request);err!=nil{
		log.Error(err)
		msg = mybeego.NewMessage(1)
		return
	}
	if b,m :=this.Valid(request);!b{
		msg = m
		return
	}
	msg = this.GenMessage({{.ControllerLowcase}}.{{.Method}}(request))
}
`

var ProtocolModel =`
/*{{.Method}} */
type {{.Method}}Request struct {
	Xxx string` +"`json:\"xxx\" valid:\"Required\"`" +`
}
type {{.Method}}Response struct {
}
`

//Method Login
var Handler =`
	func {{.Method}}(request *protocol.{{.Method}}Request)(rsp *protocol.{{.Method}}Response,err error){
	var (

	)
	rsp =&protocol.{{.Method}}Response{}
	return
}
`

var Router =`
/*{{.MethodLowcase}} controller*/
{
	{{.ControllerLowcase}} :=&v1.{{.Controller}}Controller{}
	nsV1.Router("/{{.ControllerLowcase}}/{{.MethodLowcase}}",{{.ControllerLowcase}},"post:{{.Method}}")
}
`

//Name Phone
//NameLowcase phone
//TypeName string
//ValidString Required;Mobile
var Param =`
{{.Name}} {{.TypeName}} `+"`json:\"{{.NameLowcase}}\" valid:\"{{.ValidString}}\"`"

