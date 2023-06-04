package enum

//go:generate stringer -type ErrCode -linecomment

type ErrCode int64 //错误码

const (
	ErrCodeOK            ErrCode = 0 // OK
	ErrCodeInvalidParams ErrCode = 1 // 无效参数
	ErrCodeTimeout       ErrCode = 2 // 超时
)

//执行 go generate 生成对应的错误码和错误消息
