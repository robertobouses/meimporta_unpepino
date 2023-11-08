package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) SearchProblemsCultivo(nameIntro string) ([]entity.Problem, error) {
	cultivos, err := s.repo.ExtractCultivos()
	if err != nil {
		return nil, err
	}

	threshold := 2
	var results []entity.Problem

	for _, cultivo := range cultivos {
		distance := levenshteinDistance(nameIntro, cultivo.InformacionCultivo.Nombre)
		if distance <= threshold {
			result := entity.Problem{
				Nombre:       cultivo.InformacionCultivo.Nombre,
				Plagas:       cultivo.ProblemasCultivo.Plagas,
				Dificultades: cultivo.ProblemasCultivo.Dificultades,
				Cuidados:     cultivo.ProblemasCultivo.Cuidados,
				Miscelanea:   cultivo.ProblemasCultivo.Miscelanea,
			}
			results = append(results, result)
		}
	}

	return results, nil
}

func levenshteinDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		d[i][0] = i
	}
	for j := 0; j <= n; j++ {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)
		}
	}

	return d[m][n]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
