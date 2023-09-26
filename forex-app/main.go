package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func fetchForexData(apiKey string, wg *sync.WaitGroup) {
	defer wg.Done()

	apiUrl := fmt.Sprintf("https://api.currencybeacon.com/v1/latest?api_key=%s", apiKey)

	// Create an HTTP client with a timeout of 1 second
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Send an HTTP GET request to fetch Forex prices
	response, err := client.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error fetching Forex prices: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Check if the request was successful (HTTP status code 200)
	if response.StatusCode == http.StatusOK {
		// Read the response body
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}

		// Process the Forex price data (assuming JSON response)
		fmt.Printf("Forex Prices:\n%s\n", string(responseBody))
	} else {
		fmt.Printf("HTTP Request failed with status code: %d\n", response.StatusCode)
	}
}

func main() {
	// Register Prometheus metrics
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "harrycorp_forex_requests_total",
			Help: "Total number of requests to HarryCorp Forex Pricing",
		},
		[]string{"endpoint"},
	)
	prometheus.MustRegister(requestCounter)

	// Get the API key from the environment variable
	apiKey := os.Getenv("CURRENCY_BEACON_API_KEY")

	// Check if the API key is empty or missing
	if apiKey == "" {
		fmt.Println("API key not found. Set the CURRENCY_BEACON_API_KEY environment variable.")
		return
	}

	// Number of concurrent fetches
	numFetches := 3 // Adjust as needed

	var wg sync.WaitGroup
	for i := 0; i < numFetches; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fetchForexData(apiKey, &wg)
			// Increment the Prometheus counter metric
			requestCounter.WithLabelValues("/metrics").Inc()
		}()
	}

	// Start an HTTP server to expose Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Printf("Server is listening on :%s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
