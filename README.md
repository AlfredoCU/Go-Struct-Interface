# Go-Struct-Interface

Crear y hacer uso de estructuras, métodos e interfaces.

- Implementar las siguientes estructuras con sus respectivos atributos:
    - Imagen: titulo, formato, canales
    - Audio: titulo, formato, duracion (seg)
    - Video: titulo, formato, frames
- Crear a cada estructura el método `mostrar()` el cual deberá de imprimir los atributos de cada estructura.
- Crear la interface *Multimedia* la cual tendrá el método `mostrar()`.
- Crear la estructura *ContenidoWeb*, la cual tendrá como atributo un *slice* de la interface *Multimedia.*
- Crear un menú en el main para capturar una imagen, audio o video y agregarlos a un objeto de la clase *ContenidoWeb*. Además de una opción para mostrar, la cual llamará al método `mostrar()` de cada elemento (*Multimedia*) almacenado en el *slice.*
