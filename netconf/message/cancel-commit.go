/*
Copyright 2025. Hannu Kamarainen

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

// CancelCommit represents the NETCONF `cancel-commit` message.
// https://datatracker.ietf.org/doc/html/rfc6241#section-8.4.4.1
type CancelCommit struct {
	RPC
	CancelCommit struct{} `xml:"cancel-commit"`
}

// NewCancelCommit can be used to create a `cancel-commit` message.
//
// Example cancel-commit message:
// <rpc message-id="102" xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
// ..<cancel-commit/>
// </rpc>
func NewCancelCommit() *CancelCommit {
	var rpc CancelCommit
	rpc.MessageID = uuid()
	return &rpc
}
