package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	model "serverless/src/models"
	service "serverless/src/services"
)

func printNewRequest(request *http.Request) {
	fmt.Printf("\n[ %s ] Path: %s\n", request.RemoteAddr, request.URL.Path)
	if len(request.URL.RawQuery) > 0 {
		fmt.Printf("[ %s ] Query: %s\n", request.RemoteAddr, request.URL.RawQuery)
	}
	if request.Header.Get("Content-Type") != "" {
		fmt.Printf("[ %s ] Content type: %s\n", request.RemoteAddr, request.Header.Get("Content-Type"))
	}
	if request.Header.Get("Accept") != "" {
		fmt.Printf("[ %s ] Accept: %s\n", request.RemoteAddr, request.Header.Get("Accept"))
	}
	if request.Header.Get("User-Agent") != "" {
		fmt.Printf("[ %s ] User agent: %s\n", request.RemoteAddr, request.Header.Get("User-Agent"))
	}
}

func buildRoutes() {
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		printNewRequest(request)

		switch request.URL.Path {
		case "/cardflix/form":
			if request.Method == http.MethodGet {
				var content, err = ioutil.ReadFile("src/public/index.html")
				if err != nil {
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte(err.Error()))
				}
				response.Header().Set("Content-Type", "text/html; charset=utf-8")
				response.Header().Set("Content-Length", fmt.Sprint(len(content)))
				response.WriteHeader(http.StatusOK)
				response.Write(content)
				return
			}

		// Adding Cardflix in database
		case "/cardflix/add":
			if request.Method == http.MethodPost {
				var data model.Cardflix
				// Get the requisition body data
				var err = json.NewDecoder(request.Body).Decode(&data)
				if err != nil {
					// Returns error from reading body data
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					return
				}

				id, err := service.AddCardflix(data)

				// var result, erro = database.AddData("cardflix", data)
				if err != nil {
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					return
				}

				var resData, _ = json.Marshal( /** map[keys type]values type */ map[string]string{
					"id": id,
				})
				response.Header().Set("Content-Type", "application/json; charset=utf-8")
				response.WriteHeader(http.StatusOK)
				response.Write(resData)
			}
			break
		// Getting all Cardflix
		case "/cardflix/all":
			if request.Method == http.MethodGet {
				var data = service.GetAllCardflix()
				var body, err = json.Marshal(data)
				if err != nil {
					// There was a problem converting data to a byte array
					response.Header().Set("Content-Type", "text/plain; charset=utf-8")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte(err.Error()))
					break
				}
				response.Header().Set("Content-Type", "application/json; charset=utf-8")
				response.Header().Set("Content-Length", fmt.Sprint(len(body)))
				response.WriteHeader(http.StatusOK)
				response.Write(body)
			}
			break
		default:
			http.NotFound(response, request)
			break
		}
	})
}
