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

func resourceToDoistTask() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content":todoist_schema.NewTaskContent(),
			"project_id": todoist_schema.NewTaskProjectId(),
			"priority":todoist_schema.NewTaskPriority(),
			"order":todoist_schema.NewTaskOrder(),
			"labels":todoist_schema.NewTaskLabels(),
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		CreateContext:resourceToDoistCreateTask,
		ReadContext:resourceToDoistReadTask,
		UpdateContext:resourceToDoistUpdateTask,
		DeleteContext:resourceToDoistDeleteTask,
	}
}

func resourceToDoistCreateTask(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	label_ids := c.GetLabelId(d.Get("labels").([]interface{}))
	nt := todo.Task{
		ProjectID:    d.Get("project_id").(int),
		Order:        d.Get("order").(int),
		Content:      d.Get("content").(string),
		LabelIds:     label_ids,
		Priority:     d.Get("priority").(int),
	}
	newTask := c.CreateNewTask(&nt)
	d.SetId(strconv.Itoa(newTask.ID))
	return resourceToDoistReadTask(ctx, d, m)
}

func resourceToDoistReadTask(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*todo.Client)
	_, err := c.GetTasks()
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceToDoistUpdateTask(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	var diags diag.Diagnostics
	id, _ := strconv.Atoi(d.Id())
	if d.HasChanges("project_id", "order", "content", "labels", "priority") {
		label_ids := c.GetLabelId(d.Get("labels").([]interface{}))
		t := todo.Task {
			ID:id,
			ProjectID:    d.Get("project_id").(int),
			Order:        d.Get("order").(int),
			Content:      d.Get("content").(string),
			LabelIds:     label_ids,
			Priority:     d.Get("priority").(int),
		}
		c.UpdateTask(&t)
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}
	return diags
}

func resourceToDoistDeleteTask(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*todo.Client)
	id, _ := strconv.Atoi(d.Id())
	c.DeleteTask(id)
	return diags
}