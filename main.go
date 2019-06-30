package main

import (
	"flag"
	"fmt"
	conf "golang-simple-crud/config"
	hdlr "golang-simple-crud/handler"
	mdw "golang-simple-crud/middleware"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// open mysql connection
	configFile := flag.String("conf", "config/config.yml", "main configuration file")

	flag.Parse()
	conf.LoadConfigFromFile(configFile)
	conn, err := conf.New(conf.Param.DBType, conf.Param.DBUrl)
	conf.Logf("Load Database Conf: %s ", conf.Param.DBType)
	conf.Logf("running App on port: %s ", conf.Param.ListenPort)

	if err != nil {
		conf.Logf("Load Database Conf: %s ", err)
		log.Fatal(err)
	}

	http.HandleFunc("/api/employeeSave", mdw.Chain(hdlr.EmployeeSaveHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST")))
	http.HandleFunc("/api/employeeUpdate", mdw.Chain(hdlr.EmployeeUpdateHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST")))
	http.HandleFunc("/api/employee", mdw.Chain(hdlr.GetEmployeeHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST")))
	http.HandleFunc("/api/employeeDetail", mdw.Chain(hdlr.GetEmployeeDetailHandler(conn), mdw.ContentType("application/json"), mdw.Method("POST")))

	var errors error
	errors = http.ListenAndServe(conf.Param.ListenPort, nil)

	if errors != nil {
		fmt.Println("error", errors)
		conf.Logf("Unable to start the server: %s ", conf.Param.ListenPort)
		os.Exit(1)
	}
}
