package models

import (
	"encoding/json"

	"github.com/apimatic/go-core-runtime/types"
)

type GetUsageReportResponse struct {
	Url              types.Optional[string] `json:"url"`
	UsageReportUrl   types.Optional[string] `json:"usage_report_url"`
	GroupedReportUrl types.Optional[string] `json:"grouped_report_url"`
}

func (g *GetUsageReportResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

func (g *GetUsageReportResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.UsageReportUrl.IsValueSet() {
		structMap["usage_report_url"] = g.UsageReportUrl.Value()
	}
	if g.GroupedReportUrl.IsValueSet() {
		structMap["grouped_report_url"] = g.GroupedReportUrl.Value()
	}
	return structMap
}

func (g *GetUsageReportResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Url              types.Optional[string] `json:"url"`
		UsageReportUrl   types.Optional[string] `json:"usage_report_url"`
		GroupedReportUrl types.Optional[string] `json:"grouped_report_url"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.UsageReportUrl = temp.UsageReportUrl
	g.GroupedReportUrl = temp.GroupedReportUrl
	return nil
}
