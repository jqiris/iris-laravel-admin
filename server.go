package main

import (
	"github.com/jqiris/iris-laravel-admin/app"
	"github.com/jqiris/iris-laravel-admin/app/Http"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/view"
	"time"
)
func registerView(api *iris.Application) {
	ve := iris.Django(config.ViewPath, ".html")
	ve.Reload(true)
	ve.AddFunc("smys", func(cond bool, suc , fail string) string {
		if  cond {
			return suc
		} else {
			return fail
		}
	})

	ve.AddFunc("getOrderStatus", func(otype, oc int, os string) string {
		if oc == otype && os == "desc" {
			return "asc"
		} else {
			return "desc"
		}
	})
	ve.AddFunc("getAddressByIP", func(ip string) string{
		return app.GetAddressByIP(ip)
	})
	ve.AddFilter("testvalue", func(in *view.Value, param *view.Value) (out *view.Value, err *view.Error) {
			ins := in.GetValue().String()
			params := param.GetValue().String()
			println(ins, params)
			return nil,nil
	})
	ve.AddFilter("timeformat", func(in *view.Value, param *view.Value) (out *view.Value, err *view.Error) {
		t := in.GetValue().Integer()
		f := param.GetValue().String()
		tt := app.IntToInt64(t)
		return view.AsValue(time.Unix(tt, 0).Format(f)), nil
	})
	api.RegisterView(ve)
}

func main() {
	api := iris.New()
	api.StaticWeb("/public", config.PUBLICPATH)
	registerView(api)
	api.Use(recover.New())
	api.Use(logger.New())
	Http.Router(api)
	addr := "0.0.0.0:" + config.APP_ADDR
	api.Run(iris.Addr(addr))
}
