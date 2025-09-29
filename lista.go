package lista
type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la Lista no tiene elementos encolados, false de lo contrario
	EstaVacia() bool

	// InsertarPrimero inserta el dato T al inicio de la lista
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato T al final de la lista
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista y devuelve el valor del elemento eliminado. Si la
	// lista se encontraba vacia entra en panic con el mensaje "La lista esta vacia"
	BorrarPrimero() T

	// VerPrimero devuelve el dato del inicio de la lista, que es el primer elemento. Si la lista se encontraba
	// vacia, entra en panic con el mensaje "La lista esta vacia"
	VerPrimero() T

	// VerUltimo devuelve el dato del final de la lista, que es el ultimo elemento. Si la lista se encontraba vacia
	// devuelve panic con el mensaje "La lista esta vacia"
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion visitar pasada por parametro que recibe un T y devuelve un bool, a todos
	// los elementos de la misma, hasta que no hayan mas elementos o la funcion devuelva false
	Iterar(visitar func(T) bool)

	// Iterador devuelve un IteradorLista para esta Lista
	Iterador() IteradorLista[T]
}


type IteradorLista[T any] interface {
	// VerActual devuelve el elemento actual sobre el que se esta iterando, sin avanzarlo. Si no hay
	// siguiente para ver, entra en panic con el mensaje "El iterador termino de iterar"
	VerActual() T

	// HaySiguiente devuelve true si hay mas datos por ver. Si en el lugar donde esta parado
	// hay un elemento al cual avanzar, si no hay devuelve false
	HaySiguiente() bool

	// Siguiente devuelve el siguiente a ver en la iteracion, avanzando el iterador a la siguiente posicion.
	// Si no hay siguiente, entra en panic con el mensaje "El iterador termino de iterar"
	Siguiente() T

	// Insertar inserta el elemento entre el anterior (si existe) y el actual (si existe)
	// El iterador queda apuntando al nuevo elemento insertado
	Insertar(T)

	// Borrar borra el elemento actual de la lista, y lo devuelve. El elemento anterior ahora apunta
	// hacia el que antes era el proximo del actual. Si no hay siguiente para ver, entra en panic con
	// el mensaje "El iterador termino de iterar"
	Borrar() T
}