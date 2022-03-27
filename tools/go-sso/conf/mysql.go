package conf

// DBConfig ...
type DBConfig struct {
	DriverName   string
	Dsn          string
	ShowSQL      bool
	ShowExecTime bool
	MaxIdle      int
	MaxOpen      int
}

var Db = map[string]DBConfig{
	"db1": {
		DriverName:   "mysql",
		Dsn:          "user:pass@tcp(127.0.0.1:3306)/ssodb?charset=utf8mb4&parseTime=true&loc=Local",
		ShowSQL:      true,
		ShowExecTime: false,
		MaxIdle:      10,
		MaxOpen:      200,
	},
}
