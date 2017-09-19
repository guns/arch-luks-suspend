package goLuksSuspend

import "testing"

func TestParseKeyfileFromCrypttabEntry(t *testing.T) {
	data := []struct {
		in   string
		name string
		key  Keyfile
	}{
		// Malformed input
		{in: "foo"},
		{in: "foo bar"},
		{in: " foo\tbar "},
		// Keyfiles with no options
		{
			in:   "crypt1 UUID=f7dd3b0e-b7ae-4f7c-8c31-4895e4c23231 /root/.keys/crypt1.key",
			name: "crypt1",
			key:  Keyfile{Path: "/root/.keys/crypt1.key"},
		},
		{
			in:   "crypt1 UUID=f7dd3b0e-b7ae-4f7c-8c31-4895e4c23231 /root/.keys/crypt1.key luks,noauto",
			name: "crypt1",
			key:  Keyfile{Path: "/root/.keys/crypt1.key"},
		},
		// Keyfiles with offset and size
		{
			in:   "crypt2 UUID=f7dd3b0e-b7ae-4f7c-8c31-4895e4c23231 /root/.keys/crypt2.key keyfile-size=512,luks,noauto,keyfile-offset=1024",
			name: "crypt2",
			key:  Keyfile{Path: "/root/.keys/crypt2.key", Size: 512, Offset: 1024},
		},
		{
			in:   "crypt2 UUID=f7dd3b0e-b7ae-4f7c-8c31-4895e4c23231 /root/.keys/crypt2.key keyfile-size=512,luks,keyfile-size=1024,noauto",
			name: "crypt2",
			key:  Keyfile{Path: "/root/.keys/crypt2.key", Size: 1024},
		},
		{
			in:   "crypt2 UUID=f7dd3b0e-b7ae-4f7c-8c31-4895e4c23231 /root/.keys/crypt2.key keyfile-size=foo,keyfile-size=4096,,keyfile-offset=1024,luks,,,noauto,keyfile-offset=bar",
			name: "crypt2",
			key:  Keyfile{Path: "/root/.keys/crypt2.key", Size: 4096, Offset: 1024},
		},
	}

	for _, row := range data {
		name, key := parseKeyfileFromCrypttabEntry(row.in)

		if name != row.name {
			t.Errorf("%#v != %#v", name, row.name)
		}

		if key != row.key {
			t.Errorf("%#v != %#v", key, row.key)
		}
	}
}