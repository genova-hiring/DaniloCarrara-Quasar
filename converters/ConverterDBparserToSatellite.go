package converters

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"strings"
)

func ConverterDBparserToSatellite(dbParser models.DBparser) models.SatelliteUnit {

	var sattelite = new(models.SatelliteUnit)
	sattelite.Name = dbParser.Name
	sattelite.Distance = dbParser.Distance
	sattelite.Message = strings.Split(dbParser.Message, " ")

	return *sattelite

}
