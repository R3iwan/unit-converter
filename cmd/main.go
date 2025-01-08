package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/r3iwan/unit-converter/pkg"
)

type ConversionResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/convert", convertHandler)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	valueStr := query.Get("value")
	from := query.Get("from")
	to := query.Get("to")
	conversionType := query.Get("type")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	var result float64
	var errorResponse string

	switch conversionType {
	case "length":
		result, err = pkg.Convert(value, from, to, pkg.LengthFactors)
	case "weight":
		result, err = pkg.Convert(value, from, to, pkg.WeightFactors)
	case "temperature":
		result, err = pkg.ConvertTemperature(value, from, to)
	default:
		http.Error(w, "Invalid conversion type", http.StatusBadRequest)
		return
	}

	if err != nil {
		errorResponse = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConversionResponse{
		Result: result,
		Error:  errorResponse,
	})
}
