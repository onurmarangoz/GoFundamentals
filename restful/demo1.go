package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Demo1() {
	//GET ISLEMLERI
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var toDo ToDo
	json.Unmarshal(bodyBytes, &toDo)
	fmt.Println(toDo)
}

func Demo2() {
	//PSOT ISLEMLERI
	toDo := ToDo{1, 1, "Go çalış", true}
	jsonToDo, err := json.Marshal(toDo)

	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonToDo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var toDoResponse ToDo
	json.Unmarshal(bodyBytes, &toDoResponse)
	fmt.Println(toDoResponse)

}

type ToDo struct {
	//Altgr + virgül + sonrasında herhangi bir tuş space gibi
	//Alternatif yol olarak elde json varsa https://mholt.github.io/json-to-go/ adresinden çevrim yapılabilir
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
