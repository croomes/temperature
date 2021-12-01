# Temperature

Reads temperature, humidity, pressure and altitude from a BME280 sensor and
exposes as a Prometheus endpoint on http://<ip>:9090/metrics.  

Metrics are updated every 10 seconds.

The BME280 should be connected via the I2C interface.  Tested on a Raspberry Pi
Zero but should work Arduino by passing the device as a command line argument,
e.g.`./temperature -device /dev/tty.usbmodem1421`.  This is not required on the
RPI.

## Building

Run `go build` on the target platform.  Download any missing dependencies with
`go get <package>`.
