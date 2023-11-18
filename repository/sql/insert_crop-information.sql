--insert_information_crop.sql


INSERT INTO meimporta_unpepino.cropInformation  (name, color, family, densidadplantacion, literspottingsoil  , associations       , cropid)
VALUES ($1, $2, $3, $4, $5, $6, $7);



-- INSERT INTO meimporta_unpepino.CropInformation  (Name, Color, Family, DensidadPlantacion, LitersPottingSoil  , Associations       , CropId)
-- VALUES ($2, $3, $4, $5, $6, $7, $8);