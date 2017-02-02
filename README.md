## CF App Tree Generator

This app simply generates a JSON document representing all orgs, spaces, and
applications in a Cloud Foundry foundation.

## How to get the application

Go to the [releases](http://github.com/ECSTeam/cf-apptree/releases) and download
the appropriate version for your operating system.

## How to build the application

1. Download the latest version of Go 1.7
1. Download the latest version of the source code
1. (Optional) Run the tests:

    ```shell
    $ ginkgo -r .
    testing: warning: no tests to run
    PASS

    Ginkgo ran 1 suite in 3.075729281s
    Test Suite Passed
    ```
1. Build the code:

   ```shell
   go build
   ```

## How to run the application

Run `cf-apptree` with the following command line arguments:

Short Name | Long Name | Environment Variable Option | Description
--- | --- | --- | ---
`-a` | `--api-endpoint` | `CF_TREE_API` | the API endpoint (`https://api.<RUN DOMAIN>`)
`-u` | `--user` | `CF_TREE_USER` | An admin user
`-p` | `--password` | `CF_TREE_PASSWORD` | The user's password
`-k` | `--skip-ssl-validation` | N/A | Skip SSL Validation
N/A | `--pretty-print` | `CF_TREE_PRETTY` | Pretty Print the JSON output

Environment variables will override flags on the command line.

### Example

```shell
$ ./cf-apptree -a https://api.system.mycf.com \
    -u admin \
	-p asdf1234 \
	-k \
	--pretty-print

{
  "orgs": [
    {
      "name": "system",
      "guid": "518aca77-cd79-4eec-a3ab-740511b14e75",
      "spaces": [
        {
          "name": "system",
          "guid": "beff4e8f-3858-4b75-b567-3d202b904fdd",
          "apps": [
            {
              "name": "php-demo",
              "guid": "2078aa6b-161b-4461-9dfb-bb9cd1d6db02"
            },
            {
              "name": "spring-music",
              "guid": "7a7ed7da-aa3a-4ae5-853d-cc06d99f9342"
            },
            {
              "name": "app-usage-server-venerable",
              "guid": "4dca2af7-7b99-463d-a8eb-647f6ddd9f09"
            },
            {
              "name": "p-invitations-venerable",
              "guid": "0002f265-697b-4d18-9d51-166eaed621d1"
            },
            {
              "name": "apps-manager-js-venerable",
              "guid": "0ca9890e-5922-435e-812e-a743ab6752ce"
            },
            {
              "name": "app-usage-scheduler-venerable",
              "guid": "ceac4434-f597-4360-bafb-ef3e3f8925e6"
            },
            {
              "name": "app-usage-worker-venerable",
              "guid": "ab06af8c-5f86-4d0a-b820-554fa456f463"
            },
            {
              "name": "p-invitations",
              "guid": "49fa2413-91f3-4a63-b30a-f4e71e4a8e7c"
            },
            {
              "name": "app-usage-server",
              "guid": "db2e544c-7ed5-4187-97a9-a991eb39c7a2"
            },
            {
              "name": "apps-manager-js",
              "guid": "8af5b7d5-ddd0-493a-a8c0-f40f31acdc89"
            },
            {
              "name": "app-usage-scheduler",
              "guid": "8933b109-b1ce-4c8b-9a35-1d4cee0a744d"
            },
            {
              "name": "app-usage-worker",
              "guid": "1a446451-e098-4860-a175-d4cbc8ffe317"
            }
          ]
        },
        {
          "name": "notifications-with-ui",
          "guid": "aa9fe0bb-77c1-4693-8f24-9bd2a362c897",
          "apps": [
            {
              "name": "notifications-ui",
              "guid": "425db0f1-927b-4335-ba1b-7b5fdcebd874"
            }
          ]
        },
        {
          "name": "pivotal-account-space",
          "guid": "3d365000-4fd0-40c3-abaa-ef9495fd537c",
          "apps": [
            {
              "name": "pivotal-account-cold",
              "guid": "5e49373c-d9d1-4dba-b293-84dc636708fb"
            },
            {
              "name": "pivotal-account",
              "guid": "5ea55f12-28b9-4458-81af-09345b1ad193"
            }
          ]
        },
        {
          "name": "autoscaling",
          "guid": "98545d4a-d1b1-4a83-9b73-34378ec3eb34",
          "apps": [
            {
              "name": "autoscale",
              "guid": "e917066b-281c-4f71-ad95-9ab6db0fc871"
            }
          ]
        },
        {
          "name": "p-spring-cloud-services",
          "guid": "cd71e6e7-8e5a-4587-83a8-409923a663a8",
          "apps": [
            {
              "name": "spring-cloud-broker",
              "guid": "767f1a64-ae5a-4962-bfd5-7d1822147505"
            },
            {
              "name": "spring-cloud-broker-worker",
              "guid": "403f7afb-a4ab-421b-8e58-1b79256274b8"
            }
          ]
        },
        {
          "name": "scs-smoke-tests",
          "guid": "0d529602-8196-441e-83b3-17e994c6cc38",
          "apps": []
        },
        {
          "name": "azure-service-broker-space",
          "guid": "88497612-46ca-4be9-a774-0e2a1a02e415",
          "apps": [
            {
              "name": "azure-service-broker-1.2.0",
              "guid": "382cf576-fd00-43de-85bb-fefb56589bde"
            }
          ]
        }
      ]
    },
    {
      "name": "redis-smoke-test-org",
      "guid": "e506b030-9762-4ca4-bcac-530431480234",
      "spaces": [
        {
          "name": "redis-smoke-test-space",
          "guid": "61a9adb6-508d-49c3-8edf-89425692da58",
          "apps": []
        }
      ]
    },
    {
      "name": "apps1",
      "guid": "b48bfdfc-4407-4329-9443-b3d66fa22a46",
      "spaces": [
        {
          "name": "space1",
          "guid": "c1ca20e6-40cb-4a5e-a254-1176518a20fd",
          "apps": [
            {
              "name": "spring-music/",
              "guid": "c97dfd92-4e0f-4d38-aee0-d732ab9f726a"
            },
            {
              "name": "spring-joker/",
              "guid": "588a8b70-14a3-4911-b051-05223dd64b2b"
            },
            {
              "name": "Downloads/spring-music.war",
              "guid": "41a965a2-119b-4271-be74-6def2736025d"
            },
            {
              "name": "sp-music",
              "guid": "1f675fd7-6aad-4d70-bd3c-8e4e1be6334d"
            },
            {
              "name": "simple-go-web-app",
              "guid": "b2ebd07c-0347-4a39-9d8f-102797963928"
            }
          ]
        }
      ]
    },
    {
      "name": "p-spring-cloud-services",
      "guid": "08aa6d22-b50f-429b-b13a-578ac7f8e2cb",
      "spaces": [
        {
          "name": "instances",
          "guid": "19597389-20ed-4127-972c-4c8d97253fb2",
          "apps": []
        }
      ]
    },
    {
      "name": "ecsteam",
      "guid": "88195490-843f-4342-8e24-bf9a7bc690f4",
      "spaces": [
        {
          "name": "development",
          "guid": "3b6185ca-9e4d-4802-a440-a3d4ba8725ba",
          "apps": [
            {
              "name": "spring-music",
              "guid": "679a803f-893e-4eb9-9cd8-f133c01e2618"
            },
            {
              "name": "willitconnect",
              "guid": "1b57d04a-4dcd-4470-95fe-5540d46a25d2"
            },
            {
              "name": "helloworld-jdg-cfdemo",
              "guid": "1a54f660-62f7-49a8-b179-cf02317352b5"
            }
          ]
        },
        {
          "name": "plugintest",
          "guid": "d33a8135-5e43-443d-a724-c5b7a677f55d",
          "apps": [
            {
              "name": "plugin-test-demo",
              "guid": "cf666462-a8b1-4497-902f-0a5b3998e03a"
            },
            {
              "name": "env-test",
              "guid": "557ac05b-3d9c-4364-9308-156e36e59928"
            }
          ]
        }
      ]
    },
    {
      "name": "spinnaker",
      "guid": "e935d950-2cdf-4941-8614-58d40b4739e3",
      "spaces": [
        {
          "name": "dev",
          "guid": "26d0b403-fe7a-4f09-9cb2-cdcc84e70363",
          "apps": [
            {
              "name": "spinnaker_app",
              "guid": "f4c0c76a-52a9-4087-a3fc-e541deee42a6"
            },
            {
              "name": "deck",
              "guid": "fcc15f20-9749-4872-94bc-3c53f354a8d6"
            },
            {
              "name": "clouddriver",
              "guid": "acc9c79e-39f5-436b-85b5-605031c25457"
            },
            {
              "name": "echo",
              "guid": "0dd4721e-1c5f-4d30-b997-aff9863d0988"
            }
          ]
        },
        {
          "name": "staging",
          "guid": "d8399f08-e292-4da0-9936-43d68156e209",
          "apps": []
        },
        {
          "name": "prod",
          "guid": "e9bbbcde-bd2f-4995-9d24-dc80fcdc61a2",
          "apps": []
        }
      ]
    },
    {
      "name": "sandbox",
      "guid": "0d2979a9-ec8d-4785-9066-333e2e3fb8fc",
      "spaces": [
        {
          "name": "sandbox",
          "guid": "2f2b6eab-293a-49d8-a4c1-37fb8fde19f4",
          "apps": [
            {
              "name": "sp-music01",
              "guid": "2b28888b-f6dc-4e34-ba0d-f2747e4d3ecf"
            },
            {
              "name": "hello-goodbye",
              "guid": "bdf346b5-ad7d-4b72-b85c-4b01e4118de1"
            },
            {
              "name": "mvc-demo02",
              "guid": "e0d87561-283c-4736-9f0a-edf4d66aeddb"
            },
            {
              "name": "helloworld-no-mvc",
              "guid": "4964d83a-71e5-4333-9968-3e067e628d02"
            },
            {
              "name": "helloworld-with-mvc",
              "guid": "a24ebffb-607e-4634-ad5a-dacc04df76d8"
            }
          ]
        }
      ]
    }
  ]
}
```
