package lista

const (
	LISTA_VACIA = 0
)

type nodo[T any] struct {
	dato    T
	proximo *nodo[T]
}
type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

// Iterador externo
type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func crearNodo[T any](dato T, proximo *nodo[T]) *nodo[T] {
	return &nodo[T]{dato: dato, proximo: proximo}
}
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == LISTA_VACIA
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato, lista.primero) //Crea el primer nodo, este apuntara a nil si no hay elementos cargados.
	lista.primero = nuevoNodo
	if lista.ultimo == nil {
		lista.ultimo = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato, nil) //El nodo apuntara a nil porque sera el ultimo y no tiene a quien apuntar.
	if lista.ultimo != nil {
		lista.ultimo.proximo = nuevoNodo
	}

	lista.ultimo = nuevoNodo

	if lista.primero == nil {
		lista.primero = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.proximo //Cambia la referencia, ahora primero apunta a lo que era el segundo.
	lista.largo--                         //Recortar la lista.

	if lista.primero == nil { // si la lista quedo vacia
		lista.ultimo = nil
	}

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero //Comienza a iterar desde el comienzo.
	for actual != nil {
		if !visitar(actual.dato) { //Si la funcion no devuelve true, no se itera mas.
			break
		}
		actual = actual.proximo //Pasa al siguiente nodo.
	}
}

// ITERADOR EXTERNO

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{actual: lista.primero, anterior: nil, lista: lista}
}

// PRIMITIVAS DEL ITERADOR EXTERNO

func (iterado *iteradorLista[T]) VerActual() T {
	if iterado.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iterado.actual.dato
}

func (iterado *iteradorLista[T]) HaySiguiente() bool {
	return iterado.actual != nil && iterado.actual.proximo != nil
}

func (iterado *iteradorLista[T]) Siguiente() T {
	if iterado.actual == nil || iterado.actual.proximo == nil {
		panic("El iterador termino de iterar")
	}
	iterado.anterior = iterado.actual
	iterado.actual = iterado.actual.proximo
	return iterado.actual.dato
}

func (iterado *iteradorLista[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato, iterado.actual)

	if iterado.anterior == nil {
		iterado.lista.primero = nuevoNodo
		if iterado.lista.ultimo == nil {
			iterado.lista.ultimo = nuevoNodo
		}
	} else {
		iterado.anterior.proximo = nuevoNodo
		if iterado.actual == nil {
			iterado.lista.ultimo = nuevoNodo
		}
	}

	iterado.actual = nuevoNodo
	iterado.lista.largo++
}

func (iterado *iteradorLista[T]) Borrar() T {
	if iterado.actual == nil {
		panic("El iterador termino de iterar")
	}
	dato := iterado.actual.dato

	if iterado.anterior == nil { // Borramos el primer nodo
		iterado.lista.primero = iterado.actual.proximo
		if iterado.lista.primero == nil { // Si la lista quedo vacia
			iterado.lista.ultimo = nil
		}
	} else { // Borrar nodo en el medio o al final
		iterado.anterior.proximo = iterado.actual.proximo
		if iterado.actual.proximo == nil { // Si borramos el anterior
			iterado.lista.ultimo = iterado.anterior
		}
	}

	iterado.actual = iterado.actual.proximo
	iterado.lista.largo--

	return dato
}
