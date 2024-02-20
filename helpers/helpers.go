package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

var helpersChannel chan int

func PrintHelperMessage() {
	fmt.Println("Hola desde helper.go dentro del paquete helpers!")
}

func RandomTask(id int, data chan int) {
	for taskId := range data {
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d executed task %d\n", id, taskId)
	}
}

func RandomTaskWithParameters(id int, data chan int, nombre string, edad int) {
	for taskId := range data {
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d executed task %d, su nombre es %s y su edad es %d\n", id, taskId, nombre, edad)
	}
}

var Funciones = map[string]interface{}{
	"RandomTask":               RandomTask,
	"RandomTaskWithParameters": RandomTaskWithParameters,
}

func SpawnTask(taskFunction interface{}, workersCount int, executionCount int, args ...interface{}) {
	fmt.Println("Ejecutando SpawnTask con función:", reflect.TypeOf(taskFunction))
	fmt.Println("Número de trabajadores:", workersCount)
	fmt.Println("Número de ejecuciones:", executionCount)
	fmt.Println("Argumentos adicionales:", args)

	// Inicializar canal
	helpersChannel = make(chan int)

	// Canal para coordinar la inicialización de goroutines
	initChannel := make(chan struct{})

	// Convertir la función a reflect.Value
	functionValue := reflect.ValueOf(taskFunction)
	if functionValue.Kind() != reflect.Func {
		fmt.Println("El primer argumento debe ser una función")
		return
	}

	// Verificar la aridad de la función
	expectedNumArgs := functionValue.Type().NumIn()
	if len(args) < expectedNumArgs-2 {
		fmt.Printf("Número insuficiente de argumentos. Se esperaban al menos %d, pero se proporcionaron %d\n", expectedNumArgs-2, len(args))
		return
	}

	// Preparar los argumentos para llamar a la función
	var reflectArgs []reflect.Value
	// Los dos primeros argumentos (workerID y helpersChannel) siempre se agregan
	reflectArgs = append(reflectArgs, reflect.ValueOf(0)) // workerID se actualizará en cada iteración del bucle
	reflectArgs = append(reflectArgs, reflect.ValueOf(helpersChannel))
	// Los argumentos adicionales se agregan si los hay
	for _, arg := range args[:expectedNumArgs-2] {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	// Usar sync.WaitGroup para esperar a que todas las goroutines finalicen
	var wg sync.WaitGroup

	// Iniciar las goroutines
	for i := 0; i < workersCount; i++ {
		go func(workerID int) {
			// Esperar la señal para empezar
			<-initChannel
			defer wg.Done() // Indicar que la goroutine ha terminado su trabajo
			for j := 0; j < executionCount; j++ {
				// Construir los argumentos para llamar a la función
				callArgs := make([]reflect.Value, len(reflectArgs))
				copy(callArgs, reflectArgs)
				callArgs[0] = reflect.ValueOf(workerID)

				// Llamar a la función con los argumentos construidos
				functionValue.Call(callArgs)
			}
		}(i)
		wg.Add(1)
	}

	// Cerrar el canal de inicialización para empezar las goroutines
	close(initChannel)

	// Enviar tareas al canal
	for taskId := 0; taskId < executionCount; taskId++ {
		helpersChannel <- taskId
	}

	// Cerrar el canal después de completar todas las tareas
	close(helpersChannel)

	// Esperar a que todas las goroutines finalicen
	wg.Wait()
}

func MarshalJSON(frase interface{}) (string, error) {
	// Serializar la estructura a JSON
	cadenaJSON, err := json.Marshal(frase)
	if err != nil {
		return "", fmt.Errorf("error al serializar a JSON: %w", err)
	}
	return string(cadenaJSON), nil
}

func UnmarshalJSON(cadenaJSON string) (string, error) {
	// Decodificar la cadena JSON en un map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(cadenaJSON), &data)
	if err != nil {
		return "", fmt.Errorf("error al decodificar JSON: %w", err)
	}

	// Inicializar una lista para almacenar los valores de los campos
	var valores []string

	// Recorrer los campos del mapa
	for _, v := range data {
		// Verificar si el valor es una cadena
		if valor, ok := v.(string); ok {
			// Agregar el valor a la lista
			valores = append(valores, valor)
		}
	}

	// Crear la frase requerida uniendo los valores de los campos
	fraseRequerida := ""
	if len(valores) > 0 {
		fraseRequerida = valores[0]
		for i := 1; i < len(valores); i++ {
			fraseRequerida += " " + valores[i]
		}
	}

	return fraseRequerida, nil
}
