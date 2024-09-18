---
title: Hemato.AI API Reference

language_tabs: # must be one of https://github.com/rouge-ruby/rouge/wiki/List-of-supported-languages-and-lexers
  - shell
  - go
  - typescript
  - python

toc_footers:
  - <a href='https://hemato.ai'>Hemato.AI</a>
  - <a href='mailto:ai@hemato.ai'>Sign Up for a Developer Key</a>

includes:
  - errors

search: true

code_clipboard: true

meta:
  - name: description
    content: Documentation for Hemato.AI Diagnostic API
---

# Introduction to Hemato.AI Diagnostic API

Welcome to the Hemato.AI Diagnostic API documentation. This powerful interface allows you to leverage Hemato.AI's state-of-the-art Diagnostic AI and Clinical Algorithms for advanced hematopathology analysis.

## Key Features

- Process high-resolution images of peripheral blood slides
- Access cutting-edge AI-driven diagnostic tools
- Generate comprehensive hematopathology reports

## Documentation Structure

Throughout this documentation, you'll find code examples demonstrating API usage in multiple programming languages:

- Shell script
- Go
- TypeScript
- Python

To enhance readability, code examples are displayed in the dark panel on the right side of the page. You can easily switch between programming languages using the tabs located in the top-right corner of the code panel.

We've designed this documentation to provide you with a clear understanding of our API's capabilities and guide you through its implementation. Whether you're a seasoned developer or new to our platform, you'll find the information you need to successfully integrate Hemato.AI's diagnostic power into your applications.


# Endpoints

Hemato.AI API provides access through region-specific endpoints to ensure compliance with local regulations. All data submitted to a regional endpoint are processed on infrastructure dedicated to that geographical or regulatory region.

## Security

All Hemato.AI API endpoints are accessible exclusively through encrypted (TLS) connections, ensuring the highest level of data protection during transmission.

## Special Purpose Endpoints

In addition to regional production endpoints, Hemato.AI offers special-purpose endpoints for specific use cases:

### Development Endpoint (DEV)
- **URL**: https://dev.api.hemato.ai
- **Purpose**: Internal development and testing
- **Access**: Requires developer credentials
- **Important Notes**:
  - Not HIPAA compliant
  - Not suitable for production use
  - Submitted information may be inspected by Hemato.AI staff
  - No privacy law protections apply

### Quality Assurance Endpoint (QA)
- **URL**: https://qa.api.hemato.ai
- **Purpose**: External quality assurance and automated testing
- **Access**: Requires QA credentials
- **Use Case**: Run automated benchmarks before significant hardware or software changes that may affect output quality or clinical validity
- **Important Notes**:
  - Not HIPAA compliant
  - Not suitable for production use
  - Submitted information may be inspected by Hemato.AI staff
  - No privacy law protections apply

## Regional Endpoints

Choose the appropriate regional endpoint based on your location and regulatory requirements:

| Region | Endpoint | Users |
|--------|----------|-------|
| USA | https://us.api.hemato.ai | Production use for customers in the United States |
| Canada | https://ca.api.hemato.ai | Production use for customers in Canada |
| EU | https://eu.api.hemato.ai | Production use for customers in the European Union |
| India | https://in.api.hemato.ai | DEMO MODE |

Please ensure you select the correct endpoint for your region to maintain compliance with local regulations and data protection laws.

# HTTP Responses

The Hemato.AI API typically returns responses in JSON format. The standard response structure is as follows:

```json
{
  "status": 202,
  "results": {},
  "error": "Error message, if applicable",
  "user_message": "",
  "debug_info": {
    // Helpful information for debugging
  }
}
```

## Response Fields

| Field | Description |
|-------|-------------|
| `status` | An integer representing the [HTTP status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) of the response. |
| `results` | The main output of the API call. This field can contain various data types, often a JSON object. |
| `error` | A string providing error details, including error codes or context, when applicable. |
| `user_message` | A human-readable message suitable for display to end users, particularly useful in error scenarios. |
| `debug_info` | An object containing additional information to assist with troubleshooting. Note: Authentication debugging information is handled separately (see Auth section). |

## Status Codes

The `status` field in the response corresponds to standard HTTP status codes. Common codes you may encounter include:

- 200: OK - The request was successful
- 201: Created - A new resource was successfully created
- 202: Accepted - The request has been accepted for processing
- 400: Bad Request - The request was invalid or cannot be processed as is
- 401: Unauthorized - Authentication is required and has failed or has not been provided
- 403: Forbidden - The server understood the request but refuses to authorize it
- 404: Not Found - The requested resource could not be found
- 500: Internal Server Error - The server encountered an unexpected condition


