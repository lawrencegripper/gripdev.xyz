---
author: gripdev
category:
  - uncategorized
date: "2021-11-18T13:13:06+00:00"
guid: https://blog.gripdev.xyz/?p=1541
title: Writing OPA rules to lint Kubernetes YAML resource and Outputting as annotations on Pull Requests with GitHub Actions
url: /2021/11/18/writing-opa-rules-to-lint-kubernetes-yaml-resource-and-outputting-as-annotations-on-pull-requests-with-github-actions/

---
Warning: This expects you already know about rego/opa and is more of a brain dump than a blog.

First up take a look at `conftest` it's a great little CLI tool which lets you take rules you've written in `rego/opa` and run them easily.

In our case we have the following:

\- `./rules` folder containing our rego rules  
\- `./yaml` folder containing yaml we want to validate

We're going to write a rule to flag duplicate resources, ie. when you have two yamls with the same kind and name.

The rule will be written in `rego` then executed by `conftest` and when a failure occurs it'll be shown as an annotation on the Pull Request using GitHub Actions.

Firstly for `conftest` we want to use the `--combine` option so we get a single array of all the yaml files passed into our rule. This allows us to compare the files against one another to determine if there are any duplicates.

The data structure you get looks a bit like this:

```
[
    {
        "content": {
            "apiVersion": "thing",
            "kind": "deployment"
....
        },
        "path": "path/to/yaml/file"
    }
]
```

As well as validating the rule we also use the `path` property to output metadata about which file generated the warning.

We can then use `jq` to parse the json output from `conftest` and convert it to ["Workflow Commands: Warning Messages"](https://docs.github.com/en/actions/learn-github-actions/workflow-commands-for-github-actions#setting-a-warning-message) these are outputted to the console and read by GitHub Actions. With the details in the message it generates an annotation on the file in the PR like so:   

[![](/wp-content/uploads/2021/11/image.png)](/wp-content/uploads/2021/11/image.png)

Here is a gist of this wrapped together.

https://gist.github.com/lawrencegripper/56aa710d788be8d0d6d1bad72a50943f
