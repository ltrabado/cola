package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.VerPrimero()})
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.Desencolar()})
}

func TestInvarianteCola(t *testing.T) {
	TAM := 10
	cola := TDACola.CrearColaEnlazada[int]()
	//Encolo enteros del 0 al 9
	for i := 0; i < TAM; i++ {
		cola.Encolar(i)
	}
	//Desencolo manteniendo el invariante de la cola
	for i := 0; i < TAM; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
}

func TestVolumen(t *testing.T) {
	TAM := 10000
	cola := TDACola.CrearColaEnlazada[int]()
	//Encolo
	for i := 0; i < TAM; i++ {
		cola.Encolar(i)
	}
	//Desencolo
	for i := 0; i < TAM; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
}

func TestDesencoloHastaVaciar(t *testing.T) {
	TAM := 15
	cola := TDACola.CrearColaEnlazada[int]()
	//Encolo
	for i := 0; i < TAM; i++ {
		cola.Encolar(i)
	}
	//Desencolo
	for i := 0; i < TAM; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.VerPrimero()})
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.Desencolar()})
}

func TestColaFloat64Vacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.VerPrimero()})
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.Desencolar()})
}

func TestInvarianteFloat64Vacia(t *testing.T) {
	TAM := 10
	PI := 3.14
	cola := TDACola.CrearColaEnlazada[float64]()
	//Encolo flotantes
	for i := 0; i < TAM; i++ {
		cola.Encolar(float64(i) * PI)
	}
	//Desencolo flotantes
	for i := 0; i < TAM; i++ {
		require.EqualValues(t, (float64(i) * PI), cola.Desencolar())
	}
}

func TestColaBooleanaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[bool]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.VerPrimero()})
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.Desencolar()})
}

func TestInvarianteColaBooleana(t *testing.T) {
	TAM := 10
	cola := TDACola.CrearColaEnlazada[bool]()
	//Encolo booleanos
	for i := 0; i < TAM; i++ {
		if i%2 == 0 {
			cola.Encolar(true)
		} else {
			cola.Encolar(false)
		}
	}
	//Desencolo booleanos
	for i := 0; i < TAM; i++ {
		if i%2 == 0 {
			require.True(t, cola.Desencolar())
		} else {
			require.False(t, cola.Desencolar())
		}
	}
}

func TestColaStringVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.VerPrimero()})
	require.PanicsWithValue(t, "La cola esta vacia", func() {cola.Desencolar()})
}

func TestInvarianteColaString(t *testing.T) {
	mensaje := []string{"Esta", "es", "una", "cola", "de", "cadenas"}
	cola := TDACola.CrearColaEnlazada[string]()
	//Encolo cadenas
	for i := range mensaje {
		cola.Encolar(mensaje[i])
	}
	//Desencolo cadenas
	for i := range mensaje {
		require.EqualValues(t, mensaje[i], cola.Desencolar())
	}
}