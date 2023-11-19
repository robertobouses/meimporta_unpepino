package repository

import "strings"

func (r *Repository) DeleteCropsAll() error {
	queries := strings.Split(DeleteCropsAllQuery, ";")

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

// func (r *Repository) DeleteCropsAll() error {
// 	queries := strings.Split(DeleteCropsAllQuery, "\n")

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

// func (r *Repository) DeleteCropsAll() error {
// 	queries := strings.Split(DeleteCropsAllQuery, "\n")

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
