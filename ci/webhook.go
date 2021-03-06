// Copyright 2018 The Pouch Robot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ci

// Webhook represents a struct in payload
type Webhook struct {
	ID                int    `json:"id"`
	Number            string `json:"number"`
	PullRequestNumber int    `json:"pull_request_number"`
	PullRequestTitle  string `json:"pull_request_title"`
	Duration          int    `json:"duration"`
	AuthorName        string `json:"author_name"`
	AuthorEmail       string `json:"author_email"`
	Type              string `json:"type"`
	State             string `json:"state"`
	BuildURL          string `json:"build_url"`
}
