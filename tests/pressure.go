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
)

var (
	users []User
	userchan []chan struct{}
)

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

func (u *User) request(userchan chan struct{}, times int) {
	if times > 1{
		<-userchan
	}
	var tb time.Time
	var el time.Duration
	wg := sync.WaitGroup{}
	wg.Add(u.QPSNum)
	for i := 0;i < u.QPSNum;i++ {
		u.mu.Lock()
		u.SBCNum++
		u.mu.Unlock()
		go func(u *User) {
			tb = time.Now()
			uindex := rand.Intn(3)
			url := api+Urls[uindex]
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
			wg.Done()
		}(u)
	}
	wg.Wait()
	userchan <- struct{}{}
}

func (u *User) show() {
	fmt.Printf("用户id：%d,并发数：%d,请求次数：%d,平均响应时间：%s,成功次数：%d,失败次数：%d\n",
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
	fmt.Printf("并发数：%d,请求次数：%d,平均响应时间：%s,成功次数：%d,失败次数：%d,写入次数:%d,更新次数:%d, 查询次数:%d\n",
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
	userchan = make([]chan struct{}, userNum)
	api  = "http://127.0.0.1:8089/"
	root,_:=os.Getwd()
	fileName := root+"/pressure.log"
	var err error
	logfile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil{
		log.Fatal(err.Error())
	}
	writetimes = 0
	readtimes = 0
	updatetimes = 0
}

func main() {
	http.Get(api+"clearyly")
	go func() {
		for range time.Tick(2 * time.Second) {
			SecNum += 2
			showAll(users)
		}
	}()
	begin := time.Now()

	go func(){
		times := 1
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				requite(times)
				times+=1
			}
		}
	}()
	<-time.After(time.Duration(Durtime)*time.Minute)
	end := time.Now()
	fmt.Printf("开始时间:%s, 结束时间:%s :\n", begin, end)
	showAll(users)
}

func requite(times int) {
	for i := 0;i < userNum;i++ {
		users[i].UserId = i
		users[i].QPSNum = RQNum
		if (times == 1) {
			userchan[i] = make(chan struct{}, 1)
		}
		go users[i].request(userchan[i], times)
		time.Sleep(45 * time.Millisecond)
	}
}