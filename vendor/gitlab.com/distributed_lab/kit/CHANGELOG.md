# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 1.6.0

### Changed
* vendor dependencies

## 1.5.0

### Removed
* graylog

### Added 
* traefik 

## v1.4.0

- make page params compatible with `urlval` package

## v1.3.0
### Added
- ValidateLazyDep - allows to ensure that none of the methods panics

## v1.2.0

### Added

* JSONScan and JSONValue methods 

## v.1.0.0

### Changed

* Switched to normal versions


## v0.12.0

### Fixed

* DB performance issue. Pass MaxIdleConnections, MaxOpenConnections to the sql.DB


## v0.11.0

### Added

* Cursor-based page params

## v0.10.1

### Fixed

* Typo

## v0.10.0

### Added

* RawDB to pgdb comfig helper

## v0.9.1

### Added

* offset-based page params

## v0.9.0

### Added

* support for `graylog` by `comfig.Log`

### Removed

* `comfig.Databaser` in favor of `pgdb.Databaser`
* `comfig.NewDatabaser` in favor of `pgdb.NewDatabaser`
* `comfig.Januser` in favor of `janus.Januser`
* `config.NewJanuser` in favor of `janus.NewJanuser`
