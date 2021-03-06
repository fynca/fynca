// Copyright 2022 Evan Hazlett
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
package services

import (
	"github.com/fynca/fynca/pkg/auth"
	"google.golang.org/grpc"
)

type Type string

const (
	AccountsService Type = "fynca.services.accounts.v1"
	JobsService     Type = "fynca.services.jobs.v1"
	WorkersService  Type = "fynca.services.workers.v1"
)

// Service is the interface that all stellar services must implement
type Service interface {
	// Type returns the type that the service provides
	Type() Type
	// Register registers the service with the GRPC server
	Register(*grpc.Server) error
	// Configure configures the service
	Configure(auth.Authenticator) error
	// Start provides a mechanism to start service specific actions
	Start() error
	// Stop provides a mechanism to stop the service
	Stop() error
}
