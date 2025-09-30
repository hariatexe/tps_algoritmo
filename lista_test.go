package lista_test

import (
	ListaEnlazada "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	pruebaInt1    int     = 5
	pruebaInt2    int     = 25
	pruebaInt3    int     = 17
	pruebaString1 string  = "a"
	pruebaString2 string  = "b"
	pruebaByte1   byte    = 10
	pruebaByte2   byte    = 15
	pruebaFloat1  float32 = 3.14
	pruebaFloat2  float32 = 2.71
	pruebaVolumen int     = 20000
)

func TestListaInt(t *testing.T) {
	// Prueba con int.
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() {
		lista.VerPrimero()
	})
	require.Panics(t, func() {
		lista.BorrarPrimero()
	})
	require.Panics(t, func() {
		lista.VerUltimo()
	})
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarUltimo(pruebaInt1)
	lista.InsertarPrimero(pruebaInt2)

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, pruebaInt2, lista.VerPrimero())
	require.EqualValues(t, pruebaInt1, lista.VerUltimo())
	require.EqualValues(t, pruebaInt2, lista.BorrarPrimero())

	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, pruebaInt1, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())

	require.EqualValues(t, 0, lista.Largo())
}

func TestListaString(t *testing.T) {
	// Prueba con string.
	lista := ListaEnlazada.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() {
		lista.VerPrimero()
	})
	require.Panics(t, func() {
		lista.BorrarPrimero()
	})
	require.Panics(t, func() {
		lista.VerUltimo()
	})
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarUltimo(pruebaString1)
	lista.InsertarPrimero(pruebaString2)

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, pruebaString2, lista.VerPrimero())
	require.EqualValues(t, pruebaString1, lista.VerUltimo())
	require.EqualValues(t, pruebaString2, lista.BorrarPrimero())

	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, pruebaString1, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())

	require.EqualValues(t, 0, lista.Largo())
}

func TestListaByte(t *testing.T) {
	// Prueba con byte.
	lista := ListaEnlazada.CrearListaEnlazada[byte]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() {
		lista.VerPrimero()
	})
	require.Panics(t, func() {
		lista.BorrarPrimero()
	})
	require.Panics(t, func() {
		lista.VerUltimo()
	})
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarUltimo(pruebaByte1)
	lista.InsertarPrimero(pruebaByte2)

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, pruebaByte2, lista.VerPrimero())
	require.EqualValues(t, pruebaByte1, lista.VerUltimo())
	require.EqualValues(t, pruebaByte2, lista.BorrarPrimero())

	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, pruebaByte1, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())

	require.EqualValues(t, 0, lista.Largo())
}

func TestListaFloat(t *testing.T) {
	// Prueba con float.
	lista := ListaEnlazada.CrearListaEnlazada[float32]()
	require.True(t, lista.EstaVacia())

	require.Panics(t, func() {
		lista.VerPrimero()
	})
	require.Panics(t, func() {
		lista.BorrarPrimero()
	})
	require.Panics(t, func() {
		lista.VerUltimo()
	})
	require.EqualValues(t, 0, lista.Largo())

	lista.InsertarUltimo(pruebaFloat1)
	lista.InsertarPrimero(pruebaFloat2)

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, pruebaFloat2, lista.VerPrimero())
	require.EqualValues(t, pruebaFloat1, lista.VerUltimo())
	require.EqualValues(t, pruebaFloat2, lista.BorrarPrimero())

	require.EqualValues(t, 1, lista.Largo())

	require.EqualValues(t, pruebaFloat1, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())

	require.EqualValues(t, 0, lista.Largo())
}

func TestListaVoluminosa(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	for i := 0; i < pruebaVolumen; i++ {
		lista.InsertarUltimo(i)
	}

	require.EqualValues(t, pruebaVolumen, lista.Largo())

	contador := 0
	for !lista.EstaVacia() {
		require.EqualValues(t, contador, lista.BorrarPrimero())
		contador++
	}

	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoPrincipio(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.False(t, iterador.HaySiguiente())
	iterador.Insertar(pruebaInt2)
	require.EqualValues(t, pruebaInt2, lista.BorrarPrimero())
	require.EqualValues(t, pruebaInt1, lista.VerPrimero())
}

func TestListaIteradorExternoInsertarFinal(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, pruebaInt2, iterador.Siguiente())
	require.False(t, iterador.HaySiguiente()) //El iterador esta al final.
	iterador.Insertar(pruebaInt3)

	require.EqualValues(t, pruebaInt1, lista.BorrarPrimero())
	require.EqualValues(t, pruebaInt2, lista.BorrarPrimero())
	require.EqualValues(t, pruebaInt3, lista.BorrarPrimero()) //Confirma que se guardo al final.

	require.True(t, lista.EstaVacia())
}

func TestListaIteradorExternoInsertar(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, pruebaInt1, iterador.Siguiente())
	iterador.Insertar(pruebaInt3) //En el medio de los 2 datos cargados.
	require.EqualValues(t, pruebaInt1, lista.BorrarPrimero())
	require.EqualValues(t, pruebaInt3, lista.BorrarPrimero()) //Verifica que se haya cargado en el medio.
	require.EqualValues(t, pruebaInt2, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
}

func TestRemoverPrimerElemento(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, pruebaInt1, lista.VerPrimero())
	require.EqualValues(t, pruebaInt2, lista.VerUltimo())
	require.EqualValues(t, pruebaInt1, iterador.Borrar()) //Borra el primer elemento sin cambiar la posicion del iterador.
	require.False(t, lista.VerPrimero() == pruebaInt1)    //Verifica que sea falso que el primer elemento cargado ya no este en la primera posicion.
}

func TestRemoverUltimoElemento(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, pruebaInt1, lista.VerPrimero())
	require.EqualValues(t, pruebaInt2, lista.VerUltimo())
	require.EqualValues(t, pruebaInt1, iterador.Siguiente())
	require.False(t, iterador.HaySiguiente()) //No hay siguiente si solo hay 2 elementos y paso del primero al segundo.
	require.EqualValues(t, pruebaInt2, iterador.Borrar())

	require.EqualValues(t, pruebaInt1, lista.VerPrimero())
	require.False(t, lista.VerUltimo() == pruebaInt2)
	require.EqualValues(t, 1, lista.Largo())
}

func TestRemoverElementoMedio(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	lista.InsertarUltimo(pruebaInt3)
	iterador := lista.Iterador() //Cuando se crea apunta al primero.
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, pruebaInt1, iterador.Siguiente())
	require.EqualValues(t, pruebaInt2, iterador.Borrar())     //Borra el segundo elemento
	require.EqualValues(t, pruebaInt1, lista.BorrarPrimero()) //Verifica que el primer elemento sea el que estaba en la posicion y lo borra.
	require.EqualValues(t, pruebaInt3, lista.VerPrimero())    //Verifica que el segundo elemento ahora sea el que antes era el tercero.
}

func TestIterarListaVacia(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	contador := 0
	lista.Iterar(func(dato int) bool {
		contador++
		return true
	})
	require.EqualValues(t, 0, contador)
}

func TestIterarListaCorte(t *testing.T) {
	lista := ListaEnlazada.CrearListaEnlazada[int]()
	lista.InsertarPrimero(pruebaInt1)
	lista.InsertarUltimo(pruebaInt2)
	lista.InsertarUltimo(pruebaInt3)
	contador := 0
	lista.Iterar(func(dato int) bool {
		if dato == pruebaInt2 {
			return false
		} else {
			contador++
			return true
		}
	})
	require.EqualValues(t, 1, contador)
}
