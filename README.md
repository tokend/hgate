# Hgate User Documentation

## Hgate Description

The purpose of hgate is to send signed requests to Horizon.

Hgate uses keypair provided in config file to sign and submit transactions into TokenD.

Alternatively, you can provide headers along with your request to sign transaction with specified keys:

* **Tokend-Source** - source of the transaction, account address.
* **Tokend-Signers** - slice of signer secret keys (seed) to sign transaction with.

## Usage

For detailed description of service API see Swagger specification [here](docs).

