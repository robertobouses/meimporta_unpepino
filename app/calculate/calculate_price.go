package calculate

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CalculatePrice() (float64, error) {
	apiUrl := "https://www.random.org/integers/?num=1&min=1&max=100&col=1&base=10&format=plain&rnd=new"

	response, err := http.Get(apiUrl)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var randomNumber float64
	_, err = fmt.Sscanf(string(body), "%f", &randomNumber)
	if err != nil {
		return 0, err
	}

	return randomNumber, nil
}

//no hay api que proporcione datos de subasta, establezco una api random
