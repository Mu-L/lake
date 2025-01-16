/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package archived

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
)

type CircleciWorkflow struct {
	ConnectionId   uint64 `gorm:"primaryKey;type:BIGINT"`
	Id             string `gorm:"primaryKey;type:varchar(100)"`
	ProjectSlug    string `gorm:"type:varchar(255)"`
	PipelineId     string `gorm:"type:varchar(100)"`
	CanceledBy     string `gorm:"type:varchar(100)"`
	Name           string `gorm:"type:varchar(255)"`
	ErroredBy      string `gorm:"type:varchar(100)"`
	Tag            string `gorm:"type:varchar(100)"`
	Status         string `gorm:"type:varchar(100)"`
	StartedBy      string `gorm:"type:varchar(100)"`
	PipelineNumber int64
	CreatedAt      *time.Time
	StoppedAt      *time.Time
	DurationSec    uint64

	archived.NoPKModel
}

func (CircleciWorkflow) TableName() string {
	return "_tool_circleci_workflows"
}
