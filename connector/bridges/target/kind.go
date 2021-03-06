package target

import "github.com/kubemq-hub/builder/survey"

type Kind struct {
	defaultKind string
}

func NewKind(defaultKind string) *Kind {
	return &Kind{
		defaultKind: defaultKind,
	}

}

func (k *Kind) Render() (string, error) {

	val := ""
	err := survey.NewString().
		SetKind("string").
		SetName("kind").
		SetMessage("Set Target kind").
		SetDefault(k.defaultKind).
		SetHelp("Set targets kind entry").
		SetRequired(true).
		SetOptions([]string{"target.queue", "target.events", "target.events-store", "target.command", "target.query"}).
		Render(&val)
	if err != nil {
		return "", err
	}
	return val, nil
}
