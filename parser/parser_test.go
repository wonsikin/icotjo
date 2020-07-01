package parser

import (
	"reflect"
	"testing"
)

func Test_jsonMarshal(t *testing.T) {
	type args struct {
		v            interface{}
		safeEncoding bool
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"disable safe enconding", args{
			v:            map[string]string{"用户名": "User,name", "密码": "Password"},
			safeEncoding: false,
		}, []byte(`{"密码":"Password","用户名":"User,name"}`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := jsonMarshal(tt.args.v, tt.args.safeEncoding)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("jsonMarshal() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
