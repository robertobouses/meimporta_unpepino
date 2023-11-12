package app

import (
	"fmt"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) SearchProblemsCultivo(nameIntro string) ([]entity.ProblemResponse, error) {
	fmt.Println("Buscando problemas de cultivo para:", nameIntro)

	cultivos, err := s.repo.ExtractCultivos()
	if err != nil {
		return nil, err
	}

	threshold := 2
	var results []entity.ProblemResponse

	for _, cultivo := range cultivos {
		distance := LevenshteinDistance(nameIntro, cultivo.InformacionCultivo.Nombre)
		if distance <= threshold {
			result := entity.ProblemResponse{
				Nombre:       cultivo.InformacionCultivo.Nombre,
				Plagas:       cultivo.ProblemasCultivo.Plagas,
				Dificultades: cultivo.ProblemasCultivo.Dificultades,
				Cuidados:     cultivo.ProblemasCultivo.Cuidados,
				Miscelanea:   cultivo.ProblemasCultivo.Miscelanea,
			}
			results = append(results, result)
		}
	}

	fmt.Println("Resultados:", results)

	return results, nil
}
