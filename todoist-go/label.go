package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetLabels() (*[]Labels, error) {
	req, err := c.newRequest("labels", "GET", nil)
	if err != nil {
		return nil, err
	}
	res, _ := c.doRequest(req)
	var l []Labels
	if err := json.Unmarshal(res, &l); err != nil {
		return nil, err
	}
	return &l, nil
}

func (c *Client) GetLabelId(labels []interface{}) []int {
	ls, _ := c.GetLabels()
	var id []int
	 for _, label := range *ls {
		for _, n := range labels {
			if n == label.Name {
				id = append(id, label.ID)
			} else {
				fmt.Sprintf("%n label not found", n)
			}
		}
	}
	return id
}

func (c *Client) CreateLabels(label *Labels) error {
	gl, _ := c.GetLabels()
	for _, l := range *gl {
		if l.Name != label.Name {
			jl, _ := json.Marshal(label)
			req, _ := c.newRequest("labels", "POST", bytes.NewBuffer(jl))
			req.Header.Add("Content-type", "application/json")
			_, err := c.doRequest(req)
			if err != nil {
				return err
			}
		}
	}
	return nil
}