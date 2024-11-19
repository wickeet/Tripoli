# Gestor de Tareas

El gestor de tareas escogido para este proyecto será [`Make`](../Makefile) por los siguientes motivos:

1. **Simplicidad y un Estándar Amplio**  
    Es una herramienta clásica ampliamente conocida, con soporte nativo en la mayoría de sistemas Unix/Linux y suele estar preinstalada, eliminando la necesidad de configurar herramientas adicionales en estos sistemas.
    Otros como Task o Mage, requieren de una instalación como herramienta adicional.

2. **Eficiencia**  
    Make solo ejecuta las tareas necesarias, es decir, para compilar un programa solo recompilará los archivos que hayan cambiado, basándose en timestamps, haciéndolo ideal para proyectos grandes donde no es eficiente reconstruir todo desde cero.
    Aunque Task realiza esto, lo hace de forma distinta. Mientras que Make se basa en timestamps por lo que se hace implícitamente, en Task hay que especificarlo usando la regla `sources`, necesitando una configuración manual en vez de automatizada.

3. **Independencia del lenguaje**  
    En cualquier proyecto real no se usa un único lenguaje de programación, por lo que Make es una herramienta fantástica ya que es independiente del lenguaje y es muy útil para cualquier lenguaje o tecnología que se quiera gestionar.
    Otros gestores como Mage, hasta para las tareas más sencillas depende de escribir código Go para ello.

4. **Extensibilidad**  
    Make se puede combinar con otras herramientas o se pueden incluir variables y scripts personalizados adaptándose muy bien a lo que necesite el proyecto, incluso ejecutando órdenes de shell directamente.
    Otros gestores como Goreleaser están enfocados en manejar despliegues, siendo incapaces de gestionar tareas como las pruebas o la limpieza.

5. **Compatibilidad Multiplataforma**  
    Funciona en la mayoría de Sistemas Operativos y con herramientas como GNU Make, se puede usar en Windows.

6. **Documentación y Comunidad**  
    Tiene una extensa documentación y una comunidad activa, por lo que encontrar soporte o ejemplos es relativamente sencillo.
    Otros como Mage ya no ofrece soporte, o la comunidad es mucho más pequeña, como Just, haciendo que encontrar soluciones a problemas que puedan surgir puede ser más tedioso.