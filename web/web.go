package main

import (
	"encoding/csv"
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
	dataset := []string{}

	// horrible, must do properly
	csvFile, err := os.Open("web/dataset.csv")
	if err != nil {
 		panic(err)
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll()
	
	for _,v := range rows {
		dataset = append(dataset,v[0])
	}

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