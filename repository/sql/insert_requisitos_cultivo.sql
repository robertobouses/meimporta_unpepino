--insert_requisitos_cultivo.sql


INSERT INTO meimporta_unpepino.requisitoscultivo (agua, tierra, nutricion, salinidad, ph, clima, profundidad, cultivoid)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);


-- INSERT INTO meimporta_unpepino.RequisitosCultivo (Agua, Tierra, Nutricion, Salinidad, PH, Clima, Profundidad, CultivoId)
-- VALUES ($9, $10, $11, $12, $13, $14, $15, $16);