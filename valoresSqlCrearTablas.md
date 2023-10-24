


CREATE SCHEMA meimporta_unpepino;




-- Tabla 'Cultivo'
CREATE TABLE meimporta_unpepino.Cultivo (
    IdCultivo serial PRIMARY KEY
);

-- Tabla 'InformacionCultivo'
CREATE TABLE meimporta_unpepino.InformacionCultivo (
    Nombre text,
    Color text[],
    Familia text,
    DensidadPlantacion text,
    LitrosTierraMaceta integer,
    Asociaciones text[], 
    CultivoId serial,
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);

-- Tabla 'RequisitosCultivo'
CREATE TABLE meimporta_unpepino.RequisitosCultivo (
    Agua text,
    Tierra text,
    Nutricion text,
    Salinidad text,
    PH real[],
    Clima text[], 
    Profundidad text,
    CultivoId serial, 
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);

-- Tabla 'FechasCultivo'
CREATE TABLE meimporta_unpepino.FechasCultivo (
    Siembra text,
    Transplante text,
    Cosecha text,
    Ciclo text,
    CultivoId serial, 
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);

-- Tabla 'FrutoCultivo'
CREATE TABLE meimporta_unpepino.FrutoCultivo (
    Produccion text,
    Nutrientes text,
    CultivoId serial, 
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);

-- Tabla 'SemillaCultivo'
CREATE TABLE meimporta_unpepino.SemillaCultivo (
    Semilla text,
    SemillasGramo text,
    VidaSemilla text,
    CultivoId serial, 
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);

-- Tabla 'ProblemasCultivo'
CREATE TABLE meimporta_unpepino.ProblemasCultivo (
    Plagas text,
    Dificultades text,
    Cuidados text,
    Miscelanea text,
    CultivoId serial, 
    FOREIGN KEY (CultivoId) REFERENCES meimporta_unpepino.Cultivo(IdCultivo)
);
