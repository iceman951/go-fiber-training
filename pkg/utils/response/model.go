package response

type ResponseMessageStatus struct {
	Status    string `json:"response_status"`
	Message   string `json:"response_message"`
	Timestamp string `json:"response_timestamp"`
}

type ResponseMessageStatusData struct {
	Status    string      `json:"response_status"`
	Message   string      `json:"response_message"`
	Data      interface{} `json:"response_data"`
	Timestamp string      `json:"response_timestamp"`
}
