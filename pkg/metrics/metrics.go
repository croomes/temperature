package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	temp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celcius",
		Help: "The temperature in degrees Celcius",
	})
	humidity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "humidity_relative_percentage",
		Help: "The current humidity in percentage of relative humidity",
	})
	pressure = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pressure_pascals",
		Help: "The current current pressure, in pascals",
	})
	altitude = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "altitude_meters",
		Help: "The current altitude in meters based on the current barometric pressure and estimated pressure at sea level",
	})
)

type Metrics struct {
	tempCh chan float32
	humidityCh chan float32
	pressureCh chan float32
	altitudeCh chan float32
	closeCh chan struct{}
}

func New(tempCh, humidityCh, pressureCh, altitudeCh chan float32) *Metrics {
	return &Metrics{
		tempCh: tempCh,
		humidityCh: humidityCh,
		pressureCh: pressureCh,
		altitudeCh: altitudeCh,
	}
}

// Run the server, blocking.
func (m *Metrics) Run() error {

	m.closeCh = make(chan struct{})

	// Process incoming updates.
	go m.update()

	// The Prometheus Handler function provides a default handler to expose
	// metrics via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":8080", nil)

}

// Shutdown the server.
func (m *Metrics) Shutdown() {

	// Stop watching for updates if started.
	if m.closeCh != nil {
		close(m.closeCh)
	}

}

// update listens on the channels for incoming measurements and updates the
// Prometheus guages.  Stops when closeCh is closed.
func (m *Metrics) update() {

	for {
		select {
		case t := <-m.tempCh:
			temp.Set(float64(t))
		case t := <-m.humidityCh:
			humidity.Set(float64(t))
		case t := <-m.pressureCh:
			pressure.Set(float64(t))
		case t := <-m.altitudeCh:
			altitude.Set(float64(t))
		case <- m.closeCh:
			break
		}
	}

}

