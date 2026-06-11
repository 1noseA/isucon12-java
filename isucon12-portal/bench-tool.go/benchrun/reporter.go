package benchrun

import (
	"encoding/json"
	"fmt"

	isuxportalResources "github.com/isucon/isucon12-portal/proto.go/isuxportal/resources"
)

type Reporter struct {
	raw bool
}

func NewReporter(raw bool) (*Reporter, error) {
	return &Reporter{raw: raw}, nil
}

func (r *Reporter) Report(res *isuxportalResources.BenchmarkResult) error {
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
