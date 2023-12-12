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