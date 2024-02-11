package constant

const (
	EVENT_TYPE_REQUEST  string = "Request"
	EVENT_TYPE_RESPONSE string = "Response"
	EVENT_TYPE_REJECT   string = "Reject"
	EVENT_TYPE_PROCESS  string = "Process"
	EVENT_TYPE_ERROR    string = "Error"
	EVENT_TYPE_WARN     string = "Warn"

	INIT_MONGO string = "INIT_MONGO"
	QUERY_ONE  string = "QUERY_ONE"
	QUERY_ALL  string = "QUERY_ALL"
	UPDATE     string = "UPDATE"
	DELETE_ONE string = "DELETE_ONE"
)
