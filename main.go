package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/croomes/temperature/pkg/adapter"
	"github.com/croomes/temperature/pkg/metrics"
	"gobot.io/x/gobot/drivers/i2c"
)

func main() {

	// Parse command-line flags
	var device string
	flag.StringVar(&device, "device", "", "i2c device")
	flag.Parse()

	// Initialise i2c Adapter
	adapter, err := adapter.New(device)
	if err != nil {
		log.Fatal(err)
	}
	adapter.Connect()

	// Initialise sensor
	sensor := i2c.NewBME280Driver(adapter)

	if err := sensor.Start(); err != nil {
		log.Fatal(err.Error())
	}

	// Setup channels.
	stopCh := make(chan os.Signal)
	tempCh := make(chan float32)
	humidityCh := make(chan float32)
	pressureCh := make(chan float32)
	altitudeCh := make(chan float32)

	// Watch for Ctrl-C
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	server := metrics.New(tempCh, humidityCh, pressureCh, altitudeCh)
	go func() {
		server.Run()
	}()

	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			// Temperature
			v, err := sensor.Temperature()
			if err != nil {
				log.Println(err.Error())
			}
			tempCh <- v

			// Humidity
			v, err = sensor.Humidity()
			if err != nil {
				log.Println(err.Error())
			}
			humidityCh <- v

			// Pressure
			v, err = sensor.Pressure()
			if err != nil {
				log.Println(err.Error())
			}
			pressureCh <- v

			// Altitude
			v, err = sensor.Altitude()
			if err != nil {
				log.Println(err.Error())
			}
			altitudeCh <- v
		case <-stopCh:
			server.Shutdown()
			os.Exit(0)
		}
	}

}
