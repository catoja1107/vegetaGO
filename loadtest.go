/*
Compile first with GCC or something ¯\_(ツ)_/¯

syntax: loadtest <command> [command flags]

attack command:
  -body string
    	Requests body file
  -cert string
    	TLS client PEM encoded certificate file
  -chunked
    	Send body with chunked transfer encoding
  -connections int
    	Max open idle connections per target host (default 10000)
  -duration duration
    	Duration of the test [0 = forever]
  -format string
    	Targets format [http, json] (default "http")
  -h2c
    	Send HTTP/2 requests without TLS encryption
  -header value
    	Request header
  -http2
    	Send HTTP/2 requests when supported by the server (default true)
  -insecure
    	Ignore invalid server TLS certificates
  -keepalive
    	Use persistent connections (default true)
  -key string
    	TLS client PEM encoded private key file
  -laddr value
    	Local IP address (default 0.0.0.0)
  -lazy
    	Read targets lazily
  -max-body value
    	Maximum number of bytes to capture from response bodies. [-1 = no limit] (default -1)
  -max-workers uint
    	Maximum number of workers (default 18446744073709551615)
  -name string
    	Attack name
  -output string
    	Output file (default "stdout")
  -proxy-header value
    	Proxy CONNECT header
  -rate value
    	Number of requests per time unit [0 = infinity] (default 50/1s)
  -redirects int
    	Number of redirects to follow. -1 will not follow but marks as success (default 10)
  -resolvers value
    	List of addresses (ip:port) to use for DNS resolution. Disables use of local system DNS. (comma separated list)
  -root-certs value
    	TLS root certificate files (comma separated list)
  -targets string
    	Targets file (default "stdin")
  -timeout duration
    	Requests timeout (default 30s)
  -unix-socket string
    	Connect over a unix socket. This overrides the host address in target URLs
  -workers uint
    	Initial number of workers (default 10)

encode command:
  -output string
    	Output file (default "stdout")
  -to string
    	Output encoding [csv, gob, json] (default "json")

plot command:
  -output string
    	Output file (default "stdout")
  -threshold int
    	Threshold of data points above which series are downsampled. (default 4000)
  -title string
    	Title and header of the resulting HTML page (default "loadtest Plot")

report command:
  -buckets string
    	Histogram buckets, e.g.: "[0,1ms,10ms]"
  -every duration
    	Report interval
  -output string
    	Output file (default "stdout")
  -type string
    	Report type to generate [text, json, hist[buckets], hdrplot] (default "text")

examples:
  echo "GET http://localhost/" | loadtest attack -duration=5s | tee results.bin | loadtest report
  loadtest report -type=json results.bin > metrics.json
  cat results.bin | loadtest plot > plot.html
  cat results.bin | loadtest report -type="hist[0,100ms,200ms,300ms]"

*/

package main

import (
	"fmt" 
	"vegeta" //http load testing
	"os"
	)

func main() {
  test := "World!"
  fmt.Println("Hello %v\n", test)
  
  argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]
  
  vegeta(os.Args[]) //library includes return.
  
}