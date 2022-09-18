# payment-api

## Comando para ejecutar el proyecto con Docker
`docker compose up`
Este comando monta la base de datos de mongo, y monta el back despues.
Esperar a que en la consola de docker salga [GIN-debug] Listening and serving HTTP on :8080
ya que ahi es cuando todo se deplego correctamente.
No es necesario crear .env ya que en el Dockerfile ya estan definidas las variables de entorno

una vez desplegado acceder a la siguente url que es la documentacion de swagger:
http://localhost:8080/docs/index.html#/

## Si se desea correr localmente o correr las pruebas unitarias es necesario tener go instalado en la maquina

## Si se desea correr localmente o quitar los errores del vscode por el go.sum correr los siguintes comandos
`go get .`
`go run .`

## Si se desean correr las pruebas unitarias es necesario haber corrido el go get . antes, y si o si tener corriendo el contenedor de docker con mongodb
para correr las pruebas unitarias utilizar el siguente comando:
`go test ./tests/...`

## Comandos de desarrollo - ignorar
### Comando para generar archivos de documentacion
`swag init -g core/swagger.go controllers/user.controller.go`
### Comando para correr con hot reload
`air`
