package structure

type YlyPress struct {
	Uid      int    `xorm:"not null pk autoincr comment('用户编号') INT(11)"`
	Username string `xorm:"not null default '' comment('用户名') VARCHAR(64)"`
}
