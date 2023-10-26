package repository

import "strings"

func (r *Repository) DeleteCultivosAll() error {
	queries := strings.Split(DeleteCultivosAllQuery, ";")

	for _, query := range queries {
		trimmedQuery := strings.TrimSpace(query)
		if trimmedQuery == "" {
			continue
		}

		_, err := r.db.Exec(trimmedQuery)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (r *Repository) DeleteCultivosAll() error {
// 	queries := strings.Split(DeleteCultivosAllQuery, "\n")

// 	for _, query := range queries {
// 		trimmedQuery := strings.TrimSpace(query)
// 		if trimmedQuery == "" {
// 			continue
// 		}

// 		_, err := r.db.Exec(trimmedQuery)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (r *Repository) DeleteCultivosAll() error {
// 	queries := strings.Split(DeleteCultivosAllQuery, "\n")

// 	for _, query := range queries {
// 		if strings.TrimSpace(query) != "" {
// 			_, err := r.db.Exec(query)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
