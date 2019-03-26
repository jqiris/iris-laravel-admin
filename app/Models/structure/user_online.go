package structure

type UserOnline struct {
	UserId   int    `xorm:"not null pk comment('在线用户编号') INT(11)"`
	UserUrl  string `xorm:"comment('当前ＵＲＬ') VARCHAR(200)"`
	UrlIp    string `xorm:"VARCHAR(45)"`
	LastTime int    `xorm:"not null default 0 comment('最后一次操作时间') INT(11)"`
}
