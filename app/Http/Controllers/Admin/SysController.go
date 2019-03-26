package Admin

import (
	"encoding/json"
	"github.com/jqiris/iris-laravel-admin/app"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/auth"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/captcha"
	"github.com/jqiris/iris-laravel-admin/app/Http/Middleware/csrf"
	"github.com/jqiris/iris-laravel-admin/app/Models"
	"github.com/jqiris/iris-laravel-admin/app/Models/structure"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris"
	"github.com/schigh/str"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func Msg(ctx iris.Context, message string) {
	ctx.ViewData("message", message)
	ctx.ViewData("title", "系统提示")
	ctx.View("errors/error.html")
}
func roleName(role_id int) string {
	if role_id == 1011 {
		return "销售员"
	} else if role_id == 1012 {
		return "推广员"
	} else {
		return "管理员"
	}
}

type MenuDetail struct {
	AppId     int
	AppName   string
	AppEname  string
	AppImg    string
	AppOrder  int
	AppStatus int
	Children  map[int]map[string]Models.ModelValue
}

func Main(ctx iris.Context){
	ctx.ViewData("title", "管理主页")
	ctx.View("admin/main.html")
}

func Index(ctx iris.Context) {
	role_id := auth.RoleAdmin(ctx)
	if role_id < 1 {
		ctx.Redirect("admin/login")
		return
	}
	title := roleName(role_id) + "管理后台"
	ctx.ViewData("title", title)
	uid := app.GetGlobalUid(ctx)
	user := app.CacheManager.GetUser(uid)
	if user.Uid < 1 || user.GroupId == 13 {
		app.RemoveCookie(ctx, "uid")
		ctx.Redirect("/admin/login")
		return

	}

	uname := user.Nickname
	if user.Nickname == "" {
		uname = user.Username
	}
	functions := Models.UserFunction(role_id)
	var menus map[int]MenuDetail
	menus = make(map[int]MenuDetail)
	for _, funcs := range functions {
		appid := funcs["app_id"].Intval
		funcid := funcs["func_id"].Intval
		funtmp := make(map[string]Models.ModelValue)
		funtmp["title"] = funcs["func_name"]
		funtmp["code"] =funcs["func_ename"]
		funtmp["url"] = funcs["func_url"]
		funtmp["img"] = funcs["func_img"]
		if _, ok := menus[appid]; !ok {
			menus[appid] = MenuDetail{
				AppName:  funcs["app_name"].Strval,
				AppEname: funcs["app_ename"].Strval,
				AppImg:   funcs["app_img"].Strval,
				Children: make(map[int]map[string]Models.ModelValue),
			}
		}
		menus[appid].Children[funcid] = funtmp

	}
	showHome := true
	if role_id == 1011 || role_id == 1012 {
		showHome = false
	}
	ctx.ViewData("showHome", showHome)
	ctx.ViewData("admin_menus", menus)
	ctx.ViewData("uid", uid)
	ctx.ViewData("uname", uname)
	ctx.ViewData("rid", role_id)
	ctx.View("admin/index.html")
}

func LoginView(ctx iris.Context) {
	if auth.IsAdmin(ctx) {
		ctx.Redirect("/admin/index")
	}
	ctx.ViewData("logins", 0)
	ctx.ViewData("title", "管理员登陆")
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/login.html")
}

func LoginPost(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	capc := ctx.FormValue("captcha")
	idkey := ctx.FormValue("captchaImg")
	t := app.StringToInt(ctx.FormValue("t"))
	key := "logins_" + username
	logins := app.CacheManager.GetInt(key)
	if username == "" || password == "" {
		ctx.ViewData("msg", "请重新登陆")
		ctx.ViewData("logins", logins)
		ctx.View("admin/login.html")
		return
	}
	title := "管理员登录"
	if t == 1 {
		title = "销售员登录"
	} else if t == 2 {
		title = "推广员登录"
	}

	if username != "" && password != "" {
		if logins > 2 {
			if captcha.Verify(idkey, capc) {
				msg := CheckLoginPost(ctx, username, password)
				if strings.Contains(msg, "/") {
					ctx.Redirect(msg)
					return
				}
				ctx.ViewData("msg", msg)
			} else {
				ctx.ViewData("msg", "登陆失败，验证码错误！")
				ctx.ViewData("password", password)
			}
		} else {
			msg := CheckLoginPost(ctx, username, password)
			if strings.Contains(msg, "/") {
				ctx.Redirect(msg)
				return
			}
			ctx.ViewData("msg", msg)
		}
	} else {
		ctx.ViewData("msg", "用户名或者密码错误！")
	}
	ctx.ViewData("title", title)
	ctx.ViewData("t", t)
	ctx.ViewData("logins", logins)
	ctx.ViewData("username", username)
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/login.html")
}

//用户登陆验证
func GetPassport(ctx iris.Context, username, password string) int {
	if !app.CheckUsername(username) {
		return -4
	}
	key := "logins_" + username
	logins := app.CacheManager.GetInt(key)
	if logins > 10 {
		return -5
	}
	app.CacheManager.Set(key, logins+1, app.Hour)
	return checkLogin(username, password)
}

func checkLogin(username, password string) int {
	var user *structure.YlyMember
	uid := app.StringToInt(username)
	if uid > 0 {
		user = Models.GetUserByID(uid)
	} else {
		user = Models.GetUserByName(username)
	}
	if user == nil {
		return -1
	}
	if user.Username == "" {
		return -1
	} else if user.Password != app.Password(password, user.UserSalt) {
		return -2
	} else if user.GroupId == 13 {
		return -3
	}
	app.CacheManager.Del("logins_" + username)
	return user.Uid
}

func CheckLoginPost(ctx iris.Context, username, password string) string {
	ret := GetPassport(ctx, username, password)
	if ret < 1000 {
		if ret == -1 {
			return "对不起，你使用的用户名尚未注册！"
		} else if ret == -2 {
			return "对不起，密码错误，请重新输入。"
		} else if ret == -3 {
			return "账号已经冻结，请联系客服"
		} else if ret == -4 {
			return "您的游戏账户名称异常请联系客服"
		} else if ret == -5 {
			return "登陆次数过多，请联系客服"
		} else {
			return "登陸失敗，請確認用戶名或者密码正確！"
		}
	} else {
		uid := ret
		app.SetClobalUid(ctx, uid)
		user := app.CacheManager.GetUser(uid)
		user.LoginIp = ctx.RemoteAddr()
		updateLoginData(user, 3)
		return "/admin/index"
	}
}

func updateLoginData(user *structure.YlyMember, utype int) {
	if user.Uid < 1 {
		return
	}
	user.LoginTimes = user.LoginTimes + 1
	user.LoginDate = app.Int64Toint(time.Now().Unix())
	Models.UpdateUser(user.Uid, user)
	app.CacheManager.RemoveUser(user.Uid)
	if utype != 1 {
		Models.LogUserLogin(&structure.LogUserLogin{
			Uid:   user.Uid,
			Utype: utype,
			Ldate: user.LoginDate,
			Ip:    user.LoginIp,
		})
	}
}

func LogOut(ctx iris.Context) {
	uid := app.GetGlobalUid(ctx)
	if uid > 0 {
		//app.Sess.Start(ctx).Clear()
		app.RemoveCookie(ctx, "uid")
	}
	ctx.Redirect("/admin/login")
}

func RoleList(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "RoleManage") {
		Msg(ctx, "无操作權限!")
		return
	}
	roles := Models.RolesAll()
	ctx.ViewData("roles", roles)
	ctx.ViewData("user_right", auth.FuncOp(ctx, "RoleManage"))
	ctx.ViewData("title", "角色管理")
	ctx.View("admin/sys/role_list.html")
}

