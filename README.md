
# Wager Service
## System
This repository was developed with the system below:
- Go 1.18
- Docker
- MacOS (or Ubuntu)
## Technical Using
|Technical| Description  |
|--|--|
| [Gin](https://github.com/gin-gonic/gin) | HTTP web framework |
|[Govalidator](github.com/asaskevich/govalidator)|validate request param|
| [GORM](https://gorm.io/) |working with database|
| [Copier](github.com/jinzhu/copier) |For copy between object request and model|
| [Viper](github.com/spf13/viper) |Set and get configuration|

## How to use

**Run project**
From root project run the script:
```
sh ./start.sh
```
**Run unit test**
From root project run the script:
```
sh ./test.sh
```

## Api Specification
**Create Wager**
- Request
```
curl --location --request POST 'localhost:8080/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
	"total_wager_value": 1150,
    "odds": 4,
    "selling_percentage": 70,
    "selling_price": 810.3
}'
```
- Response: Header: `HTTP 201` Body:
```json
{
   "data":{
      "id":2,
      "total_wager_value":1150,
      "odds":4,
      "selling_percentage":70,
      "selling_price":810.3,
      "current_selling_price":810.3,
      "percentage_sold":null,
      "amount_sold":null,
      "placed_at":"2022-10-23T11:06:57.452Z"
   },
   "error":null,
   "success":true
}
```

**List wager**
- Request
```
curl --location --request GET 'localhost:8080/wagers?page=1&limit=3' \
--header 'Content-Type: application/json' \
--data-raw ''
```
- Response: Header: `HTTP 200` Body:
```json
{
    "data": {
        "data": [
            {
                "id": 2,
                "total_wager_value": 1150,
                "odds": 4,
                "selling_percentage": 70,
                "selling_price": 810.3,
                "current_selling_price": 810.3,
                "percentage_sold": null,
                "amount_sold": null,
                "placed_at": "2022-10-23T11:06:57.452Z"
            },
            {
                "id": 1,
                "total_wager_value": 1150,
                "odds": 4,
                "selling_percentage": 70,
                "selling_price": 810.3,
                "current_selling_price": 20.5,
                "percentage_sold": 1.78261,
                "amount_sold": 1,
                "placed_at": "2022-10-23T10:19:57.785Z"
            }
        ],
        "total": 2
    },
    "error": null,
    "success": true
}
```
**Buy wager**
- Request
```
curl --location --request POST 'localhost:8080/wagers/buy/1' \
--header 'Content-Type: application/json' \
--data-raw '{
	"buying_price": 20.5
}'
```
- Response: header: `HTTP 201` Body:
```json
{
    "data": {
        "id": 2,
        "wager_id": 1,
        "buying_price": 20.5,
        "bought_at": "2022-10-23T11:12:52.001Z"
    },
    "error": null,
    "success": true
}
```
## Questions / Feedbacks / Bugs
Feel free to reach out to me if you have any questions or feedback on how my code can be improved.

### TODO

- [x] REST APIs

- [x] Docker build

- [x] Unit test

- [ ] Swagger documentation