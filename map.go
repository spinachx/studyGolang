package main

import (
     "fmt"
)

type PersonInfo struct {
     ID      string
     Name    string
     Address string
}

func main() {
     var personDB map[string]PersonInfo
     personDB = make(map[string]PersonInfo)
     //personDB = make(map[string]PersonInfo, 100)

     personDB["12345"] = PersonInfo{"12345", "Tom", "Room 123"}
     personDB["1"] = PersonInfo{"1", "Jack", "Room 456"}

     person, ok := personDB["1234"]

     if ok {
          fmt.Println("Found person:", person.Name, "with ID:", person.ID)
     } else {
          fmt.Println("Not found.")
     }

     delete(personDB, "1234")
}
