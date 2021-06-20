package main

import "fmt"
import "strconv"


type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipAddr IPAddr) String() string {
     ipDottedFormat := ""
     for i := range ipAddr {
	strtest := strconv.Itoa(int(ipAddr[i]))
	// fmt.Println(i , ipAddr[i], i)
	ipDottedFormat += strtest
	if i != len(ipAddr) - 1{
		ipDottedFormat +=  "."
	}
     }
     return ipDottedFormat
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
