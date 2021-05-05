# pulsar-producer-example
Pulsar producer in Go

## Build and Run
Build: `go build producer.go`
Run: go run .  

Output:
```INFO[0000] [Connecting to broker]                        remote_addr="pulsar://localhost:6650"
INFO[0000] [TCP connection established]                  local_addr="127.0.0.1:52809" remote_addr="pulsar://localhost:6650"
INFO[0000] [Connection is ready]                         local_addr="127.0.0.1:52809" remote_addr="pulsar://localhost:6650"
INFO[0000] [Created producer]                            cnx="127.0.0.1:52809 -> 127.0.0.1:6650" producerID=1 producer_name=standalone-125-8 topic="persistent://public/default/inventory-service"
Published message
INFO[0000] [Closing producer]                            producerID=1 producer_name=standalone-125-8 topic="persistent://public/default/inventory-service"
INFO[0000] [Closed producer]                             producerID=1 producer_name=standalone-125-8 topic="persistent://public/default/inventory-service"
```
