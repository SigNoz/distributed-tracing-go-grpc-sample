package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/SigNoz/distributed-tracing-go-grpc-sample/config"
	employeepb "github.com/SigNoz/distributed-tracing-go-grpc-sample/employee"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file for client", err)
	}

	tp := config.Init()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	fmt.Println("Employee Client")

	serverUrl := os.Getenv("SERVER_URL")

	cc, err := grpc.NewClient(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := employeepb.NewEmployeeServiceClient(cc)

	// create Employee
	fmt.Println("Creating the employee")
	employee := &employeepb.Employee{
		EmployeeId:  "Employee01",
		Name:        "John",
		Level:       "Engineer",
		Description: "Software Engineer",
	}

	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampMilli),
		"employeeId", employee.EmployeeId,
	)

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	createEmployeeRes, err := c.CreateEmployee(ctx, &employeepb.CreateEmployeeRequest{Employee: employee})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Employee has been created: %v", createEmployeeRes)
	employeeID := createEmployeeRes.GetEmployee().GetId()

	// Read Employee
	fmt.Println("Reading the employee")
	readEmployeeReq := &employeepb.ReadEmployeeRequest{EmployeeId: employeeID}
	readEmployeeRes, readEmployeeErr := c.ReadEmployee(ctx, readEmployeeReq)
	if readEmployeeErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readEmployeeErr)
	}

	fmt.Printf("Employee was read: %v \n", readEmployeeRes)

	// Update Employee
	newEmployee := &employeepb.Employee{
		Id:          employeeID,
		EmployeeId:  "Employee01",
		Name:        "John",
		Level:       "Leader",
		Description: "Team Lead",
	}
	updateRes, updateErr := c.UpdateEmployee(ctx, &employeepb.UpdateEmployeeRequest{Employee: newEmployee})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Employee was updated: %v\n", updateRes)

	// Delete Employee
	deleteRes, deleteErr := c.DeleteEmployee(ctx, &employeepb.DeleteEmployeeRequest{EmployeeId: employeeID})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Employee was deleted: %v \n", deleteRes)

	// List Employees
	stream, err := c.ListEmployee(ctx, &employeepb.ListEmployeeRequest{})
	if err != nil {
		log.Fatalf("error while calling ListEmployee RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println("List Employee Response", res.GetEmployee())
	}
}
