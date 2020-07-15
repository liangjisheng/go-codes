package models

// Device ...
type Device struct {
	ID     int64  `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	UID    int64  `json:"uid" xorm:"not null default 0 comment('用户主键') index BIGINT(20)"`
	Client string `json:"client" xorm:"not null default '' comment('客户端') VARCHAR(50)"`
	Model  string `json:"model" xorm:"not null default '' comment('设备型号') VARCHAR(50)"`
	IP     int    `json:"ip" xorm:"not null default 0 comment('ip地址') INT(10)"`
	Ext    string `json:"ext" xorm:"not null default '' comment('扩展信息') VARCHAR(1000)"`
	Ctime  int    `json:"ctime" xorm:"not null default 0 comment('注册时间') INT(10)"`
}