## Error Handling

When an error occurs, the `error` field will contain a descriptive message to help diagnose the issue. The `user_message` field may provide a more user-friendly explanation of the error, which can be directly displayed in your application's user interface.

## Debugging

The `debug_info` object contains additional technical details that can be valuable for troubleshooting issues with your API integration. This information is primarily intended for developers and should not be displayed to end users.

Note: For security reasons, authentication-related debugging information is handled separately. Please refer to the Auth section of this documentation for details on debugging authentication issues.


# Authentication

All authenticated calls to the Hemato.AI API require an `Authorization` header containing a signed and valid token. There are two authentication methods available:

1. **Session-Based Authentication**
2. **Public-Private Key Authentication**

## Session-Based Authentication

This method is ideal for interactive applications where users have individual Hemato.AI accounts. Examples include web applications, mobile apps, and desktop applications.

### Obtaining a Token

To obtain an authorization token:

1. Call the `/auth/login` endpoint for a specific region.
2. Provide a username and password hash (SHA-256 of the actual password).

> Note: Never send raw passwords. Always use the SHA-256 hash for security.

### Generating Password Hash

Here are code examples showing how to generate the password hash (sha256) that you will need to send along a user name to the auth endpoint.


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

### Login Request Example

Here are examples codes that uses the password hash from above to login for an interactive session.

```shell
export PASS_SHA_256=`echo -n YOUR_ACTUAL_PASSWORD | shasum -a 256 | cut -d ' ' -f 1`
echo '{"user":"ali@example.com","pass_hash":"'${PASS_SHA_256}'"}' | \
http -F POST https://in.api.hemato.ai/auth/login
```

```go
// Not implemented yet
```

```typescript
// Not implemented yet
```


### Sample Response

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

## Public-Private Key Authentication

This method is suitable for automated systems or devices that need to make API calls without user interaction. It's also useful for applications where users don't have individual Hemato.AI accounts.

### Key Registration

A self-service dashboard for registering public keys is coming soon. In the meantime, please contact Hemato.AI support to register your public keys.

### JWT Specifications

