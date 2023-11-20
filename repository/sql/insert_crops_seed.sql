--insert_seed_crop.sql

INSERT INTO meimporta_unpepino.cropseed(seed, seedspergram, seedlifespan, cropid)
VALUES ($1, $2, $3, $4);
