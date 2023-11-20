SELECT
			C.idcrop, C.abbreviation,
			IC.name, ARRAY_TO_STRING(IC.color, ',') as color, IC.family, IC.plantingdensity, IC.literspottingsoil, IC.associations,
			RC.water, RC.soil, RC.nutrition, RC.salinity, RC.ph, RC.climate, RC.depth,
			FC.planting, FC.transplant, FC.harvest, FC.cycle,
			FUC.production, FUC.nutrients,
			SC.seed, SC.seedsPerGram, SC.seedLifespan,
			PC.pests, PC.difficulties, PC.care, PC.miscellaneous
		FROM
			meimporta_unpepino.crop AS C
			JOIN meimporta_unpepino.cropinformation  AS IC ON C.idcrop = IC.cropid
			JOIN meimporta_unpepino.croprequirements  AS RC ON C.idcrop = RC.cropid
			JOIN meimporta_unpepino.cropdates AS FC ON C.idcrop = FC.cropid
			JOIN meimporta_unpepino.cropfruit AS FUC ON C.idcrop = FUC.cropid
			JOIN meimporta_unpepino.cropseed AS SC ON C.idcrop = SC.cropid
			JOIN meimporta_unpepino.cropissues AS PC ON C.idcrop = PC.cropid
		WHERE
			(IC.name = $1 OR ($1 = '' AND IC.name IS NOT NULL)) AND
			(ARRAY_TO_STRING(IC.color, ',') = $2 OR ($2 = '' AND IC.color IS NOT NULL)) AND
			(IC.plantingdensity = $3 OR ($3 = '' AND IC.plantingdensity IS NOT NULL)) AND
			(RC.water = $4 OR ($4 = '' AND RC.water IS NOT NULL)) AND
			(RC.soil = $5 OR ($5 = '' AND RC.soil IS NOT NULL)) AND
			(RC.nutrition = $6 OR ($6 = '' AND RC.nutrition IS NOT NULL)) AND
			(RC.salinity = $7 OR ($7 = '' AND RC.salinity IS NOT NULL)) AND
			(FC.cycle = $8 OR ($8 = '' AND FC.cycle IS NOT NULL))