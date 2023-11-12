module github.comcast.com/vabira200/go-grpc-crud-example

go 1.16

require (
	github.com/joho/godotenv v1.4.0
	go.mongodb.org/mongo-driver v1.9.0
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.31.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.0
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.6.3
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)
