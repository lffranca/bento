package http

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/lffranca/bento/processor/handler"
)

func TestProcessor_Handler(t *testing.T) {
	type fields struct {
		URL         string
		QueryParams map[string]string
		Headers     map[string]string
		Cookies     []*http.Cookie
	}
	type args struct {
		in0 context.Context
		in1 handler.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    handler.Response
		wantErr bool
	}{
		{
			name: "01",
			fields: fields{
				URL: "https://google.com",
				QueryParams: map[string]string{
					"limit":  "20",
					"offset": "10",
				},
				Headers: map[string]string{
					"Authentication": "Bearer tygfvgtyhbgvtyhunhbg.sadq23dwqe3d32dasd23dwqd.32d2wd23d23e2",
					"Accept":         "application/json",
				},
			},
			args: args{
				in0: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &Processor{
				URL:         tt.fields.URL,
				QueryParams: tt.fields.QueryParams,
				Headers:     tt.fields.Headers,
				Cookies:     tt.fields.Cookies,
			}
			got, err := item.Handler(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Processor.Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Processor.Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
