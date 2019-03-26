package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jqiris/iris-laravel-admin/config"
	"log"
	"regexp"
	"strconv"
	"time"
)

var (
	DB *dbm
	dbr *xorm.Engine
)

type dbm struct {
	db *xorm.Engine
}

func init() {
	config := config.Connections[config.DefaultConnect]
	args := fmt.Sprintf("charset=%s", config.Charset)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		args,
	)
	if dbx, err := xorm.NewEngine(config.Driver, dsn); err != nil {
		log.Fatal(err)
	} else {
		DB = &dbm{
			db: dbx,
		}
		dbr = DB.db
		dbr.ShowSQL(true)
	}
}

func GetDb() *xorm.Engine{
	return DB.db
}

func (this *dbm) Select(sql string) []map[string]interface{} {
	vals, err := this.db.QueryInterface(sql)
	if err != nil {
		log.Println(err.Error())
		return nil
	} else {
		return vals
	}
}

func (this *dbm) SelectOne(sql string) map[string]interface{} {
	vals, err := this.db.QueryInterface(sql)
	if err != nil {
		log.Println(err.Error())
		return nil
	} else {
		if vals == nil{
			return nil
		}
		return vals[0]
	}
}

func (this *dbm) Insert(sql string) int64 {
	res, err := this.db.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		return 0
	} else {
		insertid, err := res.LastInsertId()
		if err != nil {
			log.Println(err.Error())
			return 0
		} else {
			return insertid
		}
	}
}

func (this *dbm) Update(sql string) int64 {
	res, err := this.db.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		return 0
	} else {
		affected, err := res.RowsAffected()
		if err != nil {
			log.Println(err.Error())
			return 0
		} else {
			return affected
		}
	}
}

func (this *dbm) Delete(sql string) bool {
	res, err := this.db.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		affected, err := res.RowsAffected()
		if err != nil {
			log.Println(err.Error())
			return false
		} else {
			if affected > 0{
				return true
			}
			return false
		}
	}
}

func IfcToString(i interface{}) string {
	switch i.(type) {
	case []byte:
		return string(i.([]byte))
	case string:
		return i.(string)
	}
	return fmt.Sprintf("%v", i)
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
func IfcToInt(i interface{}) int {
	switch i.(type) {
	case []byte:
		n := StringToInt(string(i.([]byte)))
		return n
	case int:
		return i.(int)
	case int64:
		return Int64Toint(i.(int64))
	}
	return 0
}

func IfcToInt64(i interface{}) int64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseInt(string(i.([]byte)), 10, 64)
		return n
	case int:
		return int64(i.(int))
	case int64:
		return i.(int64)
	}
	return 0
}

func IfcToFloat64(i interface{}) float64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseFloat(string(i.([]byte)), 64)
		return n
	case float64:
		return i.(float64)
	case float32:
		return float64(i.(float32))
	}
	return 0
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

type ModelValue struct {
	Intval int
	Strval string
}

func byteIsNumber(b []byte) bool{
	str := string(b)
	ret, err := regexp.MatchString(`^([0-9]|\+|\-)[0-9]*$`, str)
	if err != nil{
		println(err.Error())
		return false
	}
	return ret
}

func FaceToMV(i interface{}) ModelValue{
	val := ModelValue{Intval:0, Strval:""}
	switch i.(type) {
	case int:
		val.Intval = i.(int)
	case string:
		val.Strval = i.(string)
	case []byte:
		if byteIsNumber(i.([]byte)){
			val.Intval = StringToInt(string(i.([]byte)))
		} else {
			val.Strval = string(i.([]byte))
		}
	}
	return val
}

func MapFaceToMV(m map[string]interface{}) map[string]ModelValue{
	val := make(map[string]ModelValue)
	for k,v := range m{
		val[k] = FaceToMV(v)
	}
	return val
}

func MapArrToMV(ms []map[string]interface{}) []map[string]ModelValue{
	len := len(ms)
	var vals []map[string]ModelValue
	for i:=0;i < len; i++{
		vals = append(vals, MapFaceToMV(ms[i]))
	}
	return vals
}



