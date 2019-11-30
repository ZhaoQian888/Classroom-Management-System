package information

type Response struct {
	Status string      `json:"status"`
	data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:error`
}
