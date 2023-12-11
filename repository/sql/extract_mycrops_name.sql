SELECT DISTINCT
MC.idmycrop, MC.name, MC.planting, MC.fieldId

FROM
    meimporta_unpepino.mycrops AS MC
WHERE    
    MC.name=$1;


