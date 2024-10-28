package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// PageData structure to store page information
type PageData struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

// List of Wikipedia pages to scrape
var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

func main() {
	// Create a new asynchronous collector
	collector := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)

	// Record start time for the scraping process
	startTime := time.Now()

	// Initialize WaitGroup to wait for all scraping tasks
	var wg sync.WaitGroup
	var fileMutex sync.Mutex

	// Create output file to save JSON results
	outputFile, err := os.Create("output.jl")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Define the scraping logic to handle page content
	collector.OnHTML("body", func(e *colly.HTMLElement) {
		pageContent := PageData{
			URL:  e.Request.URL.String(),
			Text: e.ChildText("p"), // Extract text from the first paragraph
		}

		// Serialize page content to JSON
		jsonData, err := json.Marshal(pageContent)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		// Use mutex to prevent race conditions when writing to the file
		fileMutex.Lock()
		defer fileMutex.Unlock()
		_, err = outputFile.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	})

	// Start a scraping task for each URL
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println("Scraping:", url)
			if err := collector.Visit(url); err != nil {
				fmt.Println("Error visiting URL:", err)
			}
		}(url)
	}

	// Wait for all scraping tasks to complete
	wg.Wait()

	// Allow the collector to finish all ongoing requests
	collector.Wait()

	// Record and print the total elapsed time for the scraping process
	fmt.Printf("Total elapsed time for scraping: %s\n", time.Since(startTime))
}
