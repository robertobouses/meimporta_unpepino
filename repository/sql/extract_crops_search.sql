SELECT DISTINCT
    C.idcrop, C.abbreviation,
    IC.name, ARRAY_TO_STRING(IC.color, ',') as color, IC.family, IC.plantingdensity, IC.literspottingsoil, IC.associations,
    RC.water, RC.soil, RC.nutrition, RC.salinity, unnest(RC.ph) as ph, RC.climate, RC.depth,
    FC.planting, FC.transplant, FC.harvest, FC.cycle,
    FUC.production, FUC.nutrients,
    SC.seed, SC.seedsPerGram, SC.seedLifespan,
    PC.pests, PC.difficulties, PC.care, PC.miscellaneous
FROM
    meimporta_unpepino.crop AS C
    JOIN meimporta_unpepino.cropinformation AS IC ON C.idcrop = IC.cropid
    JOIN meimporta_unpepino.croprequirements AS RC ON C.idcrop = RC.cropid
    JOIN meimporta_unpepino.cropdates AS FC ON C.idcrop = FC.cropid
    JOIN meimporta_unpepino.cropfruit AS FUC ON C.idcrop = FUC.cropid
    JOIN meimporta_unpepino.cropseed AS SC ON C.idcrop = SC.cropid
    JOIN meimporta_unpepino.cropissues AS PC ON C.idcrop = PC.cropid
WHERE    
    (IC.name=$1 OR $2) AND
    (ARRAY_TO_STRING(IC.color, ',')=$3 OR $4) AND
    (IC.plantingdensity=$5 OR $6) AND
    (RC.water=$7 OR $8) AND
    (RC.soil=$9 OR $10) AND
    (RC.nutrition=$11 OR $12) AND
    (RC.salinity=$13 OR $14) AND
    (FC.cycle=$15 OR $16);
