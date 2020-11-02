package model

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-model/go/datasource"
	"strings"
)

type Query struct {
	RefID	string	`json:"refId,omitempty"`
	Format	string	`json:"format"`
	Query	string	`json:"query"`
}

func CreateQueryFrom(query datasource.Query) (*Query, error) {
	model := &Query{}
	err := json.Unmarshal([]byte(query.ModelJson), &model)
	if err != nil {
		return nil, err
	}
	model.RefID = query.RefId

	model.Query = strings.TrimSpace(model.Query)
	model.Format = strings.TrimSpace(model.Format)

	if len(model.Format) == 0 {
		model.Format = "table"
	}

	return model, nil
}

func (m *Query) String() string {
	jsonBytes, _ := json.Marshal(m)
	return string(jsonBytes)
}

func (m *Query) IsEmpty() bool {
	return len(m.Query) == 0
}
