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

Currently the user can submit operations of 3 kinds by sending HTTP requests to Hgate.

KYC creation request can be submitted in the following way:

```http
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

Asset update operation is submitted as shown:

```http
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

Sale update request subission is similar:

```http
PATCH http://hgateurl/sales/23
```

```json
{
   "details":{
      "description":"BLOB_ID",
      "logo":{
         "key":"...",
         "mime_type":"image/png",
         "name":"logo.png"
      },
      "name":"Test",
      "short_description":"Test token represents shares of investments in stocks and cryptocurrencies",
      "youtube_video_id":""
   }
}
```

The user can send another request, then Hgate works as a proxy to Horizon.

For example, the user can watch account creation or update entries in the ledger.

```GET http://hgateurl/ledger_changes```

This endpoint outputs multiple records. The user is interested to find out what records have `created` type in `changes` array and have `type` of entity inside of `created` equal to `account`.

The account can be updated by KYC creation request. In this case the record has `updated` type.

Here is what JSON output with account created and updated entries looks like:

```json

{
  "_links": {
    "self": {
      "href": "https://localhost:8844/ledger_changes?order=asc\u0026limit=10\u0026cursor="
    },
    "next": {
      "href": "https://localhost:8844/ledger_changes?order=asc\u0026limit=10\u0026cursor=103079219200"
    },
    "prev": {
      "href": "https://localhost:8844/ledger_changes?order=desc\u0026limit=10\u0026cursor=64424513536"
    }
  },
  "_embedded": {
    "records": [
      {
        "id": "64424513536",
        "paging_token": "64424513536",
        "ledger": 15,
        "created_at": "2018-09-07T10:28:16Z",
        "changes": [
          {
            "type_i": 0,
            "type": "created",
            "created": {
              "last_modified_ledger_seq": 15,
              "type_i": 0,
              "type": "account",
              "account": {
                "account_id": "GBQD4EV3PWJE55HX3BRK3KNLXQMA5YBJSBG55LNKY63FOZQ622N3DQUX",
                "account_type_i": 6,
                "account_type": "AccountTypeSyndicate",
                "block_reasons_i": 0,
                "block_reasons": [],
                "limits": null,
                "policies": {
                  "account_policies_type_i": 0,
                  "account_policies_types": null
                },
                "signers": [],
                "thresholds": {
                  "low_threshold": 0,
                  "med_threshold": 0,
                  "high_threshold": 0
                }
              },
              "asset": null,
              "balance": null
            },
            "updated": null,
            "removed": null,
            "state": null
          }
        ]
      },
      {
        "id": "9298604199936",
        "paging_token": "9298604199936",
        "ledger": 2165,
        "created_at": "2018-09-07T13:36:08Z",
        "changes": [
          {
            "type_i": 1,
            "type": "updated",
            "created": null,
            "updated": {
              "last_modified_ledger_seq": 2165,
              "type_i": 0,
              "type": "account",
              "account": {
                "account_id": "GCM2HLFJKEGYG3E3KXWL7V4HSGLN5EDKLQROLJPMOKUVB42K5XE7JM3G",
                "account_type_i": 2,
                "account_type": "AccountTypeGeneral",
                "block_reasons_i": 0,
                "block_reasons": [],
                "limits": null,
                "policies": {
                  "account_policies_type_i": 0,
                  "account_policies_types": null
                },
                "signers": [],
                "thresholds": {
                  "low_threshold": 0,
                  "med_threshold": 0,
                  "high_threshold": 0
                }
              },
              "asset": null,
              "balance": null
            },
            "removed": null,
            "state": null
          }
        ]
      },
    ]
  }
}

```

The KYC data of the account can be discovered by the following request.

```http

GET http://hgateurl/accounts/GC3JVXSDZ4NI4VEKPAWF45C6RAX3QNFSSHP6KABSXJZ57N6LXBHEQ7UB

```

```json

{
  "_links": {
    ...
  },
  "id": "GC3JVXSDZ4NI4VEKPAWF45C6RAX3QNFSSHP6KABSXJZ57N6LXBHEQ7UB",
  "account_id": "GC3JVXSDZ4NI4VEKPAWF45C6RAX3QNFSSHP6KABSXJZ57N6LXBHEQ7UB",
  "is_blocked": false,
  "block_reasons_i": 0,
  "block_reasons": [],
  "account_type_i": 5,
  "account_type": "AccountTypeNotVerified",
  "referrer": "",
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "balances": [
    {
      "balance_id": "BDRZCSIBEOHZ464BA25SGPVA5NZP6YY3FEOAUVDEQ4EAOTIRF4Q5QVJ6",
      "account_id": "GC3JVXSDZ4NI4VEKPAWF45C6RAX3QNFSSHP6KABSXJZ57N6LXBHEQ7UB",
      "asset": "USD666",
      "balance": "0.000000",
      "locked": "0.000000",
      "require_review": false
    }
  ],
  "signers": [
    {
      "public_key": "GBUXKQEYAEYGL5WUTWFLVJSJE5NJPRBSQFNKEPZIFOCSHJ7NAYXFOTKQ",
      "weight": 255,
      "signer_type_i": 1073741823,
      "signer_types": [
        ...
      ],
      "signer_identity": 0,
      "signer_name": ""
    },
    {
      "public_key": "GC3JVXSDZ4NI4VEKPAWF45C6RAX3QNFSSHP6KABSXJZ57N6LXBHEQ7UB",
      "weight": 1,
      "signer_type_i": 1073741823,
      "signer_types": [
        ...
      ],
      "signer_identity": 0,
      "signer_name": ""
    }
  ],
  "policies": {
    "account_policies_type_i": 0,
    "account_policies_types": null
  },
  "account_kyc": {
    "KYCData": {
      "v2": {
        "address": {
          "city": "City",
          "country": "Afghanistan",
          "line_1": "l1",
          "line_2": "l2",
          "postal_code": "123",
          "state": "CA"
        },
        "date_of_birth": "2018-04-01T00:00:00+03:00",
        "documents": {
          "kyc_id_document": {
            "front": {
              "key": "dmybeg4inhoebjhcost7fvmjfxxy35wb74dofgigzd45ifafagxm6o22"
            },
            "type": "passport"
          },
          "kyc_poa": {
            "front": {
              "key": "dmybef4inhoebjhcost7fvmyqilt54okeu45vqa7k3cbkutsba2jrdkk"
            }
          },
          "kyc_selfie": {
            "front": {
              "key": "dmybed4inhoebjhcost7fvifzza7z5rr5n3k5tg2ht7plazwkawfbwcu"
            }
          }
        },
        "first_name": "F",
        "id_document_type": "passport",
        "id_expiration_date": "2018-04-30T00:00:00+03:00",
        "last_name": "L"
      },
      "version": "v2"
    }
  },
  "external_system_accounts": [
    {
      "type": {
        "name": "Bitcoin",
        "value": 1
      },
      "data": "1M8LXH3RpbtDTughUYRtEsUNLCWBbQaXk2",
      "asset_code": "BTC"
    },
    {
      "type": {
        "name": "Ethereum",
        "value": 2
      },
      "data": "0x3A9D26C7721b2d6a7855F3935A650490F2fadbcb",
      "asset_code": "ETH"
    }
  ],
  "referrals": []
}

```