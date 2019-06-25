package models

import (
	conf "golang-simple-crud/config"
	dts "golang-simple-crud/datastruct"
)

//GetEmployee as
func GetEmployee(conn *conf.Connection) ([]dts.Employee, error) {
	arrEmployee := []dts.Employee{}
	strEmployee := dts.Employee{}
	strQuery := `select * from employee`
	rows, err := conn.Query(strQuery)

	conf.Logf("Query Employee : %s", strQuery)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(
			&strEmployee.ID,
			&strEmployee.Name,
			&strEmployee.Phone,
			&strEmployee.Email,
			&strEmployee.Address,
			&strEmployee.CreatedDate,
			&strEmployee.UpdatedDate,
		)

		if err != nil {
			return nil, err
		}

		arrEmployee = append(arrEmployee, strEmployee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrEmployee, nil
}

//GetEmployeeDetail ..
func GetEmployeeDetail(conn *conf.Connection, ID string) ([]dts.EmployeeDetail, error) {
	arrEmployeeDetail := []dts.EmployeeDetail{}
	strEmployeeDetail := dts.EmployeeDetail{}

	strQuery := `select * from employee where Id=?`

	rows, err := conn.Query(strQuery, ID)

	conf.Logf("Query Employee  detail: %s %s", strQuery, ID)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&strEmployeeDetail.ID,
			&strEmployeeDetail.Name,
			&strEmployeeDetail.Phone,
			&strEmployeeDetail.Email,
			&strEmployeeDetail.Address,
			&strEmployeeDetail.CreatedDate,
			&strEmployeeDetail.UpdatedDate,
		)

		if err != nil {
			return nil, err
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		arrEmployeeDetail = append(arrEmployeeDetail, strEmployeeDetail)
	}

	return arrEmployeeDetail, nil
}
