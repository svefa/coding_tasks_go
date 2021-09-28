package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hi..")
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	k := numUniqueEmails(emails)
	fmt.Println(k)
}

func numUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, e := range emails {
		x := strings.Split(e, "@")
		l := strings.Split(x[0], "+")
		a := strings.Split(l[0], ".")
		r := strings.Join(a, "") + "@" + x[1]
		set[r] = true

	}
	return len(set)
}
