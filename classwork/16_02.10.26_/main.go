// Магазин по продаже телефонов
// Phone {
//     id int
//     vendor string
//     model string
//     price int
// }
// =====================
// / - главная страница
//     GET - отобразить ассортимент (все телефоны) HTML
// /about - страница про магазин телефонов
//     GET - Получить html с инфо о магазине
// /contacts - страница про контакты
//     GET - Получить html с инфо о контактах в магазине
//     POST - с JSON
//         {
//             email: "user@example.com",
//             message: "текст отзыва"
//         }
// /api/phones/ - JSON со всеми телефонами
//     GET - получить JSON со всеми телефонами
//     POST - добавляет новый телефон
// =====================

package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
}

func allPhonesHandler(w http.ResponseWriter, r *http.Request) {
}

func onePhoneHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contacts", contactsHandler)
	mux.HandleFunc("/api/phones", allPhonesHandler)
	mux.HandleFunc("/api/phones/", onePhoneHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
		return
	}
}
