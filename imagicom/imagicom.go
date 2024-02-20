package imagicom

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"yui/database"

	"github.com/fatih/color"
)

// Idea representa una idea compuesta por un nombre, un adjetivo y un verbo.
type Idea struct {
	Nombre   string
	Adjetivo string
	Verbo    string
	Genero   string // Nuevo campo
	Numero   string // Nuevo campo
}

// IdeasBase contiene las ideas base disponibles para generar ideas.
var IdeasBase = []string{
	"nombre adjetivo",
	"nombre adjetivo verbo",
	"nombre verbo",
}

// Palabra representa una palabra con sus atributos.
type Palabra struct {
	Palabra   string
	Tipo      string
	Genero    string
	Numero    string
	Raiz      string
	Afijo     string
	Tonica    string
	Silabas   string
	Locale    string
	Origen    string
	Sinonimos string
	// Otros campos...
}

// GenerarIdea genera una nueva idea a partir de las palabras base proporcionadas.
func GenerarIdea() (Idea, error) {
	// Seleccionar una idea base aleatoria de IdeasBase
	ideaBase := SeleccionarIdeaBaseAleatoria()
	color.Cyan(ideaBase)

	// Separar la idea base en sus componentes individuales
	componentes := strings.Split(ideaBase, " ")

	idea := Idea{}

	for _, palabra := range componentes {
		switch palabra {
		case "nombre":
			palabraNombre := SeleccionarPalabraPorTipo("nombre")
			idea.Nombre = palabraNombre.Palabra
			idea.Genero = palabraNombre.Genero // Nuevo campo
			idea.Numero = palabraNombre.Numero // Nuevo campo

		case "adjetivo":
			adjetivo, err := SeleccionarPalabraPorTipoYNumeroYGenero("adjetivo", idea.Numero, idea.Genero) // Utilizamos el género del nombre
			if err != nil {
				return Idea{}, err
			}
			idea.Adjetivo = adjetivo.Palabra

		case "verbo":
			verbo, err := SeleccionarPalabraPorTipoYNumeroYGenero("verbo", idea.Numero, idea.Genero) // Utilizamos el número del nombre
			if err != nil {
				return Idea{}, err
			}
			idea.Verbo = verbo.Palabra
		}
	}

	return idea, nil
}

// SeleccionarIdeaBaseAleatoria selecciona una idea base aleatoria de IdeasBase.
func SeleccionarIdeaBaseAleatoria() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return IdeasBase[r.Intn(len(IdeasBase))]
}

// SeleccionarPalabraPorTipo selecciona aleatoriamente una palabra de la base de datos del tipo especificado.
func SeleccionarPalabraPorTipo(tipo string) Palabra {
	db := database.GetDB()

	var palabra Palabra
	//debug := fmt.Sprintf("%s%%", tipo)
	//color.Green(debug)
	query := fmt.Sprintf("SELECT * FROM palabras WHERE tipo LIKE '%s%%' AND genero  <> '' AND numero  <> '' ORDER BY RANDOM() LIMIT 1", tipo)
	color.Red(query)
	row := db.QueryRow(query)
	err := row.Scan(&palabra.Palabra, &palabra.Tipo, &palabra.Genero, &palabra.Numero, &palabra.Raiz, &palabra.Afijo, &palabra.Tonica, &palabra.Silabas, &palabra.Locale, &palabra.Origen, &palabra.Sinonimos)
	if err != nil {
		// Manejar el error apropiadamente
		return Palabra{} // Devolver el error
	}

	return palabra
}

/*
// SeleccionarPalabraPorTipoYGenero selecciona aleatoriamente una palabra del tipo especificado y con el género del nombre proporcionado.
func SeleccionarPalabraPorTipoYGenero(tipo string, genero string) (Palabra, error) {
	debug := fmt.Sprintf("Genero: %s", genero)
	color.Red(debug)
	db := database.GetDB()
	var palabra Palabra

	query := fmt.Sprintf("SELECT palabra FROM palabras WHERE tipo LIKE '%s%%' AND genero = ? ORDER BY RANDOM() LIMIT 1", tipo)
	row := db.QueryRow(query, genero)
	err := row.Scan(&palabra.Palabra)
	if err != nil {
		return Palabra{}, err
	}

	return palabra, nil
}

// SeleccionarPalabraPorTipoYNumero selecciona aleatoriamente una palabra del tipo especificado y con el número del nombre proporcionado.
func SeleccionarPalabraPorTipoYNumero(tipo string, numero string) (Palabra, error) {
	db := database.GetDB()
	var palabra Palabra

	query := fmt.Sprintf("SELECT palabra FROM palabras WHERE tipo LIKE '%s%%' AND numero = ? ORDER BY RANDOM() LIMIT 1", tipo)
	row := db.QueryRow(query, numero)
	err := row.Scan(&palabra.Palabra)
	if err != nil {
		return Palabra{}, err
	}

	return palabra, nil
}
*/

// SeleccionarPalabraPorTipoYNumeroYGenero selecciona aleatoriamente una palabra del tipo especificado,
// con el género del nombre proporcionado y el número del nombre proporcionado.
func SeleccionarPalabraPorTipoYNumeroYGenero(tipo, numero string, genero string) (Palabra, error) {
	db := database.GetDB()
	var palabra Palabra

	query := fmt.Sprintf("SELECT palabra FROM palabras WHERE tipo LIKE '%s%%' AND numero = $1 AND genero = $2 ORDER BY RANDOM() LIMIT 1", tipo)
	color.Yellow(query)

	row := db.QueryRow(query, numero, genero)
	err := row.Scan(&palabra.Palabra)
	if err != nil {
		return Palabra{}, err
	}

	return palabra, nil
}
