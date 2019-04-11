package loader

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsonnet "github.com/google/go-jsonnet"
)

type Loader struct {
	Queues map[string]func() Queue
	Tasks  map[string]func() TaskConfig
}

type Runbook struct {
	Queues map[string]Queue
	Tasks  map[string]*Task
}

type runbook struct {
	Queues map[string]map[string]json.RawMessage
	Tasks  map[string]*task
}

type Queue interface {
	ConcurrencyLimit() int
	VerifyConfig() error
	VerifyTask(config interface{}) error
}

type Task struct {
	After  []string
	Queue  string
	Config TaskConfig
}

type task struct {
	After  []string
	Queue  string
	Type   string
	Config json.RawMessage
}

type TaskConfig interface {
	VerifyConfig() error
}

func (l *Loader) Load(filename string) (*Runbook, error) {
	directory := filepath.Dir(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vm := jsonnet.MakeVM()
	importer := &jsonnet.FileImporter{
		JPaths: []string{directory},
	}
	vm.Importer(importer)

	evaluated, err := vm.EvaluateSnippet(filename, string(data))
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(
		strings.NewReader(evaluated),
	)
	dec.DisallowUnknownFields()

	tmp := &runbook{}
	err = dec.Decode(tmp)
	if err != nil {
		return nil, err
	}

	book := &Runbook{
		Queues: map[string]Queue{},
		Tasks:  map[string]*Task{},
	}

	if err = l.loadQueues(tmp, book); err != nil {
		return nil, err
	}

	if err = l.loadTasks(tmp, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (l *Loader) loadQueues(tmp *runbook, book *Runbook) error {
	for queueType, queues := range tmp.Queues {
		for name, config := range queues {
			dec := json.NewDecoder(
				bytes.NewReader(config),
			)
			dec.DisallowUnknownFields()

			factory, ok := l.Queues[queueType]
			if !ok {
				return errors.New("invalid queue type")
			}
			queue := factory()

			if err := dec.Decode(queue); err != nil {
				return err
			}
			if err := queue.VerifyConfig(); err != nil {
				return err
			}

			if name == "" {
				book.Queues[queueType] = queue
			} else {
				book.Queues[queueType+"/"+name] = queue
			}
		}
	}
	for queueType, factory := range l.Queues {
		if _, ok := book.Queues[queueType]; !ok {
			book.Queues[queueType] = factory()
		}
	}
	return nil
}

func (l *Loader) loadTasks(tmp *runbook, book *Runbook) error {
	for id, raw := range tmp.Tasks {
		taskType := raw.Type
		if taskType == "" {
			taskType = raw.Queue
			if idx := strings.Index(taskType, "/"); idx >= 0 {
				taskType = taskType[:idx]
			}
		}

		factory, ok := l.Tasks[taskType]
		if !ok {
			return errors.New("invalid task type")
		}
		task := factory()

		if len(raw.Config) > 0 {
			dec := json.NewDecoder(
				bytes.NewReader(raw.Config),
			)
			dec.DisallowUnknownFields()
			if err := dec.Decode(task); err != nil {
				return err
			}
		}
		if err := task.VerifyConfig(); err != nil {
			return err
		}

		if queue, ok := book.Queues[raw.Queue]; ok {
			if err := queue.VerifyTask(task); err != nil {
				return err
			}
		} else {
			return errors.New("invalid queue ID")
		}

		book.Tasks[id] = &Task{
			Queue:  raw.Queue,
			After:  raw.After,
			Config: task,
		}
	}
	return nil
}
