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
	fmt.Print("If you write 1, you get information about seven last day."+ "\n" +
		"If you write 2, you get information about average figures for the last 7 days." + "\n" +
		"If you write 3, you get information about random day." + "\n" +
		"Please, write number:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	numb, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if numb == 1 {
		Summary()
	} else if numb == 2 {
		Average()
	} else if numb == 3 {
		OneDay()
	} else {
		fmt.Println("Incorrect input")
	}
}
