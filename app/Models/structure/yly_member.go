package structure

type YlyMember struct {
	Uid        int    `xorm:"not null pk autoincr comment('用户编号') INT(11)"`
	Pid        string `xorm:"not null default '' comment('平台id') CHAR(100)"`
	Username   string `xorm:"not null default '' comment('用户名') unique VARCHAR(64)"`
	Password   string `xorm:"not null default '' comment('密码') CHAR(32)"`
	Tel        string `xorm:"not null default '' comment('手机') VARCHAR(20)"`
	Email      string `xorm:"not null default '' comment('邮箱') CHAR(32)"`
	Birthday   string `xorm:"not null default '' comment('生日') VARCHAR(50)"`
	RegIp      string `xorm:"not null default '' comment('注册ip') CHAR(15)"`
	RegDate    int    `xorm:"not null default 0 comment('注册时间') index INT(11)"`
	Gender     string `xorm:"not null default 'm' comment('性别') ENUM('f','m')"`
	Utype      string `xorm:"not null default '' comment('用户类型') CHAR(10)"`
	Nickname   string `xorm:"not null default '' comment('昵称') CHAR(20)"`
	GroupId    int    `xorm:"not null default 5 comment('管理级别') TINYINT(3)"`
	Locale     string `xorm:"not null comment('语言') VARCHAR(20)"`
	Avatar     string `xorm:"not null comment('头像') VARCHAR(255)"`
	Upuid      int    `xorm:"not null default 0 comment('上层uid，是哪个用户推荐过来的') INT(11)"`
	Ad         string `xorm:"not null default '' comment('用户来源，哪个广告') index CHAR(100)"`
	LoginIp    string `xorm:"not null default '' comment('登录ip') CHAR(15)"`
	LoginTimes int    `xorm:"not null default 0 comment('登录次数') INT(11)"`
	LoginDate  int    `xorm:"not null default 0 comment('最后一次登陆时间') index INT(10)"`
	Remark     string `xorm:"not null default '' comment('封号备注') VARCHAR(255)"`
	Sign       string `xorm:"not null default '' comment('签名') VARCHAR(255)"`
	Regcity    string `xorm:"not null default '' comment('城市') VARCHAR(45)"`
	Regarea    string `xorm:"not null default '' comment('区域') VARCHAR(100)"`
	NewUser    int    `xorm:"not null default 1 comment('新用户') TINYINT(1)"`
	Ps         string `xorm:"comment('pass') VARCHAR(100)"`
	UserSalt   string `xorm:"not null default '' comment('salt') CHAR(6)"`
	Appid      int    `xorm:"not null default 0 comment('appid') TINYINT(4)"`
}
