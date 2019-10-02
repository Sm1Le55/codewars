package kata

import (
	"strconv"
	"strings"
)

//Is_valid_ip ...
func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")

	if len(octets) != 4 {
		return false
	}

	for _, octet := range octets {
		if num, err := strconv.ParseInt(octet, 10, 16); err == nil {
			if len(octet) > 2 && octet[0] == '0' {
				return false
			}
			if num < 0 || num > 255 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
