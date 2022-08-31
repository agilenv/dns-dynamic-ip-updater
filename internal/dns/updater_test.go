package dns

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/agilenv/linkip/internal/dns/track"

	"github.com/golang/mock/gomock"
)

func TestUpdater_SearchForChanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		provider    DNSProvider
		publicIPAPI PublicIPAPI
		stats       StatsUsecase
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		changed bool
		ip      string
		wantErr bool
	}{
		{
			name: "should return new IP founded from repository (first time -> no stats log)",
			fields: fields{
				provider: nil,
				publicIPAPI: func() PublicIPAPI {
					repo := NewMockPublicIPAPI(ctrl)
					repo.EXPECT().Get(gomock.Any()).Return("1.2.3.4", nil)
					return repo
				}(),
				stats: func() StatsUsecase {
					st := NewMockStatsUsecase(ctrl)
					st.EXPECT().LastExecution().Return(&track.Event{})
					return st
				}(),
			},
			args:    args{ctx: context.Background()},
			changed: true,
			ip:      "1.2.3.4",
			wantErr: false,
		},
		{
			name: "should return last execution if the public IP has not changed",
			fields: fields{
				provider: nil,
				publicIPAPI: func() PublicIPAPI {
					repo := NewMockPublicIPAPI(ctrl)
					repo.EXPECT().Get(gomock.Any()).Return("1.2.3.4", nil)
					return repo
				}(),
				stats: func() StatsUsecase {
					e := track.NewEvent(time.Now(), "1.2.3.4", "Test")
					st := NewMockStatsUsecase(ctrl)
					st.EXPECT().LastExecution().Return(&e)
					return st
				}(),
			},
			args:    args{ctx: context.Background()},
			changed: false,
			ip:      "1.2.3.4",
			wantErr: false,
		},
		{
			name: "should select the public IP over the last data saved in log",
			fields: fields{
				provider: nil,
				publicIPAPI: func() PublicIPAPI {
					repo := NewMockPublicIPAPI(ctrl)
					repo.EXPECT().Get(gomock.Any()).Return("1.2.3.4", nil)
					return repo
				}(),
				stats: func() StatsUsecase {
					e := track.NewEvent(time.Now(), "4.3.2.1", "Test")
					st := NewMockStatsUsecase(ctrl)
					st.EXPECT().LastExecution().Return(&e)
					return st
				}(),
			},
			args:    args{ctx: context.Background()},
			changed: true,
			ip:      "1.2.3.4",
			wantErr: false,
		},
		{
			name: "should return an error when it fails to get the public IP",
			fields: fields{
				provider: nil,
				publicIPAPI: func() PublicIPAPI {
					repo := NewMockPublicIPAPI(ctrl)
					repo.EXPECT().Get(gomock.Any()).Return("", errors.New("failed"))
					return repo
				}(),
				stats: nil,
			},
			args:    args{ctx: context.Background()},
			changed: false,
			ip:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUpdater(tt.fields.provider, tt.fields.stats, tt.fields.publicIPAPI)
			got, got1, err := u.SearchForChanges(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchForChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.changed {
				t.Errorf("SearchForChanges() got = %v, want %v", got, tt.changed)
			}
			if got1 != tt.ip {
				t.Errorf("SearchForChanges() got1 = %v, want %v", got1, tt.ip)
			}
		})
	}
}

func TestUpdater_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		provider    DNSProvider
		publicIPAPI PublicIPAPI
		stats       StatsUsecase
	}
	type args struct {
		ctx context.Context
		ip  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should update the ip to the dns provider",
			fields: fields{
				provider: func() DNSProvider {
					p := NewMockDNSProvider(ctrl)
					p.EXPECT().UpdateRecord(gomock.Any(), "1.2.3.4").Return(nil)
					return p
				}(),
				publicIPAPI: func() PublicIPAPI {
					a := NewMockPublicIPAPI(ctrl)
					a.EXPECT().Name().Return("Test")
					return a
				}(),
				stats: func() StatsUsecase {
					s := NewMockStatsUsecase(ctrl)
					s.EXPECT().Save(track.Event{IP: "1.2.3.4", PublicAPI: "Test"})
					return s
				}(),
			},
			args:    args{ctx: context.Background(), ip: "1.2.3.4"},
			wantErr: false,
		},
		{
			name: "should return an error when the update fails",
			fields: fields{
				provider: func() DNSProvider {
					p := NewMockDNSProvider(ctrl)
					p.EXPECT().UpdateRecord(gomock.Any(), "1.2.3.4").Return(errors.New("failed"))
					return p
				}(),
				publicIPAPI: nil,
				stats:       nil,
			},
			args:    args{ctx: context.Background(), ip: "1.2.3.4"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Updater{
				provider:    tt.fields.provider,
				publicIPAPI: tt.fields.publicIPAPI,
				stats:       tt.fields.stats,
			}
			if err := u.Update(tt.args.ctx, tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
