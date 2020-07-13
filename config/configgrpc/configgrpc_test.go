// Copyright 2019, OpenTelemetry Authors
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

package configgrpc

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicGrpcSettings(t *testing.T) {

	_, err := GrpcSettingsToDialOptions(GRPCSettings{
		Headers:             nil,
		Endpoint:            "",
		Compression:         "",
		CertPemFile:         "",
		UseSecure:           false,
		ServerNameOverride:  "",
		KeepaliveParameters: nil,
	})

	assert.Nil(t, err)
}

func TestInvalidPemFile(t *testing.T) {

	_, err := GrpcSettingsToDialOptions(GRPCSettings{
		Headers:             nil,
		Endpoint:            "",
		Compression:         "",
		CertPemFile:         "/doesnt/exist",
		UseSecure:           false,
		ServerNameOverride:  "",
		KeepaliveParameters: nil,
	})

	// don't validate the specific error code as this differs on windows/unix
	pathErr := err.(*os.PathError)
	assert.Equal(t, pathErr.Op, "open")
	assert.Equal(t, pathErr.Path, "/doesnt/exist")
	assert.NotNil(t, pathErr.Err)
}

func TestUseSecure(t *testing.T) {
	dialOpts, err := GrpcSettingsToDialOptions(GRPCSettings{
		Headers:             nil,
		Endpoint:            "",
		Compression:         "",
		CertPemFile:         "",
		UseSecure:           true,
		ServerNameOverride:  "",
		KeepaliveParameters: nil,
	})

	assert.Nil(t, err)
	assert.Equal(t, len(dialOpts), 1)
}
