package main

import (
	"net/http"
    _ "github.com/lib/pq"
    "encoding/json"
)

func ErrorHandler(w http.ResponseWriter, err interface{}, code int){
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(err)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var booking_id int
    room_id := r.FormValue("room_id")
    date_start := r.FormValue("date_start")
    date_end := r.FormValue("date_end")
    err := database.QueryRow("insert into booking (room_id, date_start, date_end) values ($1, $2, $3) returning booking_id", room_id, date_start, date_end).Scan(&booking_id)
    if err != nil {
        ErrorHandler(w,err,400)
       return
    }
    json_return := map[string]int{"booking_id": booking_id}
    json_data, errno := json.Marshal(json_return)
    if errno != nil {
        ErrorHandler(w,errno,400)
        return
    }
    w.Write(json_data)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
    room_id := r.URL.Query().Get("room_id")
    book, erro := database.Query("select * from booking where room_id = $1 order by date_start",room_id)
    if erro != nil {
        ErrorHandler(w,erro,400)
        return
    }
    defer book.Close()
    booking := []Booking_js{}
    for book.Next(){
        b := Booking{}
        b_js := Booking_js{}
        err := book.Scan(&b.Room_Id, &b.Booking_Id, &b.Date_start, &b.Date_end)
        b_js.Date_start = b.Date_start.Format("2006-01-02")
        b_js.Date_end = b.Date_end.Format("2006-01-02")
        b_js.Room_Id = b.Room_Id
        b_js.Booking_Id = b.Booking_Id          
        if err != nil{
            ErrorHandler(w,err,400)
            continue
        }
        booking = append(booking, b_js)
    }
    json_data, errno := json.Marshal(booking)
    if errno != nil {     
        ErrorHandler(w,errno,400)
        return
    }
    w.Write(json_data)      
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    booking_id := r.URL.Query().Get("booking_id")
    _, err := database.Exec("delete from booking where booking_id = $1", booking_id)
    if err != nil{
        ErrorHandler(w,err,400)
        return
    }
     
}
