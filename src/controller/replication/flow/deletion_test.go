// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flow

import (
	"context"
	"testing"

	"github.com/goharbor/harbor/src/replication/model"
	"github.com/goharbor/harbor/src/testing/pkg/task"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type deletionFlowTestSuite struct {
	suite.Suite
}

func (d *deletionFlowTestSuite) TestRun() {
	taskMgr := &task.Manager{}
	taskMgr.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

	policy := &model.Policy{
		SrcRegistry: &model.Registry{
			Type: model.RegistryTypeHarbor,
		},
		DestRegistry: &model.Registry{
			Type: model.RegistryTypeHarbor,
		},
	}
	resources := []*model.Resource{
		{
			Metadata: &model.ResourceMetadata{
				Repository: &model.Repository{
					Name: "library/hello-world",
				},
				Artifacts: []*model.Artifact{
					{
						Tags: []string{"latest"},
					},
				},
			},
		},
	}
	flow := &deletionFlow{
		executionID: 1,
		policy:      policy,
		taskMgr:     taskMgr,
		resources:   resources,
	}
	err := flow.Run(context.Background())
	d.Require().Nil(err)
}

func TestDeletionFlowTestSuite(t *testing.T) {
	suite.Run(t, &deletionFlowTestSuite{})
}