Hemato.AI uses JSON Web Tokens (JWT) for authorization headers. For detailed information, visit [jwt.io](https://jwt.io/).

#### Supported Signing Algorithms

- `RSA384` (RSASSA-PKCS-v1.5 using SHA-384)
- `RS256` (RSASSA-PKCS-v1.5 using SHA-256)

#### Required JWT Fields

| JWT Claim   | Abbreviated | Data Type | Description |
|-------------|-------------|-----------|-------------|
| Token ID    | `jti`       | String    | Unique ID for each token (e.g., UUID or high-resolution timestamp) |
| Issuer      | `iss`       | String    | Identifier of your organization |
| Issued at   | `iat`       | IntDate   | Token issue time (Unix timestamp) |
| Expiration  | `exp`       | IntDate   | Token expiration time (Unix timestamp, max 1 hour from issue) |
| Audience    | `aud`       | String    | API domain and region (e.g., `us.api.hemato.ai`) |
| User ID     | `sub`       | String    | User ID from the organization owning the signing key |
| Role        | `role`      | String    | User role (`admin`, `user`, `device`, or `service`) |
| Key ID *    | `kid`       | String    | Required for `RS256` algorithm. ID of the public key used for signing |

> Note: All fields are case-sensitive and required. JWTs must be BASE64 encoded when used as Authorization headers.

### Creating Key Pairs

Here is an exmaple shell command to create a key pair. You can then share ONLY THE PUBLIC key with Hemato.AI. Keep the private key safe. You will be using the private key to sign your own authorization headers.

Anyone with access to your private key can impersonate you and your users, devices and servers.


```shell
ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key
# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
cat jwtRS256.key
cat jwtRS256.key.pub
```


### Sample JWT Structure

```json
{
  "alg": "RS256",
  "typ": "JWT"
}
.
{
  "aud": ["us.api.hemato.ai"],
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

- Decode and inspect JWTs at [jwt.io](https://jwt.io/).
- Use the `/auth/hello` endpoint for detailed JWT information and potential rejection reasons.


```shell
http -f POST https://dev.api.hemato.ai/auth/hello Authorization:HEMATO_AI_AUTH_TOKEN
```

# Peripheral Blood Smear (PBS) Detection Workflow

This guide outlines the process of submitting microscopic images of peripheral blood samples and receiving diagnostic responses from Hemato.AI.

## Overview

![Peripheral Blood Study Flow 1](PBSStudyWorkflow1.png)

The workflow consists of the following steps:

1. Create a new PBS study
2. Upload image files
3. Request a detection task
4. Wait for processing
5. Retrieve the detection report

## Detailed Workflow

### 1. Create a New Study

Create a study to obtain a Study ID (PBS-ID) for subsequent operations.

#### HTTP Request

`POST /pbs`

```shell
echo '{"tags":{"patient_proxy_id":"PPID12345","sample_id":"XYZ12345678","device_id":"SABCD"}}' | \
http -f POST https://in.api.hemato.ai/pbs \
'Authorization:HEMATO_AI_AUTH_TOKEN'
```

#### Response

```json
{
  "status": 201,
  "results": {
    "pbs_study_id": "c7d453a9-676d-4b9f-a505-8a9021b76dfd",
    "tags": {
      "patient_proxy_id": "PPID12345",
      "sample_id":"XYZ12345678",
      "device_id":"SABCD"
    }
  }
}
```

> **Warning**: Do not include Personally Identifiable Information (PII) or Personal Health Information (PHI) in tags.

### 2. Upload Image Files

Upload images associated with the PBS study.
For each image, you are required to provide `rbc_diameter` parameter as a URL query param.
`rbc_diameter` is the number of pixels in the diameter of a normal human Red Blood Cell in the image. This is a measure of magnification that includes digital and optical and sensor size related magnifications.
You do NOT need to measure this for every image. You can do the measurements once and use the same number for all image prepared with the same device settings.
Alternatively you can share samples with Hemato.AI and Hemato.AI will provide the numbers for your use.

#### HTTP Request

`POST /pbs/{PBS_ID}/files?file_name={filename}&rbc_diameter={size}`

```shell
http -f POST \
https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/files?file_name=MS12_12.jpg&rbc_diameter=85 \
'Authorization:HEMATO_AI_AUTH_TOKEN' < /path/to/MS12_12.jpg
```

#### Response

```json
{
  "status": 201,
  "results": {
    "pbs_study_id": "c7d453a9-676d-4b9f-a505-8a9021b76dfd",
    "file_id": "sha224-i-4d58b0cee1e38f75f93e2e605d1e6ad44c24a975d0c3b53367de7c86"
  }
}
```

> **Note**: This is not a multipart upload. The entire request body should contain the image data.

### 3. Request a Detection Task

After uploading all files, request a detection task to be performed on the PBS.

#### Available Detection Tasks

| Task | Type | Identifier | Description |
|------|------|------------|-------------|
| Detect Any Signs of Malaria | Detection | `MALARIA_ANY_ANY` | Reports any sign of malaria parasites, including all species, active infections, and infections under treatment |

#### HTTP Request

`POST /pbs/{PBS_ID}/tasks`

```shell
echo '{"diagnostic_tasks":["MALARIA_ANY_ANY"], "callback_url":"https://example.com/pbs_report_is_ready/{PBS_STUDY_ID}/"}' | \
http -f POST https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/tasks \
'Authorization:HEMATO_AI_AUTH_TOKEN'
```

#### Response

```json
{
  "results": {
    "pbs_study_id": "48ae6352-d15d-440a-83f3-05a1f68aa11b",
    "tasks": [
      "TASK_1_IDENTIFIER"
    ]
  }
}
```

### 4. Wait for Processing

Processing time varies based on number of files and their size, number of tasks, and system workload and available resources. You can either:

- Wait for a callback if you provided a `callback_url`
- Check the status manually

#### Check Status

`GET /pbs/{PBS_ID}/status`

```shell
http https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/status \
'Authorization:HEMATO_AI_AUTH_TOKEN'
```

#### Response

```json
{
  "status": 200,
  "results": {
    "pbs_study_id": "62dd0300-880e-4069-bea8-66c9ca73c207",
    "report_status": "ready",
    "progress": 1.0,
    "progress_message": "Reports are ready"
  }
}
```

### 5. Retrieve the Detection Report

Once processing is complete, retrieve the report for the requested detection task.

#### HTTP Request

`GET /pbs/{PBS_ID}/report/{TASK_IDENTIFIER}`

```shell
http https://in.api.hemato.ai/pbs/YOUR_NEW_PBS_ID/report/MALARIA_ANY_ANY \
Authorization:HEMATO_AI_AUTH_TOKEN
```

#### Response

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
  }
}
```

## Best Practices

1. Store and track the PBS-ID (Study ID) throughout the process.
2. Ensure accurate "rbc diameter" when uploading images.
3. Use HTTPS for callback URLs and consider including a cryptographic signature for security.
4. Do not include PII or PHI in tags or callback URLs.

For any issues or questions, please contact Hemato.AI support.


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
