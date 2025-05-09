package calculate

import (
	"time"
)

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
		result = result.Add(-7 * 24 * time.Hour)
	} else if climate == "atlantico" {
	} else if climate == "montaña" {
		result = result.Add(7 * 24 * time.Hour)
	}

	if result.Month() >= time.March && result.Month() <= time.July {
		if result.Day() > 8 {
			result = result.Add(-8 * 24 * time.Hour)
		}
	}
	if result.Month() == time.February && result.Month() <= time.September {
		if result.Day() > 5 {
			result = result.Add(-5 * 24 * time.Hour)
		}
	}
	if result.Month() == time.October {
		if result.Day() > 2 {
			result = result.Add(2 * 24 * time.Hour)
		}
	}

	if result.Month() >= time.November {
		if result.Day() > 5 {
			result = result.Add(16 * 24 * time.Hour)
		}
	}
	minimum := 32 * 24 * time.Hour
	maximum := 150 * 24 * time.Hour

	if result.Before(plantingDate.Add(minimum)) {
		result = plantingDate.Add(minimum)
	} else if result.After(plantingDate.Add(maximum)) {
		result = plantingDate.Add(maximum)
	}

	return result
}
