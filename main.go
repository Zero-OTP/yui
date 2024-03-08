package main

import (
	"fmt"
	"yui/apicaller"
	"yui/config"
	"yui/database"
	"yui/helpers"
	"yui/imagicom"

	"github.com/fatih/color"
)

// Define los canales
var lexicomChannel chan int

func main() {
	// Cargar la configuración desde el archivo
	if err := config.Load(); err != nil {
		fmt.Println("Error al cargar la configuración:", err)
		return
	}

	cfg := config.GetConfig()

	color.Blue(cfg.API_URL)

	// Inicializar la conexión a la base de datos
	err := database.InitDB(cfg.DB_NAME, cfg.DB_USER, cfg.DB_PASS)
	if err != nil {
		panic(err)
	}

	// Iniciar un canal para lexicom
	lexicomChannel = make(chan int)
	//helpers.SpawnTask(lexicom.ShowParticleTypes, 1, 1)

	// Generar una frase aleatoria
	frase, err := imagicom.GenerarIdea()
	if err != nil {
		panic(err)
	}

	// Imprimir la frase generada
	//fmt.Println("Frase generada:", frase)
	color.Green("Idea generada: %+q", frase)

	// Transformar frase
	cadena, err := helpers.MarshalJSON(frase)

	// Limpiar frase
	fraselimpia, err := helpers.UnmarshalJSON(cadena)

	color.Red(fraselimpia)

	// Realiza la solicitud POST a la API con el cuerpo JSON
	respBody, err := apicaller.GenerateText(cfg.API_URL, fraselimpia, 20)
	if err != nil {
		panic(err)
	}

	// Imprime la respuesta de la API
	color.Magenta("Respuesta de la API:", string(respBody))
}

//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
