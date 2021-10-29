![Golang React Postgres App](https://i.ibb.co/PhZcWVR/golang.jpg)

## - Descripción
Es una aplicación básica escrita en `golang`, donde se hace uso del framework `fiver` para su desarrollo, la misma tiene una naturaleza SPA debido a que en el front-end encontramos `React Js`, y además para la minipulación de los datos se usa como ORM `gorm` y `postgres`. 

Esto es una práctica básica la cual nace del deseo de incursionar en esta tecnologia la cual tiene muy buena reputación. El código se puede mejorar de muchisimas maneras ya que le dediqué corto tiempo, sin embargo cumple con la naturaleza de `api rest` para el back.

## - Estructura planteada:

- api ( Aplicación backend desarrollada en GOLANG )
  - controllers
    - main.go
    - product.go
  - database
    - connection.go
  - models
    - price.go
    - product.go
  - routes
    - main.go
    - product.go
  - .env
  - data.json
  - docker-compose.yml
  - Dockerfile
  - go.mod
  - go.sum
  - main.go
- client (Aplicaciión desarrollada en React JS + Typesrcript)
  - node_modules
  - public
  - src
    - api
      - ReqResapi.tsx
    - hooks
      - useProduct.tsx
    - interfaces
      - Response.tsx
    - screens
      - HomeScreen.tsx
  - App.css
  - App.test.tsx
  - etc...

## - Cliente (Front-End)

- Está desarrollado con `React Js & typesript`, tiene lo necesario para mostrar el historial de precios y su historial de cambios.
- Para ejecutar de se debe ir a la carpera `client` y ejecutar `yarn install` y luego `yarn start`

## - API

- Está desarrollada con `golang` bajo el framework `fiber` y cuenta con `gorm` como ORM para las consultas y las transacciones con la DB.
- Para ejecutar el api se debe acceder a la carpeta `api` y ejecutar `go run main.go`.

### - Conexión con la base de datos

- La base de datos utilizada en la app desarrollada fué `postgres` esta se empleó bajo el manejador `psql`, asi que se debe descargar (en dado caso de no tener) y crear un usuario llamado **Priceomatic** con el comando `CREATE ROLE priceomatic LOGIN PASSWORD 'priceomatic';`
- El segundo paso es crear una base de datos llamada **priceomatic** con el comando `CREATE DATABASE priceomatic;`
- El ejecutar la app, la misma contará con dos tablas `products y prices` está se crearán automaticamente al migrar los modelos creados.
- Adicionalmente se leerá el archivo `data.json` ubicado en la raiz del proyecto, este insertará contenido inicial en la tabla `products` y así asi como se da el proceso de `seeds`.

### - Controladores

- `prices.go` (Controlador de la relación **precios**): este posee los siguienes métodos.
  - **GetProducts:** Obtiene los productos insertados en la base de datos.
  - **GetProduct:** Obtiene un producto especifico, este recibe el `id` del producto.
  - **NewProduct:** Crea o inserta un nuevo producto en la base de datos.
  - **DeleteProduct:** Borra un producto, este recibe el `id` del producto a borrar
  - **ModifyProduct:** Modifica la data de un producto, si el precio enviado desde el cliente es dintinto al que posee la tupla.
  - **GetPrices:** Obtiene el historial de precios de un producto, recibe el id del producto.

### - Rutas

Todas las rutas se complementan de la uri `/api/v1`, sin embargo hay dos grupos, uno que puede usarse para rutas de un `landigPage`, dado por el grupo `api.Group("/")` y otro especifico del grupo de precios, dado por `api.Group("/product")`, de estos se detallan los siguientes endpoitns

* `get -> /`: Obtiene todos los produtos
* `get -> /:id` Recibe un `id` de producto y lo retorna
* `put -> /:id` Recibe un `id` de producto y modifica sus datos, de acuerdo a lo que se reciba del cliente
* `patch -> /:id` Recibe un `id` de producto y modifica sus datos, de acuerdo a lo que se reciba del cliente siempre y cuando el precio sea distinto
* `delete -> /:id` Recibe un `id` de producto y lo borra
* `get -> /:id/prices` Recibe un `id` de producto y devuelve el historial de precios de este

### - Mejoras futuras

- Middlewares de seguridad.
- Segmentación de configuración para variables o datos globales
- Mejor manejo de las peticiones `PATCH`
- Comprensión a profundidad del framework `FIBER`
- Creación de helpers que puedan usarse en cualquier parte de la app (`utils`)

### Algunas capturas de pantallas

- Mostrando productos en cliente
  ![Mostrando productos en cliente](https://i.ibb.co/x6NFNTf/image.png)

- Mostrando historial de un producto en cliente
  ![Mostrando historial de un producto en cliente](https://i.ibb.co/5R0fyGF/image.png)  

- Logger de peticiones en api
  ![Logger de peticiones en api](https://i.ibb.co/BGHzJQD/image.png)

- Data inicial insertada luego de leer el archivo `data.json`
  ![Data inicial insertada luego de leer el archivo data.json](https://i.ibb.co/93DbsM8/image.png)

- Data inicial en un cliente grafico
  ![Data inicial en un cliente grafico](https://i.ibb.co/9pTXVRy/image.png)

- Consultando el historial de precios de un producto
  ![Consultando el historial de precios de un producto](https://i.ibb.co/QQpK1r1/image.png)

- Consultado el historial de productos
  ![Consultado el historial de productos](https://i.ibb.co/9VnMQcP/image.png)
