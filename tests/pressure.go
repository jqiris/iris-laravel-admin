package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	BeginTime time.Time // 开始时间
	SecNum    int       // 秒数
	RQNum int    // 最大并发数，由命令行传入
	Urls   []string // url，由命令行传入
	userNum    int      // 用户数
	Durtime  int //持续分钟
	writetimes int //写入次数
	updatetimes int //更新次数
	readtimes int //查询次数
	donemu sync.Mutex
	api string
	logfile *os.File
	logger *log.Logger
)

var users []User

type User struct {
	UserId      int          // 用户id
	SBCNum     int           // 并发连接数
	QPSNum     int           // 总请求次数
	RTNum      time.Duration // 响应时间
	RTTNum     time.Duration // 平均响应时间
	SuccessNum int           // 成功次数
	FailNum    int           // 失败次数
	mu         sync.Mutex
}

func (u *User) request() {
	var tb time.Time
	var el time.Duration
	for i := 0;i < u.QPSNum;i++ {
		u.mu.Lock()
		u.SBCNum++
		u.mu.Unlock()
		go func(u *User) {
			tb = time.Now()
			uindex := rand.Intn(3)
			url := api+Urls[uindex]
			logger.Printf("%s,请求用户ID:%d, 请求URL:%s\n", time.Now(),u.UserId, url)
			_, err := http.Get(url)
			if err == nil {
				el = time.Since(tb)
				u.mu.Lock() // 上锁
				u.SuccessNum++
				u.RTNum += el
				u.mu.Unlock() // 解锁
			} else {
				u.mu.Lock() // 上锁
				u.FailNum++
				u.mu.Unlock() // 解锁
			}
			donemu.Lock()
			if uindex == 0{
				writetimes+=1000
			} else if (uindex == 1){
				updatetimes+=100
			} else if (uindex == 2){
				readtimes += 2000
			}
			donemu.Unlock()
		}(u)
	}
}

func (u *User) show() {
	logger.Printf("用户id：%d,并发数：%d,请求次数：%d,平均响应时间：%s,成功次数：%d,失败次数：%d\n",
		u.UserId,
		u.SBCNum,
		u.SuccessNum + u.FailNum,
		u.RTNum/(time.Duration(SecNum)*time.Second),
		u.SuccessNum,
		u.FailNum)
}

func showAll(us []User) {
	uLen := len(us)

	var SBCNum     int           // 并发连接数
	var RTNum      time.Duration // 响应时间
	var SuccessNum int           // 成功次数
	var FailNum    int           // 失败次数

	for i := 0;i < uLen;i++ {
		SBCNum += us[i].SBCNum
		SuccessNum += us[i].SuccessNum
		FailNum += us[i].FailNum
		RTNum += us[i].RTNum
		us[i].show()
	}
	if SecNum < 1{
		SecNum  = 1
	}
	logger.Printf("并发数：%d,请求次数：%d,平均响应时间：%s,成功次数：%d,失败次数：%d,写入次数:%d,更新次数:%d, 查询次数:%d\n",
		SBCNum,
		SuccessNum+FailNum,
		RTNum/(time.Duration(SecNum)*time.Second),
		SuccessNum,
		FailNum,
		writetimes,
		updatetimes,
		readtimes)
}

func init() {
	if len(os.Args) != 4 {
		log.Fatal("用户数 每秒请求次数 持续时间")
	}
	userNum, _ = strconv.Atoi(os.Args[1])
	RQNum, _ = strconv.Atoi(os.Args[2])
	Durtime, _ = strconv.Atoi(os.Args[3])
	Urls = []string{"addyly", "modyly", "readyly"}
	users = make([]User, userNum)
	api  = "http://192.168.6.168:8088/"
	root,_:=os.Getwd()
	fileName := root+"/pressure.log"
	fmt.Print(fileName)
	var err error
	logfile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil{
		log.Fatal(err.Error())
	}
	logger = log.New(logfile,"[iris]", log.LstdFlags)
	writetimes = 0
	readtimes = 0
	updatetimes = 0
}

func main() {
	go func() {
		for range time.Tick(2 * time.Second) {
			SecNum += 2
			showAll(users)
		}
	}()
	begin := time.Now()

	go func(){
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				requite()
			}
		}
	}()
	<-time.After(time.Duration(Durtime)*time.Minute)
	end := time.Now()
	logger.Printf("开始时间:%s, 结束时间:%s :\n", begin, end)
	showAll(users)
}

func requite() {
	c := make(chan int)
	for i := 0;i < userNum;i++ {
		users[i].UserId = i
		users[i].QPSNum = RQNum
		go users[i].request()
		time.Sleep(2 * time.Millisecond)
	}
	<- c    // 阻塞
}