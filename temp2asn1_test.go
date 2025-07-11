package temp2asn1_test

import (
	"testing"
	"time"

	t2a "github.com/takanoriyanagitani/go-sensor-temperature2asn1"
)

func TestToASN1DER(t *testing.T) {
	// Create a dummy sensor data
	sensorData := t2a.TemperatureSensorData{
		SensorID:    "a1b2c3d4-e5f6-7890-1234-567890abcdef",
		LocalID:     123,
		Timestamp:   time.Now().Unix(),
		Temperature: 25,
		Unit:        0, // Celsius
	}

	// Serialize the struct to ASN.1 DER bytes
	derBytes, err := sensorData.ToASN1DER()

	// Ensure the error is nil
	if err != nil {
		t.Fatalf("Serialization failed: %v", err)
	}

	// Ensure the returned bytes are not empty (basic check)
	if len(derBytes) == 0 {
		t.Fatal("Serialized bytes are empty")
	}

	t.Logf("Successfully serialized data to %d bytes: %x", len(derBytes), derBytes)
}
