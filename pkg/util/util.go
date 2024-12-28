package util

import (
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 获取服务器IP
func GetServerIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

// 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return reg.MatchString(email)
}

// 是否是email
func IsMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	reg := regexp.MustCompile(`^1[356789][0-9]{9}$`)
	return reg.MatchString(mobile)
}

// Html过滤
func HTML2str(html string) string {
	src := string(html)
	//替换HTML的空白字符为空格
	re := regexp.MustCompile(`\s`) //ns*r
	src = re.ReplaceAllString(src, " ")
	//将HTML标签全转换成小写
	re, _ = regexp.Compile(`\\<[\\S\\s]+?\\>`)
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile(`\\<style[\\S\\s]+?\\</style\\>`)
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile(`\\<script[\\S\\s]+?\\</script\\>`)
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码,并换成换行符
	re, _ = regexp.Compile(`\\<[\\S\\s]+?\\>`)
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile(`\\s{2,}`)
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// 解析布尔
func ParseBool(s string) (bool, error) {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "on", "yes", "ok":
		return true, nil
	case "", "0", "f", "F", "false", "FALSE", "False", "off", "no":
		return false, nil
	}

	// strconv.NumError mimicing exactly the strconv.ParseBool(..) error and type
	// to ensure compatibility with std library and beyond.
	return false, &strconv.NumError{Func: "ParseBool", Num: s, Err: strconv.ErrSyntax}
}

// struct转map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println(t.NumField())
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 首字母大写
func FirstToUpper(str string) string {
	var (
		s string
		a []string
	)

	aStr := strings.Split(str, " ")
	for _, v := range aStr {
		tmp := strings.ToUpper(v[:1])
		tmp += v[1:]
		a = append(a, tmp)
	}
	s = strings.Join(a, " ")

	return s
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 返回一个固定长度的随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// 获取传入的时间所在年份的第一天，即某年第一天的0点
// 如传入time.Now(), 返回当前年份的第一天0点时间
func FirstDayYear(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return ZeroTime(d)
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点
// 如传入time.Now(), 返回当前月份的第一天0点时间
func FirstDayMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return ZeroTime(d)
}

// 获取传入的时间所在年份的最后一天，即某年最后一天的23点59分59秒
// 如传入time.Now(), 返回当前年份的最后一天23:59:59时间
func LastDayYear(d time.Time) time.Time {
	d = FirstDayYear(d).AddDate(0, 12, -1)
	return TwentyThreeTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的23点59分59秒
// 如传入time.Now(), 返回当前月份的最后一天0点时间。
func LastDayMonth(d time.Time) time.Time {
	d = FirstDayMonth(d).AddDate(0, 1, -1)
	return TwentyThreeTime(d)
}

// 获取某一天的0点时间
func ZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取某一天的23:59:59时间
func TwentyThreeTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}
