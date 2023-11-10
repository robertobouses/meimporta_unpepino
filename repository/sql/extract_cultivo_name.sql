SELECT IC.nombre, IC.densidadplantacion
FROM meimporta_unpepino.informacioncultivo AS IC
WHERE IC.nombre = $1;
