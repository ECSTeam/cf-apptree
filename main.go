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
	"os"

	"github.com/cloudfoundry-community/go-cfclient"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apiEndpoint       = kingpin.Flag("api-endpoint", "Address of Cloud Controller API [$CF_TREE_API]").Short(byte('a')).OverrideDefaultFromEnvar("CF_TREE_API").Required().String()
	user              = kingpin.Flag("user", "Name of user with cloud_controller.admin or cloud_controller.admin_read_only scope [$CF_TREE_USER]").Short(byte('u')).OverrideDefaultFromEnvar("CF_TREE_USER").Required().String()
	password          = kingpin.Flag("password", "Password of user with cloud_controller.admin or cloud_controller.admin_read_only scope [$CF_TREE_PASSWORD]").Short(byte('p')).OverrideDefaultFromEnvar("CF_TREE_PASSWORD").Required().String()
	skipSSLValidation = kingpin.Flag("skip-ssl-validation", "Please don't").Default("false").Short(byte('k')).Bool()
	prettyPrint       = kingpin.Flag("pretty-print", "Output the JSON in a slightly more human-readable format [$CF_TREE_PRETTY]").OverrideDefaultFromEnvar("CF_TREE_PRETTY").Default("false").Bool()

	version = "0.0.0"
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	c := cfclient.Config{
		ApiAddress:        *apiEndpoint,
		Username:          *user,
		Password:          *password,
		SkipSslValidation: *skipSSLValidation,
	}

	cfClient, err := cfclient.NewClient(&c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging into Cloud Foundry: %v\n", err)
		os.Exit(1)
	}

	appTree := ListFoundation(cfClient)

	var resp []byte
	var jsonError error

	if *prettyPrint {
		resp, jsonError = json.MarshalIndent(appTree, "", "  ")
	} else {
		resp, jsonError = json.Marshal(appTree)
	}

	if jsonError != nil {
		fmt.Fprintf(os.Stderr, "Error outputting json! %v\n", err)
	}

	os.Stdout.Write(resp)
}
