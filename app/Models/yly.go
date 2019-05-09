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


func ModifyYly(uid int64, press *structure.YlyPress){
	dbr.Where("uid=?", uid).AllCols().Update(press)
}


func ReadYly(limit int) []map[string]ModelValue{
	sql := fmt.Sprintf("select * from yly_press order by rand() limit %d", limit)
	res := DB.Select(sql)
	return MapArrToMV(res)
}

func ClearYly(){
	sql := fmt.Sprintf("delete from yly_press where uid>=%d", 1100000)
	DB.Delete(sql)
}

func GetPressCount() int64 {
	count, err := dbr.Count(&structure.YlyPress{})
	if err != nil {
		println(err.Error())
	}
	return count
}
