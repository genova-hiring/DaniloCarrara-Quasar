package tests

import (
	"fmt"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"
	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/services/messageParser"
	"testing"
)

var sattelitesTestmessage = []models.SatelliteUnit{

	{Name: "kenobi",
		Distance: 100,
		Message:  []string{"", "de", "teste"},
	},

	{Name: "skywalker",
		Distance: 200,
		Message:  []string{"mensagem", "", "teste"},
	},

	{Name: "sato",
		Distance: 300,
		Message:  []string{"", "de", ""},
	}}

var satelliteUnitTestMessage = models.SatelliteUnit{
	Name:     "skywalker",
	Distance: 200,
	Message:  []string{"mensagem", "do", "skywalker"},
}

func TestMessageParserAll(t *testing.T) {

	stringExpected := "mensagem de teste"
	funcParser := messageParser.MessageParserAll(sattelitesTestmessage)

	if messageParser.MessageParserAll(sattelitesTestmessage) != stringExpected {
		t.Error("Erro no parser de Mensagens!!")
		fmt.Errorf("valor esperado: %s\nvalor recebido %s: \n\n", stringExpected, funcParser)
	}

}

func TestMessageParserUnit(t *testing.T) {
	stringExpected := "mensagem do skywalker"
	funcParser := messageParser.MessageParserUnit(satelliteUnitTestMessage)

	if messageParser.MessageParserUnit(satelliteUnitTestMessage) != stringExpected {
		t.Error("Erro no parser de Mensagens!!")
		t.Errorf("valor esperado: %s\nvalor recebido: %s \n\n", stringExpected, funcParser)
	}

}
