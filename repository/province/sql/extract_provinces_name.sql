SELECT PR.nameprovince, PR.climateprovince
FROM meimporta_unpepino.province AS PR
WHERE PR.nameprovince=$1;