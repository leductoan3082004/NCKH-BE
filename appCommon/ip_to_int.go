package appCommon

import (
	"math/big"
	"net"
)

func InetAton(_ip string) int64 {
	ip := net.ParseIP(_ip)
	ipv4Int := big.NewInt(0)
	ipv4Int.SetBytes(ip.To4())
	return ipv4Int.Int64()
}
