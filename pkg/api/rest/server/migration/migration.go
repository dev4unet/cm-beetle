/*
Copyright 2019 The Cloud-Barista Authors.
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

// Package migration is to handle REST API for migration
package migration

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
}

type Infrastructure struct {
	Network        string
	Disk           string
	Compute        string
	SecurityGroup  string
	VirtualMachine string
}

type MigrateInfraRequest struct {
	Infrastructure
}

type MigrateInfraResponse struct {
	Infrastructure
}

// MigrateInfra godoc
// @Summary Migrate an infrastructure on a cloud platform
// @Description It migrates an infrastructure on a cloud platform. Infrastructure includes network, storage, compute, and so on.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param InfrastructureInfo body MigrateInfraRequest true "Specify network, disk, compute, security group, virtual machine, etc."
// @Success 200 {object} MigrateInfraResponse "Successfully migrated infrastructure on a cloud platform"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /migration/infra [post]
func (rh *Handlers) MigrateInfra(c echo.Context) error {

	// Input
	req := &MigrateInfraRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Print(req.Network)
	fmt.Print(req.Disk)
	fmt.Print(req.Compute)
	fmt.Print(req.SecurityGroup)
	fmt.Print(req.VirtualMachine)

	res := &MigrateInfraResponse{}
	// Process

	// Call CB-Tumblebug API, which can be "/mcisDynamic"

	// Ouput

	// if err != nil {
	// 	common.CBLog.Error(err)
	// 	mapA := map[string]string{"message": err.Error()}
	// 	return c.JSON(http.StatusInternalServerError, &mapA)
	// }

	return c.JSON(http.StatusOK, res)

}
