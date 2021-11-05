package localizer

import (
	// 引入internal/translations包，确保 init()函数被调用
	_ "bookstore.example.com/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 定义 Localizer 类型保存相关本地化ID (和URL中使用类似)
// 创建不可导出的 message.Printer 实例
type Localizer struct {
	ID      string
	printer *message.Printer
}

// 初始化一个片，其中保存了我们支持的每个区域设置的初始化本地化器类型
var locales = []Localizer{
	{
		// 德国
		ID:      "de-de",
		printer: message.NewPrinter(language.MustParse("de-DE")),
	},
	{
		// 中国
		ID:      "zh-cn",
		printer: message.NewPrinter(language.MustParse("zh-CN")),
	},
	{
		//英国
		ID:      "en-gb",
		printer: message.NewPrinter(language.MustParse("en-GB")),
	},
}

// Get() 函数接收一个本地化ID，并返回对应本地化实例
// 如果本地化ID不支持将返回空，false作为第二个参数值
func Get(id string) (Localizer, bool) {
	for _, locale := range locales {
		if id == locale.ID {
			return locale, true
		}
	}

	return Localizer{}, false
}

// 为本地化类型增加Translate()方法，该方法对消息和参数进行包装
func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}