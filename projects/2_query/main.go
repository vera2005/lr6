package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// через метод URL получаем запрошенный адрес
	// у адреса вызыввааем метод Query(), возвращает строку запроса
	// Get() для получения значения отдельного параметра
	name := r.URL.Query().Get("name")
	ans := "Hello," + name + "!"
	w.Write([]byte(ans))
}

func main() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("error")
	}

}
