# btsensor
scans for bluetooth devices

## Build command

### currentlabs
mkdir -p executables && rm -rf executables/* && GOARM=6 GOARCH=arm GOOS=linux go build -o executables/ble_currentlabs && scp executables/ble_currentlabs pi@1x.x.x.x:poc_tibco

### gatt
mkdir -p executables && rm -rf executables/* && GOARM=6 GOARCH=arm GOOS=linux go build -o executables/ble_gatt && scp executables/ble_gatt pi@x.x.x.x:poc_tibco
