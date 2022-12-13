package main

import(
	"fmt"
	"os"
	"encoding/json"
	"sort"
	"log"
)
	/*	Declarations	*/
	
//Structs
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

//Global
var people []Person
var relations []Relation

	/*	Functions	*/
	
//Main
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
	//fmt.Printf("%+v\n", temp)
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

func addPerson(name string, age int64) {
	var newPerson Person
	newPerson.ID = getNextPersonID()
	newPerson.Name = name
	newPerson.Age = age
	people = append(people, newPerson)	
	//saveDB("Person")
}

func removePerson(name string) {
	filtered := make([]Person, 0)
	
	for _, person := range people {
		if person.Name != name {
			filtered = append(filtered, person)
		}
	}
	
    people = filtered
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

func addRelation(person string, other string) {
	personData := getPersonByName(person)
	personID := personData.ID
	
	otherData := getPersonByName(person)
	otherID := otherData.ID
	
	var newRelation Relation
	newRelation.ID = getNextRelationID()
	newRelation.Person = personID
	newRelation.Other = otherID
	
	append(relations, newRelation)
	//saveDB("Relations")
}

func removeRelation(id int64) {
	filtered := make([]Relation, 0)
	
	for _, relation := range relations {
		if relation.ID != id {
			filtered = append(filtered, relation)
		}
	}
	
    relations = filtered
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

func saveDB(database string) bool {
	var result = true
	var data []byte
	var err error
	
	switch database {
		case "People":
			data, err = json.Marshal(people)
		case "Relations":
			data, err = json.Marshal(relations)
		default:
			break;
	}
	
	if err != nil {
		errorLog("Can not marshall data.\nError: " + err.Error())
	}
	
	fmt.Println(data, people)
	
	//dir, _ := os.Getwd()
	
	return result
}

func getNextPersonID() int64 {
	var nextID int64 
	sortedData := people
	
	sort.Slice(sortedData, func(a, b int) bool {
		return sortedData[a].ID < sortedData[b].ID
	})
	
	length := len(sortedData)
	
	nextID = sortedData[length - 1].ID + 1
	
	return nextID
}

func getNextRelationID() int64 {
	var nextID int64 
	sortedData := relations
	
	sort.Slice(sortedData, func(a, b int) bool {
		return sortedData[a].ID < sortedData[b].ID
	})
	
	length := len(sortedData)
	
	nextID = sortedData[length - 1].ID + 1
	
	return nextID
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