package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "log"
  "net/http"

)


const (
  DB_USER     = "postgres"
  DB_PASSWORD = "dhaya"
  DB_NAME     = "User"
)

// DB set up
func setupDB() *sql.DB {
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)

  checkErr(err)

  return DB
  type User struct{
    id int 
    Usename_name string
    Password string
  
  }

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []User `json:"data"`
    Message string `json:"message"`
}

func main(){
  router.HandleFunc("/User/", GetUser).Methods("GET")
  router.HandleFunc("/User/", CreateUser).Methods("POST")
  router.HandleFunc("/User/{Userid}", DeleteUser).Methods("DELETE")
  router.HandleFunc("/User/", DeleteUser).Methods("DELETE")
  fmt.Println("Server at 8080")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func printMessage(message string) {
  fmt.Println("")
  fmt.Println(message)
  fmt.Println("")
}

func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  db := setupDB()

  printMessage("Getting User...")

  // Get all User from User table that don't have UserID = "1"
  rows, err := db.Query("SELECT * FROM User")

  checkErr(err)
  var User []User
  for rows.Next() {
      var id int
      var UserID string
      var UserName string

      err = rows.Scan(&id, &UserID, &UserName)
      checkErr(err)

      User = append(User, User{UserID: UserID, UserName: UserName})
  }

  var response = JsonResponse{Type: "success", Data: User}

  json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
  UserID := r.FormValue("Userid")
  UserName := r.FormValue("Username")

  var response = JsonResponse{}
  
  if UserID == "" || UserName == "" {
      response = JsonResponse{Type: "error", Message: "You are missing UserID or UserName parameter."} //Inserting new record.
  } else {
      db := setupDB()

      printMessage("Inserting User into DB")

      fmt.Println("Inserting new User with ID: " + UserID + " and name: " + UserName)

      var lastInsertID int
  err := db.QueryRow("INSERT INTO Users(UserID, UserName) VALUES($1, $2) returning id;", UserID, UserName).Scan(&lastInsertID)

  // check errors
  checkErr(err)

  response = JsonResponse{Type: "success", Message: "The User has been inserted successfully!"}
  }

  json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  UserID := params["Userid"]

  var response = JsonResponse{}

  if UserID == "" {
      response = JsonResponse{Type: "error", Message: "You are missing UserID parameter."}
  } else {
      db := setupDB()

      printMessage("Deleting User from DB")

      _, err := db.Exec("DELETE FROM Users where UserID = $1", UserID)

      checkErr(err) //checking for error

      response = JsonResponse{Type: "success", Message: "The User has been deleted successfully!"}
  }

  json.NewEncoder(w).Encode(response)
}

