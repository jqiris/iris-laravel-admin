package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/iris-laravel-admin/app"
	"github.com/jqiris/iris-laravel-admin/app/Http"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris/view"
	"gitlab.com/go-box/pongo2gin"
	"html/template"
	"net/http"
	"time"
)
func smys(cond bool, suc , fail string) string{
	if  cond {
		return suc
	} else {
		return fail
	}
}
func getOrderStatus(otype, oc int, os string) string{
	if oc == otype && os == "desc" {
		return "asc"
	} else {
		return "desc"
	}
}

func getAddressByIP(ip string) string{
	return app.GetAddressByIP(ip)
}

func testvalue(in *view.Value, param *view.Value)(out *view.Value, err *view.Error) {
	ins := in.GetValue().String()
	params := param.GetValue().String()
	println(ins, params)
	return nil,nil
}

func timeformat(in *view.Value, param *view.Value) (out *view.Value, err *view.Error) {
	t := in.GetValue().Integer()
	f := param.GetValue().String()
	tt := app.IntToInt64(t)
	return view.AsValue(time.Unix(tt, 0).Format(f)), nil
}

func registerView(api *gin.Engine) {
	api.HTMLRender = pongo2gin.Default()
	funcmap := template.FuncMap{
		"smys":smys,
		"getOrderStatus":getOrderStatus,
		"getAddressByIP":getAddressByIP,
		"testvalue":testvalue,
		"timeformat":timeformat,
	}
	api.SetFuncMap(funcmap)
}

func main() {
	api := gin.Default()
	api.StaticFS("/public", http.Dir(config.PUBLICPATH))
	registerView(api)
	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	Http.Router(api)
	addr := "0.0.0.0:" + config.APP_ADDR
	api.Run(addr)
}
