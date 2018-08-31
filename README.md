# Hgate User Documentation


## Hgate Description

The purpose of hgate is to send signed requests to Horizon.

## Setup

Assuming the user has working Golang-setup, the installation process is as simple as:

```sh
go get gitlab.com/tokend/hgate
go install gitlab.com/tokend/hgate/...
```

Before starting the Hgate, the user should provide a configuration file in order to make Hgate work properly. Example configuration file is called `config.yaml` and can be found in the root directory of Hgate repo.

Assuming the user has created a config called `local-config.yaml`, he can launch the Hgate server in the following way:

```bash
./hgate --config="local-config.yaml"
```

If `config` argument is not provided Hgate looks for `config.yaml` file by default.

## Usage

Currently the user can submit 2 kind of operations by sending HTTP requests to Hgate. If the user sends another request, then Hgate works as a proxy to Horizon.

```
POST http://hgateurl/create_kyc_request
```
```json
{
	"request_id": 0,
	"account_to_update_kyc": "G...",
	"account_type_to_set": "general",
	"kyc_level_to_set": 0,
	"kyc_data": "{\"version\":\"v2\",\"v2\":{\"id_document_type\":\"passport\",\"documents\":{\"kyc_poa\":{\"front\":{\"key\":\"...\"}},\"kyc_selfie\":{\"front\":{\"key\":\"...\"}},\"kyc_id_document\":{\"front\":{\"key\":\"...\"},\"type\":\"passport\"}},\"first_name\":\"John\",\"last_name\":\"Doe\",\"address\":{\"line_1\":\"l1\",\"line_2\":\"l2\",\"city\":\"New York\",\"country\":\"USA\",\"state\":\"CA\",\"postal_code\":\"123\"},\"date_of_birth\":\"2018-04-01T00:00:00+03:00\",\"id_expiration_date\":\"2018-04-30T00:00:00+03:00\"}}\r\n",
	"all_tasks": 2
}
```

```
PATCH http://hgateurl/assets/USD
```
```json
{
		"policies": 1,
		"details":  {
			"external_system_type": "12",
			"name": ""
		}
}
```

```GET http://hgateurl/assets/```

```GET http://hgateurl/accounts/```