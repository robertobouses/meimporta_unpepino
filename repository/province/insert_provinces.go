package province

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

func (r *RepositoryProvince) InsertProvinces(province province.Province) error {
	_, err := r.insertProvincesStmt.Exec(
		province.NameProvince,
		province.ClimateProvince,
	)
	if err != nil {
		log.Print("error en InsertProvinces repo", err)
		return err
	}
	return nil
}
