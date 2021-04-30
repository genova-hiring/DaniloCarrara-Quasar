package coordinateCalculation

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"math"
	"strconv"
	"strings"
)

func SpaceShipCoordinatesAllSattelites(satellites []models.SatelliteUnit) models.ShipCoordinates {

	var shipCoordinates models.ShipCoordinates
	satoAngle := 30.00
	skywalkerAngle := 45.00
	kenobiAngle := 10.00

	type kenobiCoordinate struct {
		x float64
		y float64
	}

	type skywalkerCoordinate struct {
		x float64
		y float64
	}

	type satoCoordinate struct {
		x float64
		y float64
	}

	var kenobi kenobiCoordinate
	var skywalker skywalkerCoordinate
	var sato satoCoordinate

	var i int
	for _, satelliteUnit := range satellites {
		if strings.ToLower(satellites[i].Name) == "kenobi" {
			kenobi.x = math.Sin(kenobiAngle) * satelliteUnit.Distance
			kenobi.y = math.Cos(kenobiAngle) * satelliteUnit.Distance

		} else if strings.ToLower(satellites[i].Name) == "skywalker" {
			skywalker.x = math.Sin(skywalkerAngle) * satelliteUnit.Distance
			skywalker.y = math.Cos(skywalkerAngle) * satelliteUnit.Distance

		} else if strings.ToLower(satellites[i].Name) == "sato" {
			sato.x = math.Sin(satoAngle) * satelliteUnit.Distance
			sato.y = math.Cos(satoAngle) * satelliteUnit.Distance

		}

		i++
	}

	shipCoordinates.X = strconv.FormatFloat(((kenobi.x + skywalker.x + sato.x) / 3), 'f', 2, 64)
	shipCoordinates.Y = strconv.FormatFloat(((kenobi.y + skywalker.y + sato.y) / 3), 'f', 2, 64)

	return shipCoordinates
}

func SpaceShipCoordinatesSattelitesUnit(satellite models.SatelliteUnit) models.ShipCoordinates {

	var shipCoordinates models.ShipCoordinates
	satoAngle := 30.00
	skywalkerAngle := 45.00
	kenobiAngle := 10.00

	if strings.ToLower(satellite.Name) == "kenobi" {
		kenobiXcoordinate := math.Sin(kenobiAngle) * satellite.Distance
		kenobiYcoordinate := math.Cos(kenobiAngle) * satellite.Distance
		shipCoordinates.X = strconv.FormatFloat((kenobiXcoordinate), 'f', 2, 64)
		shipCoordinates.Y = strconv.FormatFloat((kenobiYcoordinate), 'f', 2, 64)

	} else if strings.ToLower(satellite.Name) == "skywalker" {
		skywalkerXcoordinate := math.Sin(skywalkerAngle) * satellite.Distance
		skywalkerYcoordinate := math.Cos(skywalkerAngle) * satellite.Distance
		shipCoordinates.X = strconv.FormatFloat((skywalkerXcoordinate), 'f', 2, 64)
		shipCoordinates.Y = strconv.FormatFloat((skywalkerYcoordinate), 'f', 2, 64)

	} else if strings.ToLower(satellite.Name) == "sato" {
		satoXcoordinate := math.Sin(satoAngle) * satellite.Distance
		satoYcoordinate := math.Cos(satoAngle) * satellite.Distance
		shipCoordinates.X = strconv.FormatFloat((satoXcoordinate), 'f', 2, 64)
		shipCoordinates.Y = strconv.FormatFloat((satoYcoordinate), 'f', 2, 64)

	}

	return shipCoordinates
}
