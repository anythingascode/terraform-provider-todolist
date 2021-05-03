package main

import (
	"encoding/json"
	"fmt"
	todo "terraform-provider-todolist/todoist-go"
)

func main() {
	apikey := "d3f7b0babb4fa120499762bee916b52e47ddbf4e"
	c := todo.NewClient(&apikey)
	//p, _ := c.GetProjects()
	//b, _ := json.Marshal(p)
	//fmt.Println(string(b))
	//p, _ := c.CreateProject("FirstProject")
	//b, _ := json.Marshal(p)
	//fmt.Println(string(b))
	//c.DeleteProject("MyFirstProject")
	//t, _ := c.GetCompletedTasks()
	//fmt.Println(t)
	//fmt.Println(c.GetCompletedTask(4768725770))

	//t := c.CreateNewTask(strings.NewReader(`{"content": "Buy Egg", "project_id": 2263974342}`))
	/*
				id := []int{2156826085}


				)
				tu := todo.Task{
					ID: 4784793696,
					Content: "TaskUpdated",
				}
				c.UpdateTask(&tu)

			task := c.GetActiveTaskById(4784793696)
			jtask, _ := json.Marshal(task)
			fmt.Println(string(jtask))

				//c.DeleteTask(4768738387)
				var taskIds = []int{4784829726, 4784787519, 4784787512, 4784787082, 4784787078, 4784774259, 4768738387, 4784547795, 4784564156, 4784626643, 4784626922, 4784772091}

				for _, id := range taskIds {
					c.DeleteTask(id)
				}

		p, _ := c.GetProjects()
		b, _ := json.Marshal(p)
		fmt.Println("===============")
		fmt.Println(string(b))
		//c.CloseOrReOpenTask(4784793696, "reopen")

	np := &todo.Project{
		ID:   2264650515,
		Name: "Project43210",
	}

	p := c.UpdateProject(np)
	fmt.Println(p)
	/*
		np := &todo.NewProject{
			Name:         "Project54321",
		}
		p, _ := c.CreateProject(np)
		b, _ := json.Marshal(p)
		fmt.Println(string(b))



	labels := []string{"test", "23"}
	lids := c.GetLabelId(labels)
	d, _ := json.Marshal(lids)
	fmt.Println(string(d))

	l := &todo.Labels{
		Name:     "MyLabel",
		Order:    4,
		Color:    47,
		Favorite: false,
	}
	c.CreateLabels(l)
	*/
	labels := []interface{}{"fake", "test"}
	id := c.GetLabelId(labels)
	nt := &todo.Task{
		Content:     "Buy Chicken",
		ProjectID:2263974342,
		LabelIds: id,
	}
	j := c.CreateNewTask(nt)
	d, _ := json.Marshal(j)
	fmt.Println(string(d))
}
