//go:generate gqlgen

package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"school/graph/client"
)

type Server struct {
	studentClient       *client.StudentClient
	teacherClient       *client.TeacherClient
}

func NewGraphQLServer(student_service_url ,teahcer_service_url string) (*Server, error) {
	studentClient, err := client.NewStudentClient(student_service_url)
	if err != nil {
		return nil, err
	}

	teacherClient, err := client.NewTeacherClient(teahcer_service_url)
	if err != nil {
		return nil, err
	}
	return &Server{
		studentClient: studentClient,
		teacherClient: teacherClient,
	},nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}
func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) TOExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}

