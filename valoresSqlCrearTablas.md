


CREATE SCHEMA meimporta_unpepino;




CREATE TABLE meimporta_unpepino.cultivo (
    idcultivo serial PRIMARY KEY,
    siglas text
);

-------------------------------------------------------

-- Tabla 'informacioncultivo'
CREATE TABLE meimporta_unpepino.informacioncultivo (
    nombre text,
    color text[],
    familia text,
    densidadplantacion text,
    litrostierramaceta integer,
    asociaciones text[]
    cultivoid integer,
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

CREATE TABLE meimporta_unpepino.informacioncultivo (
    nombre text,
    color text[],
    familia text,
    densidadplantacion text,
    litrostierramaceta integer,
    asociaciones text[],
    cultivoid integer REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);



--------------------------------------------------------

-- Tabla 'Requisitoscultivo'
CREATE TABLE meimporta_unpepino.Requisitoscultivo (
    agua text,
    tierra text,
    nutricion text,
    salinidad text,
    ph real[],
    clima text[], 
    profundidad text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

-- Tabla 'fechascultivo'
CREATE TABLE meimporta_unpepino.fechascultivo (
    siembra text,
    transplante text,
    cosecha text,
    ciclo text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

-- Tabla 'frutocultivo'
CREATE TABLE meimporta_unpepino.frutocultivo (
    produccion text,
    nutrientes text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

-- Tabla 'semillacultivo'
CREATE TABLE meimporta_unpepino.semillacultivo (
    semilla text,
    semillasGramo text,
    vidaSemilla text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

-- Tabla 'problemascultivo'
CREATE TABLE meimporta_unpepino.problemascultivo (
    plagas text,
    dificultades text,
    cuidados text,
    miscelanea text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino.cultivo(idcultivo) ON DELETE CASCADE
);

