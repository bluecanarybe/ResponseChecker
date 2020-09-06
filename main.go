package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

// function to get http status
func httpstatus(url string) string {

	// initalize tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	w.Init(os.Stdout, 70, 0, 1, ' ', tabwriter.AlignRight)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	response, error := http.Get(url)

	if error != nil {
		// fmt.Println(url, " - domain not available")
		fmt.Fprintln(w, url, "\t", "Request error\t")
		w.Flush()
		return ("Domain available!")
	} else {
		//fmt.Println(url, " - HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))
		fmt.Fprintln(w, url, "\t", response.StatusCode, http.StatusText(response.StatusCode), "\t")
		w.Flush()
		return ("Domain not available!")
	}
}

func main() {

	// throw some ascii art as intro

	asciiArt :=
		`
		_____                                              _____  _                  _               
		|  __ \                                            / ____|| |                | |              
		| |__) | ___  ___  _ __    ___   _ __   ___   ___ | |     | |__    ___   ___ | | __ ___  _ __ 
		|  _  / / _ \/ __|| '_ \  / _ \ | '_ \ / __| / _ \| |     | '_ \  / _ \ / __|| |/ // _ \| '__|
		| | \ \|  __/\__ \| |_) || (_) || | | |\__ \|  __/| |____ | | | ||  __/| (__ |   <|  __/| |   
		|_|  \_\\___||___/| .__/  \___/ |_| |_||___/ \___| \_____||_| |_| \___| \___||_|\_\\___||_|   
						  | |                                                                         
						  |_|                                                                         
	   
BY BLUECANARY.BE	   
`

	println(asciiArt)

	// initalize tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	w.Init(os.Stdout, 70, 0, 1, ' ', tabwriter.AlignRight)

	// read text file as input
	inputfile := os.Args[1]
	readFile, err := os.Open(inputfile)

	if err != nil {
		log.Fatalf("failed to open file: %s - Please check the path to your input file", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	fmt.Fprintln(w, "URL\t", "HTTP Response Code\t")
	w.Flush()

	// for loop for each line in text file
	for _, eachline := range fileTextLines {
		httpstatus(eachline)
	}

}