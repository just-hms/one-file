# List of all available APIs divided in

- [text](#free-to-use) : blank
- [Require Login](#🔑-require-login) : 🔑
- [Require Admin](#🛠️-require-admin) : 🛠️

# Free to use

### `GET` /user/login

request

```json
{
	"username" : "sus",
	"password" : "super_secret",
}
```

response ✔️ -> status : `200`

```json
{
	"token" : "token",
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🔑 Require login

Add the token to the header request as a Bearer token, ex:

```http
GET / HTTP/1.1
Authorization: Bearer "the_actual_token"
Host: url/user/:userID/file
```

### 🔑 GET /user/:userID/file

request

response ✔️ -> status : `200`

```json
{
	"file" : "content" 
}
```


response ❌ -> status : `401` | `400` 

```json
{
	"error" : "errorMessage",
}
```

### 🔑 PUT /user/:userID/file

```json
{
	"file" : "content",
}
```

response ✔️ -> status : `200`

```json
{
	"file" : "content" 
}
```

response ❌ -> status : `401` | `400`

```json
{
	"error" : "errorMessage",
}
```

## 🛠️ Require admin

### 🛠️ POST /user

```json
{
	"username" : "new_username",
	"password" : "super_secret",
}
```

response ✔️ -> status : `201`

response ❌ -> status : `403` | `401` | `400` 

```json
{
	"error" : "errorMessage",
}
```



