package Models
import (
	"fmt"
	"github.com/jqiris/iris-laravel-admin/app/Models/structure"
	"log"
)

func AddYly( press []*structure.YlyPress){
	_, err:=dbr.Insert(press)
	if err != nil{
		log.Fatal(err.Error())
	}
}


func ModifyYly(uid int, press *structure.YlyPress){
	dbr.Where("uid=?", uid).AllCols().Update(press)
}


func ReadYly(limit int){
	sql := fmt.Sprintf("select * from yly_press order by rand() limit %d", limit)
	DB.Select(sql)
}

func ClearYly(){
	sql := fmt.Sprintf("delete from yly_press where uid>=%d", 1100000)
	DB.Delete(sql)
}

func GetMaxUid() int{
 sql := "select max(uid) uid from yly_press";
 res := DB.SelectOne(sql)
 value := MapFaceToMV(res)
 return value["uid"].Intval
}

func GetPressCount() int64 {
	count, err := dbr.Count(&structure.YlyPress{})
	if err != nil {
		println(err.Error())
	}
	return count
}
