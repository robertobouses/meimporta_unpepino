package app

import (
	"fmt"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) ProcessCropsIssues(nameIntro string) ([]entity.IssueResponse, error) {
	fmt.Println("Buscando problemas de crop para:", nameIntro)

	crops, err := s.repo.ExtractCrops()
	if err != nil {
		return nil, err
	}

	threshold := 2
	var results []entity.IssueResponse

	for _, crop := range crops {
		distance := LevenshteinDistance(nameIntro, crop.CropInformation.Name)
		if distance <= threshold {
			result := entity.IssueResponse{
				Name:          crop.CropInformation.Name,
				Pests:         crop.CropIssues.Pests,
				Difficulties:  crop.CropIssues.Difficulties,
				Care:          crop.CropIssues.Care,
				Miscellaneous: crop.CropIssues.Miscellaneous,
			}
			results = append(results, result)
		}
	}

	fmt.Println("Resultados:", results)

	return results, nil
}
