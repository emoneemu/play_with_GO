package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//fmt.Println("Hello Go!")
	var intNum int16 = 9
	fmt.Println(intNum)

	var floatNum float32 = 5.57
	fmt.Println(floatNum)

	var floatnum1 float64 = 10.05

	var result float64 = floatnum1 + float64(intNum)
	fmt.Println(result)

	var intNum1 int16 = 9
	var intNum2 int16 = 7
	fmt.Println(float64(intNum1 / intNum2))
	fmt.Println(intNum1 * intNum2)

	var mystring string = "hello" + " " + "string"
	fmt.Println(mystring)

	var myboolean bool = false
	fmt.Println(myboolean)

	var cartype = "BMW"
	fmt.Println(cartype)

	carcolor := "Blue"
	fmt.Println(carcolor)

	car1, car2, car3 := "ferrari", "suzuki", "gixer"
	fmt.Println(car1, car2, car3)

	const myCOnst string = "s biaaaaaaaatch"
	fmt.Println(myCOnst)

	//you cant change const later
	//myCOnst = "not so sexy"

	printMe(myCOnst)

	fmt.Println(intDivision(6, 3))

	fmt.Printf("my name is emon.The const above is %v", myCOnst)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))

	//===========================built a server on go===============================================

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.Handle("/form", formhandler)
	http.Handle("/hello", hellohandler)

	fmt.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
func printMe(printvalue string) {
	fmt.Println(printvalue)
}

func intDivision(numerator int16, denominator int16) string {
	if denominator == 0 {
		myString := "denominator can not be zero"
		return myString
	} else {
		var result int16 = numerator / denominator
		var finResult string = strconv.Itoa(int(result))
		return finResult

	}

}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from database is :", dbData[i])
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form parse error", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Form parsed successfully")
	name := r.PostFormValue("name")
	address := r.PostFormValue("address")
	fmt.Fprintln(w, "Name: %s\n", name)
	fmt.Fprintln(w, "Address: %s\n", address)
}
