package config

// Notifyer 通知接口
type Notifyer interface {
	Callback(*Config)
}
