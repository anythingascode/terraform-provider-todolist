terraform {
  required_providers {
    todoist = {
      version = "0.1.0"
      source  = "todoist.com/dev/todoist"
    }
  }
}

provider "todoist" {
}

data "todoist_project" "this" {
    name = "Inbox"
}
output "project" {
  value = data.todoist_project.this
}

data "todoist_projects" "this" {
}

output "projects" {
  value = data.todoist_projects.this
}


data "todoist_tasks" "this" {
  is_completed = false
}

output "tasks" {
  value = data.todoist_tasks.this
}

data "todoist_task" "this" {
  id = 4784515043
}
output "task" {
  value = data.todoist_task.this
}


resource "todoist_task" "this"{
  content = "MyTaskWithLabels"
  project_id = 2263974343
  labels = ["MyLabel", "test", "23"]
}

