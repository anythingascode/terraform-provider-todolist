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

func dataSourceToDoistProjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceToDoistProjectsRead,
		Schema: map[string]*schema.Schema{
			"projects": todoist_schema.Projects(),
		},
	}
}

func dataSourceToDoistProjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)

	var diags diag.Diagnostics
	projects, err := c.GetProjects()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("projects", todoist_schema.FlattenProjectsSchema(projects)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
