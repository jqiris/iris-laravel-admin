package auth

import (
	"github.com/jqiris/iris-laravel-admin/app"
	"github.com/kataras/iris"
	"strings"
)

//是否管理员
func IsAdmin(ctx iris.Context) bool{
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


func RoleAdmin(ctx iris.Context) int{
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
func FuncOp(ctx iris.Context, funstr string) map[string]bool{
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
func CheckFunction(ctx iris.Context, funstr string, op ...string)bool{
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


func CheckAdmin(ctx iris.Context){
	isadmin := IsAdmin(ctx)
	path :=ctx.Path()
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