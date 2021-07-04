package main

import (
	"flag"
	"os"

	"github.com/abtiwary/goexifapi/exifserver"
	log "github.com/sirupsen/logrus"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	initLogger()

	ipaddr := flag.String("ipaddr", "", "the IP address to listen on")
	port := flag.Int("port", 8888, "the Port to listen on")

	flag.Parse()

	srv := exifserver.NewServer(*ipaddr, *port)

	log.WithFields(log.Fields{
		"server_ip":   srv.IP,
		"server_port": srv.Port,
	}).Info("starting the http server")
	srv.StartHTTPServer()

	// todo: graceful stuff here
	// ...

}
