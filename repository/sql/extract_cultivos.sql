--extract_cultivos.sql

SELECT
    C.idcultivo,
    IC.nombre, IC.color, IC.familia, IC.densidadplantacion, IC.litrostierramaceta, IC.asociaciones,
    RC.agua, RC.tierra, RC.nutricion, RC.salinidad, RC.ph, RC.clima, RC.profundidad,
    FC.siembra, FC.transplante, FC.cosecha, FC.ciclo,
    FUC.produccion, FUC.nutrientes,
    SC.semilla, SC.semillasGramo, SC.vidaSemilla,
    PC.plagas, PC.dificultades, PC.cuidados, PC.miscelanea
FROM
    meimporta_unpepino.cultivo AS C
    JOIN meimporta_unpepino.informacioncultivo AS IC ON C.idcultivo = IC.cultivoid
    JOIN meimporta_unpepino.requisitoscultivo AS RC ON C.idcultivo = RC.cultivoid
    JOIN meimporta_unpepino.fechascultivo AS FC ON C.idcultivo = FC.cultivoid
    JOIN meimporta_unpepino.frutocultivo AS FUC ON C.idcultivo = FUC.cultivoid
    JOIN meimporta_unpepino.semillacultivo AS SC ON C.idcultivo = SC.cultivoid
    JOIN meimporta_unpepino.problemascultivo AS PC ON C.idcultivo = PC.cultivoid;
