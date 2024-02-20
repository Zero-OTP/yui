package lexicom // Define el paquete "lexicom"
import (
	"github.com/fatih/color"
)

var particle_types = [...]string{"a", "b", "c"}

// Función para mostrar las particulas disponibles
//func ShowParticleTypes(id int, data chan int) {
//fmt.Printf("Particulas disponibles >> %+q\n", particle_types)
//	color.Green("Particulas disponibles >> %q\n", particle_types)
//}

func ShowParticleTypes(id int, data chan int) {
	for range data {
		color.Green("Particulas disponibles >> %q\n", particle_types)
	}
}
