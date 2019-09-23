package circuit

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var cb *Breaker

func init() {
	var st Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func TestCircuit(t *testing.T) {
	body, err := Get("https://cn.bing.com/?mkt=zh-CN")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
