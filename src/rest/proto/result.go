package proto

type Result struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

type Error struct {
	Result string
	Err    string
}

func NewResult(result string, data interface{}) (out *Result, err error) {
	out = &Result{Result: result, Data: data}
	return
}

func NewError(err string) (out *Error) {
	out = &Error{Result: "false", Err: err}
	return
}
