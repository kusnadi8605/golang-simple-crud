package handler

import (
	"encoding/json"
	conf "golang-simple-crud/config"
	dts "golang-simple-crud/datastruct"
	mdl "golang-simple-crud/models"
	"io/ioutil"
	"net/http"
	"time"
)

//GetEmployeeHandler with paging
func GetEmployeeHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var EmployeeResponse dts.EmployeeResponse

		logDate := time.Now().Format("20060102")
		conf.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["Employee"] + logDate + ".txt")
		listEmployee, err := mdl.GetEmployee(conn)

		if err != nil {
			EmployeeResponse.ResponseCode = "301"
			EmployeeResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(EmployeeResponse)

			conf.Logf("Response Employee : %v", EmployeeResponse.ResponseDesc)

			return
		}

		if len(listEmployee) < 1 {
			EmployeeResponse.ResponseCode = "301"
			EmployeeResponse.ResponseDesc = "data not found"
			json.NewEncoder(w).Encode(EmployeeResponse)

			conf.Logf("Response Employee : %v", EmployeeResponse.ResponseDesc)

			return
		}

		EmployeeResponse.ResponseCode = "000"
		EmployeeResponse.ResponseDesc = "Success"
		EmployeeResponse.Payload = listEmployee
		json.NewEncoder(w).Encode(EmployeeResponse)

		conf.Logf("Response Employee : %v", EmployeeResponse.ResponseDesc)
	}
}

//GetEmployeeDetailHandler single Employee
func GetEmployeeDetailHandler(conn *conf.Connection) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqEmployeeDetail dts.EmployeeDetailRequest
		var EmployeeDetailResponse dts.EmployeeDetailResponse

		logDate := time.Now().Format("20060102")
		conf.SetFilename(conf.Param.LogDir + conf.Param.LogsFile["Employee"] + logDate + ".txt")

		body, err := ioutil.ReadAll(req.Body)

		err = json.Unmarshal(body, &reqEmployeeDetail)
		if err != nil {
			EmployeeDetailResponse.ResponseCode = "500"
			EmployeeDetailResponse.ResponseDesc = err.Error()

			json.NewEncoder(w).Encode(EmployeeDetailResponse)

			conf.Logf("Decode Employee : %s", err)

			return
		}

		ID := reqEmployeeDetail.ID

		detailEmployee, err := mdl.GetEmployeeDetail(conn, ID)

		if err != nil {
			EmployeeDetailResponse.ResponseCode = "301"
			EmployeeDetailResponse.ResponseDesc = err.Error()
			EmployeeDetailResponse.Payload = detailEmployee
			json.NewEncoder(w).Encode(EmployeeDetailResponse)

			conf.Logf("Response Employee : %v", EmployeeDetailResponse.ResponseDesc)

			return
		}

		if len(detailEmployee) < 1 {
			EmployeeDetailResponse.ResponseCode = "301"
			EmployeeDetailResponse.ResponseDesc = "data not found"
			//EmployeeDetailResponse.Payload = detailEmployee
			json.NewEncoder(w).Encode(EmployeeDetailResponse)

			conf.Logf("Response Employee : %v", EmployeeDetailResponse.ResponseDesc)

			return
		}

		EmployeeDetailResponse.ResponseCode = "000"
		EmployeeDetailResponse.ResponseDesc = "Success"
		EmployeeDetailResponse.Payload = detailEmployee
		json.NewEncoder(w).Encode(EmployeeDetailResponse)

		conf.Logf("Response Employee : %v", EmployeeDetailResponse.ResponseDesc)
	}
}
