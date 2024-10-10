# Tripoli - Milestones

### [M0] Modelo del problema
+ **Descripción**: Realizar un análisis exhaustivo de las Historias de Usuario (HU) para identificar y definir los elementos clave del dominio relacionados con la gestión de turnos de enfermería, estableciendo la metodología adecuada para su modelado y representación en el sistema. El objetivo es crear un modelo que represente fielmente el contexto y los procesos implicados en la gestión de turnos de enfermería.
+ **Producto mínimamente viable**: Código que encapsule los elementos fundamentales del dominio, con clases que representen roles como Jefe de Enfermería, Enfermero, y otros elementos esenciales como Turno, Disponibilidad y Horario, donde se reflejen adecuadamente las relaciones entre estos y ejemplos de como interactúan entre ellos, por ejemplo la asignación de turnos a un grupo de enfermeros respetando las normativas del convenio.
+ **Requisitos para la Validación**:
- El modelo será considerado válido si refleja de manera precisa todas las entidades clave del dominio y sus relaciones correctamente implementadas en el código.
- El diseño debe ser modular y extensible, permitiendo ajustes futuros para incorporar cambios en las normativas o en las reglas de asignación de turnos.
- Se considera viable cuando las reglas del convenio, como los límites de horas o los descansos, están correctamente integradas en el modelo y pueden ser validadas mediante pruebas simples.
- La documentación debe incluir una descripción clara de cada entidad y su propósito dentro del dominio.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M1] Documento sobre las Reglas del Convenio
+ **Descripción**: Identificar, analizar y documentar todas las normativas aplicables que afectan la gestión de turnos en el hospital, incluyendo aspectos como el número máximo de horas que puede trabajar un enfermero por semana, los descansos obligatorios entre turnos, los límites en la asignación de turnos nocturnos y festivos, así como las reglas para la compensación de horas extra. También se tendrán en cuenta otros aspectos importantes, como las leyes laborales y los acuerdos sindicales específicos.
+ **Producto mínimamente viable**: Un documento detallado que recopile y explique todas las reglas y normativas del convenio que afectan la planificación de los turnos. Este documento servirá como referencia para el desarrollo del algoritmo de asignación.
+ **Requisitos para la Validación**: El entregable será validado cuando el revisor lo haya verificado y aprobado, confirmando que incluye todas las normativas relevantes y que están correctamente documentadas y entendidas.
+ **HUs asociadas**: [HU001]

### [M2] Algoritmo de Asignación de Turnos
+ **Descripción**: Implementar un algoritmo que asigne turnos de manera automática, respetando las reglas del convenio laboral y las preferencias de los enfermeros, asegurando un reparto justo y equitativo de los turnos más demandados, como los nocturnos o festivos. El algoritmo debe optimizar la distribución de las horas de trabajo, garantizando que no se excedan los límites legales y que se mantenga un equilibrio en las cargas de trabajo entre los enfermeros.
+ **Producto mínimamente viable**: Código funcional del algoritmo de asignación de turnos, que será utilizado para gestionar los turnos de forma automática, considerando las reglas del convenio y la disponibilidad del personal.
+ **Criterio de Validación**: El algoritmo será validado mediante pruebas con conjuntos de datos de prueba, asegurando que se respetan las reglas del convenio y que el reparto de los turnos es justo. El entregable será considerado viable cuando el revisor apruebe los resultados de las pruebas.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

## Milestones Adicionales

### [M3] Informe de la validación del Algoritmo en Casos Reales
+ **Descripción**: Probar el algoritmo de asignación de turnos utilizando datos reales de los enfermeros y turnos del hospital, asegurando que el reparto sea justo y conforme al convenio. El foco estará en que el personal enfermero revise y valide los horarios generados por el algoritmo, garantizando que el reparto sea aceptado por los empleados y no solo por los gestores de turnos.
+ **Producto mínimante viable**: Un informe detallado con los resultados de la validación, que incluya análisis sobre la equidad en el reparto de turnos, el cumplimiento de las normativas del convenio, y la aceptación del personal. El informe también deberá incluir cualquier ajuste en el algoritmo tras la retroalimentación con el personal.
+ **Requisitos para la Validación**: El informe será considerado válido cuando, además de la aprobación del revisor, se haya recogido el feedback de una muestra significativa del personal enfermero, confirmando que los horarios asignados cumplen con los requisitos del convenio y las preferencias del equipo.
+ **HUs asociadas**: [HU001], [HU002]
