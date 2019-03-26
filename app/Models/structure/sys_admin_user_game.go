package structure

type SysAdminUserGame struct {
	UserId int `xorm:"not null pk INT(11)"`
	GameId int `xorm:"not null pk INT(11)"`
}
