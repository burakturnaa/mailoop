# Mailoop App (GoFiber - Mongodb - JWT)

A REST API service built to register users, companies, and email templates to the database via APIs and send the registered email templates as bulk emails to the registered companies.

## API Reference

### Authentication
#### Login
```http
POST /api/auth/login
```
#### Register
```http
POST /api/auth/register
```

### CRUD Company
```http
GET /api/company/
GET /api/company/:id
POST /api/company/
PUT /api/company/:id
DELETE /api/company/:id
```
