package calculate

import (
	"time"
)

func CalculateWater(plantingDate time.Time, family, climate, soil string) []time.Time {
	result := make([]time.Time, 13)

	switch family {
	case "liliaceas":
		result[0] = plantingDate
		result[1] = plantingDate.Add(7 * 24 * time.Hour)
		result[2] = plantingDate.Add(7 * 24 * time.Hour)
		result[3] = plantingDate.Add(20 * 24 * time.Hour)
		result[4] = plantingDate.Add(20 * 24 * time.Hour)
		result[5] = plantingDate.Add(20 * 24 * time.Hour)
		result[6] = plantingDate.Add(14 * 24 * time.Hour)
		result[7] = plantingDate.Add(10 * 24 * time.Hour)
		result[8] = plantingDate.Add(10 * 24 * time.Hour)
		result[9] = plantingDate.Add(7 * 24 * time.Hour)
		result[10] = plantingDate.Add(7 * 24 * time.Hour)
		result[11] = plantingDate.Add(7 * 24 * time.Hour)
		result[12] = plantingDate.Add(7 * 24 * time.Hour)

	case "umberifelas":
		result[0] = plantingDate
		result[1] = plantingDate.Add(2 * 24 * time.Hour)
		result[2] = plantingDate.Add(3 * 24 * time.Hour)
		result[3] = plantingDate.Add(4 * 24 * time.Hour)
		result[4] = plantingDate.Add(3 * 24 * time.Hour)
		result[5] = plantingDate.Add(3 * 24 * time.Hour)

	case "solanaceas":
		result[0] = plantingDate
		result[1] = plantingDate.Add(2 * 24 * time.Hour)
		result[2] = plantingDate.Add(3 * 24 * time.Hour)
		result[3] = plantingDate.Add(5 * 24 * time.Hour)
		result[4] = plantingDate.Add(7 * 24 * time.Hour)
		result[5] = plantingDate.Add(5 * 24 * time.Hour)
		result[6] = plantingDate.Add(7 * 24 * time.Hour)
		result[7] = plantingDate.Add(4 * 24 * time.Hour)
		result[8] = plantingDate.Add(4 * 24 * time.Hour)
		result[9] = plantingDate.Add(4 * 24 * time.Hour)
		result[10] = plantingDate.Add(4 * 24 * time.Hour)
		result[11] = plantingDate.Add(7 * 24 * time.Hour)
		result[12] = plantingDate.Add(7 * 24 * time.Hour)
	}

	if climate == "mediterraneo" || climate == "tropical" {
		for i := range result {
			result[i] = result[i].Add(-1 * 24 * time.Hour)
		}
	} else if climate == "montaña" {
		for i := range result {
			result[i] = result[i].Add(2 * 24 * time.Hour)
		}
	}

	if plantingDate.Month() >= time.June && plantingDate.Month() <= time.August {
		if plantingDate.Day() > 8 {
			for i := range result {
				result[i] = result[i].Add(-2 * 24 * time.Hour)
			}
		}
	}

	if soil == "magra" {
		for i := range result {
			result[i] = result[i].Add(2 * 24 * time.Hour)
		}
	}

	if soil == "pobre" {
		for i := range result {
			result[i] = result[i].Add(-1 * 24 * time.Hour)
		}
	}

	return result
}
