package repository

func (r *Repository) CheckExistCropId(id int) (bool, error) {
	var count int
	err := r.db.QueryRow(CheckExistCropIdQuery, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
