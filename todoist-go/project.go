package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetProjects() (*[]Project, error) {
	req, err := c.newRequest("projects", "GET", nil)
	if err != nil {
		fmt.Sprintf("Error during request:")
		return nil, err
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during do request:")
		return nil, err
	}
	var projects []Project
	err = json.Unmarshal(res, &projects)
	if err != nil {
		fmt.Sprintf("Error during do unmarshel:")
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProjectByName(name string) (*Project, error) {
	projects, err := c.GetProjects()
	if err != nil {
		fmt.Sprintf("Error during getting project list: %s", err)
	}

	for _, project := range *projects {
		if project.Name == name {
			return &project, nil
		}
	}
	return nil, fmt.Errorf("Project with name %s not found", name)
}

func (c *Client) CreateProject(np *NewProject) (*Project, error) {
	jnp, err := json.Marshal(np)
	if err != nil {
		fmt.Sprintf("Error during marshal: %s", err)
	}
	req, err := c.newRequest("projects", "POST", bytes.NewBuffer(jnp))
	if err != nil {
		fmt.Sprintf("Error during creating project %s", err)
	}
	req.Header.Add("Content-type", "application/json")
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during request:%s", err)
	}
	var project Project
	json.Unmarshal(res, &project)

	return &project, nil
}

func (c *Client) UpdateProject(up *Project) error {
	jp, err := json.Marshal(up)
	if err != nil {
		fmt.Sprintf("Error during marshal: %s", err)
	}
	project, err := c.GetProjects()
	if err != nil {
		return err
	}
	for _, p := range *project {
		if p.ID == up.ID {
			req, _ := c.newRequest(fmt.Sprintf("projects/%d", up.ID), "POST", bytes.NewBuffer(jp))
			req.Header.Add("Content-type", "application/json")
			_, err = c.doRequest(req)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) DeleteProject(name string) ([]byte, error) {
	projects, err := c.GetProjects()
	var Id int
	if err != nil {
		fmt.Sprintf("Error during getting all projects: %s", err)
	}
	for _, p := range *projects {
		if p.Name == name {
			Id = p.ID
		}
	}
	req, err := c.newRequest(fmt.Sprintf("projects/%d", Id), "DELETE", nil)
	if err != nil {
		fmt.Sprintf("Error during deleting project %", name)
	}
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Sprintf("Error during deleting project request:%s", err)
	}
	return res, nil
}
