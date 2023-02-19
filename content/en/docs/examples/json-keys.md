---
title: "JSON Keys"
description: "JSON Keys"
date: 2020-11-12T13:26:54+01:00
lastmod: 2020-11-12T13:26:54+01:00
draft: false
toc: true
weight: 200
---

Assert if all the required keys exists in our JSON.
Here we are not worried about values here, but only `keys`

**jq operation**

```
. | keys
```


**Example:**

```yaml
name: validate json keys
stages:
  - name: get product one
    request:
      url: "https://dummyjson.com/products/1"
      method: GET
    assert:
      status: 200
      body:
        - select: . | keys
          constrain: json
          want: |
            [
              "brand",
              "category",
              "description",
              "discountPercentage",
              "id",
              "images",
              "price",
              "rating",
              "stock",
              "thumbnail",
              "title"
            ]
```


**Tip**

You can also sort keys using jq operations:

```
. | keys | sort
```
