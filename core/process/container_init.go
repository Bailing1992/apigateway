package process

import (
	"context"
	"github.com/Bailing1992/apigateway/consts"
	"github.com/Bailing1992/apigateway/core/config"
	"github.com/Bailing1992/apigateway/core/errors"
	"github.com/Bailing1992/apigateway/core/server/fasthttp"
	"github.com/Bailing1992/apigateway/core/service_proxy"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	container *Container
)

func init() {
	if v := os.Getenv("GOMAXPROCS"); v == "" {
		if v := os.Getenv("MY_CPU_LIMIT"); v != "" {
			n, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				runtime.GOMAXPROCS(int(n))
			}
		}
	}

	container = GetSingleInstance()

}
func Init() {
	container.InitConfig(consts.DefaultConfigFile)
	//container.InitLoaders(ctx)
	container.RegisterProxy()
	container.RegisterServer()
	//container.RegisterRouter()
	//container.SetReady()
}

func (c *Container) InitConfig(path string) error {
	configFilePath, err := filepath.Abs(path)
	if err != nil {
		return errors.NewInternalError("Get abs main config file failed, error: %v", err)
	}
	containerConfig, err := config.NewContainerConfigFromFile(configFilePath)
	if err != nil {
		return errors.NewInternalError("Can not parse main config file, path: %s, err: %v", configFilePath, err)
	}
	println("GlobalConfig: ", containerConfig)
	c.config = containerConfig
	return nil
}

func (c *Container) RegisterProxy() error {
	ctx := context.Background()
	proxyConfigList := container.config.ProxyListConfig
	for i := range proxyConfigList {
		proxyConfig := proxyConfigList[i]
		proxyInstance, err := service_proxy.NewServiceProxy(ctx, proxyConfig)
		if err != nil {
			return errors.NewInternalError("Register proxy failed, psm: %s, error: %s", proxyConfig.PSM, err)
		}
		container.setProxy(proxyConfig.PSM, proxyInstance)

	}
	return nil
}

func (c *Container) RegisterServer() error {
	serverConfig := c.config.Server

	serverInstance := fasthttp.NewServer(serverConfig, c.handleRequest)
	c.SetServer(serverInstance)
	return nil
}

func (c *Container) RegisterRouter() (returnError error) {
	//newRouter := router.NewRouter(processContext.Logger())

	//proxyConfigList := c.config.ProxyListConfig
	//for _, proxyConfig := range proxyConfigList {
	//	p := c.GetProxy(proxyConfig.PSM)
	//}
	return
}
