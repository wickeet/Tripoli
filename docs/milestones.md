# Tripoli - Milestones

### [M0] Modelado del problema
+ **Objetivo**: Descomponer y analizar las historias de usuario para identificar los elementos clave del dominio, tales como roles, reglas y restricciones. El objetivo es crear un modelo que represente fielmente el contexto y los procesos implicados en la gestión de turnos de enfermería.
+ **Entregable**: Código que encapsule los elementos del dominio de forma clara y estructurada, reflejando las relaciones entre los actores (jefe de enfermería, enfermeros, etc.) y las reglas de negocio (horarios, nóminas, convenios).
+ **Criterio de Validación**: La solución será considerada viable cuando el revisor la haya verificado y aprobado, asegurando que el modelo cubra adecuadamente todos los aspectos descritos en las historias de usuario.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M1] Identificación de Reglas del Convenio
+ **Objetivo**: Identificar, analizar y documentar todas las normativas aplicables que afectan la gestión de turnos en el hospital, incluyendo aspectos como el número máximo de horas que puede trabajar un enfermero por semana, los descansos obligatorios entre turnos, los límites en la asignación de turnos nocturnos y festivos, así como las reglas para la compensación de horas extra. También se tendrán en cuenta otros aspectos importantes, como las leyes laborales y los acuerdos sindicales específicos.
+ **Entregable**: Un documento detallado que recopile y explique todas las reglas y normativas del convenio que afectan la planificación de los turnos. Este documento servirá como referencia para el desarrollo del algoritmo de asignación.
+ **Criterio de Validación**: El entregable será validado cuando el revisor lo haya verificado y aprobado, confirmando que incluye todas las normativas relevantes y que están correctamente documentadas y entendidas.
+ **HUs asociadas**: [HU001]

### [M2] Análisis de Disponibilidad del Personal
+ **Objetivo**: Recopilar la disponibilidad y preferencias de los enfermeros con respecto a sus turnos, considerando las preferencias sobre trabajar en determinados horarios, fines de semana o festivos. Se recopilarán estos datos a través de formularios o entrevistas para tener una visión completa de la situación de cada enfermero. El análisis también incluirá la capacidad para cubrir turnos urgentes o cambios de última hora.
+ **Entregable**: Una base de datos o lista estructurada que detalle la disponibilidad y preferencias de cada enfermero de la plantilla, lo que permitirá alimentar el algoritmo de asignación de turnos.
+ **Criterio de Validación**: El entregable será considerado válido cuando se haya recopilado la información completa de toda la plantilla, y el formato utilizado para la recopilación de datos sea aprobado por el revisor.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M3] Desarrollo de Algoritmo de Asignación de Turnos
+ **Objetivo**: Implementar un algoritmo que asigne turnos de manera automática, respetando las reglas del convenio laboral y las preferencias de los enfermeros, asegurando un reparto justo y equitativo de los turnos más demandados, como los nocturnos o festivos. El algoritmo debe optimizar la distribución de las horas de trabajo, garantizando que no se excedan los límites legales y que se mantenga un equilibrio en las cargas de trabajo entre los enfermeros.
+ **Entregable**: Código funcional del algoritmo de asignación de turnos, que será utilizado para gestionar los turnos de forma automática, considerando las reglas del convenio y la disponibilidad del personal.
+ **Criterio de Validación**: El algoritmo será validado mediante pruebas con conjuntos de datos de prueba, asegurando que se respetan las reglas del convenio y que el reparto de los turnos es justo. El entregable será considerado viable cuando el revisor apruebe los resultados de las pruebas.
+ **HUs asociadas**: [HU001], [HU002], [HU003]

### [M4] Validación del Algoritmo con Datos Reales
+ **Objetivo**: Probar el algoritmo utilizando datos reales de los enfermeros y los turnos del hospital, asegurando que el reparto sea justo y conforme al convenio.Realizar pruebas del algoritmo de asignación de turnos utilizando datos reales de los enfermeros y los turnos del hospital, verificando que el reparto de turnos sea justo, equitativo y en conformidad con el convenio. El objetivo es asegurar que el algoritmo funcione correctamente en un entorno real, y que los resultados sean aceptados tanto por el personal como por los gestores de turnos.
+ **Entregable**: Un informe detallado que recoja los resultados de la validación con datos reales, incluyendo análisis de la equidad en el reparto de turnos, cumplimiento de las normativas y ajuste a las preferencias del personal. El informe también incluirá cualquier ajuste o mejora realizada en el algoritmo como resultado de la validación.
+ **Criterio de Validación**: El informe será considerado válido cuando el revisor lo apruebe, confirmando que el algoritmo cumple con los objetivos establecidos y se ajusta a los requisitos del convenio y las preferencias del personal.
+ **HUs asociadas**: [HU001], [HU002]

