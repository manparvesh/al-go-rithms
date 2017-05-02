package main

import (
    "os"
	"math/rand"
	"strconv"
)

func main(){
	file, err := os.OpenFile("Numbers.txt", os.O_APPEND|os.O_WRONLY, 0600)

	if(err != nil){
		panic(err)
	}

	// after doing everything, this will be executed and we'll know if there's a problem in closing the file. clever
	defer func(){
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i <= 1e5; i++ {
		file.WriteString(strconv.Itoa(rand.Int()) + "\n")
	}
}