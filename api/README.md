# Price o' Matic

En Klare necesitamos añadir un nuevo microservicio que se encargue de listar los precios de cada uno de nuestos productos.
Este debe ser simple y debe considerar lo siguiente:

* **Debe** ser escrito en Go
* **Debe documentar** la solución propuesta, considerar lo siguiente:
  * **Debe incluir** lo necesario para ejecutar y verificar los requerimientos.
  * **Debe incluir** un TODO, que indique que aspectos cree que puede mejorar o agregar en el futuro de la solución entregada.
* Debe modelar la base datos considerando las siguientes tablas `Product` y `Price`.
* Debe crear los siguientes endpoints:
    * GET `/api/v1/products` <-- Lista todos los productos
    * GET `/api/v1/products/:id` <-- Obtiene el producto :id
    * POST `/api/vi/products` <-- Inserta un producto en la base de datos
    * PATCH `/api/vi/products/:id` <-- Actualiza uno o mas campos del producto :id
    * DELETE `/api/vi/products/:id` <-- Elimina el producto :id
    * GET `/api/v1/products/:id/prices` <-- Devuelve lista de la historia de los precios del producto :id

    Nota: Puede proponer el modelo de API que usted considere.

* **Debe escribir** un script que permita insertar el json incluido (`data.json`) como seed inicial a la base de datos, puede utilizar lo que considere necesario.
* **Debe incluir el siguiente requisito**: Cada vez que se modifique un precio, debe quedar un historial.
* [Bonus] Crear una ruta `GET /`, que renderice un frontend simple mostrando un resumen del contenido en la base de datos (Puede usar lo que considere necesario, react, vuejs, libreria de render de Go, etc). Hint: No importa el aspecto visual.

## Notas

* Dudas y/o consultas por correo a ernesto.olivares@klare.cl o al whatsapp de mi firma.
* Revisar los siguientes conceptos de Go: structs, structs tags, error handling.
* Se sugiere utilizar GORM (https://gorm.io/index.html) como ORM para manejar el acceso a la base de datos, pero puede ocupar lo que considere.
* El proyecto usa fiber (https://github.com/gofiber/fiber) como framework web, se recomienda revisar
* Puede utilizar la base de datos que considere, el proyecto viene con un setup por defecto de `postgres`.
* El proyecto se probará en entorno dockerizado, por ende, debe levantar correctamente usando el comando `docker-compose up`.
* Una vez levantado el proyecto, se ejecutará el script para inicializar la base de datos con los datos provistos.
* [Bonus] Se permite hardcodear los strings de configuración (base de datos, etc), pero sumará puntos si carga las variables de entorno desde el archivo local.env.
  * Utilizar `postgres:5432` como dns de conexión al utilizar docker-compose

## A Evaluar

* Plazo de entrega (Avisar si requiere más tiempo)
* Correctitud de lo entregado.
* Planteamiento de la solución.
* Orden y estructura del código.
* Documentación entregada.
* Bonus.