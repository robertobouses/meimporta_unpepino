SELECT DISTINCT
MC.idmycrop, MC.name, MC.planting
FI.squaremeters, FI.soil, FI.climate

FROM
    meimporta_unpepino.mycrops AS MC
    JOIN meimporta_unpepino.fields AS FI ON MC.idcrop = FI.cropid   
WHERE    
    MC.name=$1


