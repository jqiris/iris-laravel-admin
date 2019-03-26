package structure

type SysRole struct {
	RoleId        int    `xorm:"not null pk autoincr comment('角色') INT(10)"`
	RoleName      string `xorm:"comment('角色名称') VARCHAR(100)"`
	RoleEname     string `xorm:"comment('角色代码') VARCHAR(50)"`
	RoleFuncnames string `xorm:"comment('角色功能') VARCHAR(3000)"`
	RoleFuncids   string `xorm:"comment('角色功能代码') VARCHAR(3000)"`
	Status        int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
}
