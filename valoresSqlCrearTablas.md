


CREATE SCHEMA meimporta_unpepino;




CREATE TABLE meimporta_unpepino.crop (
    idcrop serial PRIMARY KEY,
    abbreviation text
);

-------------------------------------------------------

-- Tabla 'cropinformation '
CREATE TABLE meimporta_unpepino.cropinformation  (
    name text,
    color text[],
    family text,
    densidadplantacion text,
    literspottingsoil   integer,
    associations        text[]
    cropid integer,
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

CREATE TABLE meimporta_unpepino.cropinformation  (
    name text,
    color text[],
    family text,
    densidadplantacion text,
    literspottingsoil   integer,
    associations        text[],
    cropid integer REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);



--------------------------------------------------------

-- Tabla 'CropRequirements '
CREATE TABLE meimporta_unpepino.CropRequirements  (
    water text,
    soil text,
    nutrition text,
    salinity text,
    ph real[],
    clima text[], 
    profundidad text,
    cropid integer, 
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

-- Tabla 'cropdates'
CREATE TABLE meimporta_unpepino.cropdates (
    planting      text,
    transplant    text,
    harvest       text,
    cycle         text,
    cropid integer, 
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

-- Tabla 'cropfruit'
CREATE TABLE meimporta_unpepino.cropfruit (
    production text,
    nutrients   text,
    cropid integer, 
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

-- Tabla 'cropseed'
CREATE TABLE meimporta_unpepino.cropseed (
    seedtext,
    seedsPerGram   text,
    seedLifespantext,
    cropid integer, 
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

-- Tabla 'cropissues'
CREATE TABLE meimporta_unpepino.cropissues (
    pests        text,
    difficulties  text,
    care text,
    miscellaneous text,
    cropid integer, 
    FOREIGN KEY (cropid) REFERENCES meimporta_unpepino.crop(idcrop) ON DELETE CASCADE
);

