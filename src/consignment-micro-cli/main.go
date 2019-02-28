package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
)

const (
	addr     = ":9927"
	filename = "consignment.json"
)

type ConsignList []*Consignment

func parseFile(file string) (*ConsignList, error) {
	var cl *ConsignList
	data, err := ioutil.ReadFile(file)
	//enc := mahonia.NewDecoder("gbk")
	//jsData := enc.ConvertString(string(data))
	//enc.Translate
	log.Println(string(data))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &cl)
	return cl, err
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	client := NewShipServiceClient(conn)
	cl, err := parseFile(filename)
	if err != nil {
		log.Fatalf("could not parse file:%v", err)
	}
	for i := 0; i < len(*cl); i++ {
		r, err := client.CreateConsignment(context.Background(), (*cl)[i])
		if err != nil {
			log.Println("GetReponse Error", err)
		} else {
			log.Println("GetReponse:", r)
		}
	}
	rpAll, err := client.GetAllConsignment(context.Background(), &RequestAll{})
	if err != nil {
		log.Println("GetAllReponse Error", err)
	} else {
		log.Println("GetAllReponse", rpAll)
	}
	//log.Println("ConsignList", cl)
}
