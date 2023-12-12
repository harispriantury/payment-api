### AUTH PAYMENT API DOCUMENTATION
## Register Customer
Endpont : POST /api/auth/register/customer
Request Body :
```json
{
  "name": "haris",
  "username": "Haris28",
  "password" : "Priantury",
  "phone" : "0812748393"
}
```
Response Success :
```json
{
  "data": "register succesfully"
}
```

## Register Merchant
Endpoint : POST /api/auth/register/merchant
Request Body :
```json
{
  "name" : "merchant",
  "username" : "merchant",
  "password" : "merchant",
  "phone" : "08986958493"
}
```

Response Success :
```json
{
  "data": "register Merchant succesfully"
}
```

## Login
Endpoint : GET /api/auth/login

Request Body :
```json
{
  "username" : "haris",
  "password" : "priantury"

}
```

Response Success :
```json
{
  "message": "login succesfully",
  "token": "$2a$14$a/JJYXhQF5yYpZ1O3P3o4.87lTZQu4ciGW3nXYtfcwQdBJF8FVPwi"
}
```
Response Failed :
```json
{
  "error": "invalid credential"
}
```



## Logout
Endpoint : DELETE /api/auth/logout

Header : "X-API-TOKEN" : $2a$14$a/JJYXhQF5yYpZ1O3P3o4.87lTZQu4ciGW3nXYtfcwQdBJF8FVPwi

Response Success :
```json
{
  "data": "logout succesfully"
}
```
