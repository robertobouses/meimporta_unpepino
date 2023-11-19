package repository

func (r *Repository) CheckExistCropsId(id int) (bool, error) {
	var count int
	err := r.db.QueryRow(CheckExistCropsIdQuery, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
