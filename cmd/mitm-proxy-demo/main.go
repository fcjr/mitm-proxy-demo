package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	proxy "github.com/fcjr/mitm-proxy-demo"
	"github.com/fcjr/mitm-proxy-demo/utils"
)

const authorityName string = "MITM Proxy Demo Authority"

func main() {

	// parse arguments
	var port int
	var autoProxy bool
	flag.IntVar(&port, "port", 9000, "port to bind to, defaults to 9000")
	flag.BoolVar(&autoProxy, "autoconfig", false, "automatically configure proxy settings")
	flag.Parse()

	// Setup safe shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nInterupt detected")
		fmt.Println("Shutting Down Proxy...")
		utils.UninstallCert(authorityName)
		if autoProxy {
			utils.DisableProxy()
		}
		os.Exit(0)
	}()

	// serve proxy
	if autoProxy {
		if err := utils.EnableProxy(port); err != nil { // todo do after serve
			log.Fatal(err)
		}
	}
	proxy.Serve(port, authorityName)
}
