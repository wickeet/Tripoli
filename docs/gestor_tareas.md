# Gestor de Tareas

Los criterios para determinar cual es el gestor de tareas más adecuado son las siguientes:

1. Mantenibilidad

Un proyecto que no sea mantenido y actualizado con regularidad puede llevar a deuda técnica, si luego por problemas de rendimiento
o de errores no solucionados tenemos que cambiar a otro gestor de tareas. Aquí también influirá la madurez del proyecto, tal vez la
herramienta no reciba actualizaciones tan asiduas, pero si está muy establecida y es más reconocida por la comunidad también se le puede
considerar una herramienta candidata.


# Posibles Candidatos a elección

## Task
Task es un proyecto con actualizaciones frecuentes y es valorado positivamente, como se puede ver en su [repositorio](https://github.com/go-task/task/).

## Make
Aunque no reciba actualizaciones constantes, Make lleva siendo utilizada desde hace décadas y se considera una herramienta estable y con muchos usuarios.

## Mage
Mage no tiene actualizaciones constantes, con largos períodos de inactividad en su [repositorio](https://github.com/magefile/mage), lo que genera ciertas dudas sobre si la herramienta puede llegar a quedarse desactualizada en un período corto de tiempo, suponiendo un problema.

## Just
Como se puede observar en el [repositorio](https://github.com/casey/just), Just recibe correcciones y actualizaciones regularmente, además de una gran valoración de su repositorio.

## Task-Ninja
Aunque el [repositorio](https://github.com/RikunjSindhwad/Task-Ninja) recibe actualizaciones regulares, es un proyecto reciente y no tan maduro como otras opciones contempladas anteriormente, lo que podría llevar a encontrar errores no identificados con más facilidad.

## Doit
Este gestor de tareas basado en Python lleva un año sin actualizar su [repositorio](https://github.com/pydoit/doit) lo que puede llevar a problemas en el caso de que se encuentren errores.

# Elección final

Tras haber descartado 3 candidatos por no cumplir el criterio, las opciones restantes son Task, Make o Just. Por familiaridad y experiencia, se ha decidido usar Make.