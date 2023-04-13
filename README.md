# Distributed Tracing Go Grpc Sample

This project demonstrates how to implement distributed tracing in go grpc application.

This application uses mongodb as the database, so make sure to create an employeedb database and employee collection in mongodb and run this app.

## Tracing flow

![Distributed tracing](go-grpc-otel.jpg)

## Running the code

Start the SigNoz server following the instructions:

```bash
git clone -b main https://github.com/SigNoz/signoz.git
cd signoz/deploy/
./install.sh
```

_*Note:* Replace OTEL_EXPORTER_OTLP_ENDPOINT environment variable with SigNoz OTLP endpoint, if SigNoz not running on host machine._

### Using docker-compose

```bash
docker-compose up -d
```

View traces and metrics at http://localhost:3301/

### Using Source

Start go grpc server and grpc client using below commands

1. Grpc-Server
```
cd server
go run server.go
```

2. Grpc-Client
```
cd client
go run client.go
```

View traces and metrics at http://localhost:3301/
