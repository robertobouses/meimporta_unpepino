--insert_requirements_crop.sql


INSERT INTO meimporta_unpepino.croprequirements  (water, soil, nutrition, salinity, ph, clima, profundidad, cropid)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);


-- INSERT INTO meimporta_unpepino.CropRequirements  (Water, Soil, Nutrition, Salinity, PH, Clima, Profundidad, CropId)
-- VALUES ($9, $10, $11, $12, $13, $14, $15, $16);