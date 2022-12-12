package main

import(
	"fmt"
	"os"
	"log"
	"encoding/json"
)

type Person struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

type Relation struct {
	ID int64 `json:"id"`
	Person int64 `json:"person"`
	Other int64 `json:"other"`
}


var people []Person
var relations []Relation

func init() {
	//People
	peopleData := openDB("People")
	people = []Person{}
	peopleError := json.Unmarshal(peopleData, &people)
	if peopleError != nil {
		errorLog("Can not parse People database.\nError: " + peopleError.Error() + "\n")
	}
	
	//Relations
	relationData := openDB("Relations")
	relations = []Relation{}
	relationError := json.Unmarshal(relationData, &relations)
	if relationError != nil {
		errorLog("Can not parse Relations database.\nError: " + relationError.Error() + "\n")
	}
}

func main() {
	fmt.Println("Go Friend DB")
	re := getPersonByName("Riley")
	fmt.Printf("%+v\n", re)
}


//Person
func getPersonByName(name string) Person {
	var person Person
	for i := range people {
		if people[i].Name == name {
			person = people[i]
		}
	}
	return person
}

func getPersonByID(id int64) Person {
	var person Person
	for i := range people {
		if people[i].ID == id {
			return people[i]
		}
	}
	return person
}

func getPersonByAge(age int64) Person {
	var person Person
	for i := range people {
		if people[i].Age == age {
			return people[i]
		}
	}
	return person
}

func addPerson(name string, age int64) bool {
	result := true
	
	
	return result
}

func removePerson(name string) bool {
	result := true
	
	
	return result
}


//Relations
func getRelations(name string) []Person {
	data := getPersonByName(name)

	var relates []int64
	for i := range relations {
		if relations[i].Person == data.ID {
			otherID := relations[i].Other
			relates = append(relates, otherID)
		}
	}
	
	var relatePeople []Person
	for i := range relates {
		peep := getPersonByID(relates[i])
		relatePeople = append(relatePeople, peep)
	}
	
	return relatePeople
}

func addRelation(person string, other int64) bool {
	result := true
	
	
	return result
}

func removeRelation(id int64) bool {
	result := true
	
	
	return result
}

//Util
func openDB(database string) []byte {
	dir, _ := os.Getwd()	
	data, err := os.ReadFile(dir + "/Data/" + database + ".sun") //[]byte, error
	
	if err != nil {
		errorLog("Can not open database " + database + "\nError: " + err.Error())
	}
	
	return data
}

func errorLog(message string) {
	logFile, err := os.OpenFile("Error", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	
	defer logFile.Close()
	
	if _, err = logFile.WriteString(message); err != nil {
		panic(err)
	}
}