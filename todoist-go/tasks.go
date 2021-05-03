package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get Tasks
func (c *Client) GetTasks() (*[]Task, error) {
	req, err := c.newRequest("tasks", "GET", nil)
	if err != nil {
		fmt.Sprintf("Error during getting Tasks: %s", err)
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during response from todoist: %s", err)
	}
	var tasks []Task
	err = json.Unmarshal(res, &tasks)
	if err != nil {
		fmt.Sprintf("Error during unmarshal: %s", err)
	}

	return &tasks, nil
}

// Get Completed Tasks
func (c *Client) GetCompletedTasks() (*[]CompletedTask, error) {
	req, err := http.NewRequest("GET", DefaultSyncurl+"/completed/get_all", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var ctr CompletedTaskResponse
	var completedTasks []CompletedTask

	err = json.Unmarshal(res, &ctr)
	if err != nil {
		return nil, err
	}

	completedTasks = ctr.Items
	return &completedTasks, nil
}

// Get Completed Task
func (c *Client) GetCompletedTask(id int) *CompletedTask {
	completedTasks, err := c.GetCompletedTasks()
	if err != nil {
		return nil
	}
	var cts *CompletedTask
	for _, ct := range *completedTasks {
		if ct.TaskId == id {
			cts = &ct
			return cts
		}

	}

	return cts
}

// Create New Task
func (c *Client) CreateNewTask(nt *Task) *Task {
	jnt, err := json.Marshal(nt)
	if err != nil {
		fmt.Sprintf("Error during marshal: %s", err)
	}
	req, err := c.newRequest("tasks", "POST", bytes.NewBuffer(jnt))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Sprintf("Error during new request: %s", err)
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during request: %s", err)
	}
	var task Task
	err = json.Unmarshal(res, &task)
	if err != nil {
		fmt.Sprintf("Error during unmarshal: %s", err)
	}
	return &task
}

// Get Active Task
func (c *Client) GetActiveTaskById(id int) *Task {
	req, err := c.newRequest(fmt.Sprintf("tasks/%d", id), "GET", nil)
	if err != nil {
		fmt.Sprintf("Error during new request: %s", err)
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during request: %s", err)
	}
	var task Task
	json.Unmarshal(res, &task)
	return &task
}

// Update task
func (c *Client) UpdateTask(t *Task) {
	jt, _ := json.Marshal(&t)
	req, err := c.newRequest(fmt.Sprintf("tasks/%d", t.ID), "POST", bytes.NewBuffer(jt))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Sprintf("Error during new request: %s", err)
	}
	_, err = c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during request: %s", err)
	}
}

// Close or Open task
func (c *Client) CloseOrReOpenTask(id int, action string) {
	req, err := c.newRequest(fmt.Sprintf("tasks/%d/%s", id, action), "POST", nil)
	if err != nil {
		fmt.Sprintf("Error during new request: %s", err)
	}
	_, err = c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during req: %s", err)
	}
}

// Delete task
func (c *Client) DeleteTask(id int) {
	req, err := c.newRequest(fmt.Sprintf("tasks/%d", id), "DELETE", nil)
	if err != nil {
		fmt.Sprintf("Error during new request: %s", err)
	}
	_, err = c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during delete: %s", err)
	}
}
