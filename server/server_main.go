package main

import (
    "flag"
    "html/template"
    "net/http"
    "fmt"
    "github.com/codegangsta/negroni"
    "./minmax"
    "./worm"
    "./palindrome"
    "strconv"
)

// Template for each request
var minmax_templ = template.Must(template.New("mm").Parse(minmax.MinmaxTemplateStr()))
var worm_templ = template.Must(template.New("mc").Parse(worm.WormTemplateStr()))
var palindrome_templ = template.Must(template.New("pl").Parse(palindrome.PalindromeTemplateStr()))

func main() {
    flag.Parse()
    
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
        fmt.Fprintf(w, "Welcome to the homepage!<br/>Choose a page")
    })

    mux.HandleFunc("/minmax/", http.HandlerFunc(MinMax))            // MinMax - Soal nomor 1
    mux.HandleFunc("/worm/", http.HandlerFunc(Worm))                // Worm - Soal nomor 2
    mux.HandleFunc("/palindrome/", http.HandlerFunc(Palindrome))    // Palindrome - Soal nomor 3

    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":8080")
}

/*
 * Running function on access to "minmax" url location
 * Executing template first then returning result(s)
 */
func MinMax(w http.ResponseWriter, req *http.Request) {

    minmax_templ.Execute(w, req.FormValue("s"))
    if req.FormValue != nil {
        min, max, sum := minmax.FindMinmax(req.FormValue("s"))
        if min != 0 && max != 0 && sum != 0 {
            fmt.Fprintf(w, "%d %d %d", min, max, sum)
        }
    }
    
}

/*
 * Running function on access to "worm" url location
 * Executing template first then returning result(s)
 */
func Worm(w http.ResponseWriter, req *http.Request) {

    req.ParseForm()

    worm_templ.Execute(w, req.Form["f"])
    if req.FormValue != nil {
        
        var nI, uI, dI int

        nS := req.FormValue("n")
        uS := req.FormValue("u")
        dS := req.FormValue("d")
        
        if n, err := strconv.Atoi(nS); err == nil {
            nI = n
        }
        if u, err := strconv.Atoi(uS); err == nil {
            uI = u
        }
        if d, err := strconv.Atoi(dS); err == nil {
            dI = d
        }

        minuteCount := worm.FindMinuteCount(nI, uI, dI)

        if minuteCount != 0 {
            fmt.Fprintf(w, "Minute Count : %d", minuteCount)
        }
    }

}

func Palindrome(w http.ResponseWriter, req *http.Request) {

    palindrome_templ.Execute(w, req.FormValue("s"))
    if req.FormValue != nil {
        palindromeState := palindrome.FindPalindrome(req.FormValue("s"))
        if palindromeState == 1 {
            fmt.Fprintf(w, "PALINDROME")
        }else if palindromeState == 0 {
            fmt.Fprintf(w, "NOT PALINDROME")
        }
    }

}
