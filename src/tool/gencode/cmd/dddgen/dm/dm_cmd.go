package dm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiptok/OFAppTest/src/tool/gencode/common"
	"github.com/tiptok/OFAppTest/src/tool/gencode/constant"
	"github.com/tiptok/OFAppTest/src/tool/gencode/model"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	readPath := path
	if !strings.Contains(readPath, "domain-model") {
		readPath = filepath.Join(path, "domain-model")
	}
	dms := ReadDomainModels(readPath)
	if len(dms) == 0 {
		return
	}
	dmGen := DomainModelGenFactory()
	dmGen.GenCommon(dms, o)
	for i := range dms {
		if err := dmGen.GenDomainModel(dms[i], o); err != nil {
			log.Println(dms[i].Name, err)
			return
		}
		// 值对象不需要生成持久模型/仓库模型
		if dms[i].ValueType != string(constant.DomainModel) {
			log.Println("jump", dms[i].ValueType, constant.DomainModel)
			continue
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
		//log.Println(path)
		if info.IsDir() {
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
	model.SvrOptions
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
	m := make(map[string]interface{})
	m["Model"] = dm.Name
	m["Items"] = buf.String()
	m["Desc"] = dm.Desc
	m["Fields"] = dm.ColumnsNeedUpdate()
	if len(dm.Desc) == 0 {
		m["Desc"] = dm.Name
	}
	m["IsDomainModel"] = dm.ValueType == string(constant.DomainModel)
	tP.Execute(bufTmpl, m)
	fileName := dm.Name
	if dm.ValueType == string(constant.DomainModel) {
		fileName = "Do" + dm.Name
	}
	if dm.ValueType == string(constant.DomainValue) {
		fileName = "Dv" + dm.Name
	}
	saveTo(o, filePath, filename(fileName, "go"), bufTmpl.Bytes())
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
	m["DBName"] = "constant.POSTGRESQL_DB_NAME"
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
	m["Desc"] = dm.Desc
	if len(dm.Desc) == 0 {
		m["Desc"] = dm.Name
	}
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
	ValueType string   `json:"value_type"` // domain-model 领域模型  domain-value值对象 property
	Property  []string `json:"property"`   // restful
	Desc      string   `json:"desc"`
	Fields    []*field `json:"fields"`
}

func (dm DomainModel) NeedRestful() bool {
	for _, v := range dm.Property {
		if strings.TrimSpace(v) == "restful" {
			return true
		}
	}
	return false
}

func (dm DomainModel) ColumnsNeedUpdate() interface{} {
	var fields []struct {
		Name   string
		Column string
		Type   string
	}
	for _, v := range dm.Fields {
		if v.Name == "Id" {
			continue
		}
		fields = append(fields, struct {
			Name   string
			Column string
			Type   string
		}{Name: v.Name, Column: common.LowFirstCase(v.Name), Type: v.TypeValue})
	}
	return fields
}

type field struct {
	Name      string `json:"name"`
	TypeValue string `json:"type"`
	Desc      string `json:"desc"`
}
