/*
Copyright 2025. Ivan Eroshkin

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

// Action represents the NETCONF `action` message.
type Action struct {
	RPC
	Action struct {
		Data interface{} `xml:",innerxml"`
	} `xml:"urn:ietf:params:xml:ns:yang:1 action"`
}

// NewAction can be used to create an `action` message.
func NewAction(msg string) *Action {
	var rpc Action
	rpc.Action.Data = msg
	rpc.MessageID = uuid()
	return &rpc
}
