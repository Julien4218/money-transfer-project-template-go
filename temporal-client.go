package app

import (
	"os"

	"go.temporal.io/sdk/client"
)

func GetTemporalClient() (client.Client, error) {
	dialOptions := client.Options{}
	hostport := os.Getenv("TEMPORAL_HOSTPORT")
	if len(hostport) > 0 {
		dialOptions.HostPort = hostport
	}
	// Create the client object just once per process
	c, err := client.Dial(dialOptions)
	return c, err
}
