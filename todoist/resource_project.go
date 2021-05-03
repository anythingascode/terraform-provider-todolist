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

func resourceToDoistProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name":          todoist_schema.NewProjectName(),
			"shared":        todoist_schema.NewSharedProject(),
			"color":         todoist_schema.NewProjectColor(),
			"favorite":      todoist_schema.NewFavProject(),
			"team_inbox":    todoist_schema.NewTeamInboxProject(),
			"inbox_project": todoist_schema.NewInboxProject(),
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		CreateContext: resourceToDoistCreateProject,
		ReadContext:   resourceToDoistReadProject,
		UpdateContext: resourceToDoistUpdateProject,
		DeleteContext: resourceToDoistDeleteProject,
	}
}

func resourceToDoistCreateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*todo.Client)
	np := &todo.NewProject{
		Name:         d.Get("name").(string),
		Color:        d.Get("color").(int),
		Shared:       d.Get("shared").(bool),
		Favorite:     d.Get("favorite").(bool),
		TeamInbox:    d.Get("team_inbox").(bool),
		InboxProject: d.Get("inbox_project").(bool),
	}
	p, err := c.CreateProject(np)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(p.ID))
	resourceToDoistReadProject(ctx, d, m)
	return diags
}

func resourceToDoistReadProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	var diags diag.Diagnostics
	_, err := c.GetProjects()
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceToDoistUpdateProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*todo.Client)
	if d.HasChanges("name", "color", "favorite") {
		task_id, _ := strconv.Atoi(d.Id())
		project := &todo.Project{
			ID:       task_id,
			Name:     d.Get("name").(string),
			Color:    d.Get("color").(int),
			Favorite: d.Get("favorite").(bool),
		}
		c.UpdateProject(project)
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceToDoistReadProject(ctx, d, m)
}

func resourceToDoistDeleteProject(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*todo.Client)
	c.DeleteProject(d.Get("name").(string))
	d.SetId("")
	return diags
}
