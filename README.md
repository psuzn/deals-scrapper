# Deals Scrapper

Small utility service as part of [Play Deals](https://github.com/psuzn/Play-Deals) to find deals from the web.

## Development

### Running tests

```shell
go test ./... -v
```

### Start service

```shell
go run ./cmd/main.go
```
### Configuration

Configuration can be done by passing environment variables listed below:

| ENV_VAR      | REQUIRED | DEFAULT | EXAMPLE                           | NOTES                                    |
|--------------|----------|---------|-----------------------------------|:-----------------------------------------|
| `SERVER_URL` | `Y`      |         | `http://localhost:3000/api/deals` | service url to report the new found deal |
| `URLS`       | `N`      | []      | `whatever`                        | web urls to look for new deals           |

## License

**GPL V3 License**

Copyright (c) 2023 Sujan Poudel
