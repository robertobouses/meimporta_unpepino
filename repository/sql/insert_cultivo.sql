--insert_cultivo.sql

BEGIN;

INSERT INTO meimporta_unpepino.Cultivo (IdCultivo) VALUES (?);


INSERT INTO meimporta_unpepino.InformacionCultivo (Nombre, Color, Familia, DensidadPlantacion, LitrosTierraMaceta, Asociaciones, CultivoId) VALUES (?, ?, ?, ?, ?, ?, ?);


INSERT INTO meimporta_unpepino.RequisitosCultivo (Agua, Tierra, Nutricion, Salinidad, Clima, Profundidad, CultivoId) VALUES (?, ?, ?, ?, ?, ?, ?);


INSERT INTO meimporta_unpepino.FechasCultivo (Siembra, Transplante, Cosecha, Ciclo, CultivoId) VALUES (?, ?, ?, ?, ?);


INSERT INTO meimporta_unpepino.FrutoCultivo (Produccion, Nutrientes, CultivoId) VALUES (?, ?, ?);


INSERT INTO meimporta_unpepino.SemillaCultivo (Semilla, SemillasGramo, VidaSemilla, CultivoId) VALUES (?, ?, ?, ?);


INSERT INTO meimporta_unpepino.ProblemasCultivo (Plagas, Dificultades, Cuidados, Miscelanea, CultivoId) VALUES (?, ?, ?, ?, ?);


COMMIT;