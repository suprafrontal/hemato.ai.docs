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

All Hemato.AI API will return a JSON encoded response in this format.

Field  | Description
-------|------------
status | This reflects the [HTTP Status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
error | If something has gone wrong, this provides an error message possibly with error codes or more context on what has happened. almost always this comes along side a 4xx or 5xx status code
results | This is the output of the call, and can be of any type, often a dictionary / object.
user_message | This is a message that can be displayed to the end users, often useful in cases there is an error
debug_info | Contains information helpful when debugging


```json
{
	"status": 201, // this reflects the HTTP Status code https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
  "error": "some error message",  // if something has gone wrong, this provides an error message possibly with error codes or more context on what has happened.
	"results": {
    // this is the output of the call, and can be of any type, often a dictionary / object.
  },
	"user_message": "", // this is a message that can be displayed to the end users, often useful in cases there is an error

	"debug_info": {

  }
}

```


# Authentication

There are 3 authentication modes available:

1. Demo Authentication
2. Session Based Authentication
3. Public-Private Key


> To authorize, using Demo Authentication this code:


```shell
# use httpie from https://httpie.io/
http https://api.hemato.ai/login/demo
# or use curl, you can just pass the correct header with each request
curl "https://api.hemato.ai/login/demo"
```

> The above command returns a JSON structured like this:

```json
{
  "results":{
    "hemato_ai_auth_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
  }
}
```


```go
// Not implemented yet
```

Hemato.AI uses API tokens to allow access to the API. You can get more information at our [developer portal](https://developer.hemato.ai/).

Hemato.AI expects for the API key to be included in majority of API requests to the server in a header that looks like the following:

`Authorization:HEMATO_AI_AUTH_TOKEN`

<aside class="notice">
You must replace <code>HEMATO_AI_AUTH_TOKEN</code> with your personal API key.
</aside>

# Peripheral Blood Smear

## How does it work?

Images of Peripheral Blood Smears (PBS) usually are very large. It is often impractical to submit an entire sample in one API call.
That's why there is a three step process to submit a PBS for review by Hemato.AI

### Step1
Get a new PBS ID. This ID will be used to upload the images, provide any other information available and afterwards get the results of Hemato.AI's review.

#### HTTP Request
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

### Step2
Upload the sample by calling the PBS upload API as many times as needed.

#### HTTP Request
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


### Step 3
When all the files for a particular PBS is uploaded, you will make a call to mark the PBS as ready to be processed. After this call you will not be able to upload additional files.

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
