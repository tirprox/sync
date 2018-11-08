package moysklad

import (
	"encoding/json"
	"fmt"
	"github.com/tirprox/sync/credentials"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// a struct to hold the Result from each request including an Index
// which will be used for sorting the results after they come in
type Result struct {
	Index int
	Res   *http.Response
	Err   error
}

type Response struct {
	Index int
	Body  []byte
}

type GenericApiResponse struct {
	Meta Meta `json:"meta"`
}

// boundedParallelGet sends requests in parallel but only up to a certain
// limit, and furthermore it's only parallel up to the amount of CPUs but
// is always concurrent up to the concurrency limit
func boundedParallelGet(urls []string, concurrencyLimit int) []Result {

	// this buffered channel will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrencyLimit)

	// this channel will not block and collect the http request results
	resultsChan := make(chan *Result)

	// make sure we close these channels when we're done with them
	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	// keen an Index and loop through every url we will send a request to
	for i, url := range urls {

		// start a go routine with the Index and url in a closure
		go func(i int, url string) {

			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			// send the request and put the response in a Result struct
			// along with the Index so we can sort them later along with
			// any error that might have occoured
			res, err := MakeRequest(url, "GET")
			result := &Result{i, res, err}

			// now we can send the Result struct through the resultsChan
			resultsChan <- result

			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(i, url)
	}

	// make a slice to hold the results we're expecting
	var results []Result

	// start listening for any results over the resultsChan
	// once we get a Result append it to the Result slice
	for {
		result := <-resultsChan
		results = append(results, *result)

		// if we've reached the expected amount of urls then stop
		if len(results) == len(urls) {
			break
		}
	}

	// let's sort these results real quick
	sort.Slice(results, func(i, j int) bool {
		return results[i].Index < results[j].Index
	})

	// now we're done we return the results
	return results
}

//TODO Return pointer maybe?
func GetAll(url string) []Response {

	res, err := Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	firstResponse := Response{}

	firstResponse.Index = 0

	firstResponse.Body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	responses := []Response{}
	responses = append(responses, firstResponse)

	meta := DecodeMeta(firstResponse.Body)

	var urls []string

	for _, v := range makeUrlList(meta.Size) {
		//TODO: Replace with Query Builder function
		urls = append(urls, url+"&"+v)
	}

	if len(urls) > 0 {
		results := boundedParallelGet(urls, 5)
		for _, result := range results {
			index := result.Index + 1
			body, err := ioutil.ReadAll(result.Res.Body)
			if err != nil {
				log.Fatal(err)
			}

			response := Response{index, body}
			responses = append(responses, response)
			defer result.Res.Body.Close()
		}

	}

	return responses
}

func Get(url string) (*http.Response, error) {
	res, err := MakeRequest(url, "GET")
	return res, err
}

var client = &http.Client{Timeout: 60 * time.Second}

func MakeRequest(url string, method string) (res *http.Response, err error) {
	fmt.Println("Requesting " + url)
	req, error1 := http.NewRequest(method, url, nil)
	if error1 != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(credentials.DreamWhite.Login, credentials.DreamWhite.Password)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err = client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	//Decode(Res)

	return res, err
}

func DecodeMeta(responseBody []byte) *Meta {
	response := GenericApiResponse{}

	json.Unmarshal(responseBody, &response)

	return &response.Meta
}

func makeUrlList(size int) []string {
	offset := 100
	limit := 100

	var postfix []string

	//TODO: < size
	for offset < size {
		postfix = append(postfix, "offset="+strconv.Itoa(offset))
		offset += limit
	}

	return postfix

}
