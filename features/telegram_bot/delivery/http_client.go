package delivery

import (
	"io/ioutil"
	"net/http"
)

func GetRequests() []byte {
	// URL to make the GET request to
	url := "https://schedule.kpi.ua/api/schedule/lessons?groupId=485a56e2-3b68-41b1-bafb-51dd1b2c2a18"

	// Create an HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Error creating request:")
	}

	resp, err := client.Do(req)
	if err != nil {
		panic("Error sending request:")
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic("Error reading request")
	}

	println(string(bytes))

	return bytes
}
