---
title: API Reference

language_tabs: # must be one of https://github.com/rouge-ruby/rouge/wiki/List-of-supported-languages-and-lexers
  - shell
  - go
  - typescript
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

Examples on this page are in shell script and language bindings in Go, Typescript, Python, C#, and TypeScript coming in future! You can view code examples in the dark area to the right, and you can switch the programming language of the examples with the tabs in the top right.

# Endpoints
Hemato.AI API is made available at specific regional endpoints to adhere to regional regulations. Any data submitted to a regional endpoint will be handled on servers located within the same geographical and regulatory region.

All Hemato.AI API endpoints are only available through encrypted (TLS) connections.

There are a few special purpose endpoints that are not region specific:
- For development and testing purposes anyone can use the Dev endpoint. To use this endpint you need to be assigned developer credentials. This endpoint is not HIPAA compliant and should not be used for any production purposes. Information submited to this service is not protected by any privacy laws and may be inspected by Hemato.AI staff.
- For external quality assurance and external automated testing you can use the QA endpoint. Specially it is imporant to run automated benchmarks agains this endpoint before any significant change to your hardware or sofrwre that has the potential to affect the quality, or clinical validity of the outputs. To use this endpoint you need QA credentials. This endpoint is not HIPAA compliant and should not be used for any production purposes. Information submited to this service is not protected by any privacy laws and may be inspected by Hemato.AI staff.


You can explicitly choose what region to use from this list:

Region | Endpoint                  | Users
-------|---------------------------|----------------------------------------
DEV    | https://dev.api.hemato.ai | for development and testing purposes
QA     | https://qa.api.hemato.ai  | for external quality assurance and external automated testing
USA    | https://us.api.hemato.ai  | Production, for customers in the United States
Canada | https://ca.api.hemato.ai  | Production, for customers in Canada
EU     | https://eu.api.hemato.ai  | Production, for customers in the European Union

# HTTP Responses

Most Hemato.AI API will return a JSON encoded response in this format.

```json
{
  "status": 202,
  "results": {

  },
  "error": "error message if any",
  "user_message": "",
  "debug_info": {
    // helpful information for debugging, except for auth debugging, see Auth section below
  }
}
```


Field  | Description
-------|------------
status | This reflects the [HTTP Status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) of the response.
error | If something has gone wrong, this provides an error message with error codes or more context on what has happened.
results | This is the output of the call, and can be of any type, often a dictionary / object.
user_message | This is a message that can be displayed to the end users. This is often useful in cases where there is an error.
debug_info | Contains helpful information for debugging




# Authentication

All authenticated calls to Hemato.AI API are expected to present an `Authorization` header containing a signed and valid token.
To obtain an authorization token you have 2 options:

1. **Session Based Authentication:** you can call the `/auth/login` endpoint for a specific region and provide a user-name and password-hash (sha256 of the actual password) to obtain a newly minted and signed token. Session Auth is best suited for interactive applications where a user is directly involved in the process of obtaining the token. Example use cases are web applications, mobile apps, and desktop applications where individual users have Hemato.AI accounts. If you have a user facing application but your users don't hold individual Hemato.AI credentials, you can use the Public-Private Key authentication method below.
2. **Public-Private Key:** you can create your own Public-Private key paris, and register your public keys with Hemato.AI. Then you can sign your own tokens using the private key associate with the public keys you registered with Hemato.AI. This is particularly helpful in situations where a user is not directly involved and a device or an automated system needs to make calls to Hemato.AI API. Examples a hardware device digitizes samples and needs to call Hemato.AI to obtain an interpretation of the sample. Or you have a mobile app and allows the users to interact with Hemato.AI but the users don't hold individual Hemato.AI accounts. To learn more read [How to register your public keys with hemato](#Register-Public-Key)


## Session Auth

