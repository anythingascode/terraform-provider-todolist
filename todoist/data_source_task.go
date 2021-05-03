package todoist

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	todo "terraform-provider-todolist/todoist-go"
	"terraform-provider-todolist/todoist/todoist_schema"
)

func dataSourceToDoistTask() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceToDoistTaskRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"project_id": todoist_schema.TaskProjectID(),
			"content":    todoist_schema.TaskContent(),
			"created":    todoist_schema.TaskCreated(),
			"url":        todoist_schema.TaskURL(),
			"priority":   todoist_schema.TaskPriority(),
			"completed":  todoist_schema.Completed(),
		},
	}
}

func dataSourceToDoistTaskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	var diags diag.Diagnostics
	task := c.GetActiveTaskById(d.Get("id").(int))
	d.Set("project_id", task.ProjectID)
	d.Set("content", task.Content)
	d.Set("priority", task.Priority)
	d.Set("creator", task.Creator)
	d.Set("created", task.Created)
	d.Set("url", task.URL)
	d.Set("completed", task.Completed)

	d.SetId(strconv.Itoa(task.ID))
	return diags
}
