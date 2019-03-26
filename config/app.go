package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initConst()
	initCache()
	initDatabase()
	initView()
	initCaptcha()
}

func getRootPath() string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error loading project rootPath")
	}
	return rootPath
}

func getEnvValueWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func GetStoragePath(path string) string {
	return filepath.Join(STORAGEPATH, path)
}

func GetResourcePath(path string) string {
	return filepath.Join(RESOURCEPATH, path)
}

var (
	GAME_TITLE,
	ROOTURL,
	BASEURI,
	RESOURCE,
	ROOTPATH,
	RESOURCEPATH,
	STORAGEPATH,
	PUBLICPATH,
	CDNSERVER,
	CDNRESOURCE,
	CDNVERSION,
	AdminAllowIps,
	AdminAllowUsers,
	AdminAllowFakeUsers,
	SystemMaintenance,
	APP_ENV,
	APP_DEBUG,
	APP_ADDR ,
	CSRF_KEY string
)

func initConst() {
	//常量定义
	GAME_TITLE = getEnvValueWithDefault("GAME_TITLE", "IRISLARAVEL")
	ROOTURL = os.Getenv("ROOTURL")
	BASEURI = ROOTURL + "/"
	RESOURCE = ROOTURL + "/client/"
	ROOTPATH = getRootPath()
	RESOURCEPATH = ROOTPATH + "/resources"
	STORAGEPATH = ROOTPATH + "/storage"
	PUBLICPATH = ROOTPATH + "/public"
	CDNSERVER = os.Getenv("CDNSERVER")
	CDNRESOURCE = getEnvValueWithDefault("CDNRESOURCE", CDNSERVER+"/public/client/")
	CDNVERSION = os.Getenv("CDNVERSION")
	AdminAllowIps = os.Getenv("AdminAllowIps")
	AdminAllowUsers = os.Getenv("AdminAllowUsers")
	AdminAllowFakeUsers = os.Getenv("AdminAllowFakeUsers")
	SystemMaintenance = getEnvValueWithDefault("SystemMaintenance", "false")
	APP_ENV = getEnvValueWithDefault("APP_ENV", "production")
	APP_DEBUG = getEnvValueWithDefault("APP_DEBUG", "false")
	APP_ADDR = getEnvValueWithDefault("APP_ADDR", "8080")
	CSRF_KEY = "9AB0F421E53A477C084477AEA06096F5"
}
