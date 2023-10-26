


CREATE SCHEMA meimporta_unpepino_pruebas;




CREATE TABLE meimporta_unpepino_pruebas.cultivo (
    idcultivo serial PRIMARY KEY,
    siglas text
);



-- Tabla 'informacioncultivo'
CREATE TABLE meimporta_unpepino_pruebas.informacioncultivo (
    nombre text,
    litrostierraMaceta integer,
     cultivoid integer,
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

--------------------------------------------------------

-- Tabla 'Requisitoscultivo'
CREATE TABLE meimporta_unpepino_pruebas.Requisitoscultivo (
    agua text,
    tierra text,
    nutricion text,
    salinidad text,
    ph real[],
    clima text[], 
    profundidad text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

-- Tabla 'fechascultivo'
CREATE TABLE meimporta_unpepino_pruebas.fechascultivo (
    siembra text,
    transplante text,
    cosecha text,
    ciclo text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

-- Tabla 'frutocultivo'
CREATE TABLE meimporta_unpepino_pruebas.frutocultivo (
    produccion text,
    nutrientes text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

-- Tabla 'semillacultivo'
CREATE TABLE meimporta_unpepino_pruebas.semillacultivo (
    semilla text,
    semillasGramo text,
    vidaSemilla text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

-- Tabla 'problemascultivo'
CREATE TABLE meimporta_unpepino_pruebas.problemascultivo (
    plagas text,
    dificultades text,
    cuidados text,
    miscelanea text,
    cultivoid integer, 
    FOREIGN KEY (cultivoid) REFERENCES meimporta_unpepino_pruebas.cultivo(idcultivo)
);

