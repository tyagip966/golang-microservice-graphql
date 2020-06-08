package graph

import (
	"context"
	"github.com/tyagip966/common-repo/models"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) GetTeachers(ctx context.Context, schoolCode int) ([]*models.Teacher, error) {
    res,err := q.server.teacherClient.GetTeachers(ctx,schoolCode)
    if err != nil {
    	return nil, err
	}
	return res,nil
}

func (q queryResolver) GetTeacherByID(ctx context.Context, id int, schoolcode int) (*models.Teacher, error) {
	res,err := q.server.teacherClient.GetTeacher(ctx,id)
	if err != nil {
		return nil, err
	}
	return res,nil
}

func (q queryResolver) GetStudents(ctx context.Context, schoolCode int) ([]*models.Student, error) {
	res,err := q.server.studentClient.GetStudents(ctx,schoolCode)
	if err != nil {
		return nil, err
	}
	return res,nil
}

func (q queryResolver) GetStudentByID(ctx context.Context, id int, schoolcode int) (*models.Student, error) {
	res,err := q.server.studentClient.GetStudent(ctx,id)
	if err != nil {
		return nil, err
	}
	return res,nil
}
