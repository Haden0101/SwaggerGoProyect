## INTEGRACIÓN SWAGGER CON APLICACION GO

Este proyecto implementa Swagger para la documentación de una aplicación "Go", para eso se usa la herramienta swaggo/swag que es un estandar.

## Paso a seguír:

# Instalar el binario de swag
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

# Comprobamos la versión
```bash
swag --version
```

## Instalar dependencias en el proyecto 

Recordar hasta este punto que los archivos go no tienen que estan ninguno corructo o con errores, ya que esto evita la creación de la carpeta "docs" necesaria para la implementación de swag.

# Para Swagger-UI
```bash
go get github.com/swaggo/http-swagger
```

# Para los archivos generados
```bash
go get github.com/swaggo/swag
```