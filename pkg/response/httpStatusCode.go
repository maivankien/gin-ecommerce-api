package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 40001
	ErrorInvalidToken     = 30001
)

var message = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Param invalid",
	ErrorInvalidToken:     "Invalid token",
}
