package dns

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/agilenv/linkip/internal/dns/track"
	"github.com/golang/mock/gomock"
)

func TestStats_LastExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEvent := track.NewEvent(time.Now(), "1.1.1.1", "Foo")
	type fields struct {
		tracks TrackRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   *track.Event
	}{
		{
			name: "should return the last event executed",
			fields: fields{
				tracks: func() TrackRepository {
					repo := NewMockTrackRepository(ctrl)
					repo.EXPECT().LastEvent().Return(mockEvent)
					return repo
				}(),
			},
			want: &mockEvent,
		},
		{
			name: "should return nil when the get an empty event from repository",
			fields: fields{
				tracks: func() TrackRepository {
					repo := NewMockTrackRepository(ctrl)
					repo.EXPECT().LastEvent().Return(track.Event{})
					return repo
				}(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewStats(tt.fields.tracks)
			if got := u.LastExecution(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastExecution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStats_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		tracks TrackRepository
	}
	type args struct {
		event track.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should save the event to the repository",
			fields: fields{
				tracks: func() TrackRepository {
					repo := NewMockTrackRepository(ctrl)
					repo.EXPECT().Save(gomock.Any()).Return(nil)
					return repo
				}(),
			},
			wantErr: false,
		},
		{
			name: "should return an error when the repository fails",
			fields: fields{
				tracks: func() TrackRepository {
					repo := NewMockTrackRepository(ctrl)
					repo.EXPECT().Save(gomock.Any()).Return(errors.New("failed"))
					return repo
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewStats(tt.fields.tracks)
			if err := u.Save(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
