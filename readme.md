# Reto tetris

## Sobre el código

Una simple prueba técnica para ver si soy capaz de hacerla más o menos rápido. No hay tests ni he tenido en cuenta DDD, o hexagonal, tampoco parecía muy necesario.

Inspirado por [Reto #33](https://github.com/mouredev/retos-programacion-2023/blob/main/Retos/Reto%20%2333%20-%20TETRIS%20%5BDif%C3%ADcil%5D/ejercicio.md).

He añadido, sin darme cuenta, la posibilidad de mover la pieza hacia arriba. Desde luego, en un _tetris_ esto sería imposible.

## El reto

### Crea un programa capaz de gestionar una pieza de Tetris

- La pantalla de juego tiene 10 filas y 10 columnas representadas por símbolos 🔲
- La pieza de tetris a manejar será la siguiente (si quieres, puedes elegir otra):

```text
🔳
🔳🔳🔳
```

- La pieza aparecerá por primera vez en la parte superior izquierda de la pantalla de juego.

```text
🔳🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔳🔳🔳🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
🔲🔲🔲🔲🔲🔲🔲🔲🔲🔲
```

- Debes desarrollar una función capaz de desplazar y rotar la pieza en el tablero,
recibiendo una acción cada vez que se llame, mostrando cómo se visualiza en la pantalla de juego.
- Las acciones que se pueden aplicar a la pieza son: derecha, izquierda, abajo, rotar.
- Debes tener en cuenta los límites de la pantalla de juego.
