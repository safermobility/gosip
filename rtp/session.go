// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// RTP Session Transport Layer.

package rtp

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

const (
	rtpBindMaxAttempts = 10
	rtpBindPortMin     = 16384
	rtpBindPortMax     = 32768
)

func Listen(host string) (sock net.PacketConn, err error) {
	if strings.Contains(host, ":") {
		return net.ListenPacket("udp", host)
	}
	for i := 0; i < rtpBindMaxAttempts; i++ {
		port := rtpBindPortMin + rand.Int63()%(rtpBindPortMax-rtpBindPortMin+1)
		if port%2 == 1 {
			port--
		}
		saddr := net.JoinHostPort(host, strconv.FormatInt(port, 10))
		sock, err = net.ListenPacket("udp", saddr)
		if err == nil || !strings.Contains(err.Error(), "address already in use") {
			break
		}
		log.Println("RTP listen congestion:", saddr)
	}
	return
}
