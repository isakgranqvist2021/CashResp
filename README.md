<!-- @format -->

all responses from the api will have this structure<br/>
Content-Type: application/json<br/>

```json
{
	"message": "String",
	"success": "Boolean",
	"status_code": "Int",
	"data": "Object"
}
```

```json
POST /api/auth/register
{
	"email": "String",
	"password": "String"
}
```

```json
POST /api/auth/login
{
	"email": "String",
	"password": "String"
}
```
