# Gestor de Tareas

Los criterios para determinar cual es el gestor de tareas más adecuado son las siguientes:

## 1. Lenguaje del propio gestor
El lenguaje o formato del archivo de configuración tiene que ser legible y fácil de mantener. Gestores con lenguajes muy complejos o que sean tediosos
pueden llevar a dificultar en vez de facilitar la tarea de la automatización.

## 2. Mantenimiento y Frecuencia de Actualizaciones
La herramienta debe de ser mantenida activamente y recibe actualizaciones frecuentes para garantizar compatibilidad futura y la corrección de errores,
y que esta no quede obsoleta.

## 3. Tolerancia a Errores
La herramienta debe manejar errores en la ejecución de tareas. Es importante que se detenga ante fallos críticos o que continue bajo condiciones específicas,
según se especifique.

# Posibles Candidatos a elección

# Task

## 1.
Task usa un archivo YAML (Taskfile.yml) que es fácil de leer y de mantener gracias a que tiene un formato simple y estructurado.
## 2.
Se mantiene regularmente en su [repositorio](https://github.com/go-task/task/tree/main) y tiene actualizaciones regulares.
## 3.
Proporciona opciones para manejar errores, como especificar si una tarea debe detenerse al fallar o continuar con otras dependencias.

# Make

## 1.
Utiliza archivos Makefile, cuyo formato no es extremadamente complejo pero hay que prestar especial atención a las tabulaciones y otras reglas.
## 2.
Make es una herramienta estable aunque su desarrollo no es tan activo porque ya cubre un amplio rango de casos.
## 3.
Puede detenerse en caso de errores o seguir adelante, dependiendo de cómo se especifiquen las reglas.

# Mage

## 1.
Utiliza Go puro para definir tareas, lo que lo hace muy flexible, aunque puede ser menos legible si no estás familiarizados con Go (no lo estoy).
## 2.
Llevaba más de un año sin actualizarse su [repositorio](https://github.com/magefile/mage) hasta hace una semana. Aún así esto es un indicio que el mantenimiento no es regular.
## 3.
Como es Go puro, permite manejar errores directamente en el código con un control detallado.

# Goreleaser

## 1.
También basa en un archivo YAML para la configuración, lo cual es simple y legible. Sin embargo, está enfocado principalmente en la distribución de aplicaciones y no en la ejecución general de tareas.
## 2.
Es una herramienta muy activa, con actualizaciones regulares en su [repositorio](https://github.com/goreleaser/goreleaser).
## 3.
Maneja errores en pasos específicos, deteniendo o reportando problemas según el caso.

# Just

## 1.
Utiliza archivos de configuración (Justfile) con una sintaxis inspirada en Makefile pero algo más legible, lo que facilita su uso.

## 2.
Es una herramienta con actualizaciones frecuentes y una comunidad activa en su [repositorio](https://github.com/casey/just).

## 3.
Permite manejar errores de forma similar a make.

# Elección final
Siguiendo los criterios establecidos y comparandolos entre sí, he llegado a la conclusión de que la herramienta que mejores valoraciones ha tenido en cada ámbito en general ha sido Task y será el gestor que se usará en el proyecto.