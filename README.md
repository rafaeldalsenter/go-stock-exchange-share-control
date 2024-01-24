# go-stock-exchange-share-control

A simple application for calculating the average sale and purchase price of stock exchange shares to practice Golang

## CURLs

Add transaction

```cmd
curl --location 'http://localhost:3000/stock/CODE/transaction' \
    --header 'Content-Type: application/json' \
    --data '{
        "quantity": 12.0,
        "value": 13.25,
        "tax": 0,
        "type": "purchase",
        "date": "2018-09-22T12:42:31Z"
    }'
```

Calculate the average purchase price 

```
curl --location 'http://localhost:3000/stock/CODE/average-purchase'
```

Calculate the average selling price

```
curl --location 'http://localhost:3000/stock/CODE/average-selling'
```