type AppDetail struct {
	AppId   int
	AppName string
	AppImg  string
	AppFunc []map[string]Models.ModelValue
}

func RoleOp(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "RoleManage") {
		Msg(ctx, "无操作權限!")
		return
	}
	role_id, err := ctx.URLParamInt("role_id")
	if err != nil {
		role_id = 0
	}
	println("role_id", role_id)
	title := "添加角色"
	if role_id > 0 {
		title = "修改角色"
	}
	role := Models.Role(role_id)
	funcs := make(map[int]AppDetail)
	apps := Models.Apps2()

	for _, app := range apps {
		appid := app["app_id"].Intval
		appname := app["app_name"].Strval
		appimg := app["app_img"].Strval
		funcs[appid] = AppDetail{
			AppId:   appid,
			AppName: appname,
			AppImg:  appimg,
			AppFunc: Models.Funcs(appid, role_id),
		}
	}
	ctx.ViewData("role_id", role_id)
	ctx.ViewData("role", role)
	ctx.ViewData("user_right", auth.FuncOp(ctx, "RoleManage"))
	ctx.ViewData("title", title)
	ctx.ViewData("apps", apps)
	ctx.ViewData("funcs", funcs)
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/role_op.html")
}

func RolePost(ctx iris.Context) {
	role_id := ctx.URLParamIntDefault("role_id", 0)
	txtName := ctx.FormValue("txtName")
	txtEName := ctx.FormValue("txtEName")
	txtFuncs := ctx.FormValue("txtFuncsName")
	txtFuncsID := strings.Replace(ctx.FormValue("txtFuncsID"), " ", "", -1)
	cboStatus := app.StringToInt(ctx.FormValue("cboStatus"))

	role := &structure.SysRole{
		RoleName:      txtName,
		RoleEname:     txtEName,
		RoleFuncnames: txtFuncs,
		RoleFuncids:   txtFuncsID,
		Status:        cboStatus,
	}
	result := false
	if role_id < 1 {
		if !auth.CheckFunction(ctx, "RoleManage", "add") {
			Msg(ctx, "无操作权限！")
			return
		}
		role_id = Models.InsertRole(role)
		if role_id > 0 {
			result = true
		}
	} else {
		if !auth.CheckFunction(ctx, "RoleManage", "edit") {
			Msg(ctx, "无操作权限！")
			return
		}
		result = Models.UpdateRole(role_id, role)
	}
	if result && len(txtFuncsID) > 1 {
		Models.DeleteRoleFuncs(role_id)
		funcs := strings.Split(txtFuncsID, ";")
		for _, func_name := range funcs {
			funcid_list := strings.Split(func_name, "-")
			if len(funcid_list) != 5 {
				continue
			}
			func_id := app.StringToInt(funcid_list[4])
			if func_id < 1 {
				continue
			}
			rolefuncs := &structure.SysRoleFunction{
				RoleId: role_id,
				FuncId: func_id,
				FuncOp: func_name,
			}
			Models.InsertRoleFuncs(rolefuncs)
		}
	}
	uid := app.GetGlobalUid(ctx)
	app.CacheManager.RemoveAdmin(uid)
	ctx.Redirect("/admin/sys/role_list")
}

