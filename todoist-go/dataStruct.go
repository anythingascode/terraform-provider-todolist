package todoist

type Project struct {
	ID           int    `json:"id"`
	Color        int    `json:"color"`
	Name         string `json:"name"`
	CommentCount int    `json:"comment_count"`
	Shared       bool   `json:"shared"`
	Favorite     bool   `json:"favorite"`
	SyncID       int    `json:"sync_id"`
	InboxProject bool   `json:"inbox_project"`
	URL          string `json:"url"`
}

type NewProject struct {
	Name         string `json:"name"`
	Color        int    `json:"color"`
	Shared       bool   `json:"shared"`
	Favorite     bool   `json:"favorite"`
	TeamInbox    bool   `json:"team_inbox"`
	InboxProject bool   `json:"inbox_project"`
}

type Task struct {
	ID           int           `json:"id"`
	Assigner     int           `json:"assigner"`
	ProjectID    int           `json:"project_id"`
	SectionID    int           `json:"section_id"`
	Order        int           `json:"order"`
	Content      string        `json:"content"`
	Completed    bool          `json:"completed"`
	LabelIds     []int		   `json:"label_ids"`
	Priority     int           `json:"priority"`
	CommentCount int           `json:"comment_count"`
	Creator      int           `json:"creator"`
	Created      string        `json:"created"`
	URL          string        `json:"url"`
}


type CompletedTask struct {
	Content       string `json:"content"`
	MetaData      string `json:"meta_data"`
	UserId        int    `json:"user_id"`
	TaskId        int    `json:"task_id"`
	ProjectId     int    `json:"project_id"`
	CompletedDate string `json:"completed_date"`
	Id            int    `json:"id"`
}

type CompletedTaskResponse struct {
	Items []CompletedTask `json:"items"`
}

type Labels struct {
	ID       int  `json:"id"`
	Name     string `json:"name"`
	Order    int    `json:"order"`
	Color    int    `json:"color"`
	Favorite bool   `json:"favorite"`
}