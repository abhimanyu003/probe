---
title: "Assertions"
description: "This article will introduce how to Write Tests."
weight: 2
---

## Overview

Probe use `jq` as JSON query processor

`jq` is like sed for JSON data - you can use it to slice and filter and map and transform structured data with the same ease that sed, awk, grep and friends let you play with text.


👉 More details on jq: [https://stedolan.github.io/jq/](https://stedolan.github.io/jq/)


**Some JQ Playgrounds**

* [JQPlay](https://jqplay.org/)
* [JQTerm](https://jqterm.com/?query=.) ( amazing one )


Lets say if you have API JSON output which looks like this.

```json
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

**Use "probe" to test**

* Create a file with name `main.yaml` ( filename can be anything )

```yaml
name: Writing Test
stages:
  - name: first stage
    request:
      url: 'https://jsonplaceholder.typicode.com/todos/1'
      method: GET
    assert:
      body:
        - select: .userId # jq query language
          want: 1
        - select: .title # jq query language
          want: "delectus aut autem"
```
> 💡 Here `select: .userId` and `select: .title` is jq query language.

## Assert Response Status Code

Assert response status code like `200`, `404`, `500` etc.

```yaml
assert:
  status: 200 # provide valid response header code
```

### Example

```yaml
name: Assert response status
stages:
  - name: test todo endpoint
    request:
      url: 'https://jsonplaceholder.typicode.com/todos/1'
      method: GET
    assert:
      status: 200 # assert of response status code
```


## Assert Response Headers

It's also possible to assert response headers.

```yaml
assert:
  status: 200
  headers:
    - select: # select header key
      want: # header value
```

### Example

```yaml
name: Assert header
stages:
  - name: test todo endpoint
    request:
      url: 'https://jsonplaceholder.typicode.com/todos/1'
      method: GET
    assert:
      status: 200
      headers:
        - select: pragma
          want: "no-cache"
```

## Assert JSON field

You can validate JSON fields, using `constrain` tags

```
constrain:json
```

### Example

```yaml
name: Test JSON field
stages:
  - name: first user
    request:
      url: 'https://jsonplaceholder.typicode.com/users/1'
      method: GET
    assert:
      body:
        - select: .address.geo # assert json field
          constrain: json # json constrain tag
          want: |
            {
              "lat": "-37.3159",
              "lng": "81.1496"
            }
```

## More Live Examples

JSON response from [https://jsonplaceholder.typicode.com/users/1](https://jsonplaceholder.typicode.com/users/1)

It looks something like this:

```json
{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz",
  "address": {
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "city": "Gwenborough",
    "zipcode": "92998-3874",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    }
  },
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org",
  "company": {
    "name": "Romaguera-Crona",
    "catchPhrase": "Multi-layered client-server neural-net",
    "bs": "harness real-time e-markets"
  }
}
```

To test the above endpoint using probe

* Create a `main.yaml` file with this

```yaml
name: Test User
stages:
  - name: test first user
    request:
      url: 'https://jsonplaceholder.typicode.com/users/1'
      method: GET
    assert:
      status: 200
      body:
        - select: .email
          want: "Sincere@april.biz"
        - select: .address.street
          want: "Kulas Light"
        - select: .address.geo
          constrain: json
          want: |
            {
              "lat": "-37.3159",
              "lng": "81.1496"
            }
```
