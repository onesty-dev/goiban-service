package main

import (
	"encoding/json"
	"net/http"

	"github.com/onesty-dev/goiban"
	"github.com/julienschmidt/httprouter"
)

type CalculateSuccess struct {
	Valid bool   `json:"valid"`
	IBAN  string `json:"iban"`
}

type CalculateError struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func calculateIBAN(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	result := goiban.CalculateIBAN(
		ps.ByName("countryCode"),
		ps.ByName("bankCode"),
		ps.ByName("accountNumber"))

	var data []byte
	var err error
	if result.Valid {
		data, err = json.Marshal(CalculateSuccess{true, result.Data})
	} else {
		data, err = json.Marshal(CalculateError{false, result.Message})
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r.Body.Close()
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func calculateAndValidateIBAN(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	iban := goiban.CalculateIBAN(
		ps.ByName("countryCode"),
		ps.ByName("bankCode"),
		ps.ByName("accountNumber"))

	if !iban.Valid {
		data, err := json.Marshal(CalculateError{false, iban.Message})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r.Body.Close()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)
		r.Body.Close()
		return
	}

	param := httprouter.Param{
		Key:   "iban",
		Value: iban.Data,
	}

	r.ParseForm()
	r.Form.Add("validateBankCode", "true")
	r.Form.Add("getBIC", "true")

	// Delegate to validation
	validationHandler(w, r, []httprouter.Param{param})
}
