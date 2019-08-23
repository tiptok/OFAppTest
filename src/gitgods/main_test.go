package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"github.com/shopspring/decimal"
	"log"
	"strconv"
	"testing"
)


// List interface that all lists implement
//type List interface {
//	Get(index int) (interface{}, bool)
//	Remove(index int)
//	Add(values ...interface{})
//	Contains(values ...interface{}) bool
//	Sort(comparator utils.Comparator)
//	Swap(index1, index2 int)
//	Insert(index int, values ...interface{})
//	Set(index int, value interface{})
//
//	containers.Container
//}

func Test_ArrayList(t *testing.T){

	//精度问题
	amount :="526.67"
	if ibudegt ,err:=strconv.ParseFloat(amount,64);err==nil{
		log.Println(ibudegt)
		log.Println(ibudegt*100)
		log.Println(int(ibudegt*float64(100)))
		log.Println(int((ibudegt*1000)/10))
		f,_:= decimal.NewFromFloat(ibudegt).Mul(decimal.New(100,0)).Float64()
		log.Println(f,int(f))
	}

	for i:=1;i<=100;i++{
		amount :=1+float64(i)/100.0
		nodecimal := int(amount*100)
		decimal,_ :=decimal.NewFromFloat(amount).Mul(decimal.New(100,0)).Float64()
		if  nodecimal!=int(decimal){
			log.Println(fmt.Sprintf("warn nodecimal:%v decimal:%v amount:%v",nodecimal,int(decimal),amount))
		}
	}

	list := arraylist.New()
	list.Add("tip")
	log.Println(list.Get(0))
	list.Remove(0)
	log.Println(list.Contains("tip"))
	list.Add("1","7","6")
	list.Sort(utils.StringComparator)
	log.Println(list)
	list.Insert(2,"8")
	list.Set(0,"2")
	log.Println(list," size:",list.Size())
	log.Println(list.String())
}
