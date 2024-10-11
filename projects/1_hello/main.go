package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
	// w - интерфейс для ответа клиенту
	// r - структура с данными о запросе

}

func main() {
	http.HandleFunc("/get", handler)         // регистрация обработчика для пути "/get"
	err := http.ListenAndServe(":8080", nil) // запуск сервера на порту 8080
	if err != nil {
		fmt.Println("error")
	}
}

//http://localhost:8080/get - url, указывает на веб ресурс
