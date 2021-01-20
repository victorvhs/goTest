package main
import "fmt"

func Ola(nome string) string {
	/*  Função que imprime ola */
	return "Olá, " +nome
}
func main(){
	fmt.Print(Ola("Andaluzia"))
}