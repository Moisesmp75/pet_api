# HairyPaws API

Esta es una API para gestionar información de mascotas y adopciones.

## Descripción

El API de HairyPaws proporciona endpoints para realizar operaciones relacionadas con mascotas, usuarios, adopciones y más. Utiliza Fiber, un framework web ligero y rápido para Go.

## Características

- **Gestión de Mascotas:** CRUD completo para administrar mascotas.
- **Adopciones:** Permite realizar adopciones de mascotas.
- **Usuarios:** Manejo de usuarios, autenticación y roles.
- **Visitas:** Registro de visitas a mascotas.

## Tecnologías Utilizadas

- Go (version 1.22.2)
- Fiber (Go Framework)
- MySQL (Base de datos)

## Instalación y Uso

<ol>
  <li>Clona este repositorio</li>
  <pre>git clone https://github.com/Moisesmp75/pet_api.git</pre>
  <li>Ejecutar el programa</li>
  <pre>En caso de tener otra version de go hacer lo sgte:
  1. Eliminar los archivos go.mod y go.sum
  2. Abrir el cmd y ejecutar los comandos:
      go mod init pet_api
      go mod tidy
  </pre>
  <pre>go run main.go</pre>
  <li>El programa se ejecuta en el puerto 3000, para ver la documentacion diriguirse al enlace</li>
  <pre>http://localhost:3000/swagger/index.html#</pre>
</ol>
