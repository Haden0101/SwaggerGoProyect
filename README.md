# INTEGRACIÓN SWAGGER CON APLICACION GO

Este proyecto implementa Swagger para la documentación de una aplicación "Go", para eso se usa la herramienta swaggo/swag que es un estandar.

# Paso a seguír:

## Instalar el binario de swag
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Comprobamos la versión
```bash
swag --version
```

# Instalar dependencias en el proyecto 

Recordar hasta este punto que los archivos .go no tienen que estar ninguno corructo, vacío o con errores, ya que esto evita la creación de la carpeta "docs" necesaria para la implementación de swag.

## Para Swagger-UI
```bash
go get github.com/swaggo/http-swagger
```

## Para los archivos generados
```bash
go get github.com/swaggo/swag
```

# Despues de la implementación del codigo swagg:

ejecutamos el comando en consola para crear los docs en swagger, esto hara que se actualice la carpeta "docs".
```bash
swag init -g main.go
```

Respuesta esperada: 
```bash
2026/01/14 13:40:40 Generate swagger docs....
2026/01/14 13:40:40 Generate general API Info, search dir:./
2026/01/14 13:40:40 Generating handlers.User
2026/01/14 13:40:40 create docs.go at docs/docs.go
2026/01/14 13:40:40 create swagger.json at docs/swagger.json
2026/01/14 13:40:40 create swagger.yaml at docs/swagger.yaml
```

# Levantamos el proyecto
```bash
go run main.go
```

## NOTA IMPORTANTE

Cuando se añada otro "EndPoint" es necesario ejecutar de nuevo el comando:
```bash
swag init -g main.go
```
Esto para que se vuelva a generar la estructura en el swagger docs, ya luego 
se ejecuta de nuevo el run para el main y poder visualizar el EP.
```bash
go run main.go
```