func RoleDelete(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "RoleManage", "delete") {
		Msg(ctx, "无操作权限！")
		return
	}
	role_id := ctx.URLParamIntDefault("role_id", 0)
	if role_id < 1 {
		Msg(ctx, "参数错误！")
		return
	}
	Models.DeleteRole(role_id)
	uid := app.GetGlobalUid(ctx)
	app.CacheManager.RemoveAdmin(uid)
	ctx.Redirect("/admin/sys/role_list")
}

func AdminList(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "AdminManage") {
		Msg(ctx, "无操作权限！")
		return
	}
	ctx.ViewData("title", "管理员管理")
	admins := Models.Admins()
	ctx.ViewData("admins", admins)
	ctx.ViewData("user_right", auth.FuncOp(ctx, "AdminManage"))
	ctx.View("admin/sys/admin_list.html")
}

func AdminOp(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "AdminManage") {
		Msg(ctx, "无操作权限！")
		return
	}
	user_id := ctx.URLParamIntDefault("user_id", 0)
	title := "添加管理员"
	if user_id > 0 {
		title = "修改管理员"
	}
	admin := Models.Admin(user_id)
	roles := Models.Roles()
	ctx.ViewData("title", title)
	ctx.ViewData("user_id", user_id)
	ctx.ViewData("admin", admin)
	ctx.ViewData("roles", roles)
	ctx.ViewData("user_right", auth.FuncOp(ctx, "AdminManage"))
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/admin_op.html")
}

