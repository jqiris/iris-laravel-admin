package Game

import (
	"github.com/jqiris/iris-laravel-admin/app/Models"
	"github.com/jqiris/iris-laravel-admin/app/Models/structure"
	"github.com/kataras/iris"
	"math/rand"
)


func AddYly(ctx iris.Context){
	num := ctx.URLParamIntDefault("num", 1)
	var tests []*structure.YlyPress
	for i:=0; i< num;i++{
		test := &structure.YlyPress{
			Username:createRandomName(),
		}
		tests = append(tests, test)
	}
	Models.AddYly(tests)
	ctx.JSONP(1)
}


func ModifyYly(ctx iris.Context){
	minuid := 1001; maxuid := Models.GetMaxUid()
	suid := minuid + rand.Intn(maxuid-minuid)
	test := &structure.YlyPress{
		Username:createRandomName(),
	}
	Models.ModifyYly(suid, test)
	ctx.JSONP(1)
}


func ReadYly(ctx iris.Context){
	num := ctx.URLParamIntDefault("num", 1)
	Models.ReadYly(num);
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