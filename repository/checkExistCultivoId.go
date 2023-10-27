package repository

func (r *Repository) CheckExistCultivoId(id int) (bool, error) {
	var count int
	err := r.db.QueryRow(CheckExistCultivoIdQuery, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
