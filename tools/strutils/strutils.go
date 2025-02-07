package strutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/asaskevich/govalidator"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/shopspring/decimal"
)

func ToInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("string parse int64 failed, err = %v", err)
	}

	return v
}

func ToInt32(s string) int32 {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Printf("string parse int64 failed, err = %v", err)
	}

	return int32(v)
}

func ToFloat64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("string parse float64 failed, err = %v", err)
	}

	return v
}

func ToInt64List(s []string) []int64 {
	y := make([]int64, len(s))
	for i, v := range s {
		y[i] = ToInt64(v)
	}

	return y
}

func FormatInt(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatFloat64(i float64, fmt byte, prec int) string {
	return strconv.FormatFloat(i, fmt, prec, 64)
}

func Float64Compare(i, j float64) int {
	f1Dec := decimal.NewFromFloat(i)
	f2Dec := decimal.NewFromFloat(j)

	return f1Dec.Cmp(f2Dec)
}

// Float32ToFloat64 float32转float64
func Float32ToFloat64(f float32) float64 {
	str := fmt.Sprintf("%f", f)
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Float32ToFloat64 err:", err, f)
	}
	return v
}

func Int64ArrayToString(a []int64, delim string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}

	return strings.Join(b, delim)
}

func RemoveDuplicates(array []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range array {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}

	return result
}

func RemoveRepeatedElement(ss []string) (result []string) {
	tmp := make(map[string]bool)
	for _, s := range ss {
		tmp[s] = true
	}
	for s := range tmp {
		result = append(result, s)
	}
	return result
}

func CompareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for i < n || j < m {
		x := 0
		for ; i < n && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++ // 跳过点号
		y := 0
		for ; j < m && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++ // 跳过点号
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

// StrToBytes string 转为byte 高效
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToStr byte 转为string 高效
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsNumber(str string) bool {
	pattern := "\\d+"
	result, _ := regexp.MatchString(pattern, str)
	return result
}

// Diff Creates a slice of slice values not included in the other given slice.
func Diff(base, exclude []string) (result []string) {
	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}
	return result
}

func CamelCaseToUnderscore(str string) string {
	return govalidator.CamelCaseToUnderscore(str)
}

func UnderscoreToCamelCase(str string) string {
	return govalidator.UnderscoreToCamelCase(str)
}

func FindString(array []string, str string) int {
	for index, s := range array {
		if str == s {
			return index
		}
	}
	return -1
}

func StringIn(str string, array []string) bool {
	return FindString(array, str) > -1
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func Reverse1(s string) string {
	return strutil.Reverse(s)
}

// AsString 转成string
func AsString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			return string(b)
		}
	}
	// return fmt.Sprintf("%v", src)
}

// UnicodeEmojiDecode Emoji表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

// UnicodeEmojiCode Emoji表情转换
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

// DbcToSbc 全角转半角
func DbcToSbc(str string) string {
	numConv := unicode.SpecialCase{
		unicode.CaseRange{
			Lo: 0x3002, // Lo 全角句号
			Hi: 0x3002, // Hi 全角句号
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x002e - 0x3002, // LowerCase 转成半角句号
				0,               // TitleCase
			},
		},
		//
		unicode.CaseRange{
			Lo: 0xFF01, // 从全角！
			Hi: 0xFF19, // 到全角 9
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x0021 - 0xFF01, // LowerCase 转成半角
				0,               // TitleCase
			},
		},
		unicode.CaseRange{
			Lo: 0xff21, // Lo: 全角 Ａ
			Hi: 0xFF5A, // Hi:到全角 ｚ
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x0041 - 0xff21, // LowerCase 转成半角
				0,               // TitleCase
			},
		},
	}

	return strings.ToLowerSpecial(numConv, str)
}

// RemoveMark 去掉标点符号
func RemoveMark(text string) string {
	// 全角转半角
	var out []rune
	tmp := []rune(DbcToSbc(text))
	for i := 0; i < len(tmp); i++ {
		// if find := strings.Contains(",:-、.;?!…", string(tmp[i])); find {
		if !isPunct(tmp[i]) {
			out = append(out, tmp[i])
		}
	}

	return string(out)
}

// 判断是否是标点 !,—.:;?…、
func isPunct(r rune) bool {
	if r == '!' || r == ',' || r == '—' || r == '.' || r == ':' || r == ';' || r == '?' || r == '…' || r == '、' {
		return true
	}

	return false
}

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, length int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, length)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

// SubstrTarget 在字符串中查找指定子串，并返回left或right部分
func SubstrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos == true {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if hasPos == false {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
	}
}

// GetStringUtf8Len 获得字符串按照uft8编码的长度
func GetStringUtf8Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// JoinStringAndOther 连接字符串和其他类型
func JoinStringAndOther(val ...interface{}) string {
	return fmt.Sprint(val...)
}

// CamelToSnake 驼峰转蛇形
func CamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// SnakeToCamel 蛇形转驼峰
func SnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func RemoveTarget(array []string, target string) []string {
	length := len(array)
	if length == 0 {
		return array
	}

	array = RemoveDuplicates(array)

	var res []string
	for i, _ := range array {
		if strings.EqualFold(array[i], target) {
			res = append(res, array[0:i]...)
			res = append(res, array[i+1:]...)
			return res
		}
	}
	return array
}

func Union(slice1, slice2 []string) []string {
	slice1 = RemoveDuplicates(slice1)
	slice2 = RemoveDuplicates(slice2)

	m := make(map[string]bool)
	res := make([]string, 0)
	for _, v := range slice1 {
		m[v] = true
		res = append(res, v)
	}

	for _, v := range slice2 {
		exist, _ := m[v]
		if !exist {
			res = append(res, v)
		}
	}
	return res
}

func Intersect(slice1, slice2 []string) []string {
	slice1 = RemoveDuplicates(slice1)
	slice2 = RemoveDuplicates(slice2)

	m := make(map[string]bool)
	for _, v := range slice1 {
		m[v] = true
	}

	res := make([]string, 0)
	for _, v := range slice2 {
		exist, _ := m[v]
		if exist {
			res = append(res, v)
		}
	}
	return res
}

func Difference(slice1, slice2 []string) []string {
	slice1 = RemoveDuplicates(slice1)
	slice2 = RemoveDuplicates(slice2)

	m := make(map[string]bool)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v] = true
	}

	res := make([]string, 0)
	for _, value := range slice1 {
		exist, _ := m[value]
		if !exist {
			res = append(res, value)
		}
	}
	return res
}
