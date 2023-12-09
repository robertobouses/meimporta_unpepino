package calculate

import (
	"time"
)

func CalculateHarvest(transplantDate time.Time, family string, climate string) time.Time {
	result := transplantDate

	switch family {
	case "liliaceas":
		result = result.Add(120 * 24 * time.Hour)
	case "umberifelas":
		result = result.Add(40 * 24 * time.Hour)
	case "solanaceas":
		result = result.Add(63 * 24 * time.Hour)
	}

	if climate == "mediterraneo" || climate == "tropical" {
		result = result.Add(-7 * 24 * time.Hour)
	} else if climate == "atlantico" {
		result = result.Add(1 * 24 * time.Hour)
	} else if climate == "montaña" {
		result = result.Add(7 * 24 * time.Hour)
	}

	if result.Month() >= time.April && result.Month() <= time.July {
		if result.Day() > 8 {
			result = result.Add(-8 * 24 * time.Hour)
		}
	}
	if result.Month() == time.March && result.Month() <= time.September {
		if result.Day() > 5 {
			result = result.Add(-5 * 24 * time.Hour)
		}
	}
	if result.Month() == time.October {
		if result.Day() > 2 {
			result = result.Add(10 * 24 * time.Hour)
		}
	}

	if result.Month() >= time.November {
		if result.Day() > 5 {
			result = result.Add(25 * 24 * time.Hour)
		}
	}
	minimum := 32 * 24 * time.Hour
	maximum := 150 * 24 * time.Hour

	if result.Before(transplantDate.Add(minimum)) {
		result = transplantDate.Add(minimum)
	} else if result.After(transplantDate.Add(maximum)) {
		result = transplantDate.Add(maximum)
	}

	return result
}
