package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetlatAndLong(query string) (float64, float64) {

	//  do a get request to the following url
	// http://api.positionstack.com/v1/forward?access_key=5307140ab6578c8237159e5cbdf5cfe5&query=query
	//  and get the latitude and longitude

	resp, err := http.Get(`http://api.positionstack.com/v1/forward?access_key=5307140ab6578c8237159e5cbdf5cfe5&query=` + query)
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
