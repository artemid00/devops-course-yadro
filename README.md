# Artem Bezpyatko && Yadro DevOps-2026

### About me:
- Student of MIET (3rd course)
- Interested in infrastructure and automation
- Prefer Vim over other text editors

### Project

This repository contains a small HTTP service written in Go that provides currency exchange rates from the Central Bank of Russia.

The service:
- fetches XML data from the official CBR API
- converts Windows-1251 encoded XML to UTF-8
- parses currency rates
- exposes them through a REST API

### Endpoints

```GET /info```
Returns service metadata (version, author, service name).

```GET /info/currency```
Returns exchange rates.  
Optional parameters:  
- currency — ISO 4217 code (example: USD)*
- date — date in format YYYY-MM-DD

```GET /metrics```
Returns Prometheus metrics.
*Custom metric:*
- `currencyapi_api_requests_total{result="success|error"}`  

Number of API requests by result (HTTP status < 400 is `success`, otherwise `error`)

*And also exposes default Go/process metrics (`go_*`, `process_*`, etc.) from the Prometheus Go client.*


### Configuration
*Configuration is provided through environment variables:*

```PORT``` (default: 8000)

```VERSION``` (default: 1.1.0)

```AUTHOR``` (default: me - a.bezpyatko)
```
    |\__/,|   (`\
  _.|o o  |_   ) )
-(((---(((--------
```
