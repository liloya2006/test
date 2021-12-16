package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Average() {
	res, err := http.Get("http://localhost:8080/average")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}

func Summary() {
	res, err := http.Get("http://localhost:8080/summary")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}

func OneDay() {
	res, err := http.Get("http://localhost:8080/oneday")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
}

func main() {
	args := os.Args

	switch args[1] {
	case "1":
		Average()

		break
	case "2":
		Summary()
		break
	case "3":
		OneDay()
		break

	}
}
