package datastruct

//Employee ..
type Employee struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID          int64
	Name        string
	Phone       *string
	Email       *string
	Address     *string
	CreatedDate *string
	UpdatedDate *string
}

type EmployeeSave struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID int64
}

type EmployeeSaveRequest struct {
	//* tanda bintang supaya tidak err jika datanya null
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

//EmployeeSaveResponse ..
type EmployeeSaveResponse struct {
	ResponseCode string         `json:"responseCode"`
	ResponseDesc string         `json:"responseDesc"`
	Payload      []EmployeeSave `json:"payload"`
}

// EmployeeUpdateRequest ..
type EmployeeUpdateRequest struct {
	//* tanda bintang supaya tidak err jika datanya null
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

//EmployeeUpdateResponse ..
type EmployeeUpdateResponse struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
}

// EmployeDeleteRequest ..
type EmployeDeleteRequest struct {
	ID string `json:"id,omitempty"`
}

//EmployeeDeleteResponse ..
type EmployeeDeleteResponse struct {
	ResponseCode string `json:"responseCode"`
	ResponseDesc string `json:"responseDesc"`
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
