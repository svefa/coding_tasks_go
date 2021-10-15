package main

import "strings"

func validIPAddress(IP string) string {
	if validIPv4(IP) {
		return "IPv4"
	}
	if validIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(IP string) bool {
	const N = 4
	s := strings.Split(IP, ".")
	if len(s) != N {
		return false
	}
	for i := 0; i < N; i++ {
		if len(s[i]) == 0 {
			return false
		}
		if len(s[i]) > 1 && s[i][0] == '0' {
			return false
		}
		if len(s[i]) > 3 {
			return false
		}
		d := 0
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] < '0' || s[i][j] > '9' {
				return false
			}
			d = d*10 + int(s[i][j]-'0')
		}
		if d > 255 {
			return false
		}
	}
	return true
}

func validIPv6(IP string) bool {
	const N = 8
	s := strings.Split(IP, ":")
	if len(s) != N {
		return false
	}
	for i := 0; i < N; i++ {
		if len(s[i]) == 0 || len(s[i]) > 4 {
			return false
		}

		for j := 0; j < len(s[i]); j++ {
			if !((s[i][j] >= '0' && s[i][j] <= '9') ||
				(s[i][j] >= 'a' && s[i][j] <= 'f') ||
				(s[i][j] >= 'A' && s[i][j] <= 'F')) {
				return false
			}
		}
	}
	return true
}
