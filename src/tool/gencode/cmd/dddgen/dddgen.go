package dddgen

import (
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
			Action: dmrun,
		},
	}
}
