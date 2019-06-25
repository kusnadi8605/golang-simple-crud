package datastruct

//MiddlewareResponse data
type MiddlewareResponse struct {
	ResponseCode string     `json:"responseCode"`
	ResponseDesc string     `json:"responseDesc"`
	Payload      []Employee `json:"payload"`
}
