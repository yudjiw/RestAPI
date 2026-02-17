package http

import (
	"encoding/json"
	"errors"
	"time"
)

//DTO == data transfer object

type CompleteTaskDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("task title is required")
	}
	if t.Description == "" {
		return errors.New("task description is required")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {

	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}
