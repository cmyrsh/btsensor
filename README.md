# btsensor
scans for bluetooth devices

## Build command

### currentlabs
```
mkdir -p executables && rm -rf executables/* && GOARM=6 GOARCH=arm GOOS=linux go build -o executables/ble_currentlabs && scp executables/ble_currentlabs pi@1x.x.x.x:poc_tibco
```
### gatt
```
mkdir -p executables && rm -rf executables/* && GOARM=6 GOARCH=arm GOOS=linux go build -o executables/ble_gatt && scp executables/ble_gatt pi@x.x.x.x:poc_tibco
```

## Configuration

### Config Gatt

```
scan_interval: 20
scan_timeout: 20
buffer_size: 100
mqtt_address: tcp://localhost:1883
mqtt_topic: sensors/bluetooth
mqtt_on_topic: lifecycle/on/btsensor
mqtt_off_topic: lifecycle/off/btsensor
scan_duplicate: false
```
### Config currentlabs
```
scan_interval: 40
scan_timeout: 39
buffer_size: 100
mqtt_address: tcp://localhost:1883
mqtt_topic: sensors/bluetooth
mqtt_on_topic: lifecycle/on/btsensor
mqtt_off_topic: lifecycle/off/btsensor
scan_duplicate: false
```
