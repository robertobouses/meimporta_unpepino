SELECT
    C.idcrop, C.abbreviation,
    IC.name, ARRAY_TO_STRING(IC.color, ',') as color, IC.family, IC.plantingdensity, IC.literspottingsoil, IC.associations,
    RC.water, RC.soil, RC.nutrition, RC.salinity, unnest(RC.ph) as ph, RC.climate, RC.depth,
    FC.planting, FC.transplant, FC.harvest, FC.cycle,
    FUC.production, FUC.nutrients,
    SC.seed, SC.seedsPerGram, SC.seedLifespan,
    PC.pests, PC.difficulties, PC.care, PC.miscellaneous
FROM
    meimporta_unpepino.crop AS C
    JOIN meimporta_unpepino.cropinformation  AS IC ON C.idcrop = IC.cropid
    JOIN meimporta_unpepino.croprequirements  AS RC ON C.idcrop = RC.cropid
    JOIN meimporta_unpepino.cropdates AS FC ON C.idcrop = FC.cropid
    JOIN meimporta_unpepino.cropfruit AS FUC ON C.idcrop = FUC.cropid
    JOIN meimporta_unpepino.cropseed AS SC ON C.idcrop = SC.cropid
    JOIN meimporta_unpepino.cropissues AS PC ON C.idcrop = PC.cropid
WHERE
    IC.name = $1
    OR (ARRAY_TO_STRING(IC.color, ',') IS NULL OR ARRAY_TO_STRING(IC.color, ',') = $2)
    OR (IC.plantingdensity IS NULL OR IC.plantingdensity = $3)
    OR (RC.water IS NULL OR RC.water = $4)
    OR (RC.soil IS NULL OR RC.soil = $5)
    OR (RC.nutrition IS NULL OR RC.nutrition = $6)
    OR (RC.salinity IS NULL OR RC.salinity = $7)
    OR (FC.cycle IS NULL OR FC.cycle = $8);
