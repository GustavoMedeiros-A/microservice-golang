# 1 - StepUp frontend

- The first thing is setup the front end, just the basic front-end to render a webpage
  and started

# 2 - Create broken service

- Create a folder to the broker-service
- Go to the folder and run **go mod init**
- After, create a main.go and add **go get github.com/go-chi/chi/v5** in the project
- **go get github.com/go-chi/chi/v5/middleware**
- **go get github.com/go-chi/chi/cors**
- Run the code of the class 12

- Run in the broker-service folder, **go run ./cmd/api/**

# 3 - Create a docker image to the Broken Service

- Create the dockerfile in the broken service directory, set all the configurantions that we write that
- After that, create a docker-compose.yaml to run the docker and create the image of the project

# 4 - Adding FRONT-END button

- Add button and run in the front-end folder **go run ./cmd/web**

# 5 - Create a helpers to read and write JSON file AND create a makefile to docker

# 6 - Create a auth-service

- Create the folder
- run **go mod init authentication**
- Add in past data the models to user
- Create the routes.go and add the chi, chi middleware and go on, same of we did in the Broker-service

- **add database postgres config** -> go get github.com/jackc/pgconn + go get github.com/jackc/pgx/v4 + go get github.com/jackc/pgx/v4/stdlib
- Config the database in the root of auth-service project
- add the users.sql to the database (run a query)

- Create route and handler to accept JSON

# 7 - Connect the broker service into the auth-service

- After finishing to config the auth service, add verify password and the Login method, you should pass in the Broker and go the the auth-service
- Doing that, you'll verify if the credentials are correct and the user can login
- Pass the broker, go to the auth and return the right Status

# - Create the microservice LoggerService (Connect to mongoDB, to salve the log)

- add mongo go **go get go.mongodb.org/mongo-driver/mongo** and **go get go.mongodb.org/mongo-driver/mongo/options**
- Create all the struct of the service, handlers, helpers and etc, add the docker_file and the mongo to the docker folder
- After done all the steps to build the logger service, we now gonna to add logger after the user is authenticated
- In the handler.go of the AuthenticationService, add some configuration, the function **logRequest** is tell you better
  what to do
