package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type ContainerConfig struct {
	Server          ServerConfig          `yaml:"Server"`
	ProxyListConfig []*ProxyServiceConfig `yaml:"ProxyListConfig"`
}

type ServerConfig struct {
	ip                  string `yaml:"Ip"`
	port                uint32 `yaml:"Port"`
	name                string `yaml:"Name"`
	maxWaitTime         uint64 `yaml:"MaxWaitTime"`
	requestTimeout      uint32 `yaml:"RequestTimeout"`
	requestReadTimeout  uint32 `yaml:"RequestReadTimeout"`
	requestWriteTimeout uint32 `yaml:"RequestWriteTimeout"`
}

func NewContainerConfigFromFile(filename string) (*ContainerConfig, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	yamlContent, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	var fullConf ContainerConfig
	err = yaml.Unmarshal(yamlContent, &fullConf)
	if err != nil {
		return nil, err
	}

	return &fullConf, nil
}

func (c *ServerConfig) GetMaxWaitTime() time.Duration {
	return time.Duration(c.maxWaitTime)
}

func (c *ServerConfig) setDefaultIP() {
	c.ip = "0.0.0.0"
}

func (c *ServerConfig) setDefaultPort() {
	c.port = 3838
}

func (c *ServerConfig) GetIP() string {
	if c.ip == "" {
		c.setDefaultIP()
	}
	return c.ip
}

func (c *ServerConfig) GetPort() uint32 {
	if c.port == 0 {
		c.setDefaultPort()
	}
	return c.port
}

func (c *ServerConfig) setDefaultName() {
	c.name = "AGW"
}

func (c *ServerConfig) GetName() string {
	if c.name == "" {
		c.setDefaultName()
	}
	return c.name
}
