### PAYMENT API DOCUMENTATION
## Create Payment
Endpont : POST /api/payment

Headers : "X-API-TOKEN" : $2a$14$HbSdMcTmRUGYJSbSO7tJ1OkvaYMa1/Ev5pcFqryuGSNC68edPiOe6

Request Body :
```json
{
  "Amount": 123000,
  "MerchantId": "merchant",
  "TransactionId": "T7",
  "Status": "pending"
}
```
Response Success :
```json
{
  "data": {
    "ID": 1,
    "Amount": 123000,
    "CustomerID": "haris",
    "MerchantId": "merchant",
    "TransactionId": "T7",
    "Status": "pending",
    "Updated": 1702402645246,
    "Created": 1702402645
  },
  "message": "succes create payment"
}
```


## Get All Payment
Endpont : GET /api/payments

Headers : "X-API-TOKEN" : $2a$14$HbSdMcTmRUGYJSbSO7tJ1OkvaYMa1/Ev5pcFqryuGSNC68edPiOe6
Response Success :
```json
{
  "data": [
    {
      "ID": 1,
      "Amount": 123000,
      "CustomerID": "haris",
      "MerchantId": "merchant",
      "TransactionId": "T7",
      "Status": "pending",
      "Updated": 1702402645246,
      "Created": 1702402645
    }
  ],
  "message": "OK"
}
```



## Get Status Payment By ID
Endpont : GET /api/payments/:id

Headers : "X-API-TOKEN" : $2a$14$HbSdMcTmRUGYJSbSO7tJ1OkvaYMa1/Ev5pcFqryuGSNC68edPiOe6
Response Success :
```json
{
  "data": {
    "ID": 1,
    "Amount": 123000,
    "CustomerID": "haris",
    "MerchantId": "merchant",
    "TransactionId": "T7",
    "Status": "pending",
    "Updated": 0,
    "Created": 1702402645
  },
  "message": "OK"
}
```


