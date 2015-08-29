package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/codeskyblue/proxylocal/pxlocal"
)

func main() {
	var addr = ":5000"
	var serverMode bool
	var serverAddr string
	var proxyPort int
	var proxyAddr string
	var subDomain string
	flag.BoolVar(&serverMode, "server", false, "run in server mode")
	flag.StringVar(&serverAddr, "addr", "localhost:5000", "server address")
	flag.IntVar(&proxyPort, "proxy-port", 0, "server proxy listen port")
	flag.StringVar(&subDomain, "subdomain", "", "proxy subdomain")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] <port | host:port>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if !serverMode && len(flag.Args()) != 1 {
		flag.Usage()
		return
	}

	proxyAddr = flag.Arg(0)
	if !strings.Contains(proxyAddr, ":") {
		if _, err := strconv.Atoi(proxyAddr); err == nil { // only contain port
			proxyAddr = "localhost:" + proxyAddr
		} else { // only contain host
			proxyAddr = proxyAddr + ":80"
		}
	}

	if serverMode {
		fmt.Println("Hello proxylocal, server listen on", addr)
		ps := pxlocal.NewProxyServer()
		log.Fatal(http.ListenAndServe(addr, ps))
	}

	pxlocal.StartAgent(proxyAddr, serverAddr, proxyPort)
	//startAgent(proxyAddr, serverAddr, proxyPort)
}
