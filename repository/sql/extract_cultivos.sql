--extract_cultivos.sql

SELECT
    C.IdCultivo,
    IC.Nombre, IC.Color, IC.Familia, IC.DensidadPlantacion, IC.LitrosTierraMaceta, IC.Asociaciones,
    RC.Agua, RC.Tierra, RC.Nutricion, RC.Salinidad, RC.PH, RC.Clima, RC.Profundidad,
    FC.Siembra, FC.Transplante, FC.Cosecha, FC.Ciclo,
    FUC.Produccion, FUC.Nutrientes,
    SC.Semilla, SC.SemillasGramo, SC.VidaSemilla,
    PC.Plagas, PC.Dificultades, PC.Cuidados, PC.Miscelanea
FROM
    meimporta_unpepino.Cultivo AS C
    JOIN meimporta_unpepino.informacioncultivo AS IC ON C.IdCultivo = IC.CultivoId
    JOIN meimporta_unpepino.requisitoscultivo AS RC ON C.IdCultivo = RC.CultivoId
    JOIN meimporta_unpepino.fechascultivo AS FC ON C.IdCultivo = FC.CultivoId
    JOIN meimporta_unpepino.frutocultivo AS FUC ON C.IdCultivo = FUC.CultivoId
    JOIN meimporta_unpepino.semillacultivo AS SC ON C.IdCultivo = SC.CultivoId
    JOIN meimporta_unpepino.problemascultivo AS PC ON C.IdCultivo = PC.CultivoId;
