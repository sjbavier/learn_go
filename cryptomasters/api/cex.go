package api

import (
	"cryptomasters/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) { // *datatypes.Rate is a pointer and a pointer can be nil
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters minimum, %v", len(currency))
	}
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))

	if err != nil {
		fmt.Printf("error, %v", err)
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body) // res.Body is a chunked stream so we are waiting for it all syncronously
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response) // using json struct map
		if err != nil {
			fmt.Println("error")
		}

	} else {
		// bad status
		return nil, fmt.Errorf("status code received %v", res.StatusCode) // can use errors. package for creating errors as well
	}
	// fmt.Println(res)

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
