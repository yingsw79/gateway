package handler

type ErrCode int64

const (
	Err_BadRequest           ErrCode = 10001
	Err_Unauthorized         ErrCode = 10002
	Err_ServerNotFound       ErrCode = 10003
	Err_ServerMethodNotFound ErrCode = 10004
	Err_RequestServerFail    ErrCode = 10005
	Err_ServerHandleFail     ErrCode = 10006
	Err_ResponseUnableParse  ErrCode = 10007
	Err_DuplicateOutOrderNo  ErrCode = 20001
)

func (p ErrCode) String() string {
	switch p {
	case Err_BadRequest:
		return "BadRequest"
	case Err_Unauthorized:
		return "Unauthorized"
	case Err_ServerNotFound:
		return "ServerNotFound"
	case Err_ServerMethodNotFound:
		return "ServerMethodNotFound"
	case Err_RequestServerFail:
		return "RequestServerFail"
	case Err_ServerHandleFail:
		return "ServerHandleFail"
	case Err_ResponseUnableParse:
		return "ResponseUnableParse"
	case Err_DuplicateOutOrderNo:
		return "DuplicateOutOrderNo"
	}
	return "<UNSET>"
}

type Err struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// New Error, the error_code must be defined in IDL.
func NewErr(errCode ErrCode) Err {
	return Err{
		ErrCode: int64(errCode),
		ErrMsg:  errCode.String(),
	}
}

func (e Err) Error() string { return e.ErrMsg }
