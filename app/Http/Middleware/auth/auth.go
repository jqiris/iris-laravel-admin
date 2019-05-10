package auth

import (
	"github.com/jqiris/iris-laravel-admin/app"
	"github.com/gin-gonic/gin"
	"strings"
)

//是否管理员
func IsAdmin(ctx *gin.Context) bool{
	uid := app.GetGlobalUid(ctx)
	if uid > 0{
		admin:=app.CacheManager.GetAdmin(uid)
		if admin == nil{
			return false
		}
		return true
	}
	return false
}


func RoleAdmin(ctx *gin.Context) int{
	uid := app.GetGlobalUid(ctx)
	if uid > 0{
		admin:=app.CacheManager.GetAdmin(uid)
		if admin == nil{
			return 0
		}
		return admin.RoleId
	}
	return 0
}

//是否有功能
func FuncOp(ctx *gin.Context, funstr string) map[string]bool{
	uid := app.GetGlobalUid(ctx)
	if uid < 1 {
		return nil
	}
	admin := app.CacheManager.GetAdmin(uid)
	if admin == nil {
		return nil
	}
	if _,ok:=admin.Permissions[funstr];!ok{
		return nil
	}
	return admin.Permissions[funstr];
}

//是否有功能的操作
func CheckFunction(ctx *gin.Context, funstr string, op ...string)bool{
	funop:=FuncOp(ctx, funstr)
	if funop == nil{
		return false
	}
	if len(op) > 0 {
		ops := op[0]
		if _, ok := funop[ops]; !ok{
			return false
		}
		return funop[ops]
	}
	return true
}


func CheckAdmin(ctx *gin.Context){
	isadmin := IsAdmin(ctx)
	path :=ctx.ContentType()
	ctx.ContentType()
	ctx.Params
	if strings.ToLower(ctx.Method()) == "get"{
		if path == "/admin/login" || path == "/admin/logout" || isadmin{
			ctx.Next()
		} else {
			ctx.Redirect("/admin/login")
		}
	} else {
		ctx.Next()
	}

}