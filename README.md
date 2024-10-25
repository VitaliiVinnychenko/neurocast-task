Simple Notes Application (Go + React + MongoDB)
----------------------------------------------------

Prerequisites
-------------

- Git
- Docker
- [Go 1.20](https://go.dev/doc/install) and above
- Node 20

_If you are not using Docker, you should also have:_

- MongoDB

Getting Started
---------------
Clone the repository

```bash
# Clone Project
git clone https://github.com/VitaliiVinnychenko/neurocast-task.git neurocast-task

# Change Directory
cd neurocast-task
```

#### Using Docker

```bash
# Build & Create Docker Containers
docker-compose up -d
```

### Frontend

#### Using Local Environment

```bash
# Install dependencies
npm install

# Start application
npm start
```

### Backend

#### Using Local Environment

```bash
# Copy Example Env file
cp ./env.example .env

# Change MongoDB/Redis URI and Database Name

# MONGO_URI=<mongo_uri>
# MONGO_DATABASE=<db_name>

# Download Modules
go mod download

# Build Project
go build -o go-starter .

# Run the Project
./go-starter
```

The application starts at port 8080:

- `POST /v1/notes` Create a new note
- `GET /v1/notes` Get paginated list of notes
- `GET /v1/notes/:id` Get a one note details
- `PUT /v1/notes/:id` Update a note
- `DELETE /v1/notes/:id` Delete a note

---

- `GET /swagger/*` Auto created swagger endpoint

You can also see: http://localhost:8080/swagger/index.html

> If you want to add new routes and swagger docs, you should run ```swag init```
> See: [Gin Swagger](https://github.com/swaggo/gin-swagger)

Project Structure
-----------------

```
├── controllers         # contains api functions and main business logic
├── docs                # swagger files 
├── logs
├── middlewares         # request/response middlewares
│   └── validators      # data/request validators
├── models              
│   └── db              # collection models
├── routes              # router initialization
└── services            # general service & database actions
```
