package main

import (
	dbDefaultEssence "./DBDefaultEssence"
	DBRequests "./DBDefaultRequests"
	Devices "./Devices"
	StandartDevices "./Devices/StandartDevices"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)


func getDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	params := mux.Vars(r)

	fmt.Fprintln(os.Stdout, "Getting info about " + params["id"])
	for _, item := range StandartDevices.DeviceList {
		if item.ID == params["id"] {
			fmt.Fprintln(os.Stdout, item.Name)
			enc := json.NewEncoder(w)
			enc.Encode(item)
			//enc2 := json.NewEncoder(os.Stdout)
			//enc2.Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Devices.Device{})
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	json.NewEncoder(w).Encode(StandartDevices.DeviceList)

}

func forGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(os.Stdout, "connection!")
	w.Write([]byte("Hello, Kirya!"))

}

func main() {

	fmt.Println("Go MySQL Tutorial")
	StandartDevices.CreateDeviceList()
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ArduinoSmartHome")
	// jdbc:mysql://localhost:3306/ArduinoSmartHome
	//if err != nil {
	//	panic(err.Error())
	//}

	DBRequests.Connect()
	_, _ = DBRequests.GetUserByLogin("tvblackman2")
	//if err != nil {
	//	panic(err.Error())
	//}
	_ = DBRequests.CreateUser(&dbDefaultEssence.User{
		ID: -1,
		Login: "tvblackman2",
		Password: "qwerty",
	})

	//if err != nil {
	//	panic(err.Error())
	//}

	// defer the close till after the main function has finished
	// executing
	defer DBRequests.Db.Close()

	fmt.Println("!!")

	r := mux.NewRouter()

	r.HandleFunc("/guest", forGuest).Methods("GET")

	r.HandleFunc("/devices", getDevices).Methods("GET")
	r.HandleFunc("/device/{id}", getDevice).Methods("GET")

	r.HandleFunc("/api/device/{id}", getDevice).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/devices", getDevices).Methods("POST", "OPTIONS")

	//r.HandleFunc("/api/")



	log.Fatal(http.ListenAndServe(":8920", r))
}

//метеостанция влажность воздуха, давление, температура
//курятник температура поилка сколько воды
//теплица температра влажность воздуха влажность почвы
