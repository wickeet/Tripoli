# Generador de Horarios para Enfermeros

## Descripción

Esta aplicación está diseñada para generar horarios de trabajo para enfermeros en un hospital, asegurando el cumplimiento de una serie de restricciones laborales. Las principales características de la aplicación incluyen:

- **Generación automática de horarios**: La aplicación asigna turnos de trabajo para los enfermeros de acuerdo a las restricciones impuestas, como el número máximo de horas trabajadas y el límite de días consecutivos que un enfermero puede trabajar.
- **Visualización de horarios**: Cada enfermero puede consultar su horario personal de forma clara y detallada.
- **Intercambio de turnos**: La aplicación permite a los enfermeros intercambiar turnos entre ellos, siempre que el intercambio no incumpla las restricciones establecidas.

## Funcionalidades

### 1. **Generación de horarios**
   - La aplicación genera automáticamente los turnos de trabajo para cada enfermero.
   - Se tienen en cuenta restricciones como:
     - Número máximo de horas trabajadas por semana.
     - Límite de días consecutivos trabajados.
     - Descanso obligatorio después de un número específico de turnos.
     - Turnos de día, tarde y noche distribuidos equitativamente entre el personal.
   
### 2. **Visualización de horarios**
   - Los enfermeros pueden iniciar sesión en la aplicación y ver su horario asignado.
   - La visualización incluye:
     - Fecha y hora del turno.
     - Tipo de turno (día, tarde, noche).
     - Duración del turno.

### 3. **Intercambio de turnos**
   - Los enfermeros tienen la posibilidad de solicitar un intercambio de turnos con otro compañero.
   - Las solicitudes de intercambio solo serán aprobadas si:
     - No se excede el número máximo de horas trabajadas por semana.
     - No se viola el límite de días consecutivos trabajados.
     - Ambos enfermeros cumplen con sus descansos obligatorios.

## Requisitos del sistema

- **Lenguaje de programación**: 
- **Base de datos**: 
- **Librerías adicionales**:

## Licencia

Este proyecto está licenciado bajo la [MIT License](./LICENSE).

