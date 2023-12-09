package calculate

import "time"

func CalculateTransplant(plantingDate time.Time, family string, climate string) time.Time {
	result := plantingDate

	switch family {
	case "liliaceas":
		result = result.Add(90 * 24 * time.Hour)
	case "umberifelas":
		result = result.Add(40 * 24 * time.Hour)
	case "solanaceas":
		result = result.Add(63 * 24 * time.Hour)
	}

	if climate == "mediterraneo" || climate == "tropical" {
		result = result.Add(7 * 24 * time.Hour)
	} else if climate == "atlantico" {
		// No action needed for "atlantico"
	} else if climate == "montaña" {
		result = result.Add(7 * 24 * time.Hour)
	}

	return result
}
