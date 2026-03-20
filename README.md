# Artem Bezpyatko && Yadro DevOps-2026

About me:
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

**Endpoints**:

```GET /info```
Returns service metadata (version, author, service name).

```GET /info/currency```
Returns exchange rates.  
Optional parameters:  
- currency — ISO 4217 code (example: USD)*
- date — date in format YYYY-MM-DD

*Configuration is provided through environment variables:*

```PORT``` (default: 8000)

```VERSION``` (default: 1.0.0)

```AUTHOR``` (default: me - a.bezpyatko)
```
    |\__/,|   (`\
  _.|o o  |_   ) )
-(((---(((--------
```
