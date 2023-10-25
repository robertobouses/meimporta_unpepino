--insert_informacion_cultivo.sql


INSERT INTO meimporta_unpepino.informacionCultivo (nombre, color, familia, densidadplantacion, litrostierramaceta, asociaciones, cultivoid)
VALUES ($1, $2, $3, $4, $5, $6, $7);



-- INSERT INTO meimporta_unpepino.InformacionCultivo (Nombre, Color, Familia, DensidadPlantacion, LitrosTierraMaceta, Asociaciones, CultivoId)
-- VALUES ($2, $3, $4, $5, $6, $7, $8);
