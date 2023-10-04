package cola

/*Definicion del struct nodo.*/
type nodo[T any] struct {
	elemento T
	siguiente *nodo
}
/* Definición del struct cola. */
type colaEnlazada[T any] struct {
	primero *nodo
	ultimo *nodo
}

/*Constructor del TAD Cola*/
func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (cola *colaEnlazada[T])EstaVacia() bool {
	return cola.primero == nil
}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (cola *colaEnlazada[T])VerPrimero() T {
	if cola.EstaVacia() == true {
		panic("La pila esta vacia")
	}
	return pila.primero.elemento
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T])Encolar(elem T) {
	nuevoNodo := new(nodo[T])
	nuevoNodo.elemento = elem
	nuevoNodo.siguiente = nil
	cola.ultimo = nuevoNodo
	return
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T])Desencolar() T {
	if cola.EstaVacia() == true {
		panic("La pila esta vacia")
	}
	elementoDesencolado := cola.primero.elemento
	cola.primero = cola.primero.siguiente
	return elementoDesencolado
}