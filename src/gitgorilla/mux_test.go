package gitgorilla

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"log"

	"github.com/gorilla/mux"
)

func TestMuxTopDB(t *testing.T) {
	r := mux.NewRouter()
	//:*/Vehicles
	r.HandleFunc("/Vehicles", GetVehiclesHandler)
	//:*/Vehicles/get  子路由
	subV := r.PathPrefix("/Vehicles").Subrouter()
	subV.HandleFunc("/get", GetVehiclesHandler)
	/*Paths can have variables*/
	subV.HandleFunc("/add/{SimNum}", AddVehiclesHandler).
		Methods("GET", "POST").
		Schemes("http")
		//r.PathPrefix("/products/")
		//r.Methods("GET", "POST")
		//r.Schemes("https")
		//r.Headers("X-Requested-With", "XMLHttpRequest")
		//subV.Queries

		/*打印已注册的路由*/
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		// queriesTemplates, err := route.GetQueriesTemplates()
		// if err == nil {
		// 	fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		// }
		// queriesRegexps, err := route.GetQueriesRegexp()
		// if err == nil {
		// 	fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		// }
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	/*
	   ROUTE: /Vehicles/add/{SimNum}
	   Path regexp: ^/Vehicles/add/(?P<v0>[^/]+)$
	   Methods: GET,POST
	*/

	if err != nil {
		fmt.Println(err)
	}

	log.Fatal(http.ListenAndServe(":8090", r))
}

type RequesVehicle struct {
	SimNum string
}
type ResponseVehicle struct {
	PlateNum string
	SimNum   string
}

func GetVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	response := ResponseVehicle{"闽A", "18860183012"}
	jsData, _ := json.Marshal(response)
	w.Write(jsData)
}

func AddVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	/*mux.Vars(r)*/
	vars := mux.Vars(r)
	response := ResponseVehicle{"闽A", vars["SimNum"]}
	jsData, _ := json.Marshal(response)
	//w.WriteHeader
	w.Write(jsData)
}
