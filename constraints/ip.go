package constraints

import (
	"net"

	"github.com/binadel/vali/core"
	"github.com/binadel/vali/errors"
)

var ipError = &errors.BasicError{
	Code:    core.RuleIP,
	Message: "value must be a valid IP address",
}

var ipv4Error = &errors.BasicError{
	Code:    core.RuleIPv4,
	Message: "value must be a valid IPv4 address",
}

var ipv6Error = &errors.BasicError{
	Code:    core.RuleIPv6,
	Message: "value must be a valid IPv6 address",
}

func ParseIP(value string) (net.IP, error) {
	ip := net.ParseIP(value)
	if ip == nil {
		return nil, ipError
	}
	return ip, nil
}

func ParseIPv4(value string) (net.IP, error) {
	ip := net.ParseIP(value)
	if ip == nil {
		return nil, ipv4Error
	}
	if ipv4 := ip.To4(); ipv4 != nil {
		return ipv4, nil
	}
	return nil, ipv4Error
}

func ParseIPv6(value string) (net.IP, error) {
	ip := net.ParseIP(value)
	if ip == nil || ip.To4() != nil {
		return nil, ipv6Error
	}
	return ip, nil
}
