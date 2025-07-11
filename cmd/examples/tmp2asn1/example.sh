#!/bin/sh

oder="./sample.d/tmp.asn1.der.dat"
cjson="./sample.d/tmp.json"

mkdir -p ./sample.d

der2fq(){
	cat /dev/stdin |
		fq \
			-d asn1_ber
}

der2jer(){
	cat /dev/stdin |
		xxd -ps |
		tr -d '\n' |
		python3 \
			-m asn1tools \
			convert \
			-i der \
			-o jer \
			./basic-temperature.asn \
			TemperatureReading \
			-
}

echo writing sample sensor data using asn1/der...
wazero \
	run \
	./tmp2asn1.wasm |
	dd if=/dev/stdin of="${oder}" bs=1048576 status=none

echo
echo converting to jer using asn1tools...
cat "${oder}" | der2jer |jq -c > "${cjson}"

cat "${cjson}" | jq

echo
echo comparing file size...' json, gzipped json, asn1/der'
cat "${cjson}" | wc -c
cat "${cjson}" | gzip --best | wc -c
cat "${oder}" | wc -c
