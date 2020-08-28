package dm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/common"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func DmRun(ctx *cli.Context) {
	var (
		path string    = ctx.String("p")
		o    DMOptions = DMOptions{}
	)
	o.ProjectPath = path
	o.SaveTo = ctx.String("st")
	o.DataPersistence = ctx.String("dp")
	o.Language = ctx.String("lang")
	o.ModulePath = common.GoModuleName(o.SaveTo)

	dms := ReadDomainModels(filepath.Join(path, "domain-model"))
	if len(dms) == 0 {
		var item = DomainModel{}
		if err := common.ReadModelFromJsonFile(path, &item); err != nil {
			fmt.Println(err)
			return
		}
		dms = append(dms, item)
	}
	dmGen := DomainModelGenFactory()
	dmGen.GenCommon(dms, o)
	for i := range dms {
		if err := dmGen.GenDomainModel(dms[i], o); err != nil {
			log.Println(dms[i].Name, err)
			return
		}
		if err := dmGen.GenPersistence(dms[i], o); err != nil {
			log.Println(dms[i].Name, err)
			return
		}
		if err := dmGen.GenRepository(dms[i], o); err != nil {
			log.Println(dms[i].Name, err)
			return
		}
	}
}

// 从描述文件里面读取模型
func ReadDomainModels(path string) (dms []DomainModel) {
	wkFunc := func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println(path, err)
			return nil
		}
		var dm DomainModel
		data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
		if err := json.Unmarshal(data, &dm); err != nil {
			log.Println(path, err)
			return nil
		}
		dms = append(dms, dm)
		return nil
	}
	filepath.Walk(path, wkFunc)
	return
}
func DomainModelGenFactory() DomainModelGen {
	gen := &GoPgDomainModelGen{}
	return gen
}

type DMOptions struct {
	ProjectPath     string
	SaveTo          string
	DataPersistence string
	Language        string
	ModulePath      string
}
type DomainModelGen interface {
	GenDomainModel(dm DomainModel, o DMOptions) error
	GenRepository(dm DomainModel, o DMOptions) error
	GenPersistence(dm DomainModel, o DMOptions) error
	GenCommon(dm []DomainModel, o DMOptions) error
}

//go pg domain model gen
type GoPgDomainModelGen struct{}

func (g *GoPgDomainModelGen) GenDomainModel(dm DomainModel, o DMOptions) error {
	filePath := "/pkg/domain"
	buf := bytes.NewBuffer(nil)
	for i := range dm.Fields {
		field := dm.Fields[i]
		buf.WriteString(fmt.Sprintf("	// %s\n", field.Desc))
		buf.WriteString(fmt.Sprintf("	%v %v `json:\"%v\"`", field.Name, field.TypeValue, common.LowFirstCase(field.Name)))
		if i != len(dm.Fields)-1 {
			buf.WriteString("\n")
		}
	}
	tP, err := template.New("controller").Parse(tmplProtocolDomainModel)
	if err != nil {
		log.Fatal(err)
	}
	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["Model"] = dm.Name
	m["Items"] = buf.String()
	tP.Execute(bufTmpl, m)

	saveTo(o, filePath, filename(dm.Name, "go"), bufTmpl.Bytes())
	return nil
}
func (g *GoPgDomainModelGen) GenRepository(dm DomainModel, o DMOptions) error {
	filePath := "/pkg/infrastructure/repository"

	tP, err := template.New("controller").Parse(tmplProtocolDomainPgRepository)
	if err != nil {
		log.Fatal(err)
	}

	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["Model"] = dm.Name
	m["Module"] = common.GoModuleName(o.SaveTo)
	tP.Execute(bufTmpl, m)

	return saveTo(o, filePath, filename("Pg"+dm.Name+"Repository", "go"), bufTmpl.Bytes())
}
func (g *GoPgDomainModelGen) GenPersistence(dm DomainModel, o DMOptions) error {
	filePath := "/pkg/infrastructure/pg/models"
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("	%v %v `pg:\"%v\"`\n", "tableName", "struct{}", common.LowCasePaddingUnderline(dm.Name)))
	for i := range dm.Fields {
		field := dm.Fields[i]
		buf.WriteString(fmt.Sprintf("	//	%s\n", field.Desc))
		buf.WriteString(fmt.Sprintf("	%v %v", field.Name, field.TypeValue))
		if i != len(dm.Fields)-1 {
			buf.WriteString("\n")
		}
	}
	tP, err := template.New("controller").Parse(tmplProtocolPgModel)
	if err != nil {
		log.Fatal(err)
	}

	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["Model"] = dm.Name
	m["Items"] = buf.String()
	tP.Execute(bufTmpl, m)

	return saveTo(o, filePath, filename("Pg"+dm.Name, "go"), bufTmpl.Bytes())
}

