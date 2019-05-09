package Http

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/iris-laravel-admin/app/Http/Controllers/Admin"
	"github.com/jqiris/iris-laravel-admin/app/Http/Controllers/Game"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/auth"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/csrf"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/captcha"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris"
	"github.com/mojocn/base64Captcha"
	"time"
)

func Router(api *gin.Engine) {
	api.OnErrorCode(iris.StatusNotFound, notFound)
	api.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	api.OnErrorCode(iris.StatusServiceUnavailable, serviceUnavailable)
	api.Use(func(ctx iris.Context) {
		ctx.ViewData("isMobile",ctx.IsMobile())
		ctx.ViewData("title", config.GAME_TITLE)
		ctx.ViewData("TIMESTAMP", time.Now().Unix())
		ctx.ViewData("CDNRESOURCE", config.CDNRESOURCE)
		ctx.ViewData("CDNVERSION", config.CDNVERSION)
		ctx.ViewData("BASEURI", config.BASEURI)
		ctx.ViewData("ROOTURL", config.ROOTURL)
		ctx.ViewData("GAME_TITLE", config.GAME_TITLE)
		ctx.Next()
	})

	//后台管理
	protect := csrf.Protect([]byte(config.CSRF_KEY), csrf.Secure(false))
	admin := api.Party("/admin", protect)
	adminRouter(admin)

	//前台管理
	frontRouter(api)
}


func adminRouter(router iris.Party){
	router.Use(auth.CheckAdmin)
	router.Get("/", Admin.Index)
	router.Get("/main", Admin.Main)
	router.Get("/index", Admin.Index)
	router.Get("/login", captcha.CaptchaImg(),Admin.LoginView)
	router.Post("/login", captcha.CaptchaImg(),Admin.LoginPost)
	router.Get("/logout",Admin.LogOut)
	router.Get("/sys/editpassword", Admin.EditPassword)
	router.Post("/sys/editpassword_post", Admin.EditpasswordPost)
	//管理员管理
	router.Get("/sys/admin_list",Admin.AdminList)
	router.Get("/sys/admin_op",Admin.AdminOp)
	router.Post("/sys/admin_post",Admin.AdminPost)
	router.Get("/sys/admin_check",Admin.AdminCheck)
	router.Get("/sys/admin_checkFunc",Admin.AdminCheckFunc)
	//角色管理
	router.Get("/sys/role_list",Admin.RoleList)
	router.Get("/sys/role_op",Admin.RoleOp)
	router.Post("/sys/role_post",Admin.RolePost)
	router.Get("/sys/role_delete",Admin.RoleDelete)
	//功能管理
	router.Get("/sys/func_list",Admin.FuncList)
	router.Get("/sys/func_op",Admin.FuncOp)
	router.Post("/sys/func_post",Admin.FuncPost)
	router.Get("/sys/func_delete",Admin.FuncDelete)
	router.Get("/sys/app_op",Admin.AppOp)
	router.Get("/sys/app_delete",Admin.AppDelete)
	router.Post("/sys/app_post",Admin.AppPost)
	//用户管理
	router.Get("/sys/user_list", Admin.UserList)
	router.Get("/sys/user_op", Admin.UserOp)
	router.Post("/sys/user_post", Admin.UserPost)
	router.Get("/sys/user_delete", Admin.UserDelete)
}

func frontRouter(api *iris.Application){
	api.Get("/captcha/{idkey string}", captchaGenerater)
	api.Get("/", func(ctx iris.Context) {
		ctx.View("game/index.html")
	})

	api.Get("/addyly", Game.AddYly)
	api.Get("/modyly", Game.ModifyYly)
	api.Get("/readyly", Game.ReadYly)
	api.Get("/clearyly", Game.ClearYly)
}
func notFound(ctx iris.Context) {
	// 当http.status=400 时向客户端渲染模板$views_dir/errors/404.html
	ctx.View("errors/404.html")
}

//当出现错误的时候，再试一次
func internalServerError(ctx iris.Context) {
	ctx.View("errors/500.html")
}

func serviceUnavailable(ctx iris.Context) {
	ctx.View("errors/503.html")
}

func captchaGenerater(ctx iris.Context){
	idkey:= ctx.Params().Get("idkey")
	if idkey == ""{
		ctx.JSON(map[string]interface{}{"ret":1, "msg":"no idkey"})
		return
	}
	configD := base64Captcha.ConfigDigit{
		Height:     config.DefaultCaptcha.Height,
		Width:      config.DefaultCaptcha.Width,
		MaxSkew:    config.DefaultCaptcha.MaxSkew,
		DotCount:   config.DefaultCaptcha.DotCount,
		CaptchaLen: config.DefaultCaptcha.CaptchaLen,
	}
	_, capD := base64Captcha.GenerateCaptcha(idkey, configD)
	//write to base64 string.
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	ctx.JSON(map[string]interface{}{"ret":0, "img":base64stringD})
}
