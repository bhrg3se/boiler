# Golang Boilerplate

## Setup and running

#### 1. Using `docker-compose`
- Run `docker-compose up`

#### 2. Running natively

- Make sure you have go:1.15+ installed
- Create a config file in /etc/boiler/config.toml (use setup/config.toml as a reference)
- Start a postgres database and configure the host,name,user and password in /etc/boiler/config.toml
- Run setup/init.sql in postgres console.
- Run `go build -o boiler && ./boiler`
> You may run into permission issues because it will try to create private key file and log file.
> You may just run the binary with sudo (not recommended)

> Add -f argument while running the binary to write log into file instead of stdout.


## File Structure

- `utils/` (utility functions)
- `setup/` (docker-compose, config samples and initial sql file)
- `models/` 
- `routes/`(API endpoint routes and middlewares) 
- `store/`(initialisation of database and other dependencies)
- `api/` (API handler functions)
- `postman/` (Postman exports)


## API endpoints

- `/api/v1/auth/signup` 
- `/api/v1/auth/login` 
- `/api/v1/auth/logout` 
- `/api/v1/profile`

> You can find the postman exports in postman/

## Used libraries

- `github.com/go-chi/chi` for routing 
- `github.com/go-playground/validator/v10` for validating json fields
- `github.com/golang-jwt/jwt` for jwt tokens 
- `github.com/google/uuid` for generating UUID 
- `github.com/sirupsen/logrus` for logging
- `github.com/spf13/viper` for reading config file
- `github.com/stretchr/testify` for mocks

## Config file reference

### database
- user: `string`   
- password: `string` 
- host: `string`
- port: `string`
- name: `string` (database name)
- ssl: `boolean` (if SSL is enabled or not)
- caCertPath: `string`  (only applies if SSL is enabled)
- userCertPath: `string` (only applies if SSL is enabled)
- userKeyPath: `string` (only applies if SSL is enabled)

### logging
- logging: `enum{"TRACE","DEBUG","ERROR","INFO"}` 

### server
listen: `string` (listen address, for example: "127.0.0.1:8081" )


## Notes
- used postgres as database.
- used JWT for authentication instead of saving tokens in database.
- added an extra `/api/v1/profile` API to show full authentication flow.
   
