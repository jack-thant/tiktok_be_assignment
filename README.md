# Tiktok Tech Immersion 2023 Backend Assigment: Instant Messaging System

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

## Project Description
Instant Messaging System is the system using **Golang** where the user is able to send and get the messages using a set of specific APIs. 
The system is implemented only in the back-end part and does not include the front end part and the account / authentication. 
In addition, the system contains two services: **HTTP server** and **RPC server**, **Redis database**.

## Installation

Requirement:

- golang 1.18+
- docker

To install dependency tools:

```` bash
make pre
````

## Run

````bash
docker-compose up -d
````

Check if it's running:

````bash
curl localhost:8080/ping
````
To send a message:
```` bash
curl --location 'localhost:8080/api/send' \
--header 'Content-Type: application/json' \
--data '{
"chat": "jack:marcus",
"text": "long time no see",
"sender": "marcus"
}'
````

To pull message from specific room:
````bash
curl --location --request GET 'localhost:8080/api/pull' \
--header 'Content-Type: application/json' \
--data '{
"chat": "jack:marcus",
"cursor": 0,
"limit": 2,
"reverse": true
}'
````

To pull multiple message
````bash
curl --location --request GET 'localhost:8080/api/pull' \
--header 'Content-Type: application/json' \
--data '{
"chat": "jack:marcus",
"cursor": 0,
"limit": 10,
"reverse": true
}'
````
## Credit
[Tiktok Tech Immersion 2023 Project Guide](https://o386706e92.larksuite.com/docx/QE9qdhCmsoiieAx6gWEuRxvWsRc)