package cluster

import (
	"fmt"
	"github.com/kubemq-hub/builder/survey"
	"math"
)

type Routing struct {
	Data       string `json:"data"`
	Url        string `json:"url"`
	AutoReload int    `json:"auto_reload"`
}

func NewRouting() *Routing {
	return &Routing{}
}
func (r *Routing) askData() error {
	err := survey.NewMultiline().
		SetKind("multiline").
		SetName("policy").
		SetMessage("Load smart routing data").
		SetDefault("").
		SetHelp("Load smart routing data").
		SetRequired(false).
		Render(&r.Data)
	if err != nil {
		return err
	}
	return nil
}
func (r *Routing) askUrl() error {
	err := survey.NewString().
		SetKind("string").
		SetName("url").
		SetMessage("Set URL routing data loading address").
		SetDefault("").
		SetHelp("Set URL routing data loading address").
		SetRequired(false).
		Render(&r.Url)
	if err != nil {
		return err
	}
	return nil
}
func (r *Routing) askAutoReload() error {
	err := survey.NewInt().
		SetKind("int").
		SetName("auto-reload").
		SetMessage("Set automatic routing data reload seconds (0 - no reload)").
		SetDefault("0").
		SetHelp("Set automatic routing data reload seconds (0 - no reload)").
		SetRequired(false).
		SetRange(0, math.MaxInt32).
		Render(&r.AutoReload)
	if err != nil {
		return err
	}
	return nil
}

func (r *Routing) Validate() error {
	if r.AutoReload < 0 {
		return fmt.Errorf("auto reload value cannot be less than 0")
	}
	return nil
}

func (r *Routing) Render() (*Routing, error) {
	if err := r.askData(); err != nil {
		return nil, err
	}
	if err := r.askUrl(); err != nil {
		return nil, err
	}

	if err := r.askAutoReload(); err != nil {
		return nil, err
	}
	return r, nil
}

var _ Validator = NewRouting()
