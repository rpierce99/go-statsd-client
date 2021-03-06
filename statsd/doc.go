/*
Package statsd provides a StatsD client implementation that is safe for
concurrent use by multiple goroutines and for efficiency can be created and
reused.

Example usage:

	// first create a client
	client, err := statsd.New("127.0.0.1:8125", "test-client")
	// handle any errors
	if err != nil {
		log.Fatal(err)
	}
	// make sure to clean up
	defer client.Close()

	// Send a stat
	err = client.Inc("stat1", 42, 1.0)
	// handle any errors
	if err != nil {
		log.Printf("Error sending metric: %+v", err)
	}

*/
package statsd
