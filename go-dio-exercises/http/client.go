package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main(){
	resp, err := http.Get("https://granlogprd.ddns.net/api/testarapi")

	if err != nil {
		fmt.Println("caiu no 1 if")
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i:=0; scanner.Scan() && i < 5; i++{
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("caiu no 2 if")
		panic(err)
	}
}