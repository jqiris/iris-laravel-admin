package Http

import (
	"github.com/gin-gonic/gin"
	"github.com/jqiris/iris-laravel-admin/app/Http/Controllers/Admin"
	"github.com/jqiris/iris-laravel-admin/app/Http/Controllers/Game"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/auth"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/captcha"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/mojocn/base64Captcha"
	"github.com/vanng822/gin-csrf"
	"time"
)

func Global() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Set("isMobile",ctx.IsMobile())
		ctx.Set("title", config.GAME_TITLE)
		ctx.Set("TIMESTAMP", time.Now().Unix())
		ctx.Set("CDNRESOURCE", config.CDNRESOURCE)
		ctx.Set("CDNVERSION", config.CDNVERSION)
		ctx.Set("BASEURI", config.BASEURI)
		ctx.Set("ROOTURL", config.ROOTURL)
		ctx.Set("GAME_TITLE", config.GAME_TITLE)
		ctx.Next()
	}
}

func Router(api *gin.Engine) {
	//api.OnErrorCode(iris.StatusNotFound, notFound)
	//api.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	//api.OnErrorCode(iris.StatusServiceUnavailable, serviceUnavailable)
	api.Use(Global())

	//后台管理
	admin := api.Group("/admin")
	options := csrf.DefaultOptions()
	options.MaxUsage = 10
	options.MaxAge = 15 * 60
	admin.Use(csrf.Csrf(options))

	adminRouter(admin)

	//前台管理
	frontRouter(api)
}


func adminRouter(router *gin.RouterGroup){
	router.Use(auth.CheckAdmin)
	router.GET("/", Admin.Index)
	router.GET("/main", Admin.Main)
	router.GET("/index", Admin.Index)
	router.GET("/login", captcha.CaptchaImg(),Admin.LoginView)
	router.POST("/login", captcha.CaptchaImg(),Admin.LoginPOST)
	router.GET("/logout",Admin.LogOut)
	router.GET("/sys/editpassword", Admin.EditPassword)
	router.POST("/sys/editpassword_POST", Admin.EditpasswordPOST)
	//管理员管理
	router.GET("/sys/admin_list",Admin.AdminList)
	router.GET("/sys/admin_op",Admin.AdminOp)
	router.POST("/sys/admin_POST",Admin.AdminPOST)
	router.GET("/sys/admin_check",Admin.AdminCheck)
	router.GET("/sys/admin_checkFunc",Admin.AdminCheckFunc)
	//角色管理
	router.GET("/sys/role_list",Admin.RoleList)
	router.GET("/sys/role_op",Admin.RoleOp)
	router.POST("/sys/role_POST",Admin.RolePOST)
	router.GET("/sys/role_delete",Admin.RoleDelete)
	//功能管理
	router.GET("/sys/func_list",Admin.FuncList)
	router.GET("/sys/func_op",Admin.FuncOp)
	router.POST("/sys/func_POST",Admin.FuncPOST)
	router.GET("/sys/func_delete",Admin.FuncDelete)
	router.GET("/sys/app_op",Admin.AppOp)
	router.GET("/sys/app_delete",Admin.AppDelete)
	router.POST("/sys/app_POST",Admin.AppPOST)
	//用户管理
	router.GET("/sys/user_list", Admin.UserList)
	router.GET("/sys/user_op", Admin.UserOp)
	router.POST("/sys/user_POST", Admin.UserPOST)
	router.GET("/sys/user_delete", Admin.UserDelete)
}

func frontRouter(api *gin.Engine){
	api.GET("/captcha/{idkey string}", captchaGenerater)
	api.GET("/", func(ctx *gin.Context) {
		ctx.View("game/index.html")
	})

	api.GET("/addyly", Game.AddYly)
	api.GET("/modyly", Game.ModifyYly)
	api.GET("/readyly", Game.ReadYly)
	api.GET("/clearyly", Game.ClearYly)
}
func notFound(ctx *gin.Context) {
	// 当http.status=400 时向客户端渲染模板$views_dir/errors/404.html
	ctx.View("errors/404.html")
}

//当出现错误的时候，再试一次
func internalServerError(ctx *gin.Context) {
	ctx.View("errors/500.html")
}

func serviceUnavailable(ctx *gin.Context) {
	ctx.View("errors/503.html")
}

func captchaGenerater(ctx *gin.Context){
	idkey:= ctx.Params().GET("idkey")
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
