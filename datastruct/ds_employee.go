package datastruct

//Employee ..
type Employee struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID          int
	Name        string
	Phone       *string
	Email       *string
	Address     *string
	CreatedDate *string
	UpdatedDate *string
}

//EmployeeRequest ..
type EmployeeRequest struct {
	ID string `json:"id,omitempty"`
}

//EmployeeResponse data
type EmployeeResponse struct {
	ResponseCode string     `json:"responseCode"`
	ResponseDesc string     `json:"responseDesc"`
	Payload      []Employee `json:"payload"`
}

//EmployeeDetail ..
type EmployeeDetail struct {
	ID          int
	Name        string
	Phone       *int32
	Email       *string
	Address     *string
	CreatedDate *string
	UpdatedDate *string
}

//EmployeeDetailRequest ..
type EmployeeDetailRequest struct {
	ID string `json:"id,omitempty"`
}

//EmployeeDetailResponse data
type EmployeeDetailResponse struct {
	ResponseCode string           `json:"responseCode"`
	ResponseDesc string           `json:"responseDesc"`
	Payload      []EmployeeDetail `json:"payload"`
}
