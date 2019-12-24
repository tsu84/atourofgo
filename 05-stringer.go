// https://tour.golang.org/methods/18

package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	ipString := make([]string, 4)
	for i,v := range ip {
		ipString[i] = strconv.Itoa(int(v))
	}
	return strings.Join(ipString, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
