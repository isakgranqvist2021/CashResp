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

URL: /api/auth/register<br/>
Method: POST<br/>

```json
{
	"email": "String",
	"password": "String"
}
```

URL: /api/auth/login<br/>
Method: POST<br/>

```json
{
	"email": "String",
	"password": "String"
}
```
