# prophet challenge

## Requirements

Followings apps / software must be installed on your machine:

- golang
- mysql
- docker
- make
- mockgen

## How to set configurations

Modify `config/config.json` and `.env`. Please keep in mind that mysql configurations in `config/config.json` are related to the ones in `.env` which are used for running mysql docker image.

## How to generate mock files

- mockgen -source=interfaces/wager.go -destination=mocks/mock_wager.go -package=mocks

## How to test

```
make test
```

## How to start service

```
make start
```

## How to run in docker

```
docker-compose up
```

## API Interface remarks

### Place wager

A validator is added to check the api request. Followings are some examples of error messages which you could receive:

1. Invalid Selling Percentage

- selling_percentage must be specified as an integer between 1 and 100. Otherwise, you will recieve an error like below:

```
{
    "error": "invalid selling percentage"
}
```

2. Invalid Selling Price Format:

- selling_price must be specified as a positive decimal value to two decimal places. selling_price of 25.123 is NOT accepted.

```
{
    "error": "Invalid selling price format. It must be specified as a positive decimal value to two decimal places"
}
```

3. Invalid Selling Price:

- selling_price must be greater than total_wager_value * (selling_percentage / 100). Otherwise, you will recieve an error like below:

```
{
    "error": "invalid selling price"
}
```

### Buy wager

1. Invalid Buying Price

- buying_price must be a positive decimal value. Futhermore, it must be lesser or equal to current_selling_price of the wager_id. Otherwise, you will recieve an error like below:

```
{
    "error": "invalid buying price"
}
```

2. Wager Not Exist

- wager_id must exist in the database. Otherwise, you will receive error like below:

```
{
    "error": "wager not exist"
}
```

### List wager

Beside page and limit, I provided another 2 params: `sort_by` and `order_by`. By default, sort_by = id and order_by = asc (ascending). If you want to list all wagers and sort them by "placed_at" descendingly, you can add sort_by=placed_a and order_by=desc in the request params.

```
curl --location --request GET 'http://localhost:8080/wagers?page=1&limit=100&order_by=desc&sort_by=placed_at' \
--header 'Content-Type: application/json'
```