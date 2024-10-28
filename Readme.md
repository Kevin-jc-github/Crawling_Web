# Wikipedia Web Scraper using Go Colly

This repository contains a web scraping application written in Go using the Colly framework. The scraper is designed to collect text data from a list of Wikipedia pages related to intelligent systems and robotics.

## Project Overview

The main goal of this project is to demonstrate concurrent web scraping capabilities using Go's goroutines and the Colly scraping library. The scraper will extract the text content from specific Wikipedia pages and save it in a JSON lines (`.jl`) format, which is a common format for large data processing and input to databases.

### Features

- Scrapes multiple Wikipedia pages concurrently for faster data collection.
- Extracts text from the `<p>` elements on each page.
- Saves the scraped data in a JSON lines file (`output.jl`), where each line represents a JSON object for a webpage.
- Utilizes a mutex to ensure safe concurrent writing to the output file.

## Requirements

- Go 1.16 or higher
- Colly library (`github.com/gocolly/colly/v2`)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/wikipedia-web-scraper.git
    cd wikipedia-web-scraper
    ```

2. Initialize a Go module and install the dependencies:

    ```sh
    go mod init wikipedia-web-scraper
    go get github.com/gocolly/colly/v2
    ```

## Usage

1. Run the scraper using the following command:

    ```sh
    go run main.go
    ```

2. The scraper will visit each Wikipedia page in the list, extract the text from the `<p>` elements, and save it to `output.jl`.

3. You will see the scraping progress printed in the terminal, along with the URL being scraped.

### Output File

The scraped data is saved in a file named `output.jl`. Each line in this file contains the JSON representation of the page content, including:
- `url`: The URL of the scraped page.
- `text`: The text extracted from the first paragraph of the page.

### Example output (JSON lines format):

```json
{"url": "https://en.wikipedia.org/wiki/Robotics", "text": "Robotics is an interdisciplinary branch of engineering and science that includes mechanical engineering, electronics, computer science, and others."}
{"url": "https://en.wikipedia.org/wiki/Robot", "text": "A robot is a machine—especially one programmable by a computer—capable of carrying out a complex series of actions automatically."}

