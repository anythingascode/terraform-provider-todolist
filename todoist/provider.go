package todoist

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	todo "terraform-provider-todolist/todoist-go"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODOIST_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"todoist_project": resourceToDoistProject(),
			"todoist_task":resourceToDoistTask(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"todoist_project":  dataSourceToDoistProject(),
			"todoist_projects": dataSourceToDoistProjects(),
			"todoist_task":     dataSourceToDoistTask(),
			"todoist_tasks":    dataSourceToDoistTasks(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var diags diag.Diagnostics
	c := todo.NewClient(&token)
	return c, diags
}