```shell
# SECURITY RISK:
# Please ONLY use this for debugging login problems.
# MAKE SURE TO REMOVE THIS FROM YOUR SHELL HISTORY.
# otherwise your plan password will remain in your shell history and possibly leak into backups and such.

echo -n YOUR_ACTUAL_PASSWORD | shasum -a 256 | cut -d ' ' -f 1
```

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	shasum := sha256.Sum256([]byte("YOUR_ACTUAL_PASSWORD"))
	fmt.Printf("%x", shasum)
}
```

```typescript
import { CryptoJS } from 'crypto'
const passSha = CtryptoJS.SHA256('YOUR_ACTUAL_PASSWORD').toString(CryptoJS.enc.Hex)
console.log(passSha)
```

For interactive use cases, where you or your users hold individual Hemato.AI accounts, you can use the Session Auth method. This method is best suited for web applications, mobile apps, and desktop applications where individual users have Hemato.AI accounts.

This call does NOT accept the raw user password. Instead you need to provide the sha256 hash of the password. This is to prevent the password from being leaked in internet trafic capture or intercepted on device or in transit. Of course all Hemato.AI endpoints are only available only through encrypted (TLS) connections. Still there are scenarios that a client browser or tools might allow for interception of traffic. You can use the following command to generate the sha256 hash of your password:



> To authorize, using Session Authentication:


```shell
# SECURITY RISK:
# Please ONLY use this for debugging login problems.
# MAKE SURE TO REMOVE THIS FROM YOUR SHELL HISTORY.
# otherwise your plan password will remain in your shell history and possibly leak into backups and such.
export PASS_SHA_256=`echo -n YOUR_ACTUAL_PASSWORD | shasum -a 256 | cut -d ' ' -f 1`
# use httpie from https://httpie.io/
echo '{"user":"ali@example.com","pass_hash":"'${PASS_SHA_256}'"}' | http -F POST https://in.api.hemato.ai/auth/login
# or use curl, you can just pass the correct header with each request
echo '{"user":"ali@example.com","pass_hash":"'${PASS_SHA_256}'"}' | curl -X POST "https://in.api.hemato.ai/auth/login"
```

```go
// Not implemented yet
```

> The above command returns a JSON structure like this:

```json
{
  "status": 202,
    "results": {
        "login_timestamp": "1682945428455621477",
        "token": {
            "HY_APP_AUTH_v1": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9E....."
        },
        "user": {
            "email": "ali@example.com",
            "id": "03349841-588d-56b3-b93c-52d0983f9c75",
            "screen_name": "Ali",
            "user_name": "ali@example.com"
        }
    },
    "debug_info": {
        "delta": "463.021938ms",
        "version": "bb_api.362.pbs_tasks.dfd119b"
    }
}
```

## Public / Private Key Auth

You can sign your own Authentications headers, see [JWT specifications](#JWT-Specifications).

A self-service dashboard to register your public keys is coming soon. In the meantime, please contact us to register your public keys.

## JWT Specifications
If you are using Public-Private Key Auth, you will be generating and signing your own Authorizaiton headers.
Hemato.AI authorization headers are JSON Web Tokens (JWT) (more info at (https://jwt.io/)[https://jwt.io/]).

### Suppoted Signing Algorithms
Curently only `RSA384` (RSASSA-PKCS-v1.5 using SHA-384) and  `RS256` (RSASSA-PKCS-v1.5 using SHA-256) is allowed.

### Required Fields

All the fields in this table are required.
All keys and values are case sensitive.
All JWTs need to be BASE64 encoded when used as Authorization headers.

JWT Claim | Abbreviated | Data Type             | Notes
----------|-------------|-----------------------|----------------------------------
Token ID  | `jti`       | String                | A unique ID for each token. Most useful for audit, debugging and investigation of issues. UUID or high resolution timestamps can be used.
Issuer    | `iss`       | String                | This is the identifier of your organization.
Issed at  | `iat`       | Unix Time aka IntDate | The time the token was issues. A token issued in future is invalid. A token more than 1 hour is considered expired.
Expiraion | `exp`       | Unix Time aka IntDate | The time this token will expire and be invalid.  A token more than 1 hour is considered expired, regardless of this value.
Audience  | `aud`       | String                | This needs to match the domain and region you are calling. Examples are: `us.api.hemato.ai` or `dev.api.hemato.ai`
          |             |                       |
User ID   | `sub`       | String                | This needs to be one of the user ids from the same organization where the signing key belongs to and this use must have permission to access this region.
Role      | `role`      | String                | This is the role of the user. Available roles are `admin`, `user`, `device`, `service`. This is used to determine the permissions of the bearer.
          |             |                       |
Key ID *  | `kid`       | String                | Only required if using `RS256` signing algorithm. This is the ID of the public key that was used to sign the token.
          |             |                       |



### How to create your own keys

```shell
ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key
# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
cat jwtRS256.key
cat jwtRS256.key.pub
```



### Example

On the right you can see an example of a Hemato.AI JWT

> This is the structure of a Hemato.AI JWT
```json
{
  "alg":"HS256", // or RS256
  "typ":"JWT"
}
.
{
  "aud": [
    "us.api.hemato.ai",
  ],
  "exp": 1683095428,
  "iat": 1682945428,
  "iss": "1stdevision.example.com",
  "jti": "2d33d1d518",
  "sub": "01234567-789d-46b7-b38c-45d4562f5c12",
  "role": "user",
  "kid": "01234567-789d-46b7-b38c-45d4562f5c12"
}
```

### Debugging JWTs

You can decode and inspect your JWTs at [https://jwt.io/](https://jwt.io/)

You can also make a call to `/auth/hello` endpoint and get more helpful information about your JWT and reasons it may be rejected.

```shell
# if you are using httpie
http -f POST https://dev.api.hemato.ai/auth/hello Authorization:HEMATO_AI_AUTH_TOKEN
# if you are using curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://in.api.hemato.ai/auth/hello
```


# Get a Detection Report

## Workflow

To get Hemato.AI's detection report on a Peripheral Blood Smear you can follow these steps:

1. Get an ID for your new PBS Study
1. Optionally provide a callback URL
1. Upload the files under the new ID
1. Request a Diagnostic Study
1. Optionally listen for a callback, with information about the sample, dwnloaded file size and formats and sanity checks.
1. Wait for it
1. Optionally listen for a callback, with information about the outcome of the detection study
1. Ask for the report

![Peripheral Blood Study Flow 1](PBSStudyWorkflow1.png)


### 1. Create a New Study
Create a study and store the Study-ID (aka PBS-ID). This ID will be used to upload the images, request diagnostic study of the sample, provide any other information available and afterwards get the results of Hemato.AI's review.
You need to store and keep track of this ID.

#### HTTP Request
Make a `POST` call to the `/pbs` endpoint. You can provide any number of tags (key value string pairs) along side this request. These tags can allow you to associate Patient Proxy Identifiers, Device Identifiers, Organization IDs or any number of other information that is important to you along side the PBS. These tags can later be used to find and retrieve PBSs easier.

Please do NOT store Personally Identifiable Information (PII) in these fields.

```shell
# use httpie from https://httpie.io/
echo '{"tags":{"patient_proxy_id":"PPID12345"}}' | http -f POST https://in.api.hemato.ai/pbs Authorization:HEMATO_AI_AUTH_TOKEN
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://in.api.hemato.ai/pbs
```
> Make sure to replace `HEMATO_AI_AUTH_TOKEN` with your API key.

> The above command returns a JSON structure like this:

```json
{
  "status": 201,
  "results":{
    "pbs_study_id":"c7d453a9-676d-4b9f-a505-8a9021b76dfd",
    "tags": {
      "patient_proxy_id": "PPID12345"
    }
  }
}
```
<aside class="warning">Please do NOT include any Personanlly Identifiable Information (PII) or Personal Health Records (PHI) in the tags</aside>

### 2. Upload the files
Upload the images by calling the PBS upload API as many times as needed.

#### HTTP Request
Make a `POST` call to the `/pbs/YOUR_NEW_PBS_ID/files?file_name=some_file_name.jpg` endpoint. Here `YOUR_NEW_PBS_ID` is the id you obtained from step 1 (`results.pbs_study_id` in the response structure returned).

To determin the file type, this endpoint expects either "file_name" or "content_type" query parameters. For example for a JPEG file you can use `file_name=some_file_name.jpg` or `content_type=image/jpeg`.


```shell

