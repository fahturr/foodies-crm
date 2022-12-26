# API SPEC

## AUTH

### POST /auth/login
Login Sales

Request Body :

```json
{
  "username": "",
  "password": ""
}
```

Response Body :

```json
{
  "status_code": 200,
  "message": "",
  "payload": {
    "token": ""
  }
}
```

## CUSTOMER

### GET /customer/list
List Semua Customer

Response Body :

```json
{
  "status_code": 200,
  "message": "",
  "payload": [
    {
      "name": "",
      "phone": "",
      "image": ""
    }
  ]
}
```

### POST /customer/add
Tambah Customer

Request Body :

```json
{
  "name": "",
  "email": "",
  "phone": "",
  "address": "",
  "urban": "",
  "birthdate": "",
  "traffic_source": "",
  "label": ""
}
```

Response Body :

```json
{
  "status_code": 201,
  "message": ""
}
```
