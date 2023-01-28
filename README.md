# httptoy 🎉
A powerful HTTP Request & Response Service. Inspired by [httpbin.org](https://httpbin.org "httpbin.org") / [Github](https://github.com/postmanlabs/httpbin "Github") Which is now outdated.

## Features

- Light Weight & Easy to setup
- Docker image avaiable 
- One command installation

## Links
- [http://httptoy.com](http://httptoy.com "http://httptoy.com")
- [https://httptoy.com](https://httptoy.com "https://httptoy.com")
- [https://hub.docker.com/r/mayankfawkes/httptoy](https://hub.docker.com/r/mayankfawkes/httptoy "https://hub.docker.com/r/mayankfawkes/httptoy")

## Run locally

```
docker pull mayankfawkes/httptoy
docker run -p 8000:8000 mayankfawkes/httptoy
```

## API Endpoints

Under `{}` of url path are path parameters, if invalid parameters are found then they will be set to default values.

### Methods

```http
GET /get
POST /post
DELETE /delete
PATCH /patch
PUT /put
```

### Auth

```http
GET /basic/{username}/{password}
GET /basic
GET /bearer
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `username` | `string` | **Required** Username for auth |
| `password` | `string` | **Required** Password for auth |


### Compression

```http
GET /deflate
GET /brotli
GET /zlib
GET /gzip
GET /utf8
```

### Compression

```http
GET /deflate
GET /brotli
GET /zlib
GET /gzip
GET /utf8
```

### Cookies

```http
GET /cookies
GET /cookies/set/{key}/{value}
GET /cookies/delete/{key}
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `key` | `string` | **Required** Cookie name |
| `value` | `string` | **Required** Cookie value |

### Images

```http
GET /image
GET /image/png
GET /image/jpeg
GET /image/svg
GET /image/webp
```

### Redirect

```http
GET /redirect/{times}
GET /redirect-delay/{delay}/{times}
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `times` | `int` | **Required** Number of time redirect  |
| `delay` | `int` | **Required** Delay in redirect in seconds  |

### Request Inspection

```http
GET /headers
GET /header/{header_name}
GET /ip
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `header_name` | `string` | **Required** Name of the request header |


### Response

```http
GET /bytes/{count}
GET /base64en/{data}
GET /base64de/{data}

GET /delay/{seconds}
POST /delay/{seconds}
DELETE /delay/{seconds}
PUT /delay/{seconds}
PATCH /delay/{seconds}

GET /uuid
GET /body/{data}
GET /text
GET /json
GET /xml
GET /yaml
GET /html

```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `count` | `int` | **Required** Number of ramdom bytes return  |
| `data` | `string` | **Required** String data  |
| `seconds` | `int` | **Required** Duration in seconds  |



### Response Status

```http
GET /status/{code}
POST /status/{code}
DELETE /status/{code}
PUT /status/{code}
PATCH /status/{code}
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `code` | `int` | **Required** Valid HTTP [Status Codes](https://www.rfc-editor.org/rfc/rfc7231#section-6.1 "Status Codes") |
