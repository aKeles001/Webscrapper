# Webscraper

## Overview

Webscraper is simple web scrapeer that fetch the given URL's HTML page and take a screenshot of the page. Fethed data is saved to Output directory but output directory can be changed with proper commandline arguments.

## Features

- Fetched the HTML source of the given URL
- Save the screenshot of the web page
- Save the extracted URL's from the HTML source

## Installation

```bash
git clone https://github.com/aKeles001/Webscraper.git
```

## Usage

```bash
go run .\Webscraper.go -h
```

## Arguments

```bash
-url             -> Target URL
-out             -> HTML output filename (Default hostname.html)
-outdir          -> HTML output directory (Default Outputs/HTML)
-screenshot      -> Screenshot file name (Default hostname.png)
-screenshotdir   ->Screenshot output directory (Default Outputs/Screenshots)
-urldir          -> Extracted url output directory (Default Outputs/URLs)
-urlfile         -> Extracted url file name (Default hostname_urls.txt)
-h               ->Usage
```
