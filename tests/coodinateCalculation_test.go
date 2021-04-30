package tests

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/coordinateCalculation"
	"testing"
)

var sattelitesTest = []models.SatelliteUnit{

	{Name: "kenobi",
		Distance: 100,
		Message:  []string{"mensagem", "do", "kenobi"},
	},

	{Name: "skywalker",
		Distance: 200,
		Message:  []string{"mensagem", "do", "skywalker"},
	},

	{Name: "sato",
		Distance: 300,
		Message:  []string{"mensagem", "do", "sato"},
	}}

var satelliteUnitTest = models.SatelliteUnit{
	Name:     "skywalker",
	Distance: 200,
	Message:  []string{"mensagem", "do", "skywalker"},
}

var shipCoordinaAllSatellitesTest = models.ShipCoordinates{
	X: "-60.21",
	Y: "22.48",
}

var shipCoordinaSatelliteUnitTest = models.ShipCoordinates{
	X: "170.18",
	Y: "105.06",
}

func TestSpaceShipCoordinatesAllSattelites(t *testing.T) {

	shipFunc := coordinateCalculation.SpaceShipCoordinatesAllSattelites(sattelitesTest)

	if (shipFunc) != shipCoordinaAllSatellitesTest {
		t.Error("Erro no Calculo de Coordenadas!!")
	}

}

func TestSpaceShipCoordinatesSattelitesUnit(t *testing.T) {

	shipFunc := coordinateCalculation.SpaceShipCoordinatesSattelitesUnit(satelliteUnitTest)
	if (shipFunc) != shipCoordinaSatelliteUnitTest {
		t.Error("Calculo de coordenadas errado")
	}

}
