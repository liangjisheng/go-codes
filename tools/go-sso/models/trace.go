package models

// Trace ...
type Trace struct {
	ID    int64  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	UID   int64  `json:"uid" xorm:"not null default 0 comment('用户主键') index(UT) BIGINT(20)"`
	Type  int    `json:"type" xorm:"not null default 0 comment('类型(0:注册1::登录2:退出3:修改4:删除)') index(UT) TINYINT(4)"`
	IP    int    `json:"ip" xorm:"not null comment('ip') INT(10)"`
	Ext   string `json:"ext" xorm:"not null comment('扩展字段') VARCHAR(1000)"`
	Ctime int    `json:"ctime" xorm:"not null default 0 comment('注册时间') INT(11)"`
}

var (
	// TraceTypeReg ...
	TraceTypeReg = 0
	// TraceTypeLogin ...
	TraceTypeLogin = 1
	// TraceTypeOut ...
	TraceTypeOut = 2
	// TraceTypeEdit ...
	TraceTypeEdit = 3
	// TraceTypeDel ...
	TraceTypeDel = 4
)
