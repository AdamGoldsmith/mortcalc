package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"math"
	"strconv"
)

const serverPort = 3001

type CalcVariables struct {
	Principal float64    // initial loan
	Interest  float64    // interest rate
	Frequency float64    // daily - 365, weekly - 52, monthly - 12, yearly - 1
	Term      float64    // term in months
}

// Compound method referencing CalcVariables
func (cv CalcVariables) Compound() float64 {
	// A = P(1 + r/n)^nt
	return math.Round((cv.Principal*math.Pow(1+cv.Interest/(100*cv.Frequency), cv.Frequency*cv.Term/12))*100)/100
}

func main() {
	// mux := http.NewServeMux()

	// Order of events - would be great to keep this all on one page (nested templates?)
	// 1. Get CalcVariables from user input forms - DONE
	// 2. Submit will populate CalcVariables struct - DONE
	// 3. Display calculated amount using Compound method of CalcVariables struct
	// 4. Brilliant to be able to show each frequency payment
	// 5. Form value data validation
	// 6. Push out from main into calc function - prepopulate with default vaules in CalcVariables struct

	params := CalcVariables{
		Principal: 200000,
		Interest: 6,
		Frequency: 365,
		Term: 300,
	}
	log.Print(params, params.Compound())

	// gp := func(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFiles("templates/input.html"))

	t, err := template.ParseFiles("templates/main.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(params)
		log.Print(r.Method)
		if r.Method != http.MethodPost {
			t.Execute(w, nil)
			return
		}

		pval, _ := strconv.ParseFloat(r.FormValue("principal"), 64)
		ival, _ := strconv.ParseFloat(r.FormValue("interest"), 64)
		fval, _ := strconv.ParseFloat(r.FormValue("frequency"), 64)
		tval, _ := strconv.ParseFloat(r.FormValue("term"), 64)

		params := CalcVariables{
			Principal:  pval,
			Interest:   ival,
			Frequency:  fval,
			Term:       tval,
		}

		log.Print(params)
		log.Print(params.Compound())

		// t.Execute(w, struct{ Success bool }{true})
		t.Execute(w, nil)
		// t, err := template.ParseFiles("templates/mortcalc.html")
		// if err != nil {
		// 	log.Print("template parsing error: ", err)
		// }
		// err = t.Execute(w, cv)
		// if err != nil {
		// 	log.Print("template executing error: ", err)
		// }
	})

	// mc := showCalc(MortCalcVars)
	// mux.Handle("/calc", mc)
	// sa := showAmount(MortCalcVars.Compound())
	// mux.Handle("/amount", sa)
	// calc := getParams()
	// mux.Handle("/", calc)

	log.Print("Listening on :", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}

// REFERENCE: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
