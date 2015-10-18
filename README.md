# Slingshot

Fast and simple uploading of static assets to Amazon S3

## Description

Slingshot is CLI tool for upload JS, CSS, PNG's and all the normal types of static assets to S3.  This is usually used in conjunction with Akamai for serving to the public.

Slingshot can also list the static assets in a bucket if you need to know what is the latest version.

Slingshot will attempt to guess the MIME type of your files and set those automatically.  Cache-Control is set to 1 year by default.  

## Usage

### Current
Options
```sh
-r region  
-b bucket path  
-p S3 folder path  
```
Example
```sh
slingshot current -r eu-west-1 -b statics -p assets

assets/
assets/56/
assets/57/
```

### Upload
Options
```sh
-r region  
-b bucket path  
-p S3 folder path  
-d local directory  
```
Example
```
slingshot upload -r eu-west-1 -b statics -p assets -d vendor

Uploading file:vendor/current.go
Uploading file:vendor/current_test.go
Uploading file:vendor/files/test.txt
Uploading file:vendor/upload.go
Uploading file:vendor/upload_test.go
Upload of directory vendor complete
```

This will upload assets resulting in the folling URL https://s3-eu-west-1.amazonaws.com/statics/assets/all.js

## Install

RPM is still WIP, this should be installed onto your Jenkins agents.


## Development

To install, use `go get`:

```bash
$ go get -d github.com/DaveBlooman/slingshot
```

## Contribution

1. Fork ([https://github.com/DaveBlooman/slingshot/fork](https://github.com/DaveBlooman/slingshot/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[DaveBlooman](https://github.com/DaveBlooman)
