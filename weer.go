package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/higker/weer/pkg/config"
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

var (
	client = &http.Client{}
)

func main() {
	fmt.Println(Banner)
	fmt.Println(AppVer)
	fmt.Println(config.ServiceApiUrl)
	fmt.Println("Weer is weather command line apps.")
	
}
