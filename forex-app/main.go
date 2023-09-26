package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    // Define the API URL for Forex price changes
    apiUrl := "https://api.example.com/forex-prices"  // Replace with the actual API URL

    // Send an HTTP GET request to fetch Forex prices
    response, err := http.Get(apiUrl)
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
