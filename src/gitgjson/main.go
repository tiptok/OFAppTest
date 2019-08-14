package main

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
)

func main(){
	data,err:=ioutil.ReadFile("data.json")
	if err!=nil{
		log.Println(err)
		return
	}
	json:=string(data)
	log.Println("json.data:",json)
	if !gjson.Valid(json){
		log.Println("invalid json")
	}
	if m,ok:=gjson.Parse(json).Value().(map[string]interface{});ok{
		log.Println("parse json to map:",m["age"])
	}

	log.Println("1.get a value:",gjson.Get(json,"age"),gjson.Get(json,"name.last"))
	log.Println("1.get a value:",gjson.Get(json,"name").Get("last"))
	log.Println("1.get a value:",gjson.Get(json,"age").Array())
	log.Println("1.get Bytes:",gjson.GetBytes(data,"name").String())
	log.Println("1.get Bytes:",gjson.GetMany(json,"age","name"))
	name :=gjson.Get(json,"age1")
	if !name.Exists(){
		log.Println("not exists key:","age1")
	}

	log.Println("2.path syntax  number:",gjson.Get(json,"friends.#"))
	log.Println("2.path syntax  :",gjson.Get(json,"friends.1"))
	log.Println("2.path syntax  :",gjson.Get(json,"friends.1.nets"))
	log.Println("2.path syntax  :",gjson.Get(json,"friends.1.nets"))

	log.Println("2.path syntax  first match:",gjson.Get(json,`friends.#(last=="Murphy").first`))
	log.Println("2.path syntax  all match:",gjson.Get(json,`friends.#(last=="Murphy")#.first`))
	log.Println("2.path syntax  all match:",gjson.Get(json,`friends.#(age>45)#.first`))

	firstArray :=gjson.Get(json,`friends.#(age>45)#.first`)
	if firstArray.Exists(){
		firstArray.ForEach(func(key,value gjson.Result)bool{
			log.Println(" value:",value)
			return value.String()=="Jane"  //keep iterating
			//return true
		})
	}
}
