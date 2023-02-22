package main

import (
	"html/template"
	"log"
	"net/http"
	"math"
)

type CalcVariables struct {
	Principal float64
	Interest  float64
	Frequency float64  // daily - 365, weekly - 52, monthly - 12, yearly - 1
	Time      float64  // term in months
}

// Compound method referencing CalcVariables
func (cv CalcVariables) Compound() float64 {
	// A = P(1 + r/n)^nt
	return math.Round((cv.Principal*math.Pow(1+cv.Interest/(100*cv.Frequency), cv.Frequency*cv.Time/12))*100) / 100
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

func main() {
	mux := http.NewServeMux()

	MortCalcVars := CalcVariables{
		Principal: 200000,
		Interest: 6,
		Frequency: 365,
		Time: 12,
	}
	// log.Print(calculateAmount(MortCalcVars))
	log.Print(MortCalcVars.Compound())
	mc := showCalc(MortCalcVars)
	mux.Handle("/calc", mc)

	log.Print("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}

// REFERENCE: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
