package dddgen

import (
	"github.com/tiptok/OFAppTest/src/tool/gencode/cmd/dddgen/api"
	"github.com/urfave/cli"
)

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "dm",
			Usage: "Create domain-model; example: gencode dm -p Path -st SaveTo -dp DataPersistence -lang Language(go)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Usage: "domain-model dsl file path(描述语言路径)   .",
					Value: "F://go//src//learn_project//ddd-project//stock",
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
			Action: dmrun,
		},
		{
			Name:  "api-dsl",
			Usage: "Create api dsl file; example: api-dsl -c Controller -url /auth/login -m post",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Usage: "api dsl file path(描述语言路径)",
					Value: "F://go//src//learn_project//ddd-project//stock",
				},
				cli.StringFlag{
					Name:  "c",
					Usage: "controller 控制器名称",
					Value: "Auth",
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "请求路径",
					Value: "/auth/login",
				},
				cli.StringFlag{
					Name:  "m",
					Usage: "m method 请求方法类型",
					Value: "post",
				},
			},
			Action: api.RunApiDSL,
		},
		{
			Name:  "api-svr",
			Usage: "Create api-svr server; example: gencode api -p Path -st SaveTo -lang Language(go) -lib beego",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "p",
					Usage: "api dsl file path(描述语言路径)   .",
					Value: "F://go//src//learn_project//ddd-project//stock",
				},
				cli.StringFlag{
					Name:  "st",
					Usage: "gen code save to file path(生成文件路径)",
					Value: "./gen",
				},
				cli.StringFlag{
					Name:  "lang",
					Usage: "gen target language code(生成指定语言代码)",
					Value: "go",
				},
				cli.StringFlag{
					Name:  "lib",
					Usage: "使用的web框架(go:beego gin)",
					Value: "beego",
				},
			},
			Action: api.RunApiSever,
		},
	}
}
