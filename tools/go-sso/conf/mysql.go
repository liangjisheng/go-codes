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

var Db = map[string]DBConfig {
	"db1": {
		DriverName: "mysql",
		Dsn: "root:ljs199711@tcp(117.51.148.112:3306)/ssodb?charset=utf8mb4&parseTime=true&loc=Local"
		ShowSQL: true,
		ShowExecTime: false,
		MaxIdle: 10,
		MaxOpen: 200,
	}
}
