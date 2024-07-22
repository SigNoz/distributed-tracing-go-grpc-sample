package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/SigNoz/distributed-tracing-go-grpc-sample/config"
	employeepc "github.com/SigNoz/distributed-tracing-go-grpc-sample/employee"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct {
	employeepc.EmployeeServiceServer
}

type employeeItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeId  string             `bson:"employeeId"`
	Name        string             `bson:"name"`
	Level       string             `bson:"level"`
	Description string             `bson:"description"`
}

func getEmployeeData(data *employeeItem) *employeepc.Employee {
	return &employeepc.Employee{
		Id:          data.ID.Hex(),
		EmployeeId:  data.EmployeeId,
		Name:        data.Name,
		Level:       data.Level,
		Description: data.Description,
	}
}

func (s *server) CreateEmployee(ctx context.Context, req *employeepc.CreateEmployeeRequest) (*employeepc.CreateEmployeeResponse, error) {

	fmt.Println("Create Employee")
	employee := req.GetEmployee()

	data := employeeItem{
		EmployeeId:  employee.GetEmployeeId(),
		Name:        employee.GetName(),
		Level:       employee.GetLevel(),
		Description: employee.GetDescription(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &employeepc.CreateEmployeeResponse{
		Employee: &employeepc.Employee{
			Id:          oid.Hex(),
			EmployeeId:  employee.GetEmployeeId(),
			Name:        employee.GetName(),
			Level:       employee.GetLevel(),
			Description: employee.GetDescription(),
		},
	}, nil
}

func (s *server) ReadEmployee(ctx context.Context, req *employeepc.ReadEmployeeRequest) (*employeepc.ReadEmployeeResponse, error) {
	fmt.Println("Read Employee")
	employeeID := req.GetEmployeeId()
	oid, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	// create an empty struct
	data := &employeeItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find employee with specified ID: %v", err),
		)
	}

	return &employeepc.ReadEmployeeResponse{
		Employee: getEmployeeData(data),
	}, nil
}

func (s *server) UpdateEmployee(ctx context.Context, req *employeepc.UpdateEmployeeRequest) (*employeepc.UpdateEmployeeResponse, error) {
	fmt.Println("Updating Employee")
	employee := req.GetEmployee()
	oid, err := primitive.ObjectIDFromHex(employee.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	// create an empty struct
	data := &employeeItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find employee with specified ID: %v", err),
		)
	}

	data.EmployeeId = employee.GetEmployeeId()
	data.Name = employee.GetName()
	data.Level = employee.GetLevel()
	data.Description = employee.GetDescription()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &employeepc.UpdateEmployeeResponse{
		Employee: getEmployeeData(data),
	}, nil

}

func (s *server) DeleteEmployee(ctx context.Context, req *employeepc.DeleteEmployeeRequest) (*employeepc.DeleteEmployeeResponse, error) {
	fmt.Println("Deleting Employee")
	oid, err := primitive.ObjectIDFromHex(req.GetEmployeeId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find employee in MongoDB: %v", err),
		)
	}

	return &employeepc.DeleteEmployeeResponse{EmployeeId: req.GetEmployeeId()}, nil
}

func (s *server) ListEmployee(_ *employeepc.ListEmployeeRequest, stream employeepc.EmployeeService_ListEmployeeServer) error {
	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &employeeItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}
		stream.Send(&employeepc.ListEmployeeResponse{Employee: getEmployeeData(data)}) // Should handle err
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return status.Error(codes.NotFound, "Internal error")
}

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file for server", err)
	}
	mongo_url := os.Getenv("MONGO_URL")

	tp := config.Init()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url).SetMonitor(otelmongo.NewMonitor()))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employee Service Started")
	collection = client.Database("employeedb").Collection("employee")

	lis, err := net.Listen("tcp", "0.0.0.0:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))

	employeepc.RegisterEmployeeServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
