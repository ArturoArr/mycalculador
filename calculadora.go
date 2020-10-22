package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//crear un struct
type calc struct{}

//agregamos el struct a la funcion, esto es receiver function
//le da la propiedad a calc de tener el metodo dentro de el, podemos llamar a calc.operate
//y podremos llamar a esta funcion que estamos creando
//parametros que resive y lo que arrojara que sera un resultado entero
func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println("La suma es: ", operador1+operador2)
		return operador1 + operador2
	case "-":
		fmt.Println("La resta es: ", operador1-operador2)
		return operador1 - operador2
	case "*":
		fmt.Println("La multiplicacion es: ", operador1*operador2)
		return operador1 * operador2
	case "/":
		fmt.Println("La division es: ", operador1/operador2)
		return operador1 / operador2
	default:
		fmt.Println(operador, "operador invalida")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

/*
func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	//fmt.Println(entrada)
	//fmt.Println(operador)
	//instanciamos el struct que creamos
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}
*/
