# Go Constant NATS Subscriber
Subscribes to the specified channel and prints out the messages it sees.

If run without any parameters, it will ask for the server and subject you wish to subscribe to

If run with two parameters <server> and <subject> it will ingest them.

## Running From Source

`go run .`  
(You will be asked for server and subject)  
or  
`go run . nats://demo.nats.io 4222`  

## Running From Binary

These are held within the bin directory

`./nats-subscribe-linux-amd64`  
(You will be asked for server and subject)  
or  
`./nats-subscribe-linux-amd64 nats://demo.nats.io 4222`  

# Building

Can be built with standard Go build process:
`go build .`

`build.sh` script provided for ease of use building a few different choice architectures.