package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/converters"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/repository"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/asteroids"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/coordinateCalculation"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/messageParser"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/topsecret", func(context *gin.Context) {

		requestBody := new(models.Satellites)
		context.BindJSON(requestBody)

		context.JSON(200, gin.H{
			"status":   "ok",
			"position": coordinateCalculation.SpaceShipCoordinatesAllSattelites(requestBody.SatellitesUnit),
			"message":  messageParser.MessageParserAll(asteroids.AsteroidInterference(requestBody.SatellitesUnit)),
		})

	})

	router.GET("/topsecret_split/name/:satellite_name/distance/:distance", func(context *gin.Context) {

		var dbItem models.DBparser
		requestBody := context.Request.Body
		x, _ := ioutil.ReadAll(requestBody)
		json.Unmarshal([]byte(x), &dbItem)

		dbItem.Name = context.Param("satellite_name")
		dbItem.Distance, _ = strconv.ParseFloat(context.Param("distance"), 64)

		dbConsult := repository.GetDBitem(dbItem)

		if (models.DBparser{}) == dbConsult {
			context.JSON(404, gin.H{
				"status": "not found",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit((converters.ConverterDBparserToSatellite(dbConsult))),
				"message":  dbConsult.Message,
			})
		}

	})

	router.POST("/topsecret_split/name/:satellite_name", func(context *gin.Context) {
		var satelliteUnit models.SatelliteUnit
		var dbInsert models.DBparser

		requestBody := context.Request.Body
		x, _ := ioutil.ReadAll(requestBody)
		json.Unmarshal([]byte(x), &satelliteUnit)

		if satelliteUnit.Distance != 0 {

			satelliteUnit.Name = context.Param("satellite_name")

			context.JSON(200, gin.H{

				"position": coordinateCalculation.SpaceShipCoordinatesSattelitesUnit(satelliteUnit),
				"message":  messageParser.MessageParserUnit(asteroids.AsteroidInterferenceUnit(satelliteUnit)),
			})

			dbInsert = converters.ConverterSatelliteToDBparser(satelliteUnit)
			repository.PostDBitem(dbInsert)

		} else {
			context.JSON(400, gin.H{
				"status": "distance should not be zero",
			})
		}
	})

	router.Run(":3000")

}
