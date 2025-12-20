package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

	// Define command-line flags
	help := flag.Bool("help", false, "Display help")
	
    Rawurl := flag.String("url", "", "URL to fetch")

	out := flag.String("out","", "output HTML filename")
	screenshotF := flag.String("screenshot","", "output Screenshot filename")
	urlF := flag.String("urlfile","", "output URL filename")
	
	outDir := flag.String("outdir", "Outputs/HTML", "output directory")
	sceenshotDir := flag.String("screenshotdir", "Outputs/Screenshots", "output Screenshot directory")
	urlDir := flag.String("urldir", "Outputs/URLs","output URL directory")
	
	flag.Parse()

	// Default filenames
	if *out == "" {
		*out = url_filename_sanitaize(*Rawurl, ".html")
	}
	if *screenshotF == "" {
		*screenshotF = url_filename_sanitaize(*Rawurl, ".png")
	}
	if *urlF == "" {
		*urlF = url_filename_sanitaize(*Rawurl, "_urls.txt")
	}
	// Directory setup
	err := os.MkdirAll(*outDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}
	err = os.MkdirAll(*sceenshotDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating screenshot directory:", err)
		os.Exit(1)
	}
	err = os.MkdirAll(*urlDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating URL directory:", err)
		os.Exit(1)
	}

	// File paths
	outPath := filepath.Join(*outDir, *out)
	screenshotPath := filepath.Join(*sceenshotDir, *screenshotF)
	urlPath := filepath.Join(*urlDir, *urlF)
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if Rawurl == nil || *Rawurl == "" {
		fmt.Println("Please provide a valid URL using the -url flag.")
		os.Exit(1)
	}
	// Data extraction
	var data string
	var screenshot []byte
	response, err := chromedp.RunResponse(ctx,
		chromedp.Navigate(*Rawurl),
	)

	if err != nil {
		fmt.Println("Error navigating to URL:", err)

	}
	if response.Status == 200 {
		chromedp.Run(ctx,
    	chromedp.Navigate(*Rawurl),
    	chromedp.OuterHTML("html", &data, chromedp.ByQuery),
		chromedp.FullScreenshot(&screenshot, 180),
		)
		outFile, err := os.Create(outPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer outFile.Close()
		n, err := io.WriteString(outFile, data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
		fmt.Printf("Wrote %d bytes to %s\n", n, *out)
		screenshotFile, err := os.Create(screenshotPath)
		if err != nil {
			fmt.Println("Error creating Screenshots file:", err)
			os.Exit(1)
		}
		defer screenshotFile.Close()
		n, err = screenshotFile.Write(screenshot)
		if err != nil {
			fmt.Println("Error writing Screenshots to file:", err)
			os.Exit(1)
		}
		fmt.Printf("Wrote %d bytes to %s\n", n, *screenshotF)

		// URL Extraction from DOM
		var Urls []string
		errUrls := chromedp.Run(ctx,
			chromedp.Evaluate(`Array.from(document.querySelectorAll("a")).map(a => a.href)`, &Urls),)
		if errUrls != nil {
			fmt.Println("Error extracting URLs:", errUrls)
		}
		UrlsFile, err := os.Create(urlPath)
		if err != nil {
			fmt.Println("Error creating URL file:", err)
			os.Exit(1)
		}
		defer UrlsFile.Close()
		for _, link := range Urls {
			_, err := UrlsFile.WriteString(link + "\n")
			if err != nil {
				fmt.Println("Error writing URL to file:", err)
			}
		}
		f, err := os.OpenFile("Outputs/scraped_urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening scraped_urls.txt file:", err)
		}
		defer f.Close()
		if _, err := f.WriteString(*Rawurl + "\n"); err != nil {
			fmt.Println("Error writing URL to file:", err)
		}
		fmt.Println(*Rawurl, "fetched successfully.")
	}else{
		fmt.Println("Failed to fetch data, status code:", response.Status)
	}
}


func url_filename_sanitaize(rawurl, ext string) string {
	parsedURL , err := url.Parse(rawurl)
	if err != nil {
		return "invalid_url" + ext
	}

	host := parsedURL.Hostname()
	host = strings.TrimSuffix(host, ".com")
	return host + ext
}
