package main

import "fmt"

func main (){
	fmt.Println("Alô, mundo")
	var NomeOne string = "Gabriel"
	var NomeTwo = "Oliveira"
	var NomeThree string
	fmt.Println(NomeOne, NomeTwo, NomeThree)
	
	nomeFour := "Abyssal"

	fmt.Println(nomeFour)

	var ageOne int = 42
	var ageTwo = 43
	ageThree := 20

	fmt.Println(ageOne, ageTwo, ageThree)

	var  scoreOne float32 = 1.5
	var scoreTwo float64 = 3333333.3
	scoreThree := 3.5

	fmt.Println(scoreOne,scoreTwo,scoreThree )

	var ages = [3]int{10, 20, 30}
	names := [4]string{"Mario", "Luigi", "deadpool","Mario2"}
	fmt.Println(ages)
	fmt.Println()


}

