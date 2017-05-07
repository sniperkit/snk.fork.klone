// Copyright © 2017 Kris Nova <kris@nivenly.com>
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
//
//  _  ___
// | |/ / | ___  _ __   ___
// | ' /| |/ _ \| '_ \ / _ \
// | . \| | (_) | | | |  __/
// |_|\_\_|\___/|_| |_|\___|
//
// config.go is how we interact with git configuration

package provider_alpha_1

type GitConfig struct {
}

func (g *GitConfig) GetConfigPath() (string, error) {
	return "", nil
}
func (g *GitConfig) GetConfigBytes() ([]byte, error) {
	return []byte(""), nil
}
func (g *GitConfig) GetConfigString() (string, error) {
	return "", nil
}
func (g *GitConfig) GetConfigDirective(key string) (string, error) {
	return "", nil
}
func (g *GitConfig) SetConfigDirective(key string, value string) error {
	return nil
}