package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)
func main()  {
	if len(os.Args) < 2{
		fmt.Println("Usage: pass <lenght> [use Symbols = Y]")
		return
	}
	
	senhaLength, err := strconv.Atoi(os.Args[1])
	if (err != nil){
		fmt.Println("Error")
		return
	}
	
	var UseSymbols string
	if len(os.Args) >= 3{
		UseSymbols = os.Args[2]
	}
	fmt.Println(GeneratePass(senhaLength, UseSymbols))
}

func GeneratePass(length int, UseSymbols string) string{
	var newPass strings.Builder

	alphabet := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"
	symbols := "!@#$%&*/.!_"

	if(strings.ToUpper(UseSymbols) == "Y"){
		alphabet = alphabet + symbols
	}

	lenAlpha := len(alphabet)
	for i := 0; i < length; i++ {
		newPass.WriteByte(alphabet[rand.Intn(lenAlpha)])
	}

	return newPass.String()
}
