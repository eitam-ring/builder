package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/kubemq-hub/builder/connector"
	"io/ioutil"
	"sort"
)

type ConnectorsFileHandler struct {
	filename   string
	connectors map[string]*connector.Connector
}

func NewConnectorsFileHandler(filename string) (*ConnectorsFileHandler, error) {
	c := &ConnectorsFileHandler{
		filename:   filename,
		connectors: map[string]*connector.Connector{},
	}
	if err := c.load(); err != nil {
		if err := c.save(); err != nil {
			return nil, err
		}
	}
	return c, nil
}
func (c *ConnectorsFileHandler) load() error {
	data, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &c.connectors)
	if err != nil {
		return err
	}
	return nil
}
func (c *ConnectorsFileHandler) save() error {
	data, err := yaml.Marshal(c.connectors)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
func (c *ConnectorsFileHandler) Add(connector *connector.Connector) error {
	c.connectors[connector.Key()] = connector
	return c.save()
}

func (c *ConnectorsFileHandler) Edit(connector *connector.Connector) error {
	c.connectors[connector.Key()] = connector
	return c.save()
}

func (c *ConnectorsFileHandler) Delete(connector *connector.Connector) error {
	delete(c.connectors, connector.Key())
	return c.save()
}

func (c *ConnectorsFileHandler) Get(namespace, name string) (*connector.Connector, error) {
	key := fmt.Sprintf("%s/%s", namespace, name)
	con, ok := c.connectors[key]
	if !ok {
		return nil, fmt.Errorf("connector not found")
	}
	return con, nil

}

func (c *ConnectorsFileHandler) List() []*connector.Connector {
	var list []*connector.Connector
	for _, con := range c.connectors {
		list = append(list, con)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Key() < list[j].Key()
	})
	return list
}
