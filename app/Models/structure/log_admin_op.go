package structure

type LogAdminOp struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Uid     int    `xorm:"default 0 comment('用户编号') index INT(11)"`
	Url     string `xorm:"default '' comment('请求ＵＲＬ') VARCHAR(255)"`
	Request string `xorm:"comment('操作内容') TEXT"`
	Ip      string `xorm:"comment('ip') VARCHAR(20)"`
	Ltime   int    `xorm:"default 0 comment('操作时间') index INT(11)"`
}
