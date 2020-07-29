package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
	"strings"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	fmt.Println("Enter your AoPS username: ")
	var username string
	fmt.Scanln(&username)
	fmt.Println("Enter your AoPS password: ")
	var password string
	fmt.Scanln(&password)

	var aops_ajax string = "https://artofproblemsolving.com/ajax.php"
	login_data := url.Values{}
	login_data.Set("a", "login")
	login_data.Set("username", username)
	login_data.Set("password", password)
	login_data.Set("stay", "true")
	login_request, err := http.NewRequest(http.MethodPost, aops_ajax, strings.NewReader(login_data.Encode()))
	login_request.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
	login_resp, err := client.Do(login_request)
	fmt.Println(login_resp.Status)
	if err != nil {
		fmt.Println(err)
	}

	var aops_web string = "https://artofproblemsolving.com/community"
	request, err := http.NewRequest("GET", aops_web, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
	resp, err := client.Do(request)
	fmt.Println(resp.Status)
	if err != nil {
		fmt.Println("Some error has occured while getting main page!!")
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error while ioutil read resp body")
		}
		f, err := os.Create("aops.html")
		if err != nil {
			fmt.Println("Error while creating output file")
		}
		f.Write(bodyBytes)
		defer f.Close()
	}
}
