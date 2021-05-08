package main

import (
	dbDefaultEssence "arduino-server/server/DBDefaultEssence"
	DBRequests "arduino-server/server/DBDefaultRequests"
	Devices "arduino-server/server/Devices"
	StandartDevices "arduino-server/server/Devices/StandartDevices"
	FrontendResponse "arduino-server/server/FrontendResponse"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"net/http"
	"os"
	"time"
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

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	//r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()


	var user dbDefaultEssence.User
	_ = dec.Decode(&user)

	var res FrontendResponse.RegisterResponse

	userInDB, err := DBRequests.GetUserByLogin(user.Login)
	if err != nil {
		// user is not exists
		res = FrontendResponse.RegisterResponse{
			Res: "Error",
			Message: "Аккаунт не существует",
		}
		json.NewEncoder(w).Encode(&res)
		return
	}

	if userInDB.Password != user.Password {
		res = FrontendResponse.RegisterResponse{
			Res: "Error",
			Message: "Неверный пароль",
		}
		json.NewEncoder(w).Encode(&res)
		return
	}

	res = FrontendResponse.RegisterResponse{
		Res: "OK",
		Message: "Вы успешно вошли в аккаунт!",
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

type Test struct {
	K string `json:"k"`
	Hum string `json:"hum"`
}

type ToGlasses struct {
	Device string `json:"device"`
}

var testElement = Test {
	K: "0",
	Hum: "0",
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	params := mux.Vars(r)
	testElement = Test {
	K: params["k"],
	Hum: params["hum"],
	}

	//dec := json.NewDecoder(r.Body)
	//dec.DisallowUnknownFields()
	//
	//_ = dec.Decode(&testElement)

	log.Printf(testElement.K)
	log.Printf(testElement.Hum)
}

func toGlasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var _toGlasses ToGlasses

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	_ = dec.Decode(&_toGlasses)
	log.Printf(_toGlasses.Device)
	sendToArduino(_toGlasses.Device)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	json.NewEncoder(w).Encode(testElement)
}

func sendToArduino(str string) {
	options := serial.OpenOptions{
		PortName: "/dev/cu.usbserial-1420",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	time.Sleep(8*time.Second)

	// Make sure to close it later.
	defer port.Close()

	b := []byte(str + "\n")
	port.Write(b)

	time.Sleep(1*time.Second)


	if str == "weather" {
		for _, item := range StandartDevices.DeviceList {
			if item.ID == "weather-station" {
				var device Devices.WeatherInfo
				strJson, _ := json.Marshal(item.Additional)
				json.Unmarshal([]byte(strJson), &device)
				fmt.Println(device.Temperature)
				b := []byte(fmt.Sprintf("%d", device.Temperature) + "\n")
				port.Write(b)
			}
		}
	}
	if str == "greenhouse" {
		for _, item := range StandartDevices.DeviceList {
			if item.ID == "weather-station" {
				var device Devices.GreenhouseInfo
				strJson, _ := json.Marshal(item.Additional)
				json.Unmarshal([]byte(strJson), &device)
				b := []byte(fmt.Sprintf("%d", device.Temperature) + "\n")
				port.Write(b)

				time.Sleep(2*time.Second)
				b = []byte(fmt.Sprintf("%d", device.AirHumidity) + "\n")
				port.Write(b)
			}
		}
	}
	time.Sleep(2*time.Second)

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
	r.HandleFunc("/api/login", login).Methods("POST", "OPTIONS")

	r.HandleFunc("/api/test/{k}/{hum}", test).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/get-test", getTest).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/to-glasses", toGlasses).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8920", r))
}

//метеостанция влажность воздуха, давление, температура
//курятник температура поилка сколько воды
//теплица температра влажность воздуха влажность почвы
