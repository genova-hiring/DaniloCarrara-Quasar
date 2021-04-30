package messageParser

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"strings"
)

func MessageParserAll(satellites []models.SatelliteUnit) string {

	messageParser := make([]string, len(satellites[0].Message))
	var i int
	for _, satelliteUnit := range satellites {
		for _, message := range satelliteUnit.Message {
			if message != "" {
				messageParser[i] = message

			}
			i++
		}
		i = 0

	}
	return strings.Join(messageParser, " ")
}

func MessageParserUnit(satellite models.SatelliteUnit) string {

	messageParser := make([]string, len(satellite.Message))

	var i int
	for _, message := range satellite.Message {
		if message != "" && strings.ToLower(satellite.Name) == "kenobi" {
			messageParser[i] = message
		} else if message != "" && strings.ToLower(satellite.Name) == "skywalker" {
			messageParser[i] = message
		} else if message != "" && strings.ToLower(satellite.Name) == "sato" {
			messageParser[i] = message
		}
		i++

	}
	return strings.Join(messageParser, " ")

}
