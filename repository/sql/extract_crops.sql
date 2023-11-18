--extract_crops.sql

SELECT
    C.idcrop,
    IC.name, IC.color, IC.family, IC.densidadplantacion, IC.literspottingsoil  , IC.associations       ,
    RC.water, RC.soil, RC.nutrition, RC.salinity, RC.ph, RC.clima, RC.profundidad,
    FC.planting     , FC.transplant   , FC.harvest      , FC.cycle        ,
    FUC.production, FUC.nutrients  ,
    SC.seed, SC.seedsPerGram  , SC.seedLifespan,
    PC.pests       , PC.difficulties , PC.care, PC.miscellaneous
FROM
    meimporta_unpepino.crop AS C
    JOIN meimporta_unpepino.cropinformation  AS IC ON C.idcrop = IC.cropid
    JOIN meimporta_unpepino.croprequirements  AS RC ON C.idcrop = RC.cropid
    JOIN meimporta_unpepino.cropdates AS FC ON C.idcrop = FC.cropid
    JOIN meimporta_unpepino.cropfruit AS FUC ON C.idcrop = FUC.cropid
    JOIN meimporta_unpepino.cropseed AS SC ON C.idcrop = SC.cropid
    JOIN meimporta_unpepino.cropissues AS PC ON C.idcrop = PC.cropid;