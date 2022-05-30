package main 

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
	"errors"
)


type day struct {
	Ate []food
}

type food struct {
	Name string `json:"name"`
	Calories int `json:"calories"`
}
var (
	data []day
)

func main() {
	err := read("test1.txt")
	if err != nil{
		log.Fatal(err)
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",handleGetAll).Methods("GET")
	myRouter.HandleFunc("/",handlePostFood).Methods("POST")

	serveErr := http.ListenAndServe(":8080",myRouter)
	if serveErr != nil{
		log.Println(serveErr)
	}

	err = save("test1.txt")
	if err != nil{
		log.Fatal(err)
	}

}

func handleGetAll(w http.ResponseWriter, r *http.Request){
	log.Println("Sending All")
	json.NewEncoder(w).Encode(data)
}

// ought to come in as name, calorie count
func handlePostFood(w http.ResponseWriter, r *http.Request){
	// body,err := (*r).GetBody()
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// var p []byte
	// _, err = body.Read(p)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// defer body.Close()

	// valid := json.Valid(r.Body)
	// if !valid {
	// 	log.Fatal("File does not contain valid json.\n")
	// }
	// var content string
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
	f := food{
		Name:foodName,
		Calories:cal,
	}
	data[len(data)-1].Ate = append(data[len(data)-1].Ate,f)

	errorResponse(w,"Success",http.StatusOK)
	return
}


func addDay(){
	today := day{
		Ate:make([]food,0),
	}
	data = append(data,today)
}


func addFood( foodName string, calCount int) {
	newFood := food{
		Name:foodName,
		Calories:calCount,
	}
	(data)[len(data)-1].Ate = append((data)[len(data)-1].Ate, newFood)
}

func save( name string) error{
	fmt.Println(data)
	toSave, err := json.Marshal(data)
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(name,toSave,0644)
	return nil
}

func read(name string) ( error){

	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		temparr := make([]day,0)
		temp, newerr := json.Marshal(temparr)
		if newerr != nil{
			return newerr
		}
		newerr = ioutil.WriteFile(name,temp,0644)
		if newerr != nil{
			return newerr
		}

	}

	fileContent, err := ioutil.ReadFile(name)
	if err != nil {
		return  err
	}

	valid := json.Valid(fileContent)
	if !valid {
		return  fmt.Errorf("File does not contain valid json.\n")
	}

	// var fileData []day
	err = json.Unmarshal(fileContent, &data)
	// fmt.Println(fileContent)
	if err != nil{
		return  err
	}

	fmt.Println(data)
	return  nil
	// return nil,nil

}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}