package Game

import (
	"github.com/jqiris/iris-laravel-admin/app/Models"
	"github.com/jqiris/iris-laravel-admin/app/Models/structure"
	"github.com/kataras/iris"
	"math/rand"
)

var (
	MinUid  int64= 1001
	MaxUid int64 = 1100000
)

func AddYly(ctx iris.Context){
	var tests []*structure.YlyPress
	for i:=0; i< 100;i++{
		test := &structure.YlyPress{
			Username:createRandomName(),
		}
		tests = append(tests, test)
	}
	Models.AddYly(tests)
	ctx.JSONP(1)
}


func ModifyYly(ctx iris.Context){
	MaxUid = Models.GetPressCount()
	suid := MinUid + rand.Int63n(MaxUid-MinUid+1)
	limit := 100 + suid;
	for i:=suid; i <= limit; i++{
		test := &structure.YlyPress{
			Username:createRandomName(),
		}
		Models.ModifyYly(i, test)
	}
	ctx.JSONP(1)
}


func ReadYly(ctx iris.Context){
	res := Models.ReadYly(200);
	for _, value:=range res{
		ctx.Writef("uid:%d, usename:%s\n", value["uid"].Intval, value["username"].Strval)
	}
	ctx.JSONP(1)
}

func ClearYly(ctx iris.Context){
	Models.ClearYly()
}

func createRandomName() string{
	names := []string{"宋", "茅", "庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁", "杜", "阮", "蓝", "闵", "季", "贾"}
	str := ""
	len := len(names)
	for i := 0; i < 3; i++{
		j := rand.Intn(len)
		str += names[j]
	}
	return str
}