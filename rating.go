package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var rating string
	// frontend
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter user name")
	name, _ = reader.ReadString('\n')
	// take input from reader and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter our rating")
	rating, _ = reader.ReadString('\n')
	userrating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	//backend
	fmt.Printf("hello %v, \n thanks for your %v rating,\n\n your rating %v recorded", name, userrating, time.Now().Format(time.Stamp))

}
