--insert_crops_fruit.sql


INSERT INTO meimporta_unpepino.cropfruit(production, nutrients, cropid)
VALUES ($1, $2, $3);
