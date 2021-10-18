package utils

import (
	"fmt"
	"os"
	"strings"
)

var (
	psm           string
	cluster       string
	idc           string
	localIP       string
	configAddr    string
	isSidecarMode bool
)

func init() {
	psm = os.Getenv("TCE_PSM")
	cluster = os.Getenv("TCE_CLUSTER")
	localIP = os.Getenv("MY_POD_IP")
	configAddr = os.Getenv("AGW_CONFIG_ADDR")

}

func Product() string {
	if isSidecarMode {
		return "sidecar"
	}
	return strings.Split(PSM(), ".")[0]
}

func PSM() string {
	return psm
}

func Cluster() string {
	return cluster
}

func IsInternalCluster() bool {
	if isSidecarMode {
		return false
	}
	if Cluster() == "internal" {
		return true
	}
	return false
}

func LocalIP() string {
	return localIP
}

func LocalAddr() string {
	return fmt.Sprintf("%s:%s", LocalIP(), os.Getenv("PORT0"))
}

func LocalAddrDebug() string {
	return fmt.Sprintf("%s:%s", LocalIP(), os.Getenv("PORT1"))
}
