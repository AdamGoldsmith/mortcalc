package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"math"
)

const serverPort = 3001

type CalcVariables struct {
	Principal float64    // initial loan
	Interest  float64    // interest rate
	Frequency float64    // daily - 365, weekly - 52, monthly - 12, yearly - 1
	Time      float64    // term in months
}

// Compound method referencing CalcVariables
func (cv CalcVariables) Compound() float64 {
	// A = P(1 + r/n)^nt
	return math.Round((cv.Principal*math.Pow(1+cv.Interest/(100*cv.Frequency), cv.Frequency*cv.Time/12))*100)/100
}

func showCalc(cv CalcVariables) http.Handler {
	sc := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/mortcalc.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, cv)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
	return http.HandlerFunc(sc)
}

func showAmount(a float64) http.Handler {
	sa := func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/mortamount.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, a)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
	return http.HandlerFunc(sa)
}

func main() {
	mux := http.NewServeMux()

	// Order of events - would be great to keep this all on one page (nested templates?)
	// 1. Get CalcVariables from user input forms
	// 2. Submit will populate CalcVariables struct
	// 3. Display calculated amount using Compound method of CalcVariables struct
	// 4. Brilliant to be able to show each frequency payment

	MortCalcVars := CalcVariables{
		Principal: 200000,
		Interest: 6,
		Frequency: 365,
		Time: 12,
	}
	log.Print(MortCalcVars.Compound())
	mc := showCalc(MortCalcVars)
	mux.Handle("/calc", mc)
	sa := showAmount(MortCalcVars.Compound())
	mux.Handle("/amount", sa)

	log.Print("Listening on :", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux)
}

// REFERENCE: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
