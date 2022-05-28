package strutils

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"

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

func RemoveDuplicates(strs []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range strs {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
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

//StrToBytes string 转为byte 高效
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//BytesToStr byte 转为string 高效
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsNumber(str string) bool {
	pattern := "\\d+"
	result, _ := regexp.MatchString(pattern, str)
	return result
}

//Diff Creates a slice of slice values not included in the other given slice.
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

func Unique(ss []string) (result []string) {
	smap := make(map[string]bool)
	for _, s := range ss {
		smap[s] = true
	}
	for s := range smap {
		result = append(result, s)
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
