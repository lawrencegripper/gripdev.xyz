---
author: gripdev
category:
  - autocomplete
  - yaml
date: "2017-11-09T09:27:27+00:00"
guid: http://blog.gripdev.xyz/?p=1103
tag:
  - autocomplete
  - yaml
title: Autocomplete Kubernetes YAML files in VSCode
url: /2017/11/09/autocomplete-kubernetes-yaml-files-in-vscode/

---
I've increasingly been working with Kubernetes and hence lots of YAML files.

It's nice and easy to get autocomplete setup for the Kubernetes YAML using this awesome extension [YAML Support by Red Hat](https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)

Setup:

- Install the Extension
- Add the following to your settings

  ```
  "yaml.schemas": {
    "Kubernetes": "*.yaml"
  }

  ```

- Reload the editor

Here is me setting it up and showing it off:
![vscodeyamlautocomplete3](/wp-content/uploads/2017/11/vscodeyamlautocomplete3.gif)

Massive thanks to the team that worked on the extension and language server to support this!
