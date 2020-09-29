package main

import "fmt"

func main() {
	var day, month float64
	var output string

	// fmt.Println("**Fecha de Nacimiento**")

	// fmt.Print("Dia: ")
	fmt.Scan(&day)
	// fmt.Print("Mes: ")
	fmt.Scan(&month)

	switch month {
	// Diciembre / Enero
	case 1:
		if day < 21 {
			output = "capricornio"
		} else {
			output = "acuario"
		}
		// Enero / Febrero
	case 2:
		if day < 19 {
			output = "acuario"
		} else {
			output = "piscis"
		}
		// Febrero / Marzo
	case 3:
		if day < 21 {
			output = "piscis"
		} else {
			output = "aries"
		}
		// Marzo / Abril
	case 4:
		if day < 21 {
			output = "aries"
		} else {
			output = "tauro"
		}
		// Abril / Mayo
	case 5:
		if day < 22 {
			output = "tauro"
		} else {
			output = "geminis"
		}
		// Mayo / Junio
	case 6:
		if day < 22 {
			output = "geminis"
		} else {
			output = "cancer"
		}
		// Junio / Julio
	case 7:
		if day < 23 {
			output = "cancer"
		} else {
			output = "leo"
		}
		// Julio / Agosto
	case 8:
		if day < 23 {
			output = "leo"
		} else {
			output = "virgo"
		}
		// Agosto / Septiembre
	case 9:
		if day < 23 {
			output = "virgo"
		} else {
			output = "libra"
		}
		// Septiembre / Octubre
	case 10:
		if day < 23 {
			output = "libra"
		} else {
			output = "escorpio"
		}
		// Octubre / Noviembre
	case 11:
		if day < 23 {
			output = "escorpio"
		} else {
			output = "sagitario"
		}
	case 12:
		if day < 22 {
			output = "sagitario"
		} else {
			output = "capricornio"
		}
	}

	fmt.Println(output)
}
