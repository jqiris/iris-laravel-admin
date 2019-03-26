package structure

type SysRoleFunction struct {
	RoleId int    `xorm:"not null pk comment('角色编号') INT(11)"`
	FuncId int    `xorm:"not null pk comment('功能编号') INT(11)"`
	FuncOp string `xorm:"comment('功能操作') VARCHAR(100)"`
}
