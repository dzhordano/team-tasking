package converter

import (
	"github.com/dzhordano/team-tasking/services/tasks/internal/domain"
	task_v1 "github.com/dzhordano/team-tasking/services/tasks/pkg/grpc/task/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Converts domain structures to grpc structures

func TaskToGRPC(task *domain.Task) *task_v1.Task {
	return &task_v1.Task{
		Id:          task.TaskID.String(),
		ProjectId:   task.ProjectID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status.String(),
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
		Deadline:    timestamppb.New(task.Deadline),
	}
}

func TasksToGRPC(tasks []*domain.Task) []*task_v1.Task {
	var grpcTasks []*task_v1.Task
	for _, task := range tasks {
		grpcTasks = append(grpcTasks, TaskToGRPC(task))
	}
	return grpcTasks
}

func ProjectToGRPC(project *domain.Project) *task_v1.Project {
	return &task_v1.Project{
		Id:   project.ProjectID.String(),
		Name: project.Name,
	}
}

func ProjectsToGRPC(projects []*domain.Project) []*task_v1.Project {
	var grpcProjects []*task_v1.Project
	for _, project := range projects {
		grpcProjects = append(grpcProjects, ProjectToGRPC(project))
	}
	return grpcProjects
}

func CommentToGRPC(comment *domain.Comment) *task_v1.Comment {
	return &task_v1.Comment{
		Id:        comment.CommentID.String(),
		TaskId:    comment.TaskID.String(),
		Content:   comment.Content,
		CreatedAt: timestamppb.New(comment.CreatedAt),
		UpdatedAt: timestamppb.New(comment.UpdatedAt),
	}
}

func CommentsToGRPC(comments []*domain.Comment) []*task_v1.Comment {
	var grpcComments []*task_v1.Comment
	for _, comment := range comments {
		grpcComments = append(grpcComments, CommentToGRPC(comment))
	}
	return grpcComments
}
