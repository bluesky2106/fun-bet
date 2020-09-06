# prophet challenge

## Requirements

- golang
- mysql
- docker

## How to set configurations

Modify `config/config.json` and `.env`. Please keep in mind that mysql configurations in `config/config.json` are related to the ones in `.env` which are used for running mysql docker image.

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