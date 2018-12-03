package utils

import (
	"net"
	"os"
)

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	var _ip string = ""
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				_ip = ipnet.IP.String()
			}
		}
	}
	return _ip
}

