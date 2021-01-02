package sel

import (
	"fmt"
	"net/http"
	"time"
)

// WebSiteRacer takes two websites and returns the one that renders first
// If 10s pass and neither render, an err is returned.
func WebSiteRacer(urlA string, urlB string) (string, error) {

	var fastest string
	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		fastest = urlA
	} else {
		fastest = urlB
	}

	return fastest, nil
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// ConfigurableRacer takes two urls and returns the one that renders first
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// Racer is our actual racer used in production. It has been
// configured to timeout after 10 seconds.
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
