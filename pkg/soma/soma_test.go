package soma

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(10, 5)
	esperado := 15
	
	if resultado != esperado {
		t.Errorf("esperado %d, obteve %d", esperado, resultado)
	}
}