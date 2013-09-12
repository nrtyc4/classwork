package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)

//Need 3 different maps
type Student struct {
 Name string
 Pawprint string
}

type Teachers struct {
  Name string
  Position string
}

type Classes struct {
  Subject string 
}

//This needs to be global
var students = make(map[string]Student)
var teachers = make(map[string]Teachers)
var classes = make(map[string]Classes)

//3 different functions to handle the different maps
func studentHandler(w http.ResponseWriter, req *http.Request) {
  
  switch req.Method {
  case "GET":
    buf, _ := json.Marshal(students)
    w.Write(buf)
  case "PUT":
    buf, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(buf, students)
  default:
    w.WriteHeader(400)
  }
}

func teachersHandler(w http.ResponseWriter, req *http.Request) {
  
  switch req.Method {
  case "GET":
    buf, _ := json.Marshal(teachers)
    w.Write(buf)
  case "PUT":
    buf, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(buf, teachers)
  default:
    w.WriteHeader(400)
  }
}

func classesHandler(w http.ResponseWriter, req *http.Request) {
 
  switch req.Method {
  case "GET":
    buf, _ := json.Marshal(classes)
    w.Write(buf)
  case "PUT":
    buf, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(buf, classes)
  default:
    w.WriteHeader(400)
  }
}

func stuItemHandler(w http.ResponseWriter, req *http.Request) {
  
  switch req.URL.Path {
  case "/students/me":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(students["/students/me"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, students["/students/me"])
    default:
      w.WriteHeader(400)
    }
   case "/students/walter":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(students["/students/walter"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, students["/students/walter"])
    default:
      w.WriteHeader(400)
    }
  }
}

func teachItemHandler(w http.ResponseWriter, req *http.Request) {
  
  switch req.URL.Path {
  case "/teachers/ryanne":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(teachers["/teachers/ryanne"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, teachers["/teachers/ryanne"])
    default:
      w.WriteHeader(400)
    }
    
  case "/teachers/sean":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(teachers["/teachers/sean"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, teachers["/teachers/sean"])
    default:
      w.WriteHeader(400)
    }
  }
}

func classItemHandler(w http.ResponseWriter, req *http.Request) {
  
  switch req.URL.Path {
  case "/classes/cs4830":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(classes["/classes/cs4830"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, classes["/classes/cs4830"])
    default:
      w.WriteHeader(400)
    }
    
  case "/classes/cs4850":
    switch req.Method {
    case "GET":
      buf, _ := json.Marshal(classes["/classes/cs4850"])
      w.Write(buf)
    case "PUT":
      buf, _ := ioutil.ReadAll(req.Body)
      json.Unmarshal(buf, classes["/classes/cs4850"])
    default:
      w.WriteHeader(400)
    }
  }
}

func main () {
  //declare fields for all maps (resources)
  students["/students/me"] = Student{"Nat Thompson", "nrtyc4"}
  students["/students/walter"] = Student{"Walter White", "wlwm6b"}
  
  teachers["/teachers/ryanne"] = Teachers{"Ryanne Dolan", "Professor"}
  teachers["/teachers/sean"] = Teachers{"Sean Lander", "Teaching Assistant"}
  
  classes["/classes/cs4830"] = Classes{"CS4830"}
  classes["/classes/cs4850"] = Classes{"CS4850"}
  
  //Handle functions for collections
  http.HandleFunc("/students", studentHandler)
  http.HandleFunc("/teachers", teachersHandler)
  http.HandleFunc("/classes", classesHandler)
  
  http.HandleFunc("/students/", stuItemHandler)
  http.HandleFunc("/teachers/", teachItemHandler)
  http.HandleFunc("/classes/", classItemHandler)

  //Handle func to set the file server at the path ./
  http.Handle("/", http.FileServer(http.Dir("./")))
  
  //Start server 
  http.ListenAndServe(":8003", nil) 
}
