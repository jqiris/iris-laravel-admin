package structure

type SysAdminUser struct {
	UserId   int    `xorm:"not null pk comment('用户') INT(11)"`
	RoleId   int    `xorm:"not null pk comment('角色') INT(11)"`
	Status   int    `xorm:"not null default 1 comment('状态') TINYINT(1)"`
	UserGame string `xorm:"comment('游戏') VARCHAR(500)"`
}
