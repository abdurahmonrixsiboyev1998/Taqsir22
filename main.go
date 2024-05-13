// 1 - misol
// package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"
// )

// type TimeResponse struct {
// 	DayOfWeek  string `json:"day_of_week"`
// 	DayOfMonth int    `lson:"day_of_month"`
// 	Month      string `json:"month"`
// 	Year       int    `json:"year"`
// 	Hour       int    `json:"hour"`
// 	Minute     int    `json:"minute"`
// 	Second     int    `json:"second"`
// }

// func main() {
// 	http.HandleFunc("/time/RFC3336", func(
// 		w http.ResponseWriter, r *http.Request) {
// 		n := time.Now()
// 		if r.Header.Get("ACCEPT") == "application/json" {
// 			resp := TimeResponse{
// 				DayOfWeek:  n.Weekday().String(),
// 				DayOfMonth: n.Day(),
// 				Month:      n.Month().String(),
// 				Year:       n.Year(),
// 				Hour:       n.Hour(),
// 				Minute:     n.Minute(),
// 				Second:     n.Second(),
// 			}
// 			json.NewEncoder(w).Encode(resp)
// 		} else {
// 			w.Write([]byte(n.Format(time.RFC3339)))
// 		}
// 	})
// 	http.ListenAndServe(":8080", nil)
// }

// 2 - misol
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {
	http.HandleFunc("/time/RFC3339", func(w http.ResponseWriter, r *http.Request) {
		Current := time.Now()

		num := Current.Format(time.RFC3339)
		nums, _ := time.Parse(time.RFC3339, num)

		res := TimeResponse{
			DayOfWeek:  nums.Weekday().String(),
			DayOfMonth: nums.Day(),
			Month:      nums.Month().String(),
			Year:       nums.Year(),
			Hour:       nums.Hour(),
			Minute:     nums.Minute(),
			Second:     nums.Second(),
		}

		Response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(Response)
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
