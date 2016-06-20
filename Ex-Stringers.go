package main

// URL -> https://tour.golang.org/methods/18
import "fmt"

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (ad IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ad[0], ad[1], ad[2], ad[3])
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
