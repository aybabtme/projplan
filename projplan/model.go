package projplan

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration struct {
	Min time.Duration
	Max time.Duration
}

func (dur *Duration) UnmarshalJSON(p []byte) error {
	var raw struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}
	err := json.Unmarshal(p, &raw)
	if err != nil {
		return err
	}
	if dur.Min, err = time.ParseDuration(raw.Min); err != nil {
		return err
	}
	if dur.Max, err = time.ParseDuration(raw.Max); err != nil {
		return err
	}
	return nil
}

type Task struct {
	Name     string    `json:"name"`
	Duration *Duration `json:"duration"`
	Depends  []string  `json:"depends"`
}

type Project struct {
	tasks []*Task

	taskRoots         map[string]*Task
	taskByName        map[string]*Task
	taskDependsByName map[string]map[string]*Task
}

func PlanProject(tasks []*Task) (*Project, error) {
	prj := new(Project)
	return prj, prj.prepare(tasks)
}

func (prj *Project) prepare(tasks []*Task) error {
	// create an index of tasks
	taskByName := make(map[string]*Task, len(prj.tasks))
	for _, task := range prj.tasks {
		if _, ok := taskByName[task.Name]; ok {
			return fmt.Errorf("duplicate task name: %q", task.Name)
		}
		taskByName[task.Name] = task
	}
	// check that all dependencies are satisfiable
	taskRoots := make(map[string]*Task)
	taskDependsByName := make(map[string]map[string]*Task, len(prj.tasks))
	for _, task := range prj.tasks {
		if len(task.Depends) == 0 {
			taskRoots[task.Name] = task
			continue
		}
		for _, depName := range task.Depends {
			if depTask, ok := taskByName[depName]; !ok {
				return fmt.Errorf("task %q depends on another task that doesn't exist: %q", task.Name, depName)
			} else {
				deps := taskDependsByName[task.Name]
				if _, ok := deps[depName]; ok {
					return fmt.Errorf("task %q has duplicate dependency: %q", task.Name, depName)
				}
				deps[depName] = depTask
			}
		}
	}
	prj.taskByName = taskByName
	prj.taskDependsByName = taskDependsByName
	prj.tasks = tasks
	return nil
}