func hasRoleText(ops []string, key int, text string) string {
	res := "×"
	if len(ops[key]) > 0 {
		res = "√"
	}
	return res + text
}
func AdminCheckFunc(ctx iris.Context) {
	role_id := ctx.URLParamIntDefault("role_id", 0)
	if !auth.IsAdmin(ctx) || role_id < 1 {
		ctx.JSONP("[]")
		return
	}
	functions := Models.UserFunction(role_id)
	var menus map[int]MenuDetail
	menus = make(map[int]MenuDetail)
	for _, funcs := range functions {
		appid := funcs["app_id"].Intval
		funcid := funcs["func_id"].Intval
		funtmp := make(map[string]Models.ModelValue)
		op := funcs["func_op"].Strval
		ops := strings.Split(op, "-")
		name := funcs["func_name"].Strval
		if len(ops) == 5 {
			opsn := []string{
				hasRoleText(ops, 0, "浏览"),
				hasRoleText(ops, 1, "添加"),
				hasRoleText(ops, 2, "修改"),
				hasRoleText(ops, 3, "删除"),
			}
			op = strings.Join(opsn, "  ")
			name = str.Pad(name, " ", 20, str.PadRight)
			name = name + op
		}
		funtmp["func_id"] = funcs["func_id"]
		funtmp["func_name"] = Models.ModelValue{Strval:name}
		funtmp["func_ename"] =funcs["func_ename"]
		funtmp["func_url"] = funcs["func_url"]
		funtmp["func_img"] = funcs["func_img"]
		if _, ok := menus[appid]; !ok {
			menus[appid] = MenuDetail{
				AppId:    funcs["app_id"].Intval,
				AppName:  funcs["app_name"].Strval,
				AppEname: funcs["app_ename"].Strval,
				AppImg:   funcs["app_img"].Strval,
				Children: make(map[int]map[string]Models.ModelValue),
			}
		}
		menus[appid].Children[funcid] = funtmp

	}
	jsonstr := ""
	tree := make(map[string]interface{})
	for _, apps := range menus {
		tree["id"] = apps.AppId
		tree["name"] = apps.AppName
		tree["pId"] = ""
		tree["open"] = true
		tree["icon"] = config.CDNRESOURCE + "images/ico/" + apps.AppImg
		if tmp, err := json.Marshal(tree); err == nil {
			jsonstr = jsonstr + string(tmp) + ","
		}
		for _, funcs := range apps.Children {
			tree["id"] = "func-" + funcs["func_id"].Strval
			tree["name"] = funcs["func_name"].Strval
			tree["pId"] = apps.AppId
			tree["icon"] = config.CDNRESOURCE + "images/ico/" + funcs["func_img"].Strval
			if tmp, err := json.Marshal(tree); err == nil {
				jsonstr = jsonstr + string(tmp) + ","
			}
		}

	}
	if len(jsonstr) > 0 {
		jsonstr = strings.TrimRight(jsonstr, ",")
	}

	ctx.JSONP("[" + jsonstr + "]")
}

func AdminPost(ctx iris.Context) {
	user_id := ctx.URLParamIntDefault("user_id", 0)
	txtGame := ctx.FormValue("txtGame")
	role_id := app.StringToInt(ctx.FormValue("cboRole"))
	if role_id < 1 {
		Msg(ctx, "参数错误！")
		return
	}
	result := false
	if user_id < 1 {
		if !auth.CheckFunction(ctx, "AdminManage", "add") {
			Msg(ctx, "无操作权限！")
			return
		}
		user_id = app.StringToInt(ctx.FormValue("txtUserID"))
		if user_id < 1 {
			Msg(ctx, "参数错误！")
			return
		}
		admin := &structure.SysAdminUser{
			UserId:   user_id,
			RoleId:   role_id,
			UserGame: txtGame,
			Status:   1,
		}
		result = Models.InsertAdmin(admin)
	} else {
		if !auth.CheckFunction(ctx, "AdminManage", "edit") {
			Msg(ctx, "无操作权限！")
			return
		}
		result = Models.UpdateAdmin(user_id, &structure.SysAdminUser{
			RoleId:   role_id,
			UserGame: txtGame,
			Status:   1,
		})
	}

	if result {
		game_ids := strings.Split(txtGame, ";")
		Models.DeleteAdminGame(user_id)
		for _, gid := range game_ids {
			igid := app.StringToInt(gid)
			if igid < 1 {
				continue
			}
			admin := &structure.SysAdminUser{
				UserId:   user_id,
				UserGame: txtGame,
			}
			Models.InsertAdmin(admin)
		}
	}
	app.CacheManager.RemoveAdmin(user_id)
	ctx.Redirect("/admin/sys/admin_list")
}

func AdminCheck(ctx iris.Context) {
	data := iris.Map{
		"user_id":    0,
		"user_email": "",
		"user_nick":  "",
	}
	key := ctx.URLParam("key")
	if !auth.IsAdmin(ctx) || len(key) == 0 {
		ctx.JSONP(data)
		return
	}
	user := Models.UserCheck(key)
	if user == nil {
		data["user_id"] = -1
		ctx.JSON(data)
		return
	}
	role_id := Models.AdminRole(user.Uid)
	if role_id > 0 {
		data["user_id"] = -2
		ctx.JSONP(data)
		return
	}
	data["user_id"] = user.Uid
	data["user_email"] = user.Email
	data["user_nick"] = user.Nickname
	ctx.JSONP(data)
}

