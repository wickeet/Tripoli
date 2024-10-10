# Tripoli - Milestones

### [M0] Modelo del problema
+ **Descripción**: Realizar un análisis exhaustivo de la HU001 para identificar y definir los elementos clave del dominio relacionados con la gestión de turnos de enfermería, estableciendo la metodología adecuada para su modelado y representación en el sistema. El objetivo es crear un modelo que represente fielmente el contexto y los procesos implicados en la gestión de turnos de enfermería.
+ **Producto mínimamente viable**: Código que encapsule los elementos fundamentales del dominio, con clases que representen roles como Jefe de Enfermería, Enfermero, y otros elementos esenciales como Turno u Horario, donde se reflejen adecuadamente las relaciones entre estos y ejemplos de como interactúan entre ellos, por ejemplo la asignación de turnos a un grupo de enfermeros respetando las normativas del convenio.
+ **Requisitos para la Validación**:
- El modelo será considerado válido si refleja de manera precisa todas las entidades clave del dominio y sus relaciones correctamente implementadas en el código.
- El diseño debe ser modular y extensible, permitiendo ajustes futuros para incorporar cambios en las normativas o en las reglas de asignación de turnos.
- Se considera viable cuando las reglas del convenio, como los límites de horas o los descansos, están correctamente integradas en el modelo y pueden ser validadas mediante pruebas simples.
- La documentación debe incluir una descripción clara de cada entidad y su propósito dentro del dominio.

### [M1] Algoritmo de Asignación de Turnos
+ **Descripción**: Implementar un algoritmo que asigne turnos de manera automática, respetando las reglas del convenio laboral y las preferencias de los enfermeros, asegurando un reparto justo y equitativo de los turnos más demandados, como los nocturnos o festivos. El algoritmo debe optimizar la distribución de las horas de trabajo, garantizando que no se excedan los límites legales y que se mantenga un equilibrio en las cargas de trabajo entre los enfermeros, atendiendo a las peticiones de estos (HU002).
+ **Producto mínimamente viable**: Código funcional del algoritmo de asignación de turnos, que será utilizado para gestionar los turnos de forma automática, considerando las reglas del convenio y la disponibilidad del personal, además de una correcta integración de este en el producto ya obtenido en el M1.
+ **Criterio de Validación**: El algoritmo será validado mediante pruebas con conjuntos de datos de prueba, asegurando que se respetan las reglas del convenio y que el reparto de los turnos es justo. El entregable será considerado viable cuando el revisor apruebe los resultados de las pruebas.

## Milestones Adicionales

## [M2] Producto mejorado con respecto a la versión requerida mínima
+ **Descripción**:Teniendo en cuenta el producto ya creado, implementar características adicionales, como la posibilidad de intercambiar turnos entre enfermeros mientras este intercambio asegure que no se incumple la normativa del hospital ni el convenio, de forma que se ataje el problema mencionado en la HU003.
+ **Producto mínimamente viable**: Código funcional que permita el intercambio de turnos entre los enfermeros considerando las reglas del convenio y que complemente lo desarrollado anteriormente.
+ **Criterio de Validación**: El código será validado mediante pruebas con situaciones simuladas, asegurando que se respetan las reglas del convenio. El entregable será considerado viable cuando el revisor apruebe los resultados de las pruebas.
