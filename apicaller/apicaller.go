package apicaller

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Options representa las opciones para una solicitud HTTP
type Options struct {
	Method      string            // Método HTTP (GET, POST, PUT, DELETE, etc.)
	URL         string            // URL de la API
	Headers     map[string]string // Cabeceras HTTP opcionales
	Body        interface{}       // Cuerpo de la solicitud (para POST y PUT)
	QueryParams map[string]string // Parámetros de consulta (para GET)
}

// Request realiza una llamada genérica a la API utilizando las opciones proporcionadas
func Request(options Options) ([]byte, error) {
	// Crear una nueva solicitud
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// Establecer el método y la URL
	req.SetRequestURI(options.URL)
	req.Header.SetMethod(options.Method)

	// Agregar cabeceras HTTP opcionales
	for key, value := range options.Headers {
		req.Header.Add(key, value)
	}

	// Agregar parámetros de consulta si son proporcionados
	if options.QueryParams != nil {
		queryArgs := req.URI().QueryArgs()
		for key, value := range options.QueryParams {
			queryArgs.Add(key, value)
		}
	}

	// Agregar cuerpo de la solicitud si es necesario
	if options.Body != nil {
		requestBody, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("error al serializar el cuerpo de la solicitud: %w", err)
		}
		req.SetBody(requestBody)
	}

	// Realizar la solicitud HTTP
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %w", err)
	}

	// Devolver el cuerpo de la respuesta
	return resp.Body(), nil
}

// PostJSON realiza una solicitud POST a la API con un cuerpo JSON
func PostJSON(url string, body interface{}) ([]byte, error) {
	// Serializar el cuerpo JSON
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error al serializar el cuerpo JSON: %w", err)
	}

	// Realizar la solicitud POST a la API
	respBody, err := Request(Options{
		Method: "POST",
		URL:    url,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: requestBody,
	})
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud POST: %w", err)
	}

	return respBody, nil
}

// GenerateText realiza una solicitud POST a la API con un cuerpo JSON
func GenerateText(url string, inputs string, maxNewTokens int) ([]byte, error) {
	// Definir la estructura de la solicitud JSON
	requestBody := struct {
		Inputs     string `json:"inputs"`
		Parameters struct {
			MaxNewTokens int `json:"max_new_tokens"`
		} `json:"parameters"`
	}{
		Inputs: inputs,
		Parameters: struct {
			MaxNewTokens int `json:"max_new_tokens"`
		}{
			MaxNewTokens: maxNewTokens,
		},
	}

	// Serializar el cuerpo JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error al serializar el cuerpo JSON: %w", err)
	}

	// Realizar la solicitud POST a la API
	respBody, err := Request(Options{
		Method:  "POST",
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    requestBodyBytes, // Pasar los bytes directamente
	})
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud POST: %w", err)
	}

	return respBody, nil
}
