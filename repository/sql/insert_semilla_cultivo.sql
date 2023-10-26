--insert_semilla_cultivo.sql

INSERT INTO meimporta_unpepino.semillacultivo (semilla, semillasgramo, vidasemilla, cultivoid)
VALUES ($1, $2, $3, $4);


-- INSERT INTO meimporta_unpepino.SemillaCultivo (Semilla, SemillasGramo, VidaSemilla, CultivoId)
-- VALUES ($25, $26, $27, $28);