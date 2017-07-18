package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"code.cloudfoundry.org/garden/server"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerflags"
	"github.com/topherbullock/garden-k8s/backend"
	restclient "k8s.io/client-go/rest"
)

func main() {
	defaultListNetwork := "unix"
	defaultListAddr := "/tmp/garden.sock"
	if os.Getenv("PORT") != "" {
		defaultListNetwork = "tcp"
		defaultListAddr = "0.0.0.0:" + os.Getenv("PORT")
	}
	var listenNetwork = flag.String(
		"listenNetwork",
		defaultListNetwork,
		"how to listen on the address (unix, tcp, etc.)",
	)
	var listenAddr = flag.String(
		"listenAddr",
		defaultListAddr,
		"address to listen on",
	)

	lagerflags.AddFlags(flag.CommandLine)
	flag.Parse()

	logger, _ := lagerflags.New("garden-k8s")

	k8s := backend.New(&restclient.Config{})

	gardenServer := server.New(*listenNetwork, *listenAddr, 10*time.Minute, k8s, logger)
	err := gardenServer.Start()
	if err != nil {
		logger.Fatal("Server Failed to Start", err)
		os.Exit(1)
	}

	logger.Info("started", lager.Data{
		"network": *listenNetwork,
		"addr":    *listenAddr,
	})

	signals := make(chan os.Signal, 1)

	go func() {
		<-signals
		gardenServer.Stop()
		os.Exit(0)
	}()

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	select {}
}
