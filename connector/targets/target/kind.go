package target

import "github.com/kubemq-hub/builder/connector/common/survey"

type Kind struct {
}

func NewKind() *Kind {
	return &Kind{}
}

func (k *Kind) Render() (string, error) {
	val := ""
	err := survey.NewInput().
		SetKind("string").
		SetName("kind").
		SetMessage("Set Target kind").
		SetDefault("target.queue").
		SetHelp("Sets targets kind entry").
		SetRequired(true).
		SetOptions([]string{"target.queue", "target.events", "target.events-store", "target.command", "target.query"}).
		Render(&val)
	if err != nil {
		return "", err
	}
	return val, nil
}