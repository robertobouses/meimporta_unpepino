--extract_cultivos.sql

SELECT
    C.idcultivo,
    IC.nombre, IC.litrostierramaceta
FROM
    meimporta_unpepino_pruebas.cultivo AS C
    JOIN meimporta_unpepino_pruebas.informacioncultivo AS IC ON C.idcultivo = IC.cultivoid;
