package config

var (
	ViewPath    string
	ViewCompild string
)

func initView() {
	ViewPath = GetResourcePath("views")
	ViewCompild = GetStoragePath("framework/views")
}
