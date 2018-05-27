# CLI

## Running this program
clone this repo, and then you can run this as CLI or as a web server.

### Run the program from CLI
```
go run main.go
```

### Run the program as a web server
This program will open port 8000
```
cd web
go run main.go
```

### example payload
``` 
{
    "loanAmount": 5000, 
    "nominalRate": 5.0,
    "duration": 12,
    "startDate": "2018-01-01T00:00:01Z"
}

e.g.:
POST http://localhost:8000/generate-plan
```