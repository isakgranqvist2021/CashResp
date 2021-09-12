<!-- @format -->

all responses from the api will have this structure
Content-Type: application/json

```json
{
    "message": String,
    "success": Boolean,
    "status_code": Int,
    "data": Object
}
```

URL: /api/auth/register
Method: POST

```json
{
    "email": String,
    "password": String
}
```

URL: /api/auth/login
Method: POST

```json
{
    "email": String,
    "password": String
}
```
