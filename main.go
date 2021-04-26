package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	"net/http"
	"os"
	//"strconv"
)

type Device struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Additional  interface{} `json:"additional"`
}

type WeatherInfo struct {
	Temperature int `json:"temperature"`
	AirPressure int `json:"air_pressure"`
	AirHumidity int `json:"air_humidity"`
}

type GreenhouseInfo struct {
	Temperature  int `json:"temperature"`
	SoilMoisture int `json:"soil_moisture"`
	AirHumidity  int `json:"air_humidity"`
}

type ChickenCoopInfo struct {
	AutoWatering bool `json:"auto_watering"`
}

var devices []Device

func getDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	params := mux.Vars(r)

	fmt.Fprintln(os.Stdout, "Getting info about " + params["id"])
	for _, item := range devices {
		if item.ID == params["id"] {
			fmt.Fprintln(os.Stdout, item.Name)
			enc := json.NewEncoder(w)
			enc.Encode(item)
			//enc2 := json.NewEncoder(os.Stdout)
			//enc2.Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Device{})
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	json.NewEncoder(w).Encode(devices)

}

func forGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(os.Stdout, "connection!")
	w.Write([]byte("Hello, Kirya!"))

}

func main() {
	r := mux.NewRouter()
	devices = append(devices, Device{
		ID: "weather-station",
		Name: "Метеостанция",
		Status: "OK",
		Description: "Общепринятый средний показатель нормальной влажности воздуха в квартире должен быть на уровне 45%. Он может варьироваться в зависимости от типа помещения и его эксплуатационных условий. Отклонение от нормы возможно, как в зимнее время года, так и в теплый период. Указанный ГОСТ в зависимости от назначения помещения определяет допустимые значения температуры воздуха в жилых помещениях: от 18 до 25 градусов Цельсия.\n\nТак, например, если, по общему правилу, температура воздуха в жилой комнате квартире должны быть не менее 18 градусов Цельсия, то температура воздуха в ванной или совмещенном помещении уборной и ванной должна быть 25 градусов Цельсия.\n\nДопустимая температура воздуха в вестибюле, лестничной клетке, общем коридоре в квартирном доме составляет 16 градусов.\n\nБатареи должны прогревать помещение до указанных выше величин. В противном случае жильцы вправе рассчитывать на перерасчет платы за коммунальные услуги. Допустимое отклонение температуры воды в батареях в дневное время: 3° С. В ночное время суток (с 00:00 до 05:00) — до 5° С. Большее снижение температуры не допускается за исключением аварийных ситуаций.",
		Additional: WeatherInfo{
			Temperature: 15,
			AirPressure: 770,
			AirHumidity: 52,
		},
	})

	devices = append(devices, Device{
		ID: "2",
		Name: "Smart-glasses",
		Status: "OK",
	})

	devices = append(devices, Device{
		ID: "chicken-coop",
		Name: "Курятник",
		Status: "Not connected",
		Description: "Потребление воды имеет непосредственное отношение к количеству корма. Так, соотношение вода: корм у кур несушек составляет 2,5 к 1, у бройлеров и молодок 1,4-1,6 к 1. ощущать нехватку жидкости в организме.\n\nПрофессиональные птицеводы выявили объем воды, который употребляют куры при определенной температуре.\n\nТак, бройлеры при +18°С выпивают примерно 170 мл жидкости на единицу массы, куры несушки – около 200 мл. Если температура воздуха составляет +30 градусов, количество потребляемой воды вырастет в 2-3 раза.\n\nОптимальная температура для содержания кур +21 градус. В этих условиях птица может съесть около 120 г корма и выпить 200 г воды (1 голова).",
		Additional: ChickenCoopInfo{
			AutoWatering: true,
		},
	})

	devices = append(devices, Device{
		ID: "greenhouse",
		Name: "Теплица",
		Status: "Not connected",
		Description: "В течение всего вегетационного периода влажность для растений в теплице играет важную роль. Влажность в теплице создаёт микроклимат, столь необходимый для роста и развития растений. Недостаток влаги, как и ее избыток, могут решающим образом повлиять на конечный урожай возделываемой культуры.\n\n \n\nДля того чтобы оценить значение влажности воздуха в теплице, нужно понять природу явления. Воздух никогда не бывает абсолютно сухим, он содержит влагу в виде водяного пара или тумана. Водяной пар находится в газообразном состоянии и невидим, лишь когда он собирается в капельки вокруг частиц пыли, становится видимым и проявляется в виде тумана, пара, росы. Насыщение воздуха влагой проходит не бесконечно, и чем выше температура воздуха, тем выше его поглотительная способность. Также на влажность в теплице влияет температура и воздухообмен, поэтому чтобы регулировать эти параметры, необходима хорошая вентиляционная система (боковое и торцевое проветривание).\n \n\n\n \nДля того чтобы поддерживать и контролировать оптимальную влажность в теплице, необходимо предусмотреть в теплице возможность искусственного увлажнения. Это можно осуществить путем увлажнения дорожек между растениями из шланга, либо установить в верхней части теплицы систему туманообразования. Суть этой системы заключается в том, чтобы получить мелкодисперсное распыление в виде тумана. Результатом этого будет повышение влажности в теплице плюс понижение температуры, что тоже очень важно для оптимального развития растений.\n\n \n\nКонтроль влажности в теплице необходимо осуществлять круглосуточно, потому что в разное время суток и в разные фазы развития у растений разная необходимость во влажности. Особенно растения чувствительны к низкой влажности в периоды прорастания семян и цветения. Результатом этого может быть низкая и не дружная всхожесть растений, а в момент опыления пересушенная пыльца не способна к опылению, пыльник не растрескивается. Если в этот период еще и высокие температуры, пыльца вообще становится стерильной. Результат этого мы видим уже через несколько дней — пустые кисти без завязи.\n\n \n\nСухость воздуха в теплице также благоприятна для развития клеща (огурец); с этой проблемой сталкиваются 99% тепличников, а предупредительная мера борьбы очень проста — повысить влажность.\n\n \n\nПоэтому оптимальная влажность воздуха в теплице должна быть в пределах 50–60%, влажность почвы 65–80% (для каждой культуры оптимальная влажность своя).\n\n \n\nТемпература воздуха, как правило, не должна превышать 30 градусов. (Все вышесказанное касается дневных показателей.) Растения, уходящие в ночь, должны быть сухими, и влажность в теплице должна быть минимальной. Это связано с тем, что, как правило, ночные температуры значительно ниже дневных и при наличии высокой влажности есть очень большая вероятность спровоцировать возникновение заболеваний, в результате которых сильно повреждаются растения, а если такие неблагоприятные условия возникают систематически, неизбежна и полная гибель растения.\n\n \n\nИтак, на основании вышесказанного можно сделать вывод, что нарушения оптимальных параметров влажности влечёт за собой дисбаланс в микроклимате теплицы, что, в свою очередь, есть нарушение технологии выращивания. А, как известно, нарушение технологии — это понижение урожайности и ухудшение качества продукции. Мало того, провоцируя благоприятную обстановку для вредителей и болезней, мы увеличиваем число химических обработок, направленных на борьбу с ними, тем самым несем дополнительные затраты, которые влияют на себестоимость выращиваемой продукции. Поэтому с уверенностью можно сказать, что фактор микроклимата очень важен и его надо принимать во внимание.\n\n \n\n\n \n\n \nЕще статьи по этой теме:\nСистема испарительного охлаждения и доувлажнения воздуха в теплицах\nВ разделе: Технологии\n\nМеханизация работ в сельском хозяйстве значительно сказывается на увеличении урожая и облегчает труд человека. Например, существуют...\nСетчатые теплицы - новый этап выращивания сельскохозяйственной продукции\nВ разделе: Технологии\n\nФермеры во всем мире ищут пути повышения эффективности сельскохозяйственного производства, повышения урожайности, качества и...\nВыращивание томатов в теплицах\nВ разделе: Овощеводство\n\nРекомендации по выращиванию индетерминантных гибридов томатов фирмы «Энза Заден» — Белле, Берберана, Буран, Эйджен,...\nЭкзотические культуры с украинским вкусом\nВ разделе: Новости\n\nСпециалисты советуют украинским аграриям переориентироваться на выращивание экзотических культур — селеры, шпината, салата. Так...\n\nДругие статьи:\nРазмножение и агротехника грецкого ореха\nИнтенсификация технологий выращивания кукурузы на зерно\nМасанобу Фукуока: Революция одной соломинки\nБотанические и биологические особенности грецкого ореха\nДиагностика винограда\nОбсудить на форуме\nТеги: теплица, влажность в теплице\nРасскажите друзьям:\n\n\n\nКомментарии (2)\nзагрузка комментариев...\n\n\tДобавить комментарий\tRSS лента\tRSS-лента комментариев\nРекомендуем книги\n\t\t\n\n \nНовые предложения от компаний\nОптические сортировочные машины\n$item.title\n1.00 грн/шт\nОптическая сортировочная машина TRIPLUS®\n$item.title\n1.00 eur/шт\nНовые объявления\nНет объявлений для отображения.\n\nАгромероприятия\nНет активных встреч!\nКалендарь мероприятий\n««\tМарт 2021\t»»\nПн\t1\t8\t15\t22\t29\nВт\t2\t9\t16\t23\t30\nСр\t3\t10\t17\t24\t31\nЧт\t4\t11\t18\t25\t \nПт\t5\t12\t19\t26\t \nСб\t6\t13\t20\t27\t \nВс\t7\t14\t21\t28\t \n\n\nВАКАНСИИ\n\nРобота в Киеве\nРазместить резюме\nО проекте / Авторам / Реклама / АгроМаркет / АгроТорги / Партнеры / Доп.услуги / Контакты / Карта сайта / Архив\nTwitter Facebook\n© 2012 — 2021\n\nАдминистрация социальной сети не несет ответственности за информацию, которая размещена на ресурсе. Вся информация публикуется пользователями. Администрация ресурса может удалить информацию в случае нарушения авторских прав; если информация содержит призывы к насилию, свержению власти, рекламу запрещенных товаров и услуг, а также другую информацию, запрещенную законодательством Украины.",
		Additional: GreenhouseInfo{
			Temperature: 24,
			AirHumidity: 43,
			SoilMoisture: 72,
		},
	})

	r.HandleFunc("/guest", forGuest).Methods("GET")

	r.HandleFunc("/devices", getDevices).Methods("GET")
	r.HandleFunc("/device/{id}", getDevice).Methods("GET")

	r.HandleFunc("/api/device/{id}", getDevice).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/devices", getDevices).Methods("POST", "OPTIONS")



	log.Fatal(http.ListenAndServe(":8000", r))
}

//метеостанция влажность воздуха, давление, температура
//курятник температура поилка сколько воды
//теплица температра влажность воздуха влажность почвы
