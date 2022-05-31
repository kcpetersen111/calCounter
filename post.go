package main
import (
	"encoding/json"
	"strings"
	"log"
	"net/http"
	"errors"
	"strconv"
)

// ought to come in as name, calorie count
func handlePostFood(w http.ResponseWriter, r *http.Request){

	var e string
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	// err = json.Unmarshal(r.Body, &content)

	msg := strings.Split(e,",")
	foodName := msg[0]
	cal,err := strconv.Atoi(msg[1])
	if err != nil{
		log.Fatal(err)
	}

	addFood(foodName,cal)
	// f := food{
	// 	Name:foodName,
	// 	Calories:cal,
	// }
	// data[len(data)-1].Ate = append(data[len(data)-1].Ate,f)

	errorResponse(w,"Success",http.StatusOK)
	return
}

func handlePostNewDay(w http.ResponseWriter, r *http.Request){
	log.Println("Request to add a new day")
	addDay()
	errorResponse(w,"Success",http.StatusOK)
	return
}