package graph

import (
	"context"
	"github.com/tyagip966/common-repo/models"
)

type mutationResolver struct {
	server *Server
}

func (m mutationResolver) InsertTeacher(ctx context.Context, input *models.TeacherInput) (*models.Teacher, error) {
	res,err := m.server.teacherClient.AddTeacher(ctx,input)
	if err != nil{
		return nil, err
	}
	return res,nil
}

func (m mutationResolver) DeleteTeacher(ctx context.Context, id int) (*models.Teacher, error) {
	res,err := m.server.teacherClient.DeleteTeacher(ctx,id)
	if err != nil{
		return nil, err
	}
	return res,nil
}

func (m mutationResolver) UpdateTeacher(ctx context.Context, id int, input *models.TeacherInput) (*models.Teacher, error) {
	res,err := m.server.teacherClient.UpdateTeacher(ctx,id,input)
	if err != nil{
		return nil, err
	}
	return res,nil
}

func (m mutationResolver) InsertStudent(ctx context.Context, input *models.StudentInput) (*models.Student, error) {
	res,err := m.server.studentClient.AddStudent(ctx,input)
	if err != nil{
		return nil, err
	}
	return res,nil
}

func (m mutationResolver) DeleteStudent(ctx context.Context, id int) (*models.Student, error) {
	res,err := m.server.studentClient.DeleteStudent(ctx,id)
	if err != nil{
		return nil, err
	}
	return res,nil
}

func (m mutationResolver) UpdateStudent(ctx context.Context, id int, input *models.StudentInput) (*models.Student, error) {
	res,err := m.server.studentClient.UpdateStudent(ctx,id,input)
	if err != nil{
		return nil, err
	}
	return res,nil
}

