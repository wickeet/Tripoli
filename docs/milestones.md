# Tripoli - Milestones

### [M0] Modelo del problema
+ **Descripción**: Realizar un análisis exhaustivo de la HU001 para identificar y definir los elementos clave del dominio relacionados con la gestión de turnos de enfermería. Se debe establecer una metodología que permita modelar y representar los procesos de gestión de turnos, roles, y restricciones del sistema, proporcionando una base sólida para el desarrollo futuro. El objetivo es crear un modelo flexible y ampliable que represente fielmente el contexto y los procesos implicados en la gestión de turnos de enfermería.
+ **Producto mínimamente viable**: Código que encapsule los elementos fundamentales del dominio, permitiendo la interacción entre roles como Jefe de Enfermería, Enfermero, y el proceso de la asignación de turnos.
+ **Requisitos para la Validación**:
- El diseño debe ser modular y extensible, permitiendo ajustes futuros para incorporar cambios en las normativas o en las reglas de asignación de turnos.
- Se considera validado mediante pruebas con conjuntos de datos de prueba, asegurando una correcta asignación de turnos.

### [M1] Algoritmo de Asignación de Turnos
+ **Descripción**: Implementar un algoritmo que asigne turnos de manera automática, respetando las reglas del convenio laboral y las preferencias de los enfermeros, asegurando un reparto justo y equitativo de los turnos más demandados, como los nocturnos o festivos. El algoritmo debe optimizar la distribución de las horas de trabajo, garantizando que no se excedan los límites legales y que se mantenga un equilibrio en las cargas de trabajo entre los enfermeros, atendiendo a las peticiones de estos (HU002).
+ **Producto mínimamente viable**: Código funcional del algoritmo de asignación de turnos, que será utilizado para gestionar los turnos de forma automática, considerando las reglas del convenio y la disponibilidad del personal, además de una correcta integración de este en el producto ya obtenido en el M1.
+ **Criterio de Validación**: El algoritmo será validado mediante pruebas con conjuntos de datos de prueba, asegurando que el reparto de los turnos es justo.

## Milestones Adicionales

## [M2] Producto mejorado con respecto a la versión requerida mínima
+ **Descripción**:Teniendo en cuenta el producto ya creado, implementar características adicionales, como la posibilidad de intercambiar turnos entre enfermeros si este intercambio asegura que no se incumple la normativa del hospital ni el convenio, de forma que se ataje el problema mencionado en la HU003.
+ **Producto mínimamente viable**: Código funcional que permita el intercambio de turnos entre los enfermeros considerando las reglas del convenio y que complemente lo desarrollado anteriormente.
+ **Criterio de Validación**: El código será validado mediante pruebas con situaciones simuladas, asegurando que se respetan las reglas del convenio.
