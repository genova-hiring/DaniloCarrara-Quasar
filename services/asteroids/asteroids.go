package asteroids

import (
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"math/rand"
	"time"
)

func AsteroidInterference(spaceShipTramission []models.SatelliteUnit) []models.SatelliteUnit {

	var i int
	for _, transmission := range spaceShipTramission {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(len(spaceShipTramission[0].Message)-0+1) + 0
		for count := 0; count < random; count++ {
			transmission.Message[i] = ""
			i++
		}
		i = 0
	}
	return spaceShipTramission

}

func AsteroidInterferenceUnit(spaceShipTramission models.SatelliteUnit) models.SatelliteUnit {
	var i int
	for _, transmission := range spaceShipTramission.Message {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(len(transmission)-0+1) + 0
		for count := 0; count < random; count++ {
			transmission = ""
			i++
		}
		i = 0
	}
	return spaceShipTramission
}
