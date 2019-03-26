package structure

type SysAppFunction struct {
	FuncId    int    `xorm:"not null pk autoincr comment('功能编号') INT(10)"`
	AppId     int    `xorm:"not null comment('应用编号') INT(11)"`
	FuncName  string `xorm:"comment('功能名称') VARCHAR(50)"`
	FuncEname string `xorm:"comment('功能代码') VARCHAR(100)"`
	FuncUrl   string `xorm:"comment('地址') VARCHAR(200)"`
	FuncImg   string `xorm:"comment('图标') VARCHAR(200)"`
	FuncOrder int    `xorm:"not null default 0 comment('排序') INT(11)"`
	Status    int    `xorm:"not null default 1 comment('状态') INT(3)"`
}
