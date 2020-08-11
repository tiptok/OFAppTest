package cmd

import (
	"bytes"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/tmpl"
	"github.com/urfave/cli"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Cmd interface {
	App() *cli.App
	Init(opts ...Option) error
	// Options set within this command
	Options() Options
}
type Option func(o *Options)
type Options struct {
	// For the Command Line itself
	Name        string
	Description string
	Version     string
}

func Name(s string) Option {
	return func(o *Options) {
		o.Name = s
	}
}
func Description(s string) Option {
	return func(o *Options) {
		o.Description = s
	}
}
func Version(s string) Option {
	return func(o *Options) {
		o.Version = s
	}
}

type cmd struct {
	opts Options
	app  *cli.App
}

func (c *cmd) App() *cli.App {
	return cli.NewApp()
}
func (c *cmd) Init(opts ...Option) error {
	for _, o := range opts {
		o(&c.opts)
	}
	c.app.Name = c.opts.Name
	c.app.Version = c.opts.Version
	c.app.HideVersion = len(c.opts.Version) == 0
	c.app.Usage = c.opts.Description

	c.app.Commands = append(c.app.Commands, Commands()...)
	return nil
}
func (c *cmd) Options() Options {
	return c.opts
}

var DefaultCmd *cmd

func newCmd() *cmd {
	return &cmd{}
}
func Init(opts ...Option) {
	DefaultCmd = newCmd()
	DefaultCmd.app = cli.NewApp()
	DefaultCmd.Init(opts...)
	err := DefaultCmd.app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "new",
			Usage: "Create a service template; example: gencode new -c Auth -m Login",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "c",
					Usage: "controller name for the service",
					Value: "Auth",
				},
				cli.StringFlag{
					Name:  "m",
					Usage: "controller handler name",
					Value: "Login",
				},
			},
			Action: run,
		},
		{
			Name:  "param",
			Usage: "crate model: gencode param -n VersionNo -t int -v Require,Mobile",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "n",
					Usage: "name of a param",
					Value: "Xxx",
				},
				cli.StringFlag{
					Name:  "t",
					Usage: "type of a param",
					Value: "string",
				},
				cli.StringFlag{
					Name:  "v",
					Usage: "valid param: Require,Mobile,Email",
					Value: "Require",
				},
			},
			Action: runParam,
		},
		{
			Name:  "dm",
			Usage: "Create domain-model; example: gencode dm -p Path -st SaveTo -dp DataPersistence -lang Language(go)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Usage: "domain-model dsl file path(描述语言路径)   .",
					Value: "F://go//src//learn_project//ddd-project//partner//domain-model",
				},
				cli.StringFlag{
					Name:  "st",
					Usage: "gen code save to file path(生成文件路径)",
					Value: "./gen",
				},
				cli.StringFlag{
					Name:  "dp",
					Usage: "data persistence(数据持久化 pg mysql redis)",
					Value: "pg",
				},
				cli.StringFlag{
					Name:  "lang",
					Usage: "gen target language code(生成指定语言代码)",
					Value: "go",
				},
			},
			Action: rundm,
		},
	}
}

func runParam(ctx *cli.Context) {
	var (
		name     string = ctx.String("n")
		tType    string = ctx.String("t")
		validStr string = ctx.String("v")
	)
	tP, err := template.New("controller").Parse(tmpl.Param)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["Name"] = name
	m["NameLowcase"] = LowFirstCase(name)
	m["TypeName"] = tType
	m["ValidString"] = strings.Replace(validStr, ",", ";", -1)
	tP.Execute(buf, m)
	fmt.Println(buf.String())
}

func run(ctx *cli.Context) {
	var (
		controller string = ctx.String("c")
		method     string = ctx.String("m")
	)
	tC, err := template.New("controller").Parse(tmpl.ControllerMethod)
	if err != nil {
		log.Fatal(err)
	}
	//param  -c Auth -m Login
	//Controller Auth
	//ControllerLowcase auth
	//Method Login
	//MethodRequest LoginRequest
	m := make(map[string]string)
	m["Controller"] = controller
	m["ControllerLowcase"] = LowFirstCase(controller)
	m["Method"] = method
	m["MethodLowcase"] = LowFirstCase(method)
	buf := bytes.NewBuffer(nil)
	tC.Execute(buf, m)

	tP, err := template.New("protocol").Parse(tmpl.ProtocolModel)
	tP.Execute(buf, m)

	tH, err := template.New("protocol").Parse(tmpl.Handler)
	tH.Execute(buf, m)

	tR, err := template.New("route").Parse(tmpl.Router)
	tR.Execute(buf, m)
	//log.Println(buf.String())
	ioutil.WriteFile("gencode.out", buf.Bytes(), os.ModePerm)
}

//单词首字母小写
func LowFirstCase(input string) string {
	array := []byte(input)
	if len(array) == 0 {
		return ""
	}
	rspArray := make([]byte, len(array))
	copy(rspArray[:1], strings.ToLower(string(array[:1])))
	copy(rspArray[1:], array[1:])
	return string(rspArray)
}

//首字母小写 多个字母用下划线拼接
// LowCase  low_case
func LowCasePaddingUnderline(input string) string {
	data := []byte(input)
	var rspData []byte
	for i := range data {
		c := data[i]
		if c >= byte('A') && c <= byte('Z') {
			if i != 0 { //首字母除外
				rspData = append(rspData, byte('_'))
			}
			rspData = append(rspData, c+32)
			continue
		}
		rspData = append(rspData, c)
	}
	return string(rspData)
}
