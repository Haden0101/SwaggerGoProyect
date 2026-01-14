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

# NOTA IMPORTANTE

Cuando se añada otro "EndPoint" es necesario ejecutar de nuevo el comando:
```bash
swag init -g main.go
```
Esto para que se vuelva a generar la estructura en el swagger docs, ya luego 
se ejecuta de nuevo el run para el main y poder visualizar el EP.
```bash
go run main.go
```

# Documentación de las anotaciones para que pinte en Swagger:

🔹 @Summary (corto) -> Descripción breve que se ve en la lista.
  + // @Summary Obtener usuario

🔹 @Description (larga) -> Explicación detallada del endpoint.
  + // @Description Retorna un usuario por ID

🔹 @Tags -> Agrupa endpoints en Swagger UI.
  + // @Tags users

🔹 @Accept -> Qué tipo de request acepta. 
  + // @Accept json

🔹 @Produce -> Qué tipo de response devuelve.
  + // @Produce json

🔹 @Param (MUY IMPORTANTE) -> Muestra los parametros a utilizar
  + // @Param id path int true "ID del usuario"

  Formato: @Param

  + // @Param <nombre> <ubicación> <tipo> <required> <descripción>
  + // @Param    id        path      int     true      "User ID"

  📍 Ubicaciones comunes:

  | Ubicación  | Uso           |
  | ---------- | ------------- |
  | `path`     | `/users/{id}` |
  | `query`    | `?page=1`     |
  | `body`     | JSON          |
  | `header`   | Headers       |
  | `formData` | Formularios   |

🔹 @Success / @Failure -> Respuesta satisfactoria / Respuesta Fallida
  + // @Success 200 {object} User
  + // @Failure 404 {string} string

  📌 Tipos:

    + {object} → struct
    + {array} → slice
    + {string} → texto plano

🔹 @Router (OBLIGATORIA) -> Para reconocer el path principal y el metodo autilizar.
  + // @Router /users/{id} [get]

  📌 Formato:

    + // @Router <ruta> [metodo]
  ❌ Sin esto → NO aparece en Swagger

## Documentar MODELOS (structs)
```go
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```
📌 Usa siempre json:"campo".

# Formas de implementación::

## Requests con BODY (POST / PUT)

// @Param user body User true "Usuario a crear"

// @Success 201 {object} User

// @Router /users [post]

## Autenticación JWT (opcional)
 + Definir seguridad (una vez)
    + // @securityDefinitions.apikey BearerAuth
    + // @in header
    + // @name Authorization
 + Usar en un endpoint:
    + // @Security BearerAuth
 + Códigos de respuesta múltiples:
    + // @Success 200 {object} User
    + // @Success 204 "No Content"
    + // @Failure 400 {string} string

## EJEMPLO MINIMO FUNCIONAL

```go
// Health godoc
// @Summary Health check
// @Tags system
// @Produce plain
// @Success 200 {string} string
// @Router /health [get]
func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
```

# 🧠 Reglas IMPORTANTES (memorízalas)

✅ Los comentarios van JUSTO encima de la función
✅ Usar // (no /* */)
✅ Ejecutar swag init después de cambios
✅ @Router debe coincidir con la ruta real
✅ El struct debe ser exportado (mayúscula)

# 🧠 Resumen ultra corto
Para que aparezca en Swagger un endpoint necesita:

+ @Summary
+ @Tags
+ @Router

Y la API necesita:

+ @title
+ @version
+ @BasePath