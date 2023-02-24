package main

import (
	"fmt"
	"log"
	"math"
	"html/template"
	"net/http"
	"strconv"
)

const serverPort = 3001

type CalcVariables struct {
	Principal float64    // initial loan
	Interest  float64    // interest rate
	Term      float64    // term in months
	Frequency int        // daily - 365, weekly - 52, monthly - 12, yearly - 1
}

// Compound method referencing CalcVariables
func (cv CalcVariables) Compound() float64 {
	// A = P(1 + r/n)^nt
	return math.Round((cv.Principal*math.Pow(1+cv.Interest/(100*float64(cv.Frequency)), float64(cv.Frequency)*cv.Term/12))*100)/100
}

type WebVariables struct {
  Calc      CalcVariables
  Amount    float64
}

// Set some default values
var MortCalcVars = CalcVariables{
	Principal: 200000,
	Interest: 6,
	Term: 12,
	Frequency: 365,
}

var WebCalcVars = WebVariables{
	Calc:   MortCalcVars,
	Amount: MortCalcVars.Compound(),
}

func showPage(wv WebVariables) http.Handler {
		sp := func(w http.ResponseWriter, r *http.Request) {
			t, err := template.ParseFiles("templates/main.html")
			if err != nil {
				log.Print("template parsing error: ", err)
			}

			if r.Method == http.MethodPost {
				pval, _ := strconv.ParseFloat(r.FormValue("principal"), 64)
				ival, _ := strconv.ParseFloat(r.FormValue("interest"), 64)
				tval, _ := strconv.ParseFloat(r.FormValue("term"), 64)
				fval, _ := strconv.Atoi(r.FormValue("frequency"))
				var cv = CalcVariables{
					Principal:   pval,
					Interest:    ival,
					Term:        tval,
					Frequency:   fval,
				}
				wv = WebVariables{
					Calc:   cv,
					Amount: cv.Compound(),
				}
			}

			err = t.Execute(w, wv)
			if err != nil {
				log.Print("template executing error: ", err)
			}

			log.Print(wv)

		}
		return http.HandlerFunc(sp)
	}

func main() {

	// Issues & TODOs:
	// 1. Multiple people using this at once can interfere with variables
	// 2. Cannot find ideal way to represent data in separated forms so
	//    added an additional amount field in a separate structure
	// 3. Consider reading port from ENV var

	mux := http.NewServeMux()

	web := showPage(WebCalcVars)
	mux.Handle("/", web)

	log.Print("Listening on :", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux))
}
