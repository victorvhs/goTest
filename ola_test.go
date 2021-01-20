package main

import "testing"
func TestOla(t *testing.T){
	resultado := Ola("Analisa")
	esperado := "Olá, Analisa"

	if (resultado != esperado){
		t.Errorf("Resultado '%s', esperado '%s' ",resultado,esperado)
	}
}