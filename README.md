# Mailoop App (GoFiber - Mongodb - JWT)

A REST API service built to register users, companies, and email templates to the database via APIs and send the registered email templates as bulk emails to the registered companies.

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
##### Send mail
```http
POST /api/mail/send
```

### LOGS
##### Get all logs
```http
GET /api/log/
```


