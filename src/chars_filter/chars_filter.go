package main

import "fmt"

func main() {
	for i := 33; i < 127; i++ {
		s := string([]byte{byte(i)})
		if s == "$" {
			s = "$dollar"
		}
		if s == "\"" {
			fmt.Printf("sub_filter '%s' \"0x%02x\";\n", s, i)
		} else {
			fmt.Printf("sub_filter \"%s\" \"0x%02x\";\n", s, i)
		}
	}
}
