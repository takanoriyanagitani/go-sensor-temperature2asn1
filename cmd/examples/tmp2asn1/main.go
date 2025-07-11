package main

import (
	"fmt"
	"io"
	"os"
	"time"

	t2a "github.com/takanoriyanagitani/go-sensor-temperature2asn1"
)

var dummyData t2a.TemperatureSensorData = t2a.TemperatureSensorData{

	SensorID:    "test-sensor-123",
	Timestamp:   time.Now().Unix(),
	Temperature: 22,
	Unit:        0, // Celsius
}

func main() {
	der, e := dummyData.ToASN1DER()
	if nil != e {
		fmt.Fprintf(os.Stderr, "Error serializing to ASN.1 DER: %v\n", e)
		os.Exit(1)
	}

	var writer io.Writer = os.Stdout
	_, err := writer.Write(der)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Successfully serialized and wrote %d bytes to stdout.\n", len(der))
}
