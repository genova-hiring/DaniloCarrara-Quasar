package converters

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"strings"
)

func ConverterSatelliteToDBparser(sattelite models.SatelliteUnit) models.DBparser {

	var dbParser = new(models.DBparser)
	dbParser.Name = sattelite.Name
	dbParser.Distance = sattelite.Distance
	dbParser.Message = strings.Join(sattelite.Message, " ")

	return *dbParser

}
