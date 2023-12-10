package calculate

func CalculateProduction(climate, soil string, production int) int {
	result := production

	if climate == "mediterraneo" || climate == "tropical" {
		result = int(float64(result) * 1.20)
	} else if climate == "atlantico" {
		result = int(float64(result) * 1.07)
	} else if climate == "montaña" {
		result = int(float64(result) * 1.04)
	} else if climate == "continental" {
		result = int(float64(result) * 1.01)
	}

	if soil == "magra" {
		result = int(float64(result) * 1.3)
	} else if soil == "media" {
		result = int(float64(result) * 1.05)
	} else if soil == "pobre" {
		result = int(float64(result) * 0.8)
	} else if soil == "muy pobre" {
		result = int(float64(result) * 0.6)
	}

	return result
}
