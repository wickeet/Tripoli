# Tripoli - Milestones

### [M0] Modelado del problema
+ **Objetivo**: Realizar un análisis exhaustivo de las Historias de Usuario (HU) para identificar y definir los elementos clave del dominio relacionados con la gestión de turnos de enfermería, estableciendo la metodología adecuada para su modelado y representación en el sistema. El objetivo es crear un modelo que represente fielmente el contexto y los procesos implicados en la gestión de turnos de enfermería.
+ **Entregable**:
- Código que encapsule los elementos fundamentales del dominio, con clases que representen roles como Jefe de Enfermería, Enfermero, y otros elementos esenciales como Turno, Disponibilidad y Horario.
- El código debe reflejar adecuadamente las relaciones entre estos elementos, incluyendo las normativas del convenio, como límites de horas semanales, descansos obligatorios, turnos nocturnos y festivos.
+ **Criterio de Validación**:
- El modelo será considerado válido si refleja de manera precisa todas las entidades clave del dominio y sus relaciones correctamente implementadas en el código.
- El diseño debe ser modular y extensible, permitiendo ajustes futuros para incorporar cambios en las normativas o en las reglas de asignación de turnos.
- Se considera viable cuando las reglas del convenio, como los límites de horas o los descansos, están correctamente integradas en el modelo y pueden ser validadas mediante pruebas simples.
- El código debe incluir ejemplos o pruebas que demuestren cómo las entidades interactúan, como un ejemplo de asignación de turnos a un grupo de enfermeros respetando las normativas del convenio.
- La documentación debe incluir una descripción clara de cada entidad y su propósito dentro del dominio.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M1] Identificación de Reglas del Convenio
+ **Objetivo**: Identificar, analizar y documentar todas las normativas aplicables que afectan la gestión de turnos en el hospital, incluyendo aspectos como el número máximo de horas que puede trabajar un enfermero por semana, los descansos obligatorios entre turnos, los límites en la asignación de turnos nocturnos y festivos, así como las reglas para la compensación de horas extra. También se tendrán en cuenta otros aspectos importantes, como las leyes laborales y los acuerdos sindicales específicos.
+ **Entregable**: Un documento detallado que recopile y explique todas las reglas y normativas del convenio que afectan la planificación de los turnos. Este documento servirá como referencia para el desarrollo del algoritmo de asignación.
+ **Criterio de Validación**: El entregable será validado cuando el revisor lo haya verificado y aprobado, confirmando que incluye todas las normativas relevantes y que están correctamente documentadas y entendidas.
+ **HUs asociadas**: [HU001]

### [M2] Desarrollo de Algoritmo de Asignación de Turnos
+ **Objetivo**: Implementar un algoritmo que asigne turnos de manera automática, respetando las reglas del convenio laboral y las preferencias de los enfermeros, asegurando un reparto justo y equitativo de los turnos más demandados, como los nocturnos o festivos. El algoritmo debe optimizar la distribución de las horas de trabajo, garantizando que no se excedan los límites legales y que se mantenga un equilibrio en las cargas de trabajo entre los enfermeros.
+ **Entregable**: Código funcional del algoritmo de asignación de turnos, que será utilizado para gestionar los turnos de forma automática, considerando las reglas del convenio y la disponibilidad del personal.
+ **Criterio de Validación**: El algoritmo será validado mediante pruebas con conjuntos de datos de prueba, asegurando que se respetan las reglas del convenio y que el reparto de los turnos es justo. El entregable será considerado viable cuando el revisor apruebe los resultados de las pruebas.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M4] Validación del Algoritmo con Datos Reales
+ **Objetivo**: Probar el algoritmo utilizando datos reales de los enfermeros y los turnos del hospital, asegurando que el reparto sea justo y conforme al convenio.Realizar pruebas del algoritmo de asignación de turnos utilizando datos reales de los enfermeros y los turnos del hospital, verificando que el reparto de turnos sea justo, equitativo y en conformidad con el convenio. El objetivo es asegurar que el algoritmo funcione correctamente en un entorno real, y que los resultados sean aceptados tanto por el personal como por los gestores de turnos.
+ **Entregable**: Un informe detallado que recoja los resultados de la validación con datos reales, incluyendo análisis de la equidad en el reparto de turnos, cumplimiento de las normativas y ajuste a las preferencias del personal. El informe también incluirá cualquier ajuste o mejora realizada en el algoritmo como resultado de la validación.
+ **Criterio de Validación**: El informe será considerado válido cuando el revisor lo apruebe, confirmando que el algoritmo cumple con los objetivos establecidos y se ajusta a los requisitos del convenio y las preferencias del personal.
+ **HUs asociadas**: [HU001], [HU002]

