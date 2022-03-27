package tracing

import (
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
	"unsafe"
)

// 包内静态变量
const gormSpanKey = "__gorm_span"

func before(db *gorm.DB) {
	span := opentracing.SpanFromContext(db.Statement.Context)

	// 利用db实例去传递span
	db.InstanceSet(gormSpanKey, span)

	return
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}

	// 注意:这里不去关闭span，是为了再获取结果并写入log后再finish
	//defer span.Finish()

	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))

	if db.Statement.Dest != nil {
		v, err := json.Marshal(db.Statement.Dest)
		if err != nil {
			db.Logger.Error(context.Background(), "could not marshal db.Statement.Dest: %v", err)
		}
		span.LogFields(tracerLog.String("result", *(*string)(unsafe.Pointer(&v))))
	}

	if db.Error != nil {
		span.SetTag("error", true)
		span.LogFields(tracerLog.String("error", db.Error.Error()))
	}

	return
}

const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

type OpentracingPlugin struct{}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {

	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

var _ gorm.Plugin = &OpentracingPlugin{}
