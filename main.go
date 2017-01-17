package main

import(
	"github.com/ooyala/go-dogstatsd"
	"log"
)

func main() {
	c, err := dogstatsd.New("127.0.0.1:8125")
	defer c.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Prefix every metric with the app name
	c.Namespace = "flubber."
	// Send the EC2 availability zone as a tag with every metric
	c.Tags = append(c.Tags, "us-east-1a")
	err = c.Gauge("request.duration", 1.2, nil, 1)

	// Post info to datadog event stream
	err = c.Info("cookie alert", "Cookies up for grabs in the kitchen!", nil)
}