func FuncList(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "FunctionManage") {
		Msg(ctx, "无操作权限")
		return
	}
	ctx.ViewData("title", "功能管理")
	functions := Models.Functions()
	var menus map[int]MenuDetail
	menus = make(map[int]MenuDetail)
	for _, funcs := range functions {
		appid := funcs["app_id"].Intval
		funcid := funcs["func_id"].Intval
		funtmp := make(map[string]Models.ModelValue)
		funtmp["func_id"] = funcs["func_id"]
		funtmp["func_name"] =funcs["func_name"]
		funtmp["func_ename"] =funcs["func_ename"]
		funtmp["func_url"] = funcs["func_url"]
		funtmp["func_img"] = funcs["func_img"]
		funtmp["func_order"] = funcs["func_order"]
		funtmp["func_status"] = funcs["func_status"]
		if _, ok := menus[appid]; !ok {
			menus[appid] = MenuDetail{
				AppId:     appid,
				AppName:   funcs["app_name"].Strval,
				AppEname:  funcs["app_ename"].Strval,
				AppImg:    funcs["app_img"].Strval,
				AppOrder:  funcs["app_order"].Intval,
				AppStatus: funcs["app_status"].Intval,
				Children:  make(map[int]map[string]Models.ModelValue),
			}
		}
		menus[appid].Children[funcid] = funtmp
	}
	ctx.ViewData("funcs", menus)
	ctx.ViewData("napps", Models.Napps())
	ctx.ViewData("user_right", auth.FuncOp(ctx, "FunctionManage"))
	ctx.View("admin/sys/func_list.html")
}
func FuncOp(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "FunctionManage") {
		Msg(ctx, "无操作权限")
		return
	}
	func_id := ctx.URLParamIntDefault("func_id", 0)
	app_id := ctx.URLParamIntDefault("app_id", 0)
	if func_id > 0 {
		ctx.ViewData("title", "修改功能")
	} else {
		ctx.ViewData("title", "添加功能")
	}
	ctx.ViewData("func_id", func_id)
	ctx.ViewData("app_id", app_id)
	ctx.ViewData("func", Models.Func(func_id))
	ctx.ViewData("apps", Models.Apps())
	ctx.ViewData("user_right", auth.FuncOp(ctx, "FunctionManage"))
	files := get_icon_files(app.PublicPath("/client/images/ico/"))
	ctx.ViewData("images",files)
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/func_op.html")

}

func get_icon_files(dir string) []string {
	var data []string
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		data = append(data,"\""+f.Name()+"\"")
	}
	return data
}

func FuncPost(ctx iris.Context) {
	func_id := ctx.URLParamIntDefault("func_id", 0)
	txtName := ctx.FormValue("txtName")
	txtEName := ctx.FormValue("txtEName")
	txtImg := ctx.FormValue("txtImg")
	txtUrl := ctx.FormValue("txtUrl")
	txtOrder := app.StringToInt(ctx.FormValue("txtOrder"))
	status := ctx.FormValue("cboStatus")
	cboStatus := app.StringToInt(status)
	cboApp := app.StringToInt(ctx.FormValue("cboApp"))
	data := &structure.SysAppFunction{
		FuncId:func_id,
		AppId:     cboApp,
		FuncName:  txtName,
		FuncEname: txtEName,
		FuncImg:   txtImg,
		FuncOrder: txtOrder,
		FuncUrl:   txtUrl,
		Status:    cboStatus,
	}
	if func_id < 1 {
		if !auth.CheckFunction(ctx, "FunctionManage", "add") {
			Msg(ctx, "无操作权限")
			return
		}
		Models.InsertFunc(data)
	} else {
		if !auth.CheckFunction(ctx, "FunctionManage", "edit") {
			Msg(ctx, "无操作权限")
			return
		}
		Models.UpdateFunc(func_id, data)
	}
	uid := app.GetGlobalUid(ctx)
	app.CacheManager.RemoveAdmin(uid)
	ctx.Redirect("/admin/sys/func_list")
}

