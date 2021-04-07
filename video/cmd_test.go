package video

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestNewProbeFormat(t *testing.T) {
	result := ProbeResult{
		Format: &ProbeFormat{
			2,

			"mov,mp4,m4a,3gp,3g2,mj2",
			"QuickTime / MOV",
			mduration{0.000000},
			mduration{6358},
			"298688",
			humanSize{"237397835"},
		},
	}
	type args struct {
		ctx      context.Context
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult ProbeResult
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"/Volumes/compatible/VIDEO/张大池直播/run1.MP4", args{context.Background(), "/Users/kemi/image/test.MP4"}, result, false},

		{"/Volumes/compatible/VIDEO/张大池直播/run1.MP4", args{context.Background(), "/Volumes/compatible/VIDEO/张大池直播/run1.MP4"}, result, false},
		{"/Users/kemi/Downloads/video/smaller人性和强势文化.mp4", args{context.Background(), "/Users/kemi/Downloads/video/smaller人性和强势文化.mp4"}, result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := NewProbeFormat(tt.args.ctx, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProbeFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NewProbeFormat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNewProbeFormat1(t *testing.T) {
	type args struct {
		ctx      context.Context
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantResult ProbeResult
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"/Users/kemi/Downloads/video/smaller人性和强势文化.mp4", args{context.Background(), "/Users/kemi/Downloads/video/smaller人性和强势文化.mp4"}, ProbeResult{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := NewProbeFormat(tt.args.ctx, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProbeFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("NewProbeFormat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_mduration_UnmarshalText(t *testing.T) {
	type fields struct {
		Duration time.Duration
	}
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &mduration{
				Duration: tt.fields.Duration,
			}
			if err := x.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
