package todoist

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	todo "terraform-provider-todolist/todoist-go"
	"terraform-provider-todolist/todoist/todoist_schema"
	"time"
)

func dataSourceToDoistTasks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceToDoistTasksRead,
		Schema: map[string]*schema.Schema{
			"is_completed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"completed_tasks": todoist_schema.CompletedTasks(),
			"active_tasks":    todoist_schema.Tasks(),
		},
	}
}

func dataSourceToDoistTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	var diags diag.Diagnostics
	switch {
	case d.Get("is_completed") == true:
		ctasks, _ := c.GetCompletedTasks()
		tasks, _ := c.GetTasks()
		d.Set("completed_tasks", todoist_schema.FlattenCompletedTasksSchema(ctasks))
		d.Set("active_tasks", todoist_schema.FlattenTasksSchema(tasks))
	default:
		tasks, _ := c.GetTasks()
		d.Set("completed_tasks", nil)
		d.Set("active_tasks", todoist_schema.FlattenTasksSchema(tasks))
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
