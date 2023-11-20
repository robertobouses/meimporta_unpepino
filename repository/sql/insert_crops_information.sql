--insert_crop_information.sql


INSERT INTO meimporta_unpepino.cropinformation(name, color, family, plantingdensity, literspottingsoil, associations, cropid)
VALUES ($1, $2, $3, $4, $5, $6, $7);
