package magnet

import (
	"reflect"
	"testing"
)

func TestMagnetParse(t *testing.T) {
	type args struct {
		magnetURI string
	}
	tests := []struct {
		name       string
		args       args
		wantMagnet *Magnet
		wantErr    bool
	}{
		{
			name: "ok_basic",
			args: args{
				magnetURI: "magnet:?xt=urn:btih:3d6da9b3e9f4e4f8e1e0b3f7b2e2e3f4&dn=example&xl=123456",
			},
			wantMagnet: &Magnet{
				ExactTopics:  []string{"urn:btih:3d6da9b3e9f4e4f8e1e0b3f7b2e2e3f4"},
				DisplayNames: []string{"example"},
				ExactLength:  123456,
			},
			wantErr: false,
		},
		{
			name: "ok_big_buck_bunny",
			args: args{
				magnetURI: "magnet:?xt=urn:btih:88594aaacbde40ef3e2510c47374ec0aa396c08e&dn=Big+Buck+Bunny+1080p+30fps&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451&tr=udp%3A%2F%2Fopen.stealth.si%3A80&tr=udp%3A%2F%2Ftracker.moeking.me%3A6969&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969&tr=udp%3A%2F%2Fopen.tracker.cl%3A1337",
			},
			wantMagnet: &Magnet{
				ExactTopics:  []string{"urn:btih:88594aaacbde40ef3e2510c47374ec0aa396c08e"},
				DisplayNames: []string{"Big Buck Bunny 1080p 30fps"},
				ExactLength:  0,
				AddressTrackers: []string{
					"udp://tracker.opentrackr.org:1337",
					"udp://tracker.torrent.eu.org:451",
					"udp://open.stealth.si:80",
					"udp://tracker.moeking.me:6969",
					"udp://opentracker.i2p.rocks:6969",
					"udp://open.tracker.cl:1337",
				},
			},
			wantErr: false,
		},
		{
			name: "ok_sintel",
			args: args{
				magnetURI: "magnet:?xt=urn:btih:08ada5a7a6183aae1e09d831df6748d566095a10&dn=Sintel&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fsintel.torrent",
			},
			wantMagnet: &Magnet{
				ExactTopics:  []string{"urn:btih:08ada5a7a6183aae1e09d831df6748d566095a10"},
				DisplayNames: []string{"Sintel"},
				ExactLength:  0,
				AddressTrackers: []string{
					"udp://explodie.org:6969",
					"udp://tracker.coppersurfer.tk:6969",
					"udp://tracker.empire-js.us:1337",
					"udp://tracker.leechers-paradise.org:6969",
					"udp://tracker.opentrackr.org:1337",
					"wss://tracker.btorrent.xyz",
					"wss://tracker.fastcast.nz",
					"wss://tracker.openwebtorrent.com",
				},
				ExactSources: []string{"https://webtorrent.io/torrents/sintel.torrent"},
				WebSeeds:     []string{"https://webtorrent.io/torrents/"},
			},
			wantErr: false,
		},
		{
			name: "fail_missing_magnet_uri_prefix",
			args: args{
				magnetURI: "not_a_magnet_uri",
			},
			wantMagnet: nil,
			wantErr:    true,
		},
		{
			name: "fail_malformed_magnet_uri",
			args: args{
				magnetURI: "magnet:?fail=%GG",
			},
			wantMagnet: nil,
			wantErr:    true,
		},
		{
			name: "fail_two_xl_fields",
			args: args{
				magnetURI: "magnet:?xt=urn:btih:3d6da9b3e9f4e4f8e1e0b3f7b2e2e3f4&dn=example&xl=123456&xl=42",
			},
			wantMagnet: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMagnet, err := Parse(tt.args.magnetURI)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotMagnet, tt.wantMagnet) {
				t.Errorf("Parse() = %+v, want %+v", gotMagnet, tt.wantMagnet)
			}

			if !tt.wantErr {
				str := gotMagnet.String()
				gotMagnet2, err := Parse(str)
				if err != nil {
					t.Errorf("Parse() error = %v", err)
				}

				if !reflect.DeepEqual(gotMagnet, gotMagnet2) {
					t.Errorf("Parse() = %+v, want %+v", gotMagnet, tt.wantMagnet)
				}
			}
		})
	}
}
