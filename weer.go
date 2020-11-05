package main

import (
	"fmt"
)

const (
	// Banner is application logo
	Banner = `
		________                   
		|  |  |  |.-----.-----.----.
		|  |  |  ||  -__|  -__|   _|
		|________||_____|_____|__| 
	`
	// AppVer is version info
	AppVer = "0.0.1 BETA"
)

func main() {
	fmt.Println(Banner)
	fmt.Println(AppVer)
	fmt.Println("Weer is weather command line apps.")
}
