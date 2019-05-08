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
	sum  int //总次数
	finish int //完成次数
	done chan struct{}
	donemu sync.Mutex
	api string
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
		u.SBCNum++
		go func(u *User) {
			for {
				tb = time.Now()
				url := api+Urls[rand.Intn(3)]
				fmt.Printf("%s,请求用户ID:%d, 请求URL:%s\n", time.Now(),u.UserId, url)
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
				finish+=1
				donemu.Unlock()
				time.Sleep(1 * time.Second)
			}
		}(u)
	}
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
	fmt.Printf("并发数：%d,请求次数：%d,平均响应时间：%s,成功次数：%d,失败次数：%d\n",
		SBCNum,
		SuccessNum+FailNum,
		RTNum/(time.Duration(SecNum)*time.Second),
		SuccessNum,
		FailNum)
	if (finish >= sum){
		done <- struct{}{}
	}
}

func init() {
	if len(os.Args) != 3 {
		log.Fatal("用户数 请求次数 url")
	}
	userNum, _ = strconv.Atoi(os.Args[1])
	RQNum, _ = strconv.Atoi(os.Args[2])
	sum =  userNum*RQNum
	finish = 0
	Urls = []string{"addyly", "modyly", "readyly"}
	users = make([]User, userNum)
	done = make(chan struct{})
	api  = "http://localhost:8081/"
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
	<-done
	end := time.Now()
	fmt.Printf("开始时间:%s, 结束时间:%s, 总耗时：%s:\n", begin, end, end.Sub(begin))
}

func requite() {
	c := make(chan int)
	temp := 0
	for i := 0;i < userNum;i++ {
		if RQNum % userNum != 0 && i < RQNum % userNum {
			temp = 1
		} else {
			temp = 0
		}
		users[i].UserId = i
		users[i].QPSNum = RQNum / userNum + temp
		go users[i].request()
		time.Sleep(45 * time.Millisecond)
	}
	<- c    // 阻塞
}