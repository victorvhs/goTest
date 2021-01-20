package main
import "fmt"

func Ola(nome string) string {
	if nome == " "{
		nome = "Mundo"
	}	
	return "Olá, " +nome
}
func main(){
	fmt.Print(Ola("Andaluzia"))
}