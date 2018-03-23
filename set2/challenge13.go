package main

import (
	"fmt"
	"strings"
)

func parseQuery(query string) map[string]string {
	m = make(map[string]string)
	query = strings.Trim(query, "&")
	kvs := strings.Split(query, "&")
	for i := 0; i < len(kvs); i++ {
		k, v := ...strings.Split(kvs, "=")
		m[k] = v
	}
	return m
}

func challenge13() {
	fmt.Println(DetectECB(randomEncrypt))
}

func main() {
	challenge13()
}
