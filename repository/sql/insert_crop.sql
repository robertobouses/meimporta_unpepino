--insert_crop.sql

INSERT INTO meimporta_unpepino.crop (abbreviation) VALUES ($1) RETURNING idcrop;