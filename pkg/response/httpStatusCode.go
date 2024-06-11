package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 40001
)

var message = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Param invalid",
}
