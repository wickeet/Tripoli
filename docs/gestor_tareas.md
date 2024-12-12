# Gestor de Tareas

Los criterios para determinar cual es el gestor de tareas más adecuado son las siguientes:

1. Mantenibilidad

Un proyecto que no sea mantenido y actualizado con regularidad puede llevar a deuda técnica, si luego por problemas de rendimiento
o de errores no solucionados tenemos que cambiar a otro gestor de tareas. Aquí también influirá la madurez del proyecto, tal vez la
herramienta no reciba actualizaciones tan asiduas, pero si está muy establecida y es más reconocida por la comunidad también se le puede
considerar una herramienta candidata.

2. Integración con el lenguaje

El gestor de tareas tiene que integrarse con el ecosistema y las herramientas Go, siendo capaz de ejecutar y gestionar tareas específicas
como compilar (go build) o el análisis de dependencias (go mod).
Algunos gestores que son específicos de otros lenguajes no serán capaces de realizar estas tareas y por tanto no serán aptos.

3. Simplicidad

Una vez comprobados los otros criterios, el objetivo principal del gestor de tareas es facilitar la gestión, y como es imposible que el
desarrollador tenga experiencia en todos los gestores y la tarea del gestor es simplificar las operaciones cotidianas, no complicarlas.
Tal vez decidir si algo es simple o no, se basa en la práctica o experiencia personal, pero valorar si algo es más complicado que otro o comparar las curvas de dificultad no lo es. Lo idóneo sería comparar entre sí los gestores que cumplan los requisitos anteriores, siendo este en cierta manera una criba final.



# Posibles Candidatos a elección

## Task
Task es un proyecto con actualizaciones frecuentes y es valorado positivamente, como se puede ver en su [repositorio](https://github.com/go-task/task/). Al estar diseñado para usarse con Go, tiene una integración excelente y puedes definir las tareas directamente en el archivo de configuración YAML.

## Make
Aunque no reciba actualizaciones constantes, Make lleva siendo utilizada desde hace décadas y se considera una herramienta estable y con muchos usuarios. Al ser una herramienta más general, no tiene una integración específica para Go y por tanto habría que realizar scripts personalizados para cada tarea.

## Mage
Mage no tiene actualizaciones constantes, con largos períodos de inactividad en su [repositorio](https://github.com/magefile/mage), lo que genera ciertas dudas sobre si la herramienta puede llegar a quedarse desactualizada en un período corto de tiempo, suponiendo un problema. Por parte de la integración, utiliza el propio Go para definir las tareas, lo que facilita mucho la tarea.

## Just
Como se puede observar en el [repositorio](https://github.com/casey/just), Just recibe correcciones y actualizaciones regularmente, además de una gran valoración de su repositorio. Al igual que Make, no tiene una integración específica para Go y va a necesitar de scripts específicos para las tareas.

## Task-Ninja
Aunque el [repositorio](https://github.com/RikunjSindhwad/Task-Ninja) recibe actualizaciones regulares, es un proyecto reciente y no tan maduro como otras opciones contempladas anteriormente, lo que podría llevar a encontrar errores no identificados con más facilidad. Al ser más general, le ocurrirá lo mismo que a Mage y Make.

## Grunt
Aunque el [repositorio](https://github.com/gruntjs/grunt) está muy valorado y recibe actualizaciones asiduamente, el problema que tiene este gestor es que tiene nula integración con Go, descartándolo completamente de este proceso de selección.

# Elección final

Tras haber visto las distintas opciones, las dos opciones más viables serían Task y Mage. Si bien teniendo el último requisito tal vez Mage sea mejor, porque no requiere que el desarrollador no aprenda nada nuevo al usar el propio Go, el tener una mantenabilidad pobre y que el archivo YAML de Task es sencillo, se ha decidido que la herramienta escogida sea Task.