
# Aplicación de Lista de Tareas (To-Do List)

## Descripción

Esta es una REST API de lista de tareas (To-Do List) construida con Go, utilizando el framework Gin para el enrutamiento HTTP y GORM para la gestión de la base de datos. La aplicación permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en una lista de tareas, almacenando los datos en una base de datos PostgreSQL. La aplicación está configurada para ejecutarse en un entorno Docker con Docker Compose.

## Características

- Crear una nueva tarea
- Consultar todas las tareas
- Actualizar una tarea existente
- Eliminar una tarea

## Requisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Instalación y Ejecución

### 1. Clona el repositorio

```bash
git clone https://github.com/Maxbeurer/SADSeminary1.git
cd SADSeminary1
```

### 2. Configura el entorno

Asegúrate de tener Docker y Docker Compose instalados en tu sistema.

### 3. Construye y ejecuta los contenedores

En la raíz del proyecto, ejecuta el siguiente comando para construir y levantar los contenedores:

```bash
docker-compose up --build
```

Este comando:
- Construirá la imagen de Docker para la aplicación Go.
- Levantará los contenedores para la aplicación y la base de datos PostgreSQL.

### 4. Verifica que la aplicación esté corriendo

Una vez que los contenedores estén levantados, puedes acceder a la API en `http://localhost:8080/todos`.

## Uso

La API ofrece las siguientes rutas para realizar operaciones en la lista de tareas:

### Crear una tarea

- **Endpoint**: `POST /todos`
- **Descripción**: Crea una nueva tarea.
- **Cuerpo de la solicitud**:
  ```json
  {
      "title": "Título de la tarea",
      "status": "incompleta"
  }
  ```
- **Ejemplo**:
  ```consola
  curl -X POST http://localhost:8080/todos -d '{"title": "Titulo de la tarea", "status": "incomplete"}'
  ```
### Consultar todas las tareas

- **Endpoint**: `GET /todos`
- **Descripción**: Obtiene todas las tareas existentes.
- **Ejemplo**:
  ```consola
  curl -X GET http://localhost:8080/todos
  ```
### Consultar la tarea por ID

- **Endpoint**: `GET /todos:id`
- **Descripción**: Obtiene una tarea existente por id.
- **Ejemplo**:
  ```consola
  curl -X GET http://localhost:8080/todos/1
  ```
### Actualizar una tarea

- **Endpoint**: `PUT /todos/:id`
- **Descripción**: Actualiza los detalles de una tarea existente.
- **Cuerpo de la solicitud**:
  ```json
  {
      "title": "Nuevo título de la tarea",
      "status": "completa"
  }
  ```
- **Ejemplo** :
  ```consola
  curl -X GET http://localhost:8080/todos/1 -d '{"title": "Titulo de la tarea", "status": "complete"}'
  ```

### Eliminar una tarea

- **Endpoint**: `DELETE /todos/:id`
- **Descripción**: Elimina una tarea específica.
