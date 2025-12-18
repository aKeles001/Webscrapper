# Webscrapper

## Overview

Webscrapper is simple web scrapeer that fetch the given URL's HTML page and take a screenshot of the page. Fethed data is saved to Output directory but output directory can be changed with proper commandline variable.

## Features

- Fetched the HTML source of the given URL
- Save the screenshot of the web page
- Save the existing URL's in the HTML source

## Installation

```bash
#git clone REPOSITORY
```

## Usage

```bash
#go run .\Webscrapper.go -url https://www.example.com -out HTML-Filename -outdir HTML-directory -screenshot screenshot.name -screenshotdir screenshor.directory -urlfile extracted.url.filename -urldir extracted.url.directory
```
