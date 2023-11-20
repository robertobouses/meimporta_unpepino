--insert_requirements_crop.sql


INSERT INTO meimporta_unpepino.croprequirements(water, soil, nutrition, salinity, ph, climate, depth, cropid)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

