package utils

import (
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

func SleepMS(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Get Fully Qualified Domain Name
// returns "unknown" or hostanme in case of error
func FQDN() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"

	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".") // return fqdn without trailing dot
		}

	}

	return hostname
}

func Split(s, sep string) []string {
	if s == "" {
		// return empty slice but not length = 1 as strings.Split
		return []string{}
	}

	return strings.Split(s, sep)
}

// s should be sorted
func InStringSlice(s []string, e string) (exist bool) {
	defer func() {
		if r := recover(); r != nil {
			exist = false
		}
	}()
	exist = s[sort.SearchStrings(s, e)] == e
	return
}
