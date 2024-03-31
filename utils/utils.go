package utils

import (
	"bytes"
	"encoding/gob"
	"os"
	"regexp"
	"strings"
)

// IsPhone 判断是否是手机号
func IsPhone(mobileNum string) bool {
	tmp := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
	reg := regexp.MustCompile(tmp)
	return reg.MatchString(mobileNum)
}

// IsIDCard 判断是否是18或15位身份证
func IsIDCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// EncodeByte 编码二进制
func EncodeByte(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeByte 解码二进制
func DecodeByte(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

// Getenv 获取本地系统变量
func Getenv(key string) string {
	return os.Getenv(key)
}

// GetLocalSystemLang 获取本地语言 (like:zh_CN.UTF-8)(simple:zh)
func GetLocalSystemLang(isSimple bool) (locale string) {
	locale = Getenv("LC_ALL")
	if locale == "" {
		locale = Getenv("LANG")
	}
	if isSimple {
		locale, _ = splitLocale(locale)
	}
	if len(locale) == 0 {
		locale = "zh"
	}
	return
}

func splitLocale(locale string) (string, string) {
	formattedLocale := strings.Split(locale, ".")[0]
	formattedLocale = strings.Replace(formattedLocale, "-", "_", -1)

	pieces := strings.Split(formattedLocale, "_")
	language := pieces[0]
	territory := ""
	if len(pieces) > 1 {
		territory = strings.Split(formattedLocale, "_")[1]
	}
	return language, territory
}
