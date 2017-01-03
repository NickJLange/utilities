package main

import "testing"

func TestListenLoop(t *testing.T) {
	cases := []struct {
		host, port, prot string; timeout,want int
	}{
  	{"localhost79","3333","udp",1,-99},
		{"localhost","99","tcp",1,-1},
		{"localhost79","3333","tcp",1,-99},
		{"localhost","3333","tcp",1,0},
	}
	for _, c := range cases {
		returnCode := ListenLoop(c.host,c.port,c.prot,c.timeout)
		if returnCode != c.want {
			t.Errorf("Listening on %v:%v/%v returned %v, should have return %v", c.host,c.port,c.prot, returnCode,c.want)
		}
	}
}
