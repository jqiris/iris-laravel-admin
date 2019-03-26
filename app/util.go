package app

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"crypto/md5"
)

var (
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey  = []byte("one-big-and-secret-fish-key-here")
	blockKey = []byte("god-secret-of-characters-pig-too")
	sc       = securecookie.New(hashKey, blockKey)
)

func GetCookie(ctx iris.Context, name string, options ...iris.CookieOption) string{
	options = append(options, iris.CookieDecode(sc.Decode))
	return ctx.GetCookie(name, options...)
}

func GetCookieInt(ctx iris.Context, name string, options ...iris.CookieOption) int{
	val := GetCookie(ctx, name, options...)
	return StringToInt(val)
}

func SetCookie(ctx iris.Context, name,value string, options ...iris.CookieOption) {
	options = append(options, iris.CookieEncode(sc.Encode))
	ctx.SetCookieKV(name, value, options...)
}

func RemoveCookie(ctx iris.Context,name string, options ...iris.CookieOption ){
	ctx.RemoveCookie(name, options...)
}

func StringToInt(valstr string) int{
	val, err := strconv.Atoi(valstr)
	if err != nil{
		val = 0
	}
	return val
}

func Int64ToString(valint int64) string{
	return strconv.FormatInt(valint, 10)
}

func Int64Toint(valint int64) int{
	valstr := Int64ToString(valint)
	return StringToInt(valstr)
}

func IntToInt64(valint int) int64{
	s := strconv.Itoa(valint)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil{
		println(err.Error())
		return 0
	}
	return i
}

func GetGlobalUid(ctx iris.Context) int{
	//uid, err := Sess.Start(ctx).GetInt("uid")
	//if err != nil{
	//	println(err.Error())
	//	return 0
	//}
	uid := GetCookieInt(ctx, "uid")
	return uid
}

func SetClobalUid(ctx iris.Context, uid int){
	//Sess.Start(ctx).Set("uid", uid)
	SetCookie(ctx,"uid",strconv.Itoa(uid))
}
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func CheckUsername(username string) bool{
	len := len(username)
	matchUser, err := regexp.MatchString(`^[0-9@A-Z_a-z]*$`, username)
	if err != nil{
		println(err.Error())
	}
	if len > 20 || len < 3 || !matchUser{
		return false
	} else if strings.Contains(username, " "){
		return false
	} else {
		return true
	}
}

func Md5(sign string) string{
	h := md5.New()
	io.WriteString(h, sign)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}

func Password(user_password, user_salt string) string{
	return Md5(Md5(user_password)+user_salt)
}


// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func Addslashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

func PublicPath(relative string) string{
	return config.PUBLICPATH+relative
}

func StoragePath(relative string) string{
	return config.STORAGEPATH+relative
}

func StrToTime(str string, tpl string) int64{
	t, err := time.Parse(tpl, str)
	if err != nil{
		println(err.Error())
		return 0
	}
	return t.Unix()
}

func CheckEmail(email string) bool{
	ret ,err :=regexp.MatchString(`([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)`, email)
	if err !=nil {
		println(err.Error())
		return false
	}
	return ret
}