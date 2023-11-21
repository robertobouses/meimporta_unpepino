package province

import "github.com/robertobouses/meimporta_unpepino/entity/province"

func (r *RepositoryProvince) InsertProvinces(province province.Province) error {
	_, err := r.insertProvinceStmt.Exec(
		province.NameProvince,
		province.ClimateProvince,
	)
	if err != nil {
		return err
	}
	return nil
}
