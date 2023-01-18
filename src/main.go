package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getlatAndLong(query string) (float64, float64) {

	//  do a get request to the following url
	//  and get the latitude and longitude

	resp, err := http.Get(`http://api.positionstack.com/v1/forward?access_key=%s&query=`+query, key)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return 0, 0
}

func main() {
	getlatAndLong("1075 Lake Washington Blvd E, Seattle, WA 98112")

}
