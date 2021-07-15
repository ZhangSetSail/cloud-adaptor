// RAINBOND, Application Management Platform
// Copyright (C) 2020-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package task

import (
	"context"
	"fmt"
	"testing"

	"goodrain.com/cloud-adaptor/internal/types"
)

func TestInitRainbondCluster(t *testing.T) {
	task, err := CreateTask(InitRainbondClusterTask, &types.InitRainbondConfig{
		Provider:  "ack",
		SecretKey: "hBsW4mlp35xQlqvqvm5Izmbt2UFR6E",
		AccessKey: "LTAI4FtxHG8A8h328zBBNMtw",
		ClusterID: "c9b4516adec504a458bdf9a6891fc74fe",
	})
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for {
			message := <-task.GetChan()
			fmt.Println(message)
		}
	}()
	task.Run(context.TODO())
}

func TestInitRKERainbondCluster(t *testing.T) {
	task, err := CreateTask(InitRainbondClusterTask, &types.InitRainbondConfig{
		Provider:  "rke",
		ClusterID: "c9b4516adec504a458bdf9a6891fc74fe",
	})
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for {
			message := <-task.GetChan()
			fmt.Println(message)
		}
	}()
	task.Run(context.TODO())
}

func TestCheckVersion(t *testing.T) {
	tests := []struct {
		version string
		isPass  bool
	}{
		{
			version: "1.15.9",
			isPass:  false,
		},
		{
			version: "1.16.0",
			isPass:  true,
		},
		{
			version: "1.19.0",
			isPass:  true,
		},
		{
			version: "1.20.0",
			isPass:  false,
		},
		{
			version: "v1.15.9-rke",
			isPass:  false,
		},
		{
			version: "v1.16.0-rke",
			isPass:  true,
		},
		{
			version: "v1.19.0-rke",
			isPass:  true,
		},
		{
			version: "v1.20.0-rke",
			isPass:  false,
		},
		{
			version: "1.15.9-aliyun.1",
			isPass:  false,
		},
		{
			version: "1.16.0-aliyun.1",
			isPass:  true,
		},
		{
			version: "1.19.0-aliyun.1",
			isPass:  true,
		},
		{
			version: "1.20.0-aliyun.1",
			isPass:  false,
		},
	}
	for _,tc:= range tests{
		if checkVersion(tc.version) != tc.isPass{
			t.Fatalf("version %v expect %v, actual %v", tc.version, tc.isPass, checkVersion(tc.version))
		}
	}
}
