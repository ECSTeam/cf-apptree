/**
 * Copyright 2017 ECS Team, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-cfclient"
)

type identifiable struct {
	Name string `json:"name"`
	GUID string `json:"guid"`
}

// Application - represents an app in the foundation
type Application struct {
	identifiable
}

// Space - a space in the foundation. Contains applications
type Space struct {
	identifiable
	Apps []Application `json:"apps"`
}

// Org - an organization in the foundation. Contains spaces
type Org struct {
	identifiable
	Spaces []Space `json:"spaces"`
}

// Foundation - represents all the orgs in a CF installation
type Foundation struct {
	Orgs []Org `json:"orgs"`
}

// ListFoundation - Get all orgs, spaces, and apps
func ListFoundation(cfClient *cfclient.Client) Foundation {
	orgs, err := cfClient.ListOrgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching orgs: %v\n", err)
		os.Exit(1)
	}

	treeOrgs := make([]Org, 0, len(orgs))
	for _, cforg := range orgs {
		treeOrg := Org{
			identifiable: identifiable{
				Name: cforg.Name,
				GUID: cforg.Guid,
			},
		}

		cfspaces, getSpaceErr := cfClient.OrgSpaces(treeOrg.GUID)
		if getSpaceErr != nil {
			fmt.Fprintf(os.Stderr, "Error fetching spaces for org %s: %v\n", treeOrg.Name, getSpaceErr)
			os.Exit(1)
		}

		orgSpaces := make([]Space, 0, len(cfspaces))
		for _, cfspace := range cfspaces {
			treeSpace := Space{
				identifiable: identifiable{
					Name: cfspace.Name,
					GUID: cfspace.Guid,
				},
			}

			cfapps, getAppErr := listAppsBySpace(cfClient, cfspace.Guid)
			if getAppErr != nil {
				fmt.Fprintf(os.Stderr, "Error fetching spaces for space %s: %v\n", treeSpace.Name, getAppErr)
				os.Exit(1)
			}

			spaceApps := make([]Application, 0, len(cfapps))
			for _, cfapp := range cfapps {
				spaceApp := Application{
					identifiable{
						Name: cfapp.Name,
						GUID: cfapp.Guid,
					},
				}

				spaceApps = append(spaceApps, spaceApp)
			}

			treeSpace.Apps = spaceApps
			orgSpaces = append(orgSpaces, treeSpace)
		}

		treeOrg.Spaces = orgSpaces
		treeOrgs = append(treeOrgs, treeOrg)
	}

	return Foundation{
		Orgs: treeOrgs,
	}
}

// adapted from cfclient.ListApps
func listAppsBySpace(c *cfclient.Client, spaceGUID string) ([]cfclient.App, error) {
	var apps []cfclient.App

	requestURL := "/v2/apps?inline-relations-depth=2&q=space_guid:" + spaceGUID
	for {
		var appResp cfclient.AppResponse
		r := c.NewRequest("GET", requestURL)
		resp, err := c.DoRequest(r)

		if err != nil {
			return nil, fmt.Errorf("Error requesting apps %v", err)
		}
		resBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading app request %v", resBody)
		}

		err = json.Unmarshal(resBody, &appResp)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshaling app %v", err)
		}

		for _, app := range appResp.Resources {
			app.Entity.Guid = app.Meta.Guid
			app.Entity.SpaceData.Entity.Guid = app.Entity.SpaceData.Meta.Guid
			app.Entity.SpaceData.Entity.OrgData.Entity.Guid = app.Entity.SpaceData.Entity.OrgData.Meta.Guid
			apps = append(apps, app.Entity)
		}

		requestURL = appResp.NextUrl
		if requestURL == "" {
			break
		}
	}

	return apps, nil
}
