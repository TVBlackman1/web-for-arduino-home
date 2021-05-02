package main

import (
	dbDefaultEssence "./DBDefaultEssence"
	DBRequests "./DBDefaultRequests"
	Devices "./Devices"
	StandartDevices "./Devices/StandartDevices"
	FrontendResponse "./FrontendResponse"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func faviconHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./favicon/favicon.ico")
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()


	var user dbDefaultEssence.User
	_ = dec.Decode(&user)

	var res FrontendResponse.RegisterResponse

	_, err := DBRequests.GetUserByLogin(user.Login)
	if err == nil {
		// user is already exists
		res = FrontendResponse.RegisterResponse{
			Res: "Error",
			Message: "Аккаунт уже существует",
		}
		json.NewEncoder(w).Encode(&res)
		return

	}

	err = DBRequests.CreateUser(&user)

	res = FrontendResponse.RegisterResponse{
		Res: "OK",
		Message: "Аккаунт успешно создан!",
	}

	json.NewEncoder(w).Encode(&res)

}

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

	StandartDevices.CreateDeviceList()

	DBRequests.Connect()
	defer DBRequests.Db.Close()


	r := mux.NewRouter()

	r.HandleFunc("/favicon.ico", faviconHandler)

	r.HandleFunc("/guest", forGuest).Methods("GET")

	r.HandleFunc("/devices", getDevices).Methods("GET")
	r.HandleFunc("/device/{id}", getDevice).Methods("GET")

	r.HandleFunc("/api/device/{id}", getDevice).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/devices", getDevices).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/register", register).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8920", r))
}

//метеостанция влажность воздуха, давление, температура
//курятник температура поилка сколько воды
//теплица температра влажность воздуха влажность почвы
