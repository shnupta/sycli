package sycli

type SendWrapper struct {
	Command []string `json:"command"`
}

type ReceiveWrapper struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type Command struct {
	CommandName string
	Params      []interface{}
}
