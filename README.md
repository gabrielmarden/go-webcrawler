# Project go-webcrawler
Just a simple project for a web crawler in Go. The idea is scraping a web page for a specific url for links and keywords inside the page.

## Usage 
The application receives some parameters, as shown below, start scraping webpages for a specific keyword and, in the end, returns a list of links found throughtout the process. This list size can be limited using the max_result constraint. The whole list is print out in console.

To run, first make sure you have Go installed in your machine, and run the code below:

```
go run main.go -url [URL] -keyword [KEYWORD] -max [MAX_RESULT]
```
Notice: The URL and KEYWORD are required, but the MAX_RESULT is optional

## Build and Run
Generate a binary with Go using the following command, but make to pass as argument the correct pameteres, as mentioned above:
```
go build main.go
```

To run the binary, use the following example for Windows:
```
./main.exe -url https://google.com -keyword search -max 10 
```

## Apache
### Apache 2.0 License
[![License](https://img.shields.io/badge/License-Apache_2.0-yellowgreen.svg)](https://opensource.org/licenses/Apache-2.0)  
`[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)`

