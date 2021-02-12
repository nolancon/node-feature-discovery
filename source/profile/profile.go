/*
Copyright 2020 The Kubernetes Authors.

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

package profile

import (
	"log"

	"sigs.k8s.io/node-feature-discovery/source"
)

type Profile struct {
	Name     string            `json:"name"`
	Features map[string]string `json:"features"`
}

type Config []Profile

// newDefaultConfig returns a new config with pre-populated defaults
func newDefaultConfig() *Config {
	return &Config{}
}

// Implements FeatureSource Interface
type Source struct {
	Config *Config
}

// Return name of the feature source
func (s Source) Name() string { return "profile" }

// NewConfig method of the FeatureSource interface
func (s *Source) NewConfig() source.Config { return newDefaultConfig() }

// GetConfig method of the FeatureSource interface
func (s *Source) GetConfig() source.Config { return s.Config }

// SetConfig method of the FeatureSource interface
func (s *Source) SetConfig(conf source.Config) {
	switch v := conf.(type) {
	case *Config:
		s.Config = v
	default:
		log.Printf("PANIC: invalid config type: %T", conf)
	}
}

// Discover features
func (s Source) Discover() (source.Features, error) {
	features := source.Features{}
	return features, nil
}
