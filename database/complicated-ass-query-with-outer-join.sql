SELECT zeiten_100_200.id, zeiten_100_200.kfz_variante, zeiten_100_200.nass, zeiten_100_200.gemessene_zeit, kfz_variante.id, kfz_variante.serie_ab_werk, kfz_variante.ps, kfz_variante.nm, kfz_variante.tuning, kfz_variante.serien_kfz, basis_kfz.id, basis_kfz.kfz_name, tuning.id, tuning.serien_kfz, tuning.tuning_name FROM
zeiten_100_200 INNER JOIN kfz_variante
ON zeiten_100_200.kfz_variante =  kfz_variante.id INNER JOIN basis_kfz
ON kfz_variante.serien_kfz = basis_kfz.id LEFT OUTER JOIN tuning 
ON kfz_variante.tuning = tuning.id;