package main

import "testing"
func TestOla(t *testing.T){
	verificaMsg := func(t *testing.T, resultado, esperado string){
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado '%s', esperado '%s'", resultado,esperado)
		}
	}
	t.Run("Diz olá para todo mundo", func(t *testing.T){
		resultado := Ola("Analisa")
		esperado := "Olá, Analisa"
		verificaMsg(t,resultado,esperado)
	})
	t.Run("Diz olá mundo, quando não tem nome", func(* testing.T){
		resultado := Ola(" ")
		esperado := "Olá, Mundo"
		verificaMsg(t,resultado,esperado)
	})

}