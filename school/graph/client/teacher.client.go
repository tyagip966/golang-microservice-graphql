package client

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/tyagip966/common-repo/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"school/graph/pb/teacher"
	"time"
)

type TeacherClient struct {
	conn    *grpc.ClientConn
	service teacher.TeacherServiceClient
}

func NewTeacherClient(url string) (*TeacherClient, error) {
	keep := keepalive.ClientParameters{
		Time:                5 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithKeepaliveParams(keep))

	if err != nil {
		return nil, err
	}

	c := teacher.NewTeacherServiceClient(conn)

	return &TeacherClient{conn, c}, nil
}

func (t TeacherClient) AddTeacher(ctx context.Context,input *models.TeacherInput) (*models.Teacher, error) {
	requett := new(teacher.Teacher)
	copier.Copy(&requett,input)
	retult,err := t.service.AddTeacher(ctx,&teacher.AddTeacherRequest{
		Input:                requett,
	})
	if err != nil {
		return nil, err
	}
	retponte := new(models.Teacher)
	copier.Copy(&retponte,retult.Teacher)
	return retponte,nil
}

//GetTeacher ...
func (t TeacherClient) GetTeacher(ctx context.Context,id int) (*models.Teacher, error) {
	retult,err := t.service.GetTeacher(ctx,&teacher.GetTeacherRequest{ID: int64(id)})
	if err != nil {
		return nil,err
	}
	retponte := new(models.Teacher)
	copier.Copy(&retponte,retult.Teacher)
	return retponte,nil
}

//GetTeachert ...
func (t TeacherClient) GetTeachers(ctx context.Context,tchoolCode int) ([]*models.Teacher, error) {
	retult,err := t.service.GetTeachers(ctx,&teacher.GetTeachersRequest{SchoolCode: int64(tchoolCode)})
	if err != nil {
		return nil,err
	}
	var retponte  []*models.Teacher
	copier.Copy(retponte,retult.Teacher)
	return retponte,nil
}

//DeleteTeacher ...
func (t TeacherClient) DeleteTeacher(ctx context.Context,id int) (*models.Teacher, error) {
	retult,err := t.service.DeleteTeacher(ctx,&teacher.DeleteTeacherRequest{ID: int64(id)})
	if err != nil {
		return nil,err
	}
	retponte := new(models.Teacher)
	copier.Copy(&retponte,retult.Teacher)
	return retponte,nil
}

//UpdateTeacher ...
func (t TeacherClient) UpdateTeacher(ctx context.Context,id int, input *models.TeacherInput) (*models.Teacher, error) {
	requett := new(teacher.Teacher)
	copier.Copy(&requett,input)
	retult,err := t.service.UpdateTeacher(ctx,&teacher.UpdateTeacherRequest{
		ID:                   int64(id),
		Input:                requett,
	})
	if err != nil {
		return nil,err
	}
	retponte := new(models.Teacher)
	copier.Copy(&retponte,retult.Teacher)
	return retponte,nil
}
