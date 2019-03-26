package app

import (
	"encoding/json"
	"github.com/jqiris/iris-laravel-admin/app/Models"
	"github.com/jqiris/iris-laravel-admin/app/Models/structure"
	"github.com/jqiris/iris-laravel-admin/config"
	"github.com/kataras/iris/sessions"
	"gopkg.in/redis.v4"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	CacheManager *cacheManager
	Sess         *sessions.Sessions
)

func init() {
	RegisterSession()
	CacheManager = newCacheManager()
	Week = 604800 * time.Second
	Day = 86400 * time.Second
	Hour = 3600 * time.Second
}

func RegisterSession() {
	Sess = sessions.New(sessions.Config{
		Cookie:       "iris_laravel_session",
		Expires:      24 * time.Hour, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	})
}

var (
	Week time.Duration
	Day  time.Duration
	Hour time.Duration
)

type cacheManager struct {
	Config string
	Client *redis.Client
	Prefix string
}

func newCacheManager() *cacheManager {
	cm := &cacheManager{
		Prefix: config.CACHE_PREFIX,
		Config: config.DefaultCache,
	}
	switch cm.Config {
	case "redis":
		rconfig := config.Caches["redis"].Servers[0]
		addr := strings.Join([]string{rconfig.Host, rconfig.Port}, ":")
		cm.Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: rconfig.Password,
			DB:       0,
		})
		_, err := cm.Client.Ping().Result()
		if err != nil {
			log.Fatal("redis", err.Error())
		}
	}
	if cm.Client == nil {
		log.Fatal("can't find the cache client")
	}
	return cm
}

func (cm *cacheManager) Set(key string, value interface{}, expiration time.Duration) bool {
	ckey := cm.Prefix + ":" + key
	err := cm.Client.Set(ckey, value, expiration).Err()
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}
func (cm *cacheManager) Get(key string) string {
	ckey := cm.Prefix + ":" + key
	val, err := cm.Client.Get(ckey).Result()
	if err != nil {
		println(err.Error())
		return ""
	}
	return val
}
func (cm *cacheManager) GetInt(key string) int {
	val := cm.Get(key)
	return StringToInt(val)
}
func (cm *cacheManager) GetSet(key string, value interface{}) string {
	ckey := cm.Prefix + ":" + key
	val, err := cm.Client.GetSet(ckey, value).Result()
	if err != nil {
		println(err.Error())
		return ""
	}
	return val
}

func (cm *cacheManager) Del(keys ...string) bool {
	for i, v := range keys {
		keys[i] = cm.Prefix + ":" + v
	}
	_, err := cm.Client.Del(keys...).Result()
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

func (cm *cacheManager) CacheKey(prefix string, params ...interface{}) string {
	key := prefix
	len := len(params)
	if len > 0 {
		for i := 0; i < len; i++ {
			param := params[i]
			switch t := param.(type) {
			case string:
				key += "_" + t
			case int:
				key += "_" + strconv.Itoa(t)
			}
		}
	}
	return key
}

func (cm *cacheManager) GetAdmin(uid int) *Models.AdminInfo {
	key := cm.CacheKey("getAdmin", uid)
	val := cm.Get(key)
	var info *Models.AdminInfo
	if val == "" {
		info = Models.GetAdmin(uid)
		val, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		cm.Set(key, val, Day)
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		panic(err)
	}
	return info
}

func (cm *cacheManager) RemoveAdmin(uid int) bool {
	key := cm.CacheKey("getAdmin", uid)
	return cm.Del(key)
}

func (cm *cacheManager) GetUser(uid int) *structure.YlyMember {
	key := cm.CacheKey("getUser", uid)
	val := cm.Get(key)
	var info *structure.YlyMember
	if val == "" {
		info = Models.GetUserByID(uid)
		val, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		cm.Set(key, val, Day)
		return info
	}
	err := json.Unmarshal([]byte(val), &info)
	if err != nil {
		panic(err)
	}
	return info
}

func (cm *cacheManager) RemoveUser(uid int) bool {
	key := cm.CacheKey("getUser", uid)
	return cm.Del(key)
}
