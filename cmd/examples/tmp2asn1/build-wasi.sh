#!/bin/sh

tinygo \
	build \
	-o ./tmp2asn1.wasm \
	-target=wasip1 \
	-opt=z \
	-no-debug \
	./main.go