func FuncDelete(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "FunctionManage", "delete") {
		Msg(ctx, "无操作权限")
		return
	}
	func_id := ctx.URLParamIntDefault("func_id", 0)
	if func_id < 1 {
		Msg(ctx, "参数错误")
		return
	}
	Models.DeleteFunc(func_id)
	uid := app.GetGlobalUid(ctx)
	app.CacheManager.RemoveAdmin(uid)
	ctx.Redirect("/admin/sys/func_list")

}

func UserList(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "UserManage") {
		Msg(ctx, "无操作权限")
		return
	}
	page := ctx.URLParamIntDefault("page", 1)
	key := ctx.URLParam("key")
	begin := ctx.URLParam("begin")
	end := ctx.URLParam("end")
	oc := ctx.URLParamIntDefault("oc", 1)
	os := ctx.URLParam("os")
	if len(os) == 0 || os == "desc" {
		os = "desc"
	} else {
		os = "asc"
	}
	ctx.ViewData("title", "用户管理")
	pageSize := 25
	users := Models.Users(begin, end, key, oc, os, page, pageSize)
	totals := Models.UsersCount(begin, end, key, oc, os)
	ctx.ViewData("users", users)
	ctx.ViewData("count", len(users))
	ctx.ViewData("totals", totals)
	ctx.ViewData("key", key)
	ctx.ViewData("oc", oc)
	ctx.ViewData("os", os)
	ctx.ViewData("begin", begin)
	ctx.ViewData("end", end)
	ctx.ViewData("page", page)
	ctx.ViewData("role_id", auth.RoleAdmin(ctx))
	ctx.ViewData("user_right", auth.FuncOp(ctx, "UserManage"))
	ctx.View("admin/sys/user_list.html")
}
func UserOp(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "UserManage") {
		Msg(ctx, "无操作权限")
		return
	}
	user_id := ctx.URLParamIntDefault("uid", 0)
	if user_id > 0 {
		ctx.ViewData("title", "修改用户")
	} else {
		ctx.ViewData("title", "添加用户")
	}
	user := app.CacheManager.GetUser(user_id)
	userExist := false
	if user != nil {
		userExist = true
	}
	ctx.ViewData("uid", user_id)
	ctx.ViewData("user", user)
	ctx.ViewData("userExist", userExist)
	ctx.ViewData("user_right", auth.FuncOp(ctx, "UserManage"))
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/user_op.html")
}

func UserPost(ctx iris.Context) {
	user_id := ctx.URLParamIntDefault("uid", 0)
	meid := app.GetGlobalUid(ctx)
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	email := ctx.FormValue("email")
	nickname := ctx.FormValue("nickname")
	gender := ctx.FormValue("gender")
	utype := ctx.FormValue("utype")
	tel := ctx.FormValue("tel")
	birthday := ctx.FormValue("birthday")
	user := &structure.YlyMember{
		Username: username,
		Email:    email,
		Nickname: nickname,
		Gender:   gender,
		Utype:    utype,
		Tel:      tel,
		Birthday: birthday,
	}
	if user_id < 1 {
		if !auth.CheckFunction(ctx, "UserManage", "add") {
			Msg(ctx, "无操作权限")
			return
		}
		if password != "" {
			salt := app.GetRandomString(6)
			user.UserSalt = salt
			user.Password = app.Password(password, salt)
		}
		user.RegIp = ctx.RemoteAddr()
		user.RegDate = app.Int64Toint(time.Now().Unix())
		Models.InsertUser(user)
		user_id = user.Uid
	} else {
		if !auth.CheckFunction(ctx, "UserManage", "edit") {
			Msg(ctx, "无操作权限")
			return
		}
		u := app.CacheManager.GetUser(user_id)
		admin := app.CacheManager.GetAdmin(meid)
		if password != "" && (meid == user_id || (admin != nil && admin.RoleId == 1006)) {
			user.Password = app.Password(password, u.UserSalt)
		}
		Models.UpdateUser(user_id, user)
	}
	app.CacheManager.RemoveUser(user_id)
	if user_id > 0 {
		ctx.ViewData("title", "修改用户")
		ctx.ViewData("userExist", true)
	} else {
		ctx.ViewData("title", "添加用户")
		ctx.ViewData("userExist", false)
	}
	ctx.ViewData("uid", user_id)
	ctx.ViewData("user", app.CacheManager.GetUser(user_id))
	ctx.ViewData("user_right", auth.FuncOp(ctx, "UserManage"))
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/user_op.html")
}

