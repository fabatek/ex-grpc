package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/faba/client/config"
	dao "github.com/faba/client/dao"
	"github.com/faba/client/domain/model"
	rep "github.com/faba/client/domain/repository"
)

func main() {
	fabaClient, err := dao.NewFabaClient(config.SERVER)
	if err != nil {
		log.Fatal(err)
	}

	employeeRequestParam := model.EmployeeRequestParam{
		Name:        "FABA company",
		Address:     "ho chi minh",
		Phone:       "787345",
		BirthOfDate: time.Now(),
	}
	data, err := RegisterEmployee(fabaClient, &employeeRequestParam)

	fmt.Println("err ---->", err)
	fmt.Println("data ---->", data)
}

// RegisterEmployee call repository to create employee
func RegisterEmployee(fabaClient rep.FabaRepository, sarFileRequestParam *model.EmployeeRequestParam) (data model.EmployeeResponseData, err error) {
	return fabaClient.CreateEmployee(context.Background(), sarFileRequestParam)
}
