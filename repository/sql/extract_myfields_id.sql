SELECT DISTINCT
FI.idmyfield, FI.catastralcode, FI.name, FI.squaremeters, FI.soil, FI.city, FI.province, FI.climate 

FROM
meimporta_unpepino.fields AS FI

WHERE
FI.idmyfield=$1;