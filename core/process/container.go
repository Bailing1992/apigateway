package process

import (
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/Bailing1992/apigateway/core/server/fasthttp"
	"github.com/Bailing1992/apigateway/core/service_proxy"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

func GetSingleInstance() *Container {
	once.Do(func() {
		instance = &Container{
			config: &config.ContainerConfig{},
		}
	})

	return instance
}

// Container
type Container struct {
	config    *config.ContainerConfig
	server    *fasthttp.Server
	handlers  []service_proxy.NamedHandler // 最终的handler列表
	initReady int32                        // 0: not_ready, 1: ready
	shutdown  int32
	proxyMap  sync.Map // psm -> *proxy.Proxy
}

var (
	once     sync.Once
	instance *Container
)

func (c *Container) SetReady() {
	atomic.StoreInt32(&c.initReady, 1)
}

func (c *Container) IsReady() bool {
	val := atomic.LoadInt32(&c.initReady)
	return val > 0
}

func (c *Container) StartServer() error {
	return c.server.StartServer()
}

func (c *Container) StopServer() error {
	atomic.StoreInt32(&c.shutdown, 1)
	return c.server.StopServer()
}

func (c *Container) Handlers() []service_proxy.NamedHandler {
	if len(c.handlers) == 0 {
		// 允许多线程同时设置，不需要同步机制。因为
		// 1. 这个方法是幂等的，并发执行不会有不一样的结果。
		// 2. 这个方法仅在初始化的时候执行，后续将不会再执行。
		// 3. 这个方法仅内存运[算，不会造成过高的负载。
		c.InitHandlers()
	}

	return c.handlers
}

func (c *Container) InitHandlers() {

}

func Stop() error {
	if err := container.StopServer(); err != nil {
		return err
	}

	return nil
}

func Start() {
	errCh := make(chan error, 1)
	go func() {
		errCh <- container.StartServer()
	}()
	waitSignal(errCh)
}

func (c *Container) setProxy(key string, proxyInstance *service_proxy.ServiceProxy) {
	c.proxyMap.Store(key, proxyInstance)
}

func (c *Container) SetServer(serverInstance *fasthttp.Server) {
	c.server = serverInstance
}

func (c *Container) GetProxy(key string) *service_proxy.ServiceProxy {
	proxyInstance, ok := c.proxyMap.Load(key)
	if !ok {
		return nil
	}
	return proxyInstance.(*service_proxy.ServiceProxy)
}

func waitSignal(errCh chan error) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	for {
		select {
		case sig := <-signals:
			switch sig {
			// exit forcely
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGHUP:
				// TODO reload server
				fallthrough
			case syscall.SIGINT:
				return Stop()
			}
		case err := <-errCh:
			return err
		}
	}

	return nil
}
