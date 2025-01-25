# Mailoop App (GoFiber - Mongodb - JWT)

A REST API service built to register users, companies, and email templates to the database via APIs and send the registered email templates as bulk emails to the registered companies.

# Getting Started

### Environment Config

##### Replace the .env.example file in the root directory with .env Then edit your environment variables

### Clone the repository
##### Clone this repository
```bash
➜ git clone https://github.com/burakturnaa/mailoop.git
```

##### Install dependencies
```bash
➜ go mod download
```

##### Run
```bash
➜ go run main.go
```

##### Build
```bash
➜ go build
```

### Working with makefile
##### Run
```bash
➜ make run
```

##### Build
```bash
➜ make build
```

##### Watch for file changes and reload
```bash
➜ make watch
```

If you had installed make utility

## API Reference

### Authentication
##### Login
```http
POST /api/auth/login
```
##### Register
```http
POST /api/auth/register
```
##### Check token
```http
POST /api/auth/check_token
```

### CRUD Company
##### Get all companies
```http
GET /api/company/
```
##### Get company by id
```http
GET /api/company/:id
```
##### Create company
```http
POST /api/company/
```
##### Update company
```http
PUT /api/company/:id
```
##### Delete company
```http
DELETE /api/company/:id
```

### CRUD Mail Template
##### Get all mail templates
```http
GET /api/mailtemp/
```
##### Get mail template by id
```http
GET /api/mailtemp/:id
```
##### Create mail template
```http
POST /api/mailtemp/
```
##### Update mail template
```http
PUT /api/mailtemp/:id
```
##### Delete mail template
```http
DELETE /api/mailtemp/:id
```

### Mail Sender
##### Send email
```http
POST /api/mail/send
```

### LOGS
##### Get all logs
```http
GET /api/log/
```


