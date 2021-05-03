package todoist_schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	todo "terraform-provider-todolist/todoist-go"
)

func Tasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id":            TaskID(),
				"project_id":    TaskProjectID(),
				"creator":       TaskCreater(),
				"content":       TaskContent(),
				"priority":      TaskPriority(),
				"created":       TaskCreated(),
				"url":           TaskURL(),
				"comment_count": TaskComment(),
				"completed":     Completed(),
			},
		},
	}
}

func CompletedTasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id":             TaskID(),
				"project_id":     TaskProjectID(),
				"content":        TaskContent(),
				"task_id":        CTaskId(),
				"user_id":        CTaskUserId(),
				"completed_date": CTaskCompletdDate(),
				"meta_data":      CTaskMetaData(),
			},
		},
	}
}
func TaskID() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func TaskProjectID() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func TaskCreater() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func TaskContent() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TaskPriority() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func TaskCreated() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TaskURL() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TaskComment() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func CTaskId() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func CTaskUserId() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func CTaskCompletdDate() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func CTaskMetaData() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func Completed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
}

// New task Schema
func NewTaskContent() *schema.Schema {
	return &schema.Schema{
		Type:schema.TypeString,
		Required:true,
	}
}

func NewTaskProjectId() *schema.Schema {
	return &schema.Schema{
		Type:schema.TypeInt,
		Required:true,
	}
}

func NewTaskOrder() *schema.Schema {
	return &schema.Schema{
		Type:schema.TypeInt,
		Optional:true,
	}
}
func NewTaskLabels() *schema.Schema {
	return &schema.Schema{
		Type:schema.TypeList,
		Optional:true,
		Elem:&schema.Schema{
			Type:schema.TypeString,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func NewTaskPriority() *schema.Schema {
	return &schema.Schema{
		Type:schema.TypeInt,
		Optional:true,
	}
}

// Flatten Task properties
func FlattenTasksSchema(tasks *[]todo.Task) interface{} {
	if tasks == nil {
		return []interface{}{}
	}
	result := make([]interface{}, 0)
	for _, task := range *tasks {
		completed := task.Completed
		id := task.ID
		project_id := task.ProjectID
		content := task.Content
		comment_count := task.CommentCount
		creator := task.Creator
		created := task.Created
		url := task.URL
		result = append(result, map[string]interface{}{
			"completed":     completed,
			"id":            id,
			"project_id":    project_id,
			"content":       content,
			"comment_count": comment_count,
			"creator":       creator,
			"created":       created,
			"url":           url,
		})
	}
	return result
}

func FlattenCompletedTasksSchema(ctasks *[]todo.CompletedTask) interface{} {
	if ctasks == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, ctask := range *ctasks {
		completed_date := ctask.CompletedDate
		content := ctask.Content
		id := ctask.Id
		meta_data := ctask.MetaData
		project_id := ctask.ProjectId
		task_id := ctask.TaskId
		user_id := ctask.UserId

		result = append(result, map[string]interface{}{
			"completed_date": completed_date,
			"content":        content,
			"id":             id,
			"meta_data":      meta_data,
			"project_id":     project_id,
			"task_id":        task_id,
			"user_id":        user_id,
		})
	}
	return result
}
