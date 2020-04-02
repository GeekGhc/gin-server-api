package server

import (
	"context"
	"fmt"
	"gin-server-api/app"
	"gin-server-api/app/setting"
	"gin-server-api/internal/logger"
	"gin-server-api/middleware"
	"gin-server-api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	go s.listenSign()
	go s.runServer()
	return s
}

// listenSign 监听退出信息
func (s *GracefulServer) listenSign() {
	//监听退出信号
	listener := make(chan os.Signal)
	signal.Notify(listener, syscall.SIGTERM)
	signal.Notify(listener, syscall.SIGINT)
	<-listener
	//正常退出
	s.stopping <- true
}

// runServer 运行http服务
func (s *GracefulServer) runServer() {
	defer func() {
		if err := recover(); err != nil {
			//非正常退出
			s.stopping <- false
		}
	}()

	//线上
	if app.Env().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()

	engine := gin.New()
	//middleware init
	s.initMiddleware(engine)
	//router init
	routers.InitRouter(engine)
	Addr := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	s.httpServer = &http.Server{
		Addr:              Addr,
		Handler:           engine,
		ReadTimeout:       6 * time.Second,
		WriteTimeout:      6 * time.Second,
		ReadHeaderTimeout: 500 * time.Millisecond,
		MaxHeaderBytes:    1 << 20, // 1 MB
	}

	err := s.httpServer.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func (s *GracefulServer) Wait() {
	if ok := <-s.stopping; ok {
		logger.Info("The application is exiting...")
	} else {
		logger.Warn("The application has an exception and is exiting...")
	}

	ctx, cancelTimer := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancelTimer()
		s.onShutDown()
	}()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		logger.Error("The application shutdown error")
	}
}

func (s *GracefulServer) initMiddleware(engine *gin.Engine) {
	//自定义access panic log日志
	engine.Use(gin.Logger())
	engine.Use(middleware.Recovery())
	//跨域
	engine.Use(cors.Default())
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
