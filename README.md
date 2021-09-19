<!-- @format -->

```json
Response Structure
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
