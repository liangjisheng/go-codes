package conf

import (
	"strings"
	"time"

	. "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/config/source/file"
)

var conf *config

type config struct {
	AppName   string
	MicroConf Config
}

// InitConfig ...
func InitConfig(etcdAddr, appName string) error {
	var err error
	conf = &config{appName, NewConfig()}

	if etcdAddr != "" {
		err = conf.MicroConf.Load(etcd.NewSource(
			etcd.WithAddress(strings.Split(etcdAddr, ",")...),
			etcd.WithPrefix(appName),
			etcd.StripPrefix(true),
			etcd.WithDialTimeout(10*time.Second),
		))
	} else {
		err = conf.MicroConf.Load(file.NewSource(file.WithPath(appName + ".yaml")))
	}
	return err
}

// GetLogPath ...
func GetLogPath() string {
	return conf.MicroConf.Get("log", "path").String(conf.AppName + ".%Y%m%d.log")
}
