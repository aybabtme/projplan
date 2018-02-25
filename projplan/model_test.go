package projplan_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aybabtme/projplan/projplan"
)

func parseFile(filename string) ([]*projplan.Task, error) {
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var out []*projplan.Task
	return out, json.Unmarshal(p, &out)
}

func TestExamplesAreValidProject(t *testing.T) {
	exampleDir := "../example"
	fis, err := ioutil.ReadDir(exampleDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, fi := range fis {
		if strings.ToLower(filepath.Ext(fi.Name())) != ".json" {
			continue
		}
		tasks, err := parseFile(filepath.Join(exampleDir, fi.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := projplan.PlanProject(tasks); err != nil {
			t.Fatal(err)
		}
	}
}

func TestParseTasks(t *testing.T) {
	tasks, err := parseFile("../example/boat_buying.json")
	if err != nil {
		t.Fatal(err)
	}
	project, err := projplan.PlanProject(tasks)
	if err != nil {
		t.Fatal(err)
	}
	_ = project
}
