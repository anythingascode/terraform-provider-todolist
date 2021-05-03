package todoist_schema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	todo "terraform-provider-todolist/todoist-go"
)

// Project Schema for Data source
func Projects() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id":     DataProjectId(),
				"name":   DataProjectName(),
				"shared": DataProjectShared(),
				"url":    DataProjectURL(),
			},
		},
	}
}
func DataProjectId() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func DataProjectName() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func DataProjectShared() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
}

func DataProjectURL() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

// New Project Schema
func NewProjectName() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
}

func NewSharedProject() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
}
func NewProjectColor() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Default:  47,
	}
}

func NewFavProject() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
}

func NewTeamInboxProject() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
}

func NewInboxProject() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	}
}

// Flatten Project properties
func FlattenProjectsSchema(projects *[]todo.Project) interface{} {
	if projects == nil {
		return []interface{}{}
	}
	result := make([]interface{}, 0)
	for _, project := range *projects {
		name := project.Name
		id := project.ID
		shared := project.Shared
		url := project.URL
		result = append(result, map[string]interface{}{
			"name":   name,
			"id":     id,
			"shared": shared,
			"url":    url,
		})

	}
	return result
}

func ToDoistProjectsRead(m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)

	var diags diag.Diagnostics
	_, err := c.GetProjects()
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
