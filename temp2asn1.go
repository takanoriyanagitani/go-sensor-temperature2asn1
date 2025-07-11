package temp2asn1

import (
	"encoding/asn1"
)

type TemperatureSensorData struct {
	SensorID    string `asn1:"optional,tag:0"`
	LocalID     int    `asn1:"optional,tag:1"`
	Timestamp   int64  `asn1:"tag:2"`
	Temperature int    `asn1:"tag:3"`
	Unit        int    `asn1:"tag:4"`
}

func (t TemperatureSensorData) ToASN1DER() ([]byte, error) {
	return asn1.Marshal(t)
}
