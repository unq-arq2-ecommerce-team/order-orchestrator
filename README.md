# order-orchestrator

## Tecnologías:

- [Golang](https://go.dev/)
- [Gin (WEB API)](https://gin-gonic.com/)
- [MongoDB](https://www.mongodb.com/)

## Prerequisitos:

- Go 1.20 or up / Docker

## Swagger

Instalar swag localmente (se necesita go 1.20 or up)

```
go install github.com/swaggo/swag/cmd/swag@v1.8.10
```

Para actualizar la api doc de swagger, ejecutar en el folder root del repo:

```
swag init -g src/infrastructure/api/app.go
```

Luego de levantar la api e ir al endpoint:

```
http://localhost:<port>/docs/index.html
```


## Inicialización y ejecución del proyecto (docker)

### Pasos:

1) Ir a la carpeta root del repositorio

2) Construir el Dockerfile (imagen) del servicio

```
docker build -t order-orchestrator .
```

3) Ejecutar la imagen construida. Es importante tener ejecutando los servicios para ejecutar su funcionalidad:

- users-service
- products-orders-service
- payment-service

Tambien, si se desea se puede cambiar las envs por otras de las que estan. Se recomienda utilizar el mismo puerto externo e interno para que funcione correctamente swagger.

```
docker run -p <port>:8082 --env-file ./resources/local.env --name order-orchestrator order-orchestrator
```

Nota: agregar "-d" si se quiere ejecutar como deamon

```
docker run -d -p <port>:8082 --env-file ./resources/local.env --name order-orchestrator order-orchestrator
```

Ejemplo:

```
docker run -d -p 8082:8082 --env-file ./resources/local.env --name order-orchestrator order-orchestrator
```

4) En un browser, abrir swagger del servicio en el siguiente url:

`http://localhost:<port>/docs/index.html`

Segun el ejemplo:

`http://localhost:8082/docs/index.html`

5) Probar el endpoint health check y debe retornar ok

6) La API esta disponible para ser utilizada

