--insert_cultivo.sql

INSERT INTO meimporta_unpepino.cultivo (siglas) VALUES ($1) RETURNING idcultivo;