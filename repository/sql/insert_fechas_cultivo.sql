--insert_fechas_cultivo.sql


INSERT INTO meimporta_unpepino.fechascultivo (siembra, transplante, cosecha, ciclo, cultivoid)
VALUES ($1, $2, $3, $4, $5);


--INSERT INTO meimporta_unpepino.FechasCultivo (Siembra, Transplante, Cosecha, Ciclo, CultivoId)
--VALUES ($17, $18, $19, $20, $21);