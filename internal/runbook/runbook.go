package runbook

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsonnet "github.com/google/go-jsonnet"

	"github.com/sjansen/pgutil/internal/runbook/tasks"
)

type Config struct {
	Databases map[string]*Database
	Tasks     map[string]*Task
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	After []string `json:"after"`
	*tasks.Exec
	*tasks.SQL
}

func Load(directory, filename string) (*Config, error) {
	pathname := filepath.Join(directory, filename)
	bytes, err := ioutil.ReadFile(pathname)
	if err != nil {
		return nil, err
	}

	vm := jsonnet.MakeVM()
	importer := &jsonnet.FileImporter{
		JPaths: []string{directory},
	}
	vm.Importer(importer)

	evaluated, err := vm.EvaluateSnippet(filename, string(bytes))
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(strings.NewReader(evaluated))
	dec.DisallowUnknownFields()

	cfg := &Config{}
	err = dec.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}