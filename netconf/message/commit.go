/*
Copyright 2021. Alexis de Talhouët

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package message

// Commit represents the NETCONF `commit` message.
// https://datatracker.ietf.org/doc/html/rfc6241#section-8.3.4.1
type Commit struct {
	RPC
	Commit commit `xml:"commit"`
}

type commit struct {
	XMLName        string  `xml:"commit"`
	Confirmed      *string `xml:"confirmed,omitempty"`
	ConfirmTimeout *int    `xml:"confirm-timeout,omitempty"`
}

// NewCommit can be used to create a `commit` message.
//
// Example commit message:
// <rpc message-id="101">
// ..<commit></commit>
// </rpc>
//
// Example commit-confirmed message:
// <rpc message-id="101">
// ..<commit>
// ....<confirmed/>
// ..</commit>
// </rpc>
//
// Example commit-confirmed message with custom timeout (seconds):
// <rpc message-id="101">
// ..<commit>
// ....<confirmed/>
// ....<confirm-timeout>120</confirm-timeout>
// ..</commit>
// </rpc>
func NewCommit(confirmed bool, confirmTimeout int) *Commit {
	var rpc Commit
	if confirmed {
		s := ""
		rpc.Commit.Confirmed = &s
	}
	if confirmTimeout > 0 {
		rpc.Commit.ConfirmTimeout = &confirmTimeout
	}
	rpc.MessageID = uuid()
	return &rpc
}