This is an idempotent call. It means you can send the same file multiple times and it will only be counted as one file.
Once submitted however, you cannot update the same file or change it or its metadata. You can only add more  files to the same PBS study.

The file type is determined by the file extension of the provided "file_name" query parameter or the "content_type" provided as a query parameter. So when sending a request to upload a jpg file for example, it is enough to have a file_name in your url ending in .jpg or .jpeg for example  https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files?file_name=MS12_12.jpg


If no file type can be determined or provided, a 400 - Bad Request error is returned.


```shell
# use httpie from https://httpie.io/
http -f POST https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files?file_name=MS12_12.jpg  Authorization:HEMATO_AI_AUTH_TOKEN  < /local/path/to/MS12_12.jpg
# alternatively use curl
curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --data-binary "@/local/path/to/MS12_12.jpg"  https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files
```

```Python
```

```Go

  imageBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", filePath, fileName))
  if err != nil {
		return ...
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID", bytes.NewReader(imageBytes))
	if err != nil {
		return ...
	}
	q := req.URL.Query()
	q.Add("file_name", fileName)
	req.URL.RawQuery = q.Encode()

  err = client.Do(q)
  if err != nil {
		return ...
	}
  ...

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
<li>Plaese note that this is not a form submission call. The entire body of the request is the image.
</ul>
</aside>


### 3. Request a Detection Task
When all the files for a particular PBS is uploaded, you can make a call to mark the PBS as ready to be processed and ask for one or more Detection Tasks to be performed on this PBS. After this call you will not be able to upload additional files for this PBS.

At this time, for this deployment, these Detection Tasks are available:


| Task                        | Type      | Identifier        | Notes |
|-----------------------------|-----------|-------------------|--------------|
| Detect Any Signs of Malaria | Detection | `MALARIA_ANY_ANY` | Reports any sign of malaria parasites, including All species, Active infections, Infections under treatment |


Optionally you can register a `callback_url` that will be called when the report is ready.
We’ll send a POST request to the callback URL with details of the report when ready.
The callback url needs to accept a POST call, and must support HTTPS with a valid server certificate. The body of the call will conform to `{"pbs_stody_id":"25f7de38-ba9c-4b45-a4d8-8c13461d7b39", "report_status": "ready", "progress": 1.0}` but it can be ignored.
So there are two approaches that users can take:
1. Include the `pbs_study_id` or any other information you need (no PII or PHI) in the url to be called.
2. Use a single generic url, but parse and use the information in the body to identify which study is the subject of the call back.

Although it is not strictly needed, we recommend you include a cryptographic signature as part of your URL, so that you can verify the authenticity of the calls to your endpoint. Example:
`{"callback_url":"https://example.com/pbs_report_is_ready/<PBS_STUDY_ID>/<timestamped_cryptographic_signature_valid_only_for_this_PBS_STUDY_ID>"`

This enables early filtration of incoming traffic, particularly in cases where the backend systems may be inundated by a surge of malicious or rogue requests.


```shell
# use httpie from https://httpie.io/
echo '{"diagnostic_tasks":["pbs_v1"], "callback_url":"https://example.com/pbs_report_is_ready/<PBS_STUDY_ID>/"}' | http -f POST https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/tasks Authorization:HEMATO_AI_AUTH_TOKEN
# alternatively use curl
echo '{"diagnostic_tasks":["pbs_v1"], "callback_url":"https://example.com/pbs_report_is_ready/<PBS_STUDY_ID>/"}' | curl -x POST --header "Authorization:HEMATO_AI_AUTH_TOKEN" --data-binary @- https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/tasks
```

> The above command returns a JSON structure like this:

```json
{
  "results":{
    "pbs_study_id":"48ae6352-d15d-440a-83f3-05a1f68aa11b",
    "tasks": [
      "TASK_1_IDENTIFIER"
    ]
  }
}
```


### 4. Wait for it
Depending on the size of the files submitted, the number Diagnostic tasks requested, and other factors like general system workload, it will take between a few seconds to 10s of seconds for the results to be ready.

If you registerd a callback url in step 3, you can wait for that url to be called.

You can check the status of the task by making a GET call to /status

> To get the status of a PBS

```shell
#
http https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/status Authorization:HEMATO_AI_AUTH_TOKEN
#
curl --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/status
```

> This returns
```json
{
  "status": 200,
  "results": {
    "pbs_study_id": "62dd0300-880e-4069-bea8-66c9ca73c207",
    "report_status": "ready",
    "progress": 1.0,
    "progress_message":"Reports are ready",
  },
  "debug_info" :{}
}
```


### 5. Get The Report for a Detection Task
You can get the latest version of the reports by making a GET call.
In cases that there are multiple revisions and you want to get an earlier version of the report you need to specify the version.


> To get the latest diagnostic report on a PBS

```shell
#
http https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/report/MALARIA_ANY_ANY Authorization:HEMATO_AI_AUTH_TOKEN
#
curl --header "Authorization:HEMATO_AI_AUTH_TOKEN" https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/report/MALARIA_ANY_ANY
```

> Which in turn returns a JSON structure like:

```json
{
	"status": 200,
	"results": {
		"pbs_study_id": "4bb7fe9e-b608-4d68-adbe-8655c991f494",
    "report": {
      "MALARIA_ANY_ANY": {
        "positive_count": 12,
        "negative_count": 3
      }
    }
	},
	"debug_info": {}
}
```

# Health Check
If you need to check the health status of the Hemato.AI api you can make a GET call to the heartbeat endpoint.

```shell
http https://in.api.hemato.ai/heartbeat
```

```json
{
    "delta": "59.078µs",
    "heartbeat": "a7ef8fba741237f693c3",
    "request_timestamp": 1680716829,
    "version": "bb_api.331.develop.18ff645"
}
```
