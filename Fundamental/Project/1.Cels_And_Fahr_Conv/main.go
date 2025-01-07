package main

import "fmt"

func main() {
	fmt.Println("<========= Celsius And Fahrenheit Converter =========>")

	var celsius float32
	var fahrenheit float32

	var option = [2]string{"Celsius", "Fahrenheit"}

	fmt.Printf("1.%s\n", option[0])
	fmt.Printf("2.%s\n", option[1])

	var selectedOptions int8
	fmt.Print("Select Option: ")
	fmt.Scanf("%d", &selectedOptions)
	fmt.Scanln()

	switch selectedOptions {
	case (1):
		fmt.Print("Enter Celsius: ")
		fmt.Scanf("%f", &celsius)

		fahrenheit = celsius*(float32(9)/float32(5)) + float32(32)

		fmt.Printf("%.2f Celsius = %.2f Fahrenheit", celsius, fahrenheit)
	case (2):
		fmt.Print("Enter Fahrenheit: ")
		fmt.Scanf("%f", &fahrenheit)

		celsius = (fahrenheit - float32(32)) * (float32(5) / float32(9))

		fmt.Printf("%.2f Fahrenheit = %.2f Celsius", fahrenheit, celsius)
	default:
		fmt.Println("Invalid Option")
	}
}
