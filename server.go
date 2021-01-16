package main
import (
    "fmt"
	"net/http"
	"database/sql"
    _ "github.com/lib/pq"
    "html/template"
    "log"
    "time"
    "encoding/json"
    "strings"

)
var database *sql.DB

type Hotel struct{
    Id int 
    Price int 
    Description string 
    Update time.Time 
}

type Hotels_js struct{
    Id int `json:"id"`
    Price int `json:"price"`
    Description string `json:"description"`
    Update string `json:"update"`
}

type Booking struct{
    Room_Id int
    Booking_Id int
    Date_start time.Time 
    Date_end time.Time
}

type Booking_js struct{
    Room_Id int `json:"room_id"`
    Booking_Id int `json:"booking_id"`
    Date_start string `json:"date_start"`
    Date_end string `json:"date_end"`
}

func IndexBookHandler(w http.ResponseWriter, r *http.Request) {
 
    book, erro := database.Query("select * from booking")
    if erro != nil {
        log.Println(erro)
    }
    defer book.Close()
    booking := []Booking{} 
   
    for book.Next(){
        b := Booking{}
        err := book.Scan(&b.Room_Id, &b.Booking_Id, &b.Date_start, &b.Date_end)
        if err != nil{
            fmt.Println(err)
            continue
        }
        booking = append(booking, b)
    }
    tmpl, _ := template.ParseFiles("view/booking.html")
    tmpl.Execute(w, booking)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
 
    rows, err := database.Query("select * from hotel")
    if err != nil {
        log.Println(err)
    }
  
    defer rows.Close()
    hotel := []Hotel{}
   
    for rows.Next(){
        h := Hotel{}
        err := rows.Scan(&h.Id, &h.Price, &h.Description, &h.Update)
        if err != nil{
            fmt.Println(err)
            continue
        }
        hotel = append(hotel, h)
    }
   
    tmpl, _ := template.ParseFiles("view/index.html")
    tmpl.Execute(w, hotel)
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
    var room_id int
    r.ParseForm()
    price := r.FormValue("price")
    description := r.FormValue("description")
    update := time.Now().Format("2006-01-02")
    err := database.QueryRow("insert into hotel (price, description, update) values ($1, $2, $3) returning room_id", price, description, update).Scan(&room_id)
    if err != nil {
        ErrorHandler(w,err,400)
        return
    }
    json_return := map[string]int{"room_id": room_id}
    json_data, errno := json.Marshal(json_return)
    if errno != nil {
        ErrorHandler(w,errno,400)
        return
    }
    w.Write(json_data)
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
    hotel, erro := database.Query("select * from hotel order by update")
    switch r.URL.Query().Get("sort"){
        case "price": hotel, erro = database.Query("select * from hotel order by price")
        case "date" : hotel, erro = database.Query("select * from hotel order by update")
        case "date_desc" : hotel, erro = database.Query("select * from hotel order by update desc")
        default : hotel, erro = database.Query("select * from hotel order by update")
    }
    if erro != nil {
        ErrorHandler(w,erro,400)
        return
    }
    defer hotel.Close()
    room := []Hotels_js{}
    for hotel.Next(){
        h := Hotel{}
        h_js := Hotels_js{}
        err := hotel.Scan(&h.Id, &h.Price, &h.Description, &h.Update)
        h_js.Id = h.Id
        h_js.Price = h.Price
        h_js.Description = strings.TrimSpace(h.Description)
        h_js.Update = h.Update.Format("2006-01-02")
        if err != nil{
            ErrorHandler(w,err,400)
            continue
        }
        room = append(room, h_js)
    }
    json_data, errno := json.Marshal(room)
    if errno != nil {
        ErrorHandler(w,errno,400)
        return
    }
    w.Write(json_data)
}

func DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
    room_id := r.URL.Query().Get("room_id")
    _, erro := database.Exec("delete from booking where room_id = $1", room_id)
    if erro != nil{
        ErrorHandler(w,erro,400)
        return
    }

    _, err := database.Exec("delete from hotel where room_id = $1", room_id)
    if err != nil{
        ErrorHandler(w,err,400)
        return
    }
}

func main() {
  connStr := "host=db port=5432 user=postgres password=1805 dbname=avito_db sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  database = db
  defer db.Close()

  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/booking", IndexBookHandler)
  http.HandleFunc("/bookings/create", CreateHandler)
  http.HandleFunc("/bookings/list", FindHandler)
  http.HandleFunc("/bookings/delete", DeleteHandler)
  http.HandleFunc("/room/create", CreateRoomHandler)
  http.HandleFunc("/room/list", RoomHandler)
  http.HandleFunc("/room/delete", DeleteRoomHandler)
  fmt.Println("Server is listening...")
  http.ListenAndServe(":9000", nil)
}