func UserDelete(ctx iris.Context) {
	if !auth.CheckFunction(ctx, "UserManage", "delete") {
		Msg(ctx, "无操作权限")
		return
	}
	user_id := ctx.URLParamIntDefault("uid", 0)
	if user_id < 1 {
		Msg(ctx, "参数错误")
		return
	}
	Models.DeleteUser(user_id)
	Models.DeleteAdmin(user_id)
	app.CacheManager.RemoveUser(user_id)
	app.CacheManager.RemoveAdmin(user_id)
	ctx.Redirect("/admin/sys/user_list")
}

func EditPassword(ctx iris.Context) {
	uid := app.GetGlobalUid(ctx)
	ctx.ViewData("title", "修改密码")
	ctx.ViewData("uid", uid)
	ctx.View("admin/sys/editpassword.html")
}

func EditpasswordPost(ctx iris.Context) {
	uid := app.GetGlobalUid(ctx)
	if !auth.IsAdmin(ctx) || uid < 1 {
		Msg(ctx, "参数错误")
		return
	}
	o_password := ctx.FormValue("o_password")
	n_password := ctx.FormValue("n_password1")
	passport := GetPassport(ctx, strconv.Itoa(uid), o_password)
	rel := 0
	msg := ""
	if passport < 1000 {
		if passport == -1 {
			msg = "对不起，你使用的用户名尚未注册！"
		} else if passport == -2 {
			msg = "对不起，密码错误，请重新输入。"
		} else if passport == -3 {
			msg = "账号已经冻结客服"
		} else if passport == -4 {
			msg = "您的游戏账户名称异常请联系客服"
		} else if passport == -5 {
			msg = "密码错误次数过多，请联系客服"
		}
	} else {
		salt := app.GetRandomString(6)
		Models.UpdateUser(uid, &structure.YlyMember{
			Password: app.Password(n_password, salt),
			UserSalt: salt,
			Ps:       "",
		})
		app.CacheManager.RemoveUser(uid)
		rel = 1
		msg = "密码修改成功！"
	}
	data := &iris.Map{
		"rel": rel,
		"msg": msg,
	}
	ctx.JSONP(data)
}

func AppOp(ctx iris.Context){
	if !auth.CheckFunction(ctx, "FunctionManage"){
		Msg(ctx,"无操作权限")
		return
	}

	app_id := ctx.URLParamIntDefault("app_id", 0)
	if app_id >0{
		ctx.ViewData("title", "修改应用")
	} else {
		ctx.ViewData("title", "添加应用")
	}
	ctx.ViewData("app_id", app_id)
	ctx.ViewData("app", Models.App(app_id))
	ctx.ViewData("user_right",auth.FuncOp(ctx,"FunctionManage"))
	images := get_icon_files(app.PublicPath("/client/images/ico/"))
	ctx.ViewData("images",images)
	ctx.ViewData(csrf.TemplateTag, csrf.TemplateField(ctx))
	ctx.View("admin/sys/app_op.html")
}

func AppPost(ctx iris.Context){
	app_id := ctx.URLParamIntDefault("app_id", 0)
	data := &structure.SysApp{
		AppName:ctx.FormValue("txtName"),
		AppEname:ctx.FormValue("txtEname"),
		AppImg:ctx.FormValue("txtImg"),
		AppOrder:app.StringToInt(ctx.FormValue("txtOrder")),
		Status:app.StringToInt(ctx.FormValue("cboStatus")),
	}
	if app_id < 1{
		if !auth.CheckFunction(ctx, "FunctionManage","add"){
			Msg(ctx,"无操作权限")
			return
		}
		Models.InsertApp(data)
	} else {
		if !auth.CheckFunction(ctx, "FunctionManage","edit"){
			Msg(ctx,"无操作权限")
			return
		}
		Models.UpdateApp(app_id, data)
	}
	ctx.Redirect("/admin/sys/func_list")
}

func AppDelete(ctx iris.Context){
	if !auth.CheckFunction(ctx, "FunctionManage","delete"){
		Msg(ctx,"无操作权限")
		return
	}

	app_id := ctx.URLParamIntDefault("app_id", 0)
	if app_id < 1{
		Msg(ctx, "参数错误")
		return
	}

	Models.DeleteApp(app_id)
	Models.DeleteFuncs(app_id)
	ctx.Redirect("/admin/sys/func_list")
}