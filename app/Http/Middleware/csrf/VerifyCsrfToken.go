package csrf

var except  = map[string]bool{
	"/admin/sys/editpassword_post":true,
}