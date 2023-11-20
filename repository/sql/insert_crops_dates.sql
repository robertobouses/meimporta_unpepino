--insert_crop_dates.sql


INSERT INTO meimporta_unpepino.cropdates (planting, transplant, harvest, cycle, cropid)
VALUES ($1, $2, $3, $4, $5);

