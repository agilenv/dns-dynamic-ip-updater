package track

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_fileStorage_Save(t *testing.T) {
	type fields struct {
		filepath string
	}
	type args struct {
		event Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		remove  bool
	}{
		{
			name: "should create a file and put a track log into it successfully",
			fields: fields{
				filepath: fmt.Sprintf("./testdata/%d", time.Now().Unix()),
			},
			args:    args{event: NewEvent(time.Now(), "1.1.1.1", "test")},
			wantErr: false,
			remove:  true,
		},
		{
			name: "should return an error when it fails to open/create the track file",
			fields: fields{
				filepath: "./testdata",
			},
			args:    args{event: NewEvent(time.Now(), "1.1.1.1", "test")},
			wantErr: true,
			remove:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.remove {
				defer os.Remove(tt.fields.filepath)
			}
			f := NewFileStorage(tt.fields.filepath)
			err := f.Save(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileStorage_LastEvent(t *testing.T) {
	type fields struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		want   Event
	}{
		{
			name:   "should return a valid event from the tracks log",
			fields: fields{filepath: "./testdata/tracks.log"},
			want: func() Event {
				time_, _ := time.Parse(timeLayout, "Mon, 29 Aug 2022 21:25:02 -0300")
				event := NewEvent(time_, "1.2.3.4", "Test")
				return event
			}(),
		},
		{
			name:   "should return an error when tracks log was not found",
			fields: fields{filepath: "./testdata/notfound.log"},
			want:   Event{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFileStorage(tt.fields.filepath)
			if got := f.LastEvent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
