package pdserver

import (
	"log"
	"os"
)

type ServerConfig struct {
	HTTPAddress        string `flag:"http-address"`
	BroadcastAddr      string `flag:"broadcast-address"`
	BroadcastInterface string `flag:"broadcast-interface"`

	ReverseProxyPort string `flag:"reverse-proxy-port"`

	ClusterID                  string   `flag:"cluster-id"`
	ClusterLeadershipAddresses string   `flag:"cluster-leadership-addresses" cfg:"cluster_leadership_addresses"`
	BalanceInterval            []string `flag:"balance-interval"`

	LogLevel int32  `flag:"log-level" cfg:"log_level"`
	LogDir   string `flag:"log-dir" cfg:"log_dir"`
}

func NewServerConfig() *ServerConfig {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return &ServerConfig{
		HTTPAddress:        "0.0.0.0:18001",
		BroadcastAddr:      hostname,
		BroadcastInterface: "eth0",

		ClusterLeadershipAddresses: "",
		ClusterID:                  "",

		LogLevel: 1,
		LogDir:   "",
	}
}