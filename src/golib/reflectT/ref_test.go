package reflectT

import (
	"encoding"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

type ReflectDemo struct {
	Key string   `dns:"key,127.0.0.1"`
	Value []int  `dns:"value,root"`
	Ptr  interface{}
}

func TestParseFlag(t *testing.T){
	demo :=ReflectDemo{
		Key:"aaa",
		Value:[]int{1,2,3},
		Ptr:"tiptok",
	}
	typeDemo := reflect.TypeOf(demo)
	valueDemo :=reflect.ValueOf(demo)
	log.Println("type.Name() :",typeDemo.Name())
	log.Println("type.PkgPath() :",typeDemo.PkgPath())
	log.Println("type.String() :",typeDemo.String())
	log.Println("type.Size() :",typeDemo.Size())


	for i:=0;i<typeDemo.NumField();i++{
		vf :=valueDemo.Field(i)//reflect.value 值
		tf :=typeDemo.Field(i)//属性信息
		if tag,ok:=tf.Tag.Lookup("dns");i==0 && ok{
			log.Println("Lookup Tag :",strings.TrimSpace(tag))
			tagopt :=parseTag(strings.TrimSpace(tag))
			jsTag,_:= json.Marshal(tagopt)
			log.Println("Parse Tag :",string(jsTag))

			//动态设置字段值
			nv :=reflect.New(vf.Type())
			SetValue(nv,"tip")
			log.Println(fmt.Sprintf("Values:%v %v",nv.String(),*(nv.Interface().(*string))))
			//nv.Set(valueDemo)
		}
		log.Println("Index:",i,"vt->",vf,tf)
	}
	switch typeDemo.Kind() {
	case reflect.Struct:
		log.Println("type is struct")
	case reflect.Interface:
	case reflect.Ptr:
		log.Println("type is Ptr")
	default:
		log.Println(valueDemo.Kind())
	}

	log.Println("ReflectDemo : ",valueDemo.Kind())
}
//SetValue 动态设置值
func SetValue(v reflect.Value, value string)error{
	var tu encoding.TextUnmarshaler
	tu, v = indirect(v)
	if tu != nil {

	}
	switch v.Kind() {
	case reflect.Bool:
		//err = d.valueBool(v, prefix, to)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		//err = d.valueInt64(v, prefix, to)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		//err = d.valueUint64(v, prefix, to)
	case reflect.Float32, reflect.Float64:
		//err = d.valueFloat64(v, prefix, to)
	case reflect.String:
		//err = d.valueString(v, prefix, to)
		v.SetString(value)
	case reflect.Slice:
		//err = d.valueSlice(v, prefix, to)
	case reflect.Struct:
		//err = d.valueStruct(v, prefix, to)
	case reflect.Ptr:
		//if !d.hasKey(combinekey(prefix, to)) {
		//	break
		//}
		//if !v.CanSet() {
		//	break
		//}
		nv := reflect.New(v.Type().Elem())
		v.Set(nv)
		SetValue(nv,value)
	}
	return nil
}
//将指针类型 指向底层数据
func indirect(v reflect.Value) (encoding.TextUnmarshaler, reflect.Value) {
	v0 := v
	haveAddr := false

	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && e.Elem().Kind() == reflect.Ptr {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		if v.Elem().Kind() != reflect.Ptr && v.CanSet() {
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 {
			if u, ok := v.Interface().(encoding.TextUnmarshaler); ok {
				return u, reflect.Value{}
			}
		}
		if haveAddr {
			v = v0
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return nil, v
}

//结构附加标签
type tagOpt struct {
	Name    string
	Default string
}
func parseTag(tag string)tagOpt{
	vs := strings.SplitN(tag, ",", 2)
	if len(vs) == 2 {
		return tagOpt{Name: vs[0], Default: vs[1]}
	}
	return tagOpt{Name: vs[0]}
}

type BindTypeError struct {
	Value string
	Type  reflect.Type
}
func (e *BindTypeError) Error() string {
	return "cannot decode " + e.Value + " into Go value of type " + e.Type.String()
}



func Bind(v interface{}) error {
	//assignFuncs := make(map[string]assignFunc)
	return nil
}
type assignFunc func(v reflect.Value, to tagOpt) error
func stringsAssignFunc(val string) assignFunc {
	return func(v reflect.Value, to tagOpt) error {
		if v.Kind() != reflect.String || !v.CanSet() {
			return &BindTypeError{Value: "string", Type: v.Type()}
		}
		if val == "" {
			v.SetString(to.Default)
		} else {
			v.SetString(val)
		}
		return nil
	}
}