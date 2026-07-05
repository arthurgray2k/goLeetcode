package all

// Problem: 1108
// Title: Defanging an IP Address
// Category: all
// Tags: all


import "strings"

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
