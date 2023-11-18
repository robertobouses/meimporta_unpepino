--insert_dates_crop.sql


INSERT INTO meimporta_unpepino.cropdates (planting     , transplant   , harvest      , cycle        , cropid)
VALUES ($1, $2, $3, $4, $5);


--INSERT INTO meimporta_unpepino.CropDates (Planting     , Transplant   , Harvest      , Cycle        , CropId)
--VALUES ($17, $18, $19, $20, $21);