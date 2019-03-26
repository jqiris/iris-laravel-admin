package structure

type SysApp struct {
	AppId       int    `xorm:"not null pk autoincr comment('应用编号') INT(10)"`
	AppEname    string `xorm:"comment('应用Code') VARCHAR(50)"`
	AppName     string `xorm:"comment('应用名称') VARCHAR(100)"`
	AppImg      string `xorm:"comment('图片') VARCHAR(200)"`
	AppOrder    int    `xorm:"not null default 0 comment('排序') INT(11)"`
	AppTreeShow int    `xorm:"not null default 1 comment('是否在导航中显示') TINYINT(1)"`
	Status      int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
}
