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
    JOIN meimporta_unpepino.InformacionCultivo AS IC ON C.IdCultivo = IC.CultivoId
    JOIN meimporta_unpepino.RequisitosCultivo AS RC ON C.IdCultivo = RC.CultivoId
    JOIN meimporta_unpepino.FechasCultivo AS FC ON C.IdCultivo = FC.CultivoId
    JOIN meimporta_unpepino.FrutoCultivo AS FUC ON C.IdCultivo = FUC.CultivoId
    JOIN meimporta_unpepino.SemillaCultivo AS SC ON C.IdCultivo = SC.CultivoId
    JOIN meimporta_unpepino.ProblemasCultivo AS PC ON C.IdCultivo = PC.CultivoId;
