package client

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/tyagip966/common-repo/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"school/graph/pb/student"
	"time"
)

type StudentClient struct {
	conn    *grpc.ClientConn
	service student.StudentServiceClient
}

func NewStudentClient(url string) (*StudentClient, error) {
	keep := keepalive.ClientParameters{
		Time:                5 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithKeepaliveParams(keep))

	if err != nil {
		return nil, err
	}

	c := student.NewStudentServiceClient(conn)

	return &StudentClient{conn, c}, nil
}

func (s StudentClient) AddStudent(ctx context.Context,input *models.StudentInput) (*models.Student, error) {
	request := new(student.Student)
	copier.Copy(&request,input)
	result,err := s.service.AddStudent(ctx,&student.AddStudentRequest{
		Input:                request,
	})
	if err != nil {
		return nil,err
	}
	response := new(models.Student)
	copier.Copy(&response,result.Student)
	return response,nil
}

//GetStudent ...
func (s StudentClient) GetStudent(ctx context.Context,id int) (*models.Student, error) {
	result,err := s.service.GetStudent(ctx,&student.GetStudentRequest{ID: int64(id)})
	if err != nil {
		return nil,err
	}
	response := new(models.Student)
	copier.Copy(&response,result.Student)
	return response,nil
}

//GetStudents ...
func (s StudentClient) GetStudents(ctx context.Context,schoolCode int) ([]*models.Student, error) {
	result,err := s.service.GetStudents(ctx,&student.GetStudentsRequest{SchoolCode: int64(schoolCode)})
	if err != nil {
		return nil,err
	}
	var response  []*models.Student
	copier.Copy(response,result.Student)
	return response,nil
}

//DeleteStudent ...
func (s StudentClient) DeleteStudent(ctx context.Context,id int) (*models.Student, error) {
	result,err := s.service.DeleteStudent(ctx,&student.DeleteStudentRequest{ID: int64(id)})
	if err != nil {
		return nil,err
	}
	response := new(models.Student)
	copier.Copy(&response,result.Student)
	return response,nil
}

//UpdateStudent ...
func (s StudentClient) UpdateStudent(ctx context.Context,id int, input *models.StudentInput) (*models.Student, error) {
	request := new(student.Student)
	copier.Copy(&request,input)
	result,err := s.service.UpdateStudent(ctx,&student.UpdateStudentRequest{
		ID:                   int64(id),
		Input:                request,
	})
	if err != nil {
		return nil,err
	}
	response := new(models.Student)
	copier.Copy(&response,result.Student)
	return response,nil
}
