package scraper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	
	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	URL string
	SourceName string
	Date string
}

func GetFromURL(url string) (body []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp,err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FAIL : %s", resp.Status)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err.Error())
	}
	return body, nil
}

func ScrapeTableFromURL(toScrapeObj Scraper) (data [][]string, err error) {
	var scrapedData [][]string

	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			data := []string{}

            // Process Headers
            row.ForEach("th", func(_ int, th *colly.HTMLElement) {
                header := th.Text
                header = strings.TrimSpace(header)
                header = strings.ToLower(header)
                header = strings.ReplaceAll(header, ".", " ")
                header = strings.ReplaceAll(header, " ", "_")
                header = strings.ReplaceAll(header, "/", "_")
                data = append(data, header)
                // fmt.Printf("\nData = %v\nHeader = %s\n", scrapedData, header)
            })
			
            // if len(scrapedData) > 0 {
            //     fmt.Println("Headers inserted")
            // }

            // Process Data
            row.ForEach("td", func(_ int, td *colly.HTMLElement) {
                text := td.Text
                text = strings.TrimSpace(text)
                text = strings.ReplaceAll(text, "%", "")
                text = strings.ReplaceAll(text, ".", "")
                text = strings.ReplaceAll(text, ",", ".")
                if text == "NA" {
                    text = "0"
                }
                data = append(data, text)
            })

			scrapedData = append(scrapedData, data)
        })
    })

    err = c.Visit(toScrapeObj.URL)
	if err != nil {
		return nil ,fmt.Errorf("error scraping %s: %s", toScrapeObj.SourceName, err.Error())
	}
	return scrapedData, nil
}
