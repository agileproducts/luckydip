package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/agileproducts/luckydip"
	"github.com/gorilla/mux"
	"os"
	"strconv"
)

func main() {
	dataset := []string{"one","two","three","four","five"}
	handler := &luckydipInstance{sourceData:&dataset}

	router := mux.NewRouter()
	router.HandleFunc("/luckydip/{n}", handler.ServeHTTP)
	http.Handle("/",router)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

}

type luckydipInstance struct {
	sourceData *[]string 
}

func (l *luckydipInstance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Path[len("/luckydip/"):]
	num,_ := strconv.Atoi(size)
	selection := l.Selection(num)
	output, _ := json.Marshal(selection)
	fmt.Fprint(w,string(output))
}

func (l *luckydipInstance) Selection(num int) []string {
	return luckydip.RandomSubset(*l.sourceData, num)
}