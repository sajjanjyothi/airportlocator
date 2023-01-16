## Airport Detector
This application will expose an endpoint to calculate the starting and end airports from a list of flight paths.

## Tests
You can run unit tests as
```
    make test
```

## Test run
Test run 

```
    make run
```
## Sample API usage

```
    curl http://localhost:8080/calculate -H 'Content-Type: application/json' -X GET  --data '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' -v

    Eg:
    curl http://localhost:8080/calculate -H 'Content-Type: application/json' -X GET  --data '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' 
    ["SFO","EWR"]
```
## API document 
APIs are documented in open API spec 3.0.0

after running server, access below url in browser
```
    http://localhost:8080/
```

[Local API doc access](./api/airportcalculator.yaml)

## API endpoint

```
http://localhost:8080/calculate

```
## Helm Charts
[Helm charts](./deployment/)

## Logging 
uber zap package used for logging

## Packages used
- https://github.com/deepmap/oapi-codegen  
- https://github.com/labstack/echo 
- https://github.com/uber-go/zap

## ToDo
- more error handling in endpoint
- Authentication addition to endpoint
- Helm chart not tested