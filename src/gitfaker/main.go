package main

import (
	"encoding/json"
	"github.com/bxcodec/faker"
	"log"
	"reflect"
)

type Some struct {
	Name string `faker:"name"`
	Age  int32
	WorkSkills string  `faker:"url"`
	Like ALike
}

type ALike struct {
	Ball string `faker:"diff"`  //custom tag
	Game string `faker:"last_name"`
	Amount float64  `faker:"amount"`
	Course []string `faker:"len=10"`
	Friend BFriend
}

type BFriend struct {
	Id string `faker:"uuid_digit"`
	Desc string `faker:"-"`
}

func main(){
	faker.AddProvider("diff",func(v reflect.Value)(interface{},error){
		return "null",nil
	})
	var s Some
	if err :=faker.FakeData(&s);err!=nil{
		log.Println(err)
	}
	jsData,_ :=json.Marshal(s)
	log.Println(string(jsData))
}
