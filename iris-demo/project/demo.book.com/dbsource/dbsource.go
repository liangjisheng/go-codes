package dbsource

import (
	"fmt"
	"log"
	"sync"

	"demo.book.com/conf"

	// mysql
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// InstanceMaster 主库单例
func InstanceMaster() *xorm.Engine {
	// 调用时需要判断是否已经初始化过
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	// 如果初始化的时候同时有2个G走到了这里,需要再次判断一次
	if masterEngine != nil {
		return masterEngine
	}

	// root:112233@tcp(127.0.0.1:3305)/mygo?charset=utf8
	driveSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", conf.SysConfMap["dbuser"], conf.SysConfMap["dbpwd"], conf.SysConfMap["dbhost"], conf.SysConfMap["dbport"], conf.SysConfMap["dbname"])
	fmt.Println("InstanceMaster数据库链接", driveSource)

	engine, err := xorm.NewEngine("mysql", driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	// engine.SetTZLocation(conf.SysTimeLocation)
	//
	// 性能优化的时候才考虑，加上本机的SQL缓存
	// cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	// engine.SetDefaultCacher(cacher)
	masterEngine = engine
	return engine
}

// InstanceSlave 从库单例
func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}
	driveSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", conf.SysConfMap["dbuser"], conf.SysConfMap["dbpwd"], conf.SysConfMap["dbhost"], conf.SysConfMap["dbport"], conf.SysConfMap["dbname"])
	fmt.Println(driveSource)
	engine, err := xorm.NewEngine("mysql", driveSource)
	if err != nil {
		log.Fatal("dbhelper", "DbInstanceMaster", err)
		return nil
	}
	// engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return engine
}
