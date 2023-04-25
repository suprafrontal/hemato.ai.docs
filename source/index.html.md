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

Welcome to the Hemato.AI Diagnostic API! You can use our API to access Hemato.AI's Diagnostic AI and Clinical Algorithms, which processes images of peripheral blood slides and generates hematopathology reports.

We have language bindings in Go, (also Python, C#, and TypeScript coming in future)! You can view code examples in the dark area to the right, and you can switch the programming language of the examples with the tabs in the top right.

# Endpoints
Hemato.AI API is made available at specific regional endpoints to adhere to regional regulations. Any data submitted to a regional endpoint will be handled on servers located within the same geographical and regulatory region.

You can explicitly choose what region to use from this list:

Region | Endpoint
-------|--------------------------
USA    | https://us.api.hemato.ai
Canada | https://ca.api.hemato.ai

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
error | If something has gone wrong, this provides an error message with error codes or more context on what has happened.
results | This is the output of the call, and can be of any type, often a dictionary / object.
user_message | This is a message that can be displayed to the end users. This is often useful in cases where there is an error.
debug_info | Contains helpful information for debugging




# Authentication

All authenticated calls to Hemato.AI API are expected to present an `Authorization` header containing a signed and not expired token.
To obtain an authorization token you have 3 options:

1. Demo Authentication: for development and testing purposes.
2. Session Based Authentication: you can provide a user name and password to obtain a newly minted and signed token.
3. Public-Private Key: if you have a public key registered with Hemato.AI you can sign your own tokens. This is particularly helpful in situations where a user is not directly involved and a device or an automated system needs to make calls to Hemato.AI API. To learn more read [How to register your public keys with hemato](#Register-Public-Key)

## Demo Auth

> To authorize, using Demo Authentication:


```shell
# use httpie from https://httpie.io/
http https://api.hemato.ai/login/demo
# or use curl, you can just pass the correct header with each request
curl "https://api.hemato.ai/login/demo"
```

```go
// Not implemented yet
```

> The above command returns a JSON structure like this:

```json
{
  "results":{
    "hemato_ai_auth_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```

## Session Auth

> To authorize, using Session Authentication:


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

> The above command returns a JSON structure like this:

```json
{
  "results":{
    "hemato_ai_auth_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```

## Public / Private Key Auth

WIP: (Detailed instructions to be added)



Hemato.AI expects for the API key to be included in majority of API requests to the server in a header that looks like the following:

`Authorization:HEMATO_AI_AUTH_TOKEN`

<aside class="notice">
You must replace <code>HEMATO_AI_AUTH_TOKEN</code> with your personal API key.
</aside>

# Peripheral Blood Smear

## How does it work?

Images of Peripheral Blood Smears (PBS) usually are very large. It is often impractical to submit an entire sample in one API call.
That's why there are a few steps to submit a PBS for review by Hemato.AI.

To get Hemato.AI's opinion on a Peripheral Blood Smear you follow this flow:

1. Get an ID for your new PBS
2. Upload the files under the new ID
3. Request a Diagnostic Study
4. Wait for it
5. Ask for the report

![Peripheral Blood Study Flow](PBSStudyFlow.png)


## 1. Get an ID
Get a new PBS ID. This ID will be used to upload the images, request diagnostic study of the sample, provide any other information available and afterwards get the results of Hemato.AI's review.

### HTTP Request
Make a `POST` call to the `/pbs` endpoint. You can provide any number of tags (key value string pairs) along side this request. These tags can allow you to associate Patient Proxy Identifiers or Organization IDs or any number of other information that is important to you with the PBS. These tags can later be used to find and retrieve PBSs easier.
<aside class="warning">Please do NOT include any Personanlly Identifiable Information (PII) or Personal Health Records (PHI) in the tags</aside>


```shell
# use httpie from https://httpie.io/
echo '{"tags":{"patient_proxy_id":"}}' | http -f POST https://api.hemato.ai/pbs Authorization:HEMATO_AI_AUTH_TOKEN
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://api.hemato.ai/pbs
```
> Make sure to replace `HEMATO_AI_AUTH_TOKEN` with your API key.

> The above command returns a JSON structure like this:

```json
{
  "status": 201,
  "results":{
    "pbs_study_id":"c7d453a9-676d-4b9f-a505-8a9021b76dfd"
  }
}
```

## 2. Upload the files
Upload the sample by calling the PBS upload API as many times as needed.

### HTTP Request
Make a `POST` call to the `/pbs/YOUR_NEW_PBS_ID/files?file_name=some_file_name.jpg` endpoint. Here `YOUR_NEW_PBS_ID` is the id you obtained from step 1 (`results.pbs_id` in the response structure returned).

This api accepts an optional "file_name" and stors this along side the file. This can be later used for audit purposes.

This is an idempotent call. It means you can send the same file multiple times and it will only be counted as one file.
Once submitted however, you cannot update the same file or change it or its metadata. You can only add more  files to the same PBS study.

The file type is determined by the `Content-Type` header of your request. So when sending a request to upload a jpg file for example, your request needs to have a header `Content-Type:image/jpeg`.
This endpoint accepts:

File Type | Request Content-Type Header
--------- | -----------
jpeg , jpg | image/jpeg
png | image/png
zip ★ | application/zip

★ Zip files can contain jpg or png files.

If the Content-Type header is missing, the file name extension for "file_name" provided is used to determine the file type.
If no file type can be determined or provided, a 400 - Bad Request error is returned.


```shell
# use httpie from https://httpie.io/
http -f POST https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files Authorization:HEMATO_AI_AUTH_TOKEN Content-Type:image/jpeg  < /path/to/filename.jpg
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --header "Content-Type:image/jpeg" --data-binary "@/path/to/filename.jpg"  https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files
```

> The above command returns a JSON structure like this:

```json
{
  "status": 201,
  "results":{
    "pbs_study_id":"c7d453a9-676d-4b9f-a505-8a9021b76dfd",
    "file_id":"sha224-i-4d58b0cee1e38f75f93e2e605d1e6ad44c24a975d0c3b53367de7c86"
  },
  "debug_info":{

  }
}
```

<aside class="warning">Notes:<ul>
<li>Please note that this is NOT an HTTP multipart call with `Content-Type: multipart/form-data` </li>
<li>Plaese note that this is not a form submission call either. The entire body of the request is the image.
</ul>
</aside>


## 3. Request a Diagnostic Study
When all the files for a particular PBS is uploaded, you will make a call to mark the PBS as ready to be processed and ask for any number of diagnostics tasks to be performed on this PBS. After this call you will not be able to upload additional files for this PBS.


```shell
# use httpie from https://httpie.io/
echo '{"diagnostic_tasks":["pbs_v1"]}' | http -f POST https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/tasks Authorization:HEMATO_AI_AUTH_TOKEN
# alternatively use curl
echo '{"diagnostic_tasks":["pbs_v1"]}' | curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --data-binary @- https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/tasks
```

> The above command returns a JSON structure like this:

```json
{
  "results":{
    "pbs_id":"HEMATO_AI_AUTH_TOKEN",
    "file_id":"296bca9e-4325-4da9-9e95-d0c84bdf1b97"
  }
}
```

## 4. Wait for it
Depending on the size of the files submitted, the number Diagnostic tasks requested, and other factors like general system workload, it will take between a few seconds to 10s of seconds for the results to be ready.

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
{
  "status": 200,
  "results": {
    "report_status": "ready",
    "progress": 1.0,
    "progress_message":"Reports are ready",
  },
  "debug_info" :{}
}
```


## 5. Get The Report for a Peripheral Blood Smear Study
You can get the latest version of the report by making a GET call.
In cases that there are multiple revisions and you want to get an earlier version of the report you need to specify the version.


> To get the latest diagnostic report on a PBS

```shell
#
http https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/report Authorization:HEMATO_AI_AUTH_TOKEN
#
curl --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://api.hemato.ai/pbs/YOUR_NEW_PBS_ID/report
```

> Which in turn returns a JSON structure like:

```json
{
	"status": 200,
	"results": {
		"pbs_study_id": "4bb7fe9e-b608-4d68-adbe-8655c991f494",
		"file_id": "sha224-i-db367ccd89fc39aa6ff9fff0cd9b11e6b5b6a41cff4bbb232bee7c93",
		"feature_statistics": {
			"rbc-density": {
				"value": "x",
				"unit": "y"
			},
			"platelet-density": {
				"value": "x",
				"unit": "y"
			},
			"blast-density": {
				"value": "x",
				"unit": "y"
			}
		},
		"morphological_findings": {
			"ring trophozoite": {
				"count": "1",
				"heat_map": [{
					"center": [435, 345],
					"radius": 35
				}]
			}
		},
		"diagnostic_report": {
			"differential_diagnosis": [{
				"score": 7,
				"confidence": "0.95",
				"clinical-significance": "D",
				"icd10-code": "B50",
				"english-long": "Malaria Infection with Plasmodium Falciparum",
				"english-short": "Malaria"
			}]
		},
		"pathology_report": "No demographics or past medical history is available about the patient. Inspection of peripheral blood slide revealed ring trophozoites ..."
	},
	"debug_info": {}
}
```

There are 3 levels of report available.

1. Morphological Findings
2. Diagnostic Information
3. Pathology Report

# Health Check
If you need to check the health status of the Hemato.AI api you can make a GET call to the heartbeat endpoint.

```shell
http https://api.hemato.ai/heartbeat
```

```json
{
    "delta": "59.078µs",
    "heartbeat": "a7ef8fba741237f693c3",
    "request_timestamp": 1680716829,
    "version": "bb_api.331.develop.18ff645"
}
```
