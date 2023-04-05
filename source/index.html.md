---
title: API Reference

language_tabs: # must be one of https://github.com/rouge-ruby/rouge/wiki/List-of-supported-languages-and-lexers
  - shell
  - go
  - python

toc_footers:
  - <a href='https://hemato.ai'>Hemato.AI</a>
  - <a href='mailto:ai@hemato.ai'>Sign Up for a Developer Key</a>
  - <a href='https://github.com/slatedocs/slate'>Documentation Powered by Slate</a>

includes:
  - errors

search: true

code_clipboard: true

meta:
  - name: description
    content: Documentation for Hemato.AI Diagnostic API
---

# Introduction

Welcome to the Hemato.AI Diagnostic API! You can use our API to access Hemato.AI's Diagnostic AI and Clinical Algorithms, which can get process images of peripheral blood slides and generate hematopathology reports.

We have language bindings in Go, ( also Python, C#, and TypeScript coming in future)! You can view code examples in the dark area to the right, and you can switch the programming language of the examples with the tabs in the top right.

# HTTP Responses

Most Hemato.AI API will return a JSON encoded response in this format.

```json
{
	"status": 401,
	"error": "some error message",
	"results": {

	},
	"user_message": "",
	"debug_info": {

	}
}

```


Field  | Description
-------|------------
status | This reflects the [HTTP Status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
error | If something has gone wrong, this provides an error message possibly with error codes or more context on what has happened. almost always this comes along side a 4xx or 5xx status code
results | This is the output of the call, and can be of any type, often a dictionary / object.
user_message | This is a message that can be displayed to the end users, often useful in cases there is an error
debug_info | Contains information helpful when debugging




# Authentication

All authenticated calls to Hemato.AI API are expected to present an `Authorization` header containing a signed and not expired token.
To obtain an authorization token you have 3 options.

1. Demo Authentication: this is only useful for development and testing purposes.
2. Session Based Authentication: you can provide a user name and password to obtain a newly minted and signed token.
3. Public-Private Key: if you have a public key registered with Hemato.AI you can sign your own tokens. This is particularly helpful in situations where a user is not directly involved and a device or an automated system needs to make calls to Hemato.AI API. To learn more read [How to register your public keys with hemato](#Register-Public-Key)

## Demo Auth

> To authorize, using Demo Authentication this code:


```shell
# use httpie from https://httpie.io/
http https://api.hemato.ai/login/demo
# or use curl, you can just pass the correct header with each request
curl "https://api.hemato.ai/login/demo"
```

```go
// Not implemented yet
```

> The above command returns a JSON structured like this:

```json
{
  "results":{
    "hemato_ai_auth_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```

## Session Auth

> To authorize, using Session Authentication this code:


```shell
# SECURITY RISK:
# Please ONLY use this for debugging login problems.
# MAKE SURE TO REMOVE THIS FROM YOUR SHELL HISTORY.
# otherwise your plan password will remain in your shell history and possibly leak into backups and such.
export PASS_SHA_256=`echo -n YOUR_ACTUAL_PASSWORD | shasum -a 256 | cut -d ' ' -f 1`
# use httpie from https://httpie.io/
echo '{"user":"ali@example.com","pass_hash":"'${PASS_SHA_256}'"}' | http -F POST https://api.hemato.ai/auth/login
# or use curl, you can just pass the correct header with each request
echo '{"user":"ali@example.com","pass_hash":"'${PASS_SHA_256}'"}' | curl -X POST "https://api.hemato.ai/auth/login"
```

```go
// Not implemented yet
```

> The above command returns a JSON structured like this:

```json
{
  "results":{
    "hemato_ai_auth_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```

## Public / Priavte Key Auth

Detailed Instructions to be added here.



Hemato.AI expects for the API key to be included in majority of API requests to the server in a header that looks like the following:

`Authorization:HEMATO_AI_AUTH_TOKEN`

<aside class="notice">
You must replace <code>HEMATO_AI_AUTH_TOKEN</code> with your personal API key.
</aside>

# Peripheral Blood Smear

## How does it work?

Images of Peripheral Blood Smears (PBS) usually are very large. It is often impractical to submit an entire sample in one API call.
That's why there are a few steps to submit a PBS for review by Hemato.AI

To get Hemato.AI's opinion on a Peripheral Blood Smear you follow this flow:
1. Get an ID
2. Upload the files
3. Request a Diagnostic Study
4. Wait for it
5. Ask for the report

## 1. Get an ID
Get a new PBS ID. This ID will be used to upload the images, request diagnostic study of the sample, provide any other information available and afterwards get the results of Hemato.AI's review.

### HTTP Request
Make a `POST` call to the `/pbs` endpoint

```shell
# use httpie from https://httpie.io/
http -f POST https://api.hemato.ai/pbs Authorization:HEMATO_AI_AUTH_TOKEN
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://api.hemato.ai/pbs
```
> Make sure to replace `HEMATO_AI_AUTH_TOKEN` with your API key.

> The above command returns a JSON structured like this:

```json
{
  "results":{
    "pbs_id":"c7d453a9-676d-4b9f-a505-8a9021b76dfd"
  }
}
```

## 2. Upload the files
Upload the sample by calling the PBS upload API as many times as needed.

### HTTP Request
Make a `POST` call to the `/pbs/YOUR_NEW_PBS_ID/files` endpoint. Here `YOUR_NEW_PBS_ID` is the id you obtained from step 1 (`results.pbs_id` in the response structure returned).

This is an idempotent call. It means you can send the same file multiple times and it will only be counted as one file.

The file type is determined by the `Content-Type` header of your request. So when sending a requet to upload a jpg file for example, your request needs to have a header `Content-Type:image/jpeg`.
This endpoint accepts:

File Type | Request Content-Type Header
--------- | -----------
jpeg , jpg | image/jpeg
png | image/png
zip ★ | application/zip

★ Zip files can containing jpg or png files.


```shell
# use httpie from https://httpie.io/
http -f POST https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files Authorization:HEMATO_AI_AUTH_TOKEN Content-Type:image/jpeg  < /path/to/filename.jpg
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --header "Content-Type:image/jpeg" --data-binary "@/path/to/filename.jpg"  https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files
```

> The above command returns a JSON structured like this:

```json
{
  "status": 201,
  "results":{
    "pbs_id":"HEMATO_AI_AUTH_TOKEN",
    "file_id":"296bca9e-4325-4da9-9e95-d0c84bdf1b97"
  },
  "debug_info":{

  }
}
```

<aside class="warning">Notes:<ul>
<li>Please note that this is NOT an HTTP multipart call with `Content-Type: multipart/form-data` </li>
<li>Plaese note that this is not a form submission call either. The entir body of the request is
</ul>
</aside>


## 3. Request a Diagnostic Study
When all the files for a particular PBS is uploaded, you will make a call to mark the PBS as ready to be processed. After this call you will not be able to upload additional files.
Here you also need to specifiy what diagnostic stydies you are interested in.

```shell
# use httpie from https://httpie.io/
http -f POST https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files Authorization:HEMATO_AI_AUTH_TOKEN Content-Type:image/jpeg  < /path/to/filename.jpg
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --header "Content-Type:image/jpeg" --data-binary "@/path/to/filename.jpg"  https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files
```

> The above command returns a JSON structured like this:

```json
{
  "results":{
    "pbs_id":"HEMATO_AI_AUTH_TOKEN",
    "file_id":"296bca9e-4325-4da9-9e95-d0c84bdf1b97"
  }
}
```

## 4. Wait for it
Depending on the size of the files submitted and the number Diagnostic tasks requested and depending other factors, like general system workload. It will take between a few seconds to 10s of seconds for the results to be ready.

You can check the status of the task by making a GET call to /status

> To get the status of a PBS

```shell
#
http https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/status Authorization:HEMATO_AI_AUTH_TOKEN
#
curl --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/status
```

> This returns
```json
```


## 5. Get The Report for a Peripheral Blood Seamr Study