func (g *GoPgDomainModelGen) GenCommon(dm []DomainModel, o DMOptions) error {
	err := g.genConstant(dm, o)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = g.genPgInit(dm, o)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if common.FileExists(filepath.Join(o.SaveTo, "/pkg/infrastructure/pg/transaction", filename("transaction", "go"))) {
		log.Println("【gen code】 jump:", o.SaveTo, "/pkg/infrastructure/pg/transaction", filename("transaction", "go"))
		return nil
	}
	return saveTo(o, "/pkg/infrastructure/pg/transaction", filename("transaction", "go"), []byte(tmplPgTransaction))
}

func (g *GoPgDomainModelGen) genConstant(dm []DomainModel, o DMOptions) error {
	var filePath = "/pkg/constant"
	tP, err := template.New("controller").Parse(tmplConstantPg)
	if err != nil {
		log.Fatal(err)
	}

	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	tP.Execute(bufTmpl, m)

	if common.FileExists(filepath.Join(o.SaveTo, filePath, filename("postgresql", "go"))) {
		log.Println("【gen code】 jump:", o.SaveTo, filePath, filename("postgresql", "go"))
		return nil
	}
	saveTo(o, filePath, filename("postgresql", "go"), bufTmpl.Bytes())

	return nil
}
func (g *GoPgDomainModelGen) genPgInit(dm []DomainModel, o DMOptions) error {
	var filePath = "/pkg/infrastructure/pg"
	tP, err := template.New("controller").Parse(tmplPgInit)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	for i := range dm {
		m := dm[i].Name
		buf.WriteString(fmt.Sprintf("			(*models.%v)(nil),", m))
		if i != len(dm)-1 {
			buf.WriteString("\n")
		}
	}
	bufTmpl := bytes.NewBuffer(nil)
	m := make(map[string]string)
	m["models"] = buf.String()
	m["Module"] = o.ModulePath
	tP.Execute(bufTmpl, m)

	saveTo(o, filePath, filename("init", "go"), bufTmpl.Bytes())

	return nil
}

//保存文件
func saveTo(o DMOptions, st string, filename string, data []byte) (err error) {
	filePath := filepath.Join(o.SaveTo, st)
	if _, e := os.Stat(filePath); e != nil {
		//log.Println(filePath,e)
		if err = os.MkdirAll(filePath, 0777); err != nil {
			return
		}
	}
	log.Println("【gen code】", "path:", filePath, "file:", filename)
	return ioutil.WriteFile(filepath.Join(filePath, filename), data, os.ModePerm)
}

//文件名称
func filename(filename, suffix string) string {
	if len(suffix) == 0 {
		suffix = "go"
	}
	return fmt.Sprintf("%v.%v", common.LowCasePaddingUnderline(filename), suffix)
}

//领域模型
type DomainModel struct {
	Name      string   `json:"name"`
	ValueType string   `json:"type"`
	Fields    []*field `json:"fields"`
}
type field struct {
	Name      string `json:"name"`
	TypeValue string `json:"type"`
	Desc      string `json:"desc"`
}
