SELECT IC.name, IC.densidadplantacion
FROM meimporta_unpepino.cropinformation  AS IC
WHERE IC.name = $1;
