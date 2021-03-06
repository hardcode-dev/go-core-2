package main

import (
	"fmt"
	"log"
	"net/rpc"

	"go-core-2/17-WS-RPC/2-RPC/pkg/books"
)

type book struct {
	ID     int
	Title  string
	Author string
}

func main() {
	// создание клиента RPC
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// получение списка книг, используя тип данных результата из пакета "books"
	var req = &books.Request{}
	var data []books.Book
	client.Call("Server.Books", req, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
	// получение книги по ID, используя локальный тип данных
	req = &books.Request{ID: 1}
	var item book
	client.Call("Server.Book", req, &item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", item)
	// неверное имя функции
	/*err = client.Call("Server.WrongName", req, &item)
	if err != nil {
		log.Fatal(err)
	}*/
	// неверные типы аргументов
	err = client.Call("Server.Book", 100, "ABC")
	if err != nil {
		log.Fatal(err)
	}
}
