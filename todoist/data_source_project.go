package todoist

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	todo "terraform-provider-todolist/todoist-go"
	"terraform-provider-todolist/todoist/todoist_schema"
)

func dataSourceToDoistProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceToDoistProjectRead,
		Schema: map[string]*schema.Schema{
			"id": todoist_schema.DataProjectId(),
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"shared": todoist_schema.DataProjectShared(),
			"url":    todoist_schema.DataProjectURL(),
		},
	}
}

func dataSourceToDoistProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	var diags diag.Diagnostics
	projectName := d.Get("name").(string)
	p, err := c.GetProjectByName(projectName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("url", p.URL)
	d.Set("shared", p.Shared)
	d.SetId(strconv.Itoa(p.ID))
	return diags
}
