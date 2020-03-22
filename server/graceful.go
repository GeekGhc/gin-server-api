package server

import (
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GracefulServer struct {
	httpServer *http.Server
	stopping   chan bool
	status     bool //server 状态
	sync.Mutex
}

func NewGracefulServer() *GracefulServer {
	return &GracefulServer{
		stopping: make(chan bool),
		status:   false,
	}
}

//Run 运行应用
func (s *GracefulServer) Run(env string) *GracefulServer {
	s.Lock()
	defer s.Unlock()
	if s.status {
		return s
	}

	s.status = true
	s.onBoot(env)
	return s
}

// listenSign 监听退出信息
func (s *GracefulServer) listenSign(){
	//监听退出信号
	listener := make(chan os.Signal)
	signal.Notify(listener,syscall.SIGTERM)
	signal.Notify(listener,syscall.SIGINT)
	<- listener
	//正常退出
	s.stopping <- true
}

// runServer 运行http服务
func (s *GracefulServer) runServer(){
	defer func(){
		if err := recover();err != nil{
			//非正常退出
			s.stopping <- false
		}
	}()
}

// onBoot 当启动时执行
func (s *GracefulServer) onBoot(env string) {
	//初始化随机数
	rand.Seed(time.Now().UnixNano())
	//初始化环境 日志 配置
}

// onShutDown 当关闭时执行
func (s *GracefulServer) onShutDown() {
	//
}
