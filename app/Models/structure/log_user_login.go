package structure

type LogUserLogin struct {
	Id     int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Uid    int    `xorm:"not null default 0 comment('用户编号') INT(11)"`
	Online int    `xorm:"not null default 0 comment('状态 0:上线;1:下线') TINYINT(3)"`
	Utype  int    `xorm:"not null default 0 comment('登陆类型 0:fb;1:mobile;2web,3admin,4模拟登录') TINYINT(3)"`
	Ldate  int    `xorm:"not null default 0 comment('时间') INT(11)"`
	Ip     string `xorm:"not null default '' comment('ip') VARCHAR(15)"`
}
