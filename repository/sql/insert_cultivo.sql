--insert_cultivo.sql

INSERT INTO meimporta_unpepino_pruebas.cultivo (siglas) VALUES ($1) RETURNING idcultivo;


