# Gestor de Dependencias

El gestor de dependencias escogido para este proyecto es el propio [`go.mod`](../go.mod) por los siguientes motivos:

1. **Resolución automática de dependencias**  
    Una de las principales ventajas de go.mod es que las dependencias se gestionan automáticamente. Al importar un paquete en el código y ejecutar comandos como `go build`, `go test` o `go run`, Go Modules descarga y resuelve las versiones adecuadas de las dependencias de forma transparente, sin necesidad de configuraciones manuales adicionales.

2. **Se puede trabajar en cualquier directorio, facilitando la portabilidad**  
    Con Go Modules, no es necesario trabajar dentro del directorio `$GOPATH`, lo que ofrece mayor libertad para organizar el proyecto en cualquier ubicación. Esto facilita la colaboración en equipos y la integración con sistemas de control de versiones como Git, haciendo que el proyecto sea más portable entre diferentes máquinas y entornos.

3. **Es propio del lenguaje**  
    Al ser parte del ecosistema oficial de Go, go.mod está diseñado específicamente para trabajar de manera óptima con las herramientas del lenguaje. Esto garantiza una experiencia uniforme y elimina la necesidad de instalar herramientas externas, simplificando el flujo de trabajo.

4. **Recomendación del profesor**
