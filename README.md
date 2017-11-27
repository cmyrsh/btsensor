# btsensor
scans for bluetooth devices

## Build command

mkdir -p executables && rm executables/btsensor && GOARM=6 GOARCH=arm GOOS=linux go build -o executables/btsensor && scp executables/btsensor raspberrypi..
