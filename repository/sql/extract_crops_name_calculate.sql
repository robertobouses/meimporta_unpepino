SELECT IC.name, IC.plantingdensity
FROM meimporta_unpepino.cropinformation  AS IC
WHERE IC.name = $1;
