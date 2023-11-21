package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProvince(city string) (string, error) {

	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", city)
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	var results []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return "", fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if len(results) > 0 {
		firstResult := results[0]

		for _, addressPart := range firstResult["address"].(map[string]interface{}) {
			if addressPart.(string) != city {
				return fmt.Sprintf("%v", addressPart), nil
			}
		}
	} else {
		return "", fmt.Errorf("no se encontraron resultados para la ciudad: %s", city)
	}

	return "", nil
}
