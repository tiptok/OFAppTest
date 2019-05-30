package new

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"github.com/tiptok/OFAppTest/src/tool/dmicro/tmpl"
	"time"
)

type config struct {
	// foo
	Alias string
	// micro new example -type
	Command string
	// go.micro
	Namespace string
	// api, srv, web, fnc
	Type string
	// go.micro.srv.foo
	FQDN string
	// github.com/micro/foo
	Dir string
	// $GOPATH/src/github.com/micro/foo
	GoDir string
	// $GOPATH
	GoPath string
	// Files
	Files []file
	// Comments
	Comments []string
	// Plugins registry=etcd:broker=nats
	Plugins []string
}

type file struct {
	Path string
	Tmpl string
}

func run(ctx *cli.Context){
	var (
		atype string=ctx.String("type")
		dir string =ctx.String("name")

		alias string=""
		command string =  fmt.Sprintf("micro new %s", dir)
		namespace string=ctx.String("namespace")
		fqdn string =""
		goDir string =""
		goPath string =""
		plugins []string

		useGoModule = os.Getenv("GO111MODULE")
	)

	curDir,err :=os.Getwd()
	if err!=nil{
		fmt.Println(err)
		return
	}else{
		fmt.Printf("-> project root:%s\n",curDir)
	}

	goDir = filepath.Join(curDir,path.Clean(dir))

	if len(dir) == 0 {
		fmt.Println("specify service name")
		return
	}
	if len(alias) == 0 {
		// set as last part
		alias = filepath.Base(dir)
	}
	if len(fqdn) == 0 {
		fqdn = strings.Join([]string{namespace, atype, alias}, ".")
	}


	var c config

	switch atype {
	case "fnc":
		// create srv config
		c = config{
			Alias:     alias,
			Command:   command,
			Namespace: namespace,
			Type:      atype,
			FQDN:      fqdn,
			Dir:       dir,
			GoDir:     goDir,
			GoPath:    goPath,
			Plugins:   plugins,
			Files: []file{
				{"main.go", tmpl.MainFNC},
				{"plugin.go", tmpl.Plugin},
				{"handler/example.go", tmpl.HandlerFNC},
				{"subscriber/example.go", tmpl.SubscriberFNC},
				{"proto/example/example.proto", tmpl.ProtoFNC},
				{"Dockerfile", tmpl.DockerFNC},
				{"Makefile", tmpl.Makefile},
				{"README.md", tmpl.ReadmeFNC},
			},
			Comments: []string{
				"\ndownload protobuf for micro:\n",
				"brew install protobuf",
				"go get -u github.com/golang/protobuf/{proto,protoc-gen-go}",
				"go get -u github.com/micro/protoc-gen-micro",
				"\ncompile the proto file example.proto:\n",
				"cd " + goDir,
				"protoc --proto_path=. --go_out=. --micro_out=. /proto/example/example.proto\n",
			},
		}
	case "srv":
		// create srv config
		c = config{
			Alias:     alias,
			Command:   command,
			Namespace: namespace,
			Type:      atype,
			FQDN:      fqdn,
			Dir:       dir,
			GoDir:     goDir,
			GoPath:    goPath,
			Plugins:   plugins,
			Files: []file{
				{"main.go", tmpl.MainSRV},
				//{"plugin.go", tmpl.Plugin},
				{"conf/app.yaml",tmpl.ConfAPP},
				{"conf/app-dev.yaml",tmpl.ConfAPP_DEV},
				{"conf/app-prod.yaml",tmpl.ConfAPP_PROD},
				{"handler/example.go", tmpl.HandlerSRV},
				//{"subscriber/example.go", tmpl.SubscriberSRV},
				{"model/proto/example/example.proto", tmpl.ProtoSRV},
				{"server/tmp",""},
				{"service/http/tmp",""},
				{"service/grpc",""},
				{"Dockerfile", tmpl.DockerSRV},
				{"Makefile", tmpl.Makefile},
				{"README.md", tmpl.Readme},
			},
			Comments: []string{
				"\ndownload protobuf for micro:\n",
				"brew install protobuf",
				"go get -u github.com/golang/protobuf/{proto,protoc-gen-go}",
				"go get -u github.com/micro/protoc-gen-micro",
				"\ncompile the proto file example.proto:\n",
				"cd " + goDir,
				"protoc --proto_path=. --go_out=. --micro_out=. model/proto/example/example.proto\n",
			},
		}
	case "api":
		// create api config
		c = config{
			Alias:     alias,
			Command:   command,
			Namespace: namespace,
			Type:      atype,
			FQDN:      fqdn,
			Dir:       dir,
			GoDir:     goDir,
			GoPath:    goPath,
			Plugins:   plugins,
			Files: []file{
				{"main.go", tmpl.MainAPI},
				{"conf/app.yaml",tmpl.ConfAPP},
				{"conf/app-dev.yaml",tmpl.ConfAPP_DEV},
				{"conf/app-prod.yaml",tmpl.ConfAPP_PROD},
				//{"plugin.go", tmpl.Plugin},
				{"handler/example.go", tmpl.HandlerAPI},
				{"model/proto/", ""},
				{"server/tmp",""},
				{"service/http/tmp",""},
				{"Makefile", tmpl.Makefile},
				{"Dockerfile", tmpl.DockerSRV},
				{"README.md", tmpl.Readme},
			},
			Comments: []string{
				"\ndownload protobuf for micro:\n",
				"brew install protobuf",
				"go get -u github.com/golang/protobuf/{proto,protoc-gen-go}",
				"go get -u github.com/micro/protoc-gen-micro",
				"\ncompile the proto file example.proto:\n",
				"cd " + goDir,
				"protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto\n",
			},
		}
	case "web":
		// create srv config
		c = config{
			Alias:     alias,
			Command:   command,
			Namespace: namespace,
			Type:      atype,
			FQDN:      fqdn,
			Dir:       dir,
			GoDir:     goDir,
			GoPath:    goPath,
			Plugins:   plugins,
			Files: []file{
				{"main.go", tmpl.MainWEB},
				{"plugin.go", tmpl.Plugin},
				{"handler/handler.go", tmpl.HandlerWEB},
				{"html/index.html", tmpl.HTMLWEB},
				{"Dockerfile", tmpl.DockerWEB},
				{"Makefile", tmpl.Makefile},
				{"README.md", tmpl.Readme},
			},
			Comments: []string{},
		}
	default:
		fmt.Println("Unknown type", atype)
		return
	}

	// set gomodule
	if useGoModule == "on" || useGoModule == "auto" {
		c.Files = append(c.Files, file{"go.mod", tmpl.Module})
	}else {
		c.Files = append(c.Files, file{"go.mod", tmpl.Module})
	}

	if err := create(c); err != nil {
		fmt.Println(err)
		return
	}
}


func create(c config) error {
	// check if dir exists
	if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.GoDir)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	fmt.Printf("-> Creating service %s in %s\n\n", c.FQDN, c.GoDir)


	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.GoDir, file.Path)
		dir := filepath.Dir(f)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		if len(file.Tmpl)==0{//部分只创建文件夹
			continue
		}
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}

	for _, comment := range c.Comments {
		fmt.Println(comment)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	return nil
}

func write(c config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": strings.Title,
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}


func Commands() []cli.Command{
	return []cli.Command{
		{
			Name:"new",
			Usage: "Create a service template; example: dmicro new -name demo -type api / dmicro new -name demo (srv)",
			Flags:[]cli.Flag{
				cli.StringFlag{
					Name:"name",
					Usage:"project name for the service",
					Value:"demo",
				},
				cli.StringFlag{
					Name:"namespace",
					Usage:"Namespace for the service e.g com.example",
					Value:"go.micro",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Type of service e.g api, fnc, srv, web",
					Value: "srv",
				},
			},
			Action:run,
		},
	}
}
