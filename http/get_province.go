package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetProvince(city string) (string, error) {

	apiURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", url.QueryEscape(city))

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	var result []NominatimResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	if len(result) > 0 {
		province := extractProvince(result[0].DisplayName)
		return province, nil
	}

	return "", fmt.Errorf("no se encontró la ciudad en la respuesta para: %s", city)
}

func extractProvince(displayName string) string {
	parts := strings.Split(displayName, ",")
	if len(parts) > 2 {
		return strings.TrimSpace(parts[len(parts)-2])
	}
	return ""
}

type NominatimResponse struct {
	DisplayName string `json:"display_name"`
}
