package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)
func main(){
	app:= cli.NewApp()
	app.Name="g-pwd"
	app.Version="1.0.0"
	app.Usage="secret pwd string."
	app.Commands = append(app.Commands,cli.Command{
		Name:"new",
		Usage:"new crpty pwd string(exp:new -from 123456)",
		Flags:[]cli.Flag{
			cli.StringFlag{
				Name:"t",
				Usage:"new crpty type",
				Value:"bcrypt",
			},
			cli.StringFlag{
				Name:"from",
				Usage:"pwd to crpty",
				Value:"bcrypt",
			},
		},
		Action:run,
	})
	err:= app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}

func run(ctx *cli.Context){
	scretType := ctx.String("t")
	pwdFrom :=ctx.String("from")

	switch scretType {
	case "bcrypt":
		bytes, err := bcrypt.GenerateFromPassword([]byte(pwdFrom), 10)
		if err!=nil{
			log.Println(err)
		}
		log.Println(fmt.Sprintf("gen pwd srect string->\n%s",string(bytes)))
		break
	default:
		break

	}
}
