---
title: "Tool: CLI for JSON Schema Validation with Line Number and detailed error information"
date: 2025-01-31T07:02:00+00:00
author: "Lawrence Gripper"
tags: ["programming", "json", "schema", "validation", "cli", "tool"]
categories: ["programming"]
url: /2025/01/31/cli-json-schema-validation-line-numbers-detailed-error/
draft: false
---

Quick post.

My goal was, I thought, simple:

**Run a CLI tool which checks if a JSON document is valid for a JSON schema and returns the line numbers with error details**

So I head over to [`json-schema.org/tools`](https://json-schema.org/tools?query=&sortBy=name&sortOrder=ascending&groupBy=toolingTypes&licenses=&languages=&drafts=&toolingTypes=&environments=&showObsolete=false) and look for one.

I try:

- `ajv-cli` ⛔ Doesn't handle field type of `date`
- `boon` ❌ Doesn't show line numbers  
- `Test-Json` ❌ built in powershell cmdlet - Doesn't handle `if-then-else` schemas or show line info
- `check-jsonschema` ❌ no line info

Along with that most of them didn't give great error messages.

Frustrated, I remembered using [jsonschemavalidator.net](https://www.jsonschemavalidator.net/) and double checked.

It outputs all the info and shows how to do it in code using the Newtonsoft library ✨!

So a few hours later there is now [github.com/lawrencegripper/gripdev-json-schema-validator](https://github.com/lawrencegripper/gripdev-json-schema-validator)

A CLI and PowerShell Module which gives detailed info and formats it nicely to show to users too. For example:

```
❌ JSON validation failed!
   Found the following errors:

  ❌ Error Details:
    └─ Message: Required properties are missing from object: city.
    └─ Location: Line 4, Position 14
    └─ Path: address
    └─ Value: city
```

The CLI and CmdLet also output structured JSON data so you can use it to annotate files on builds or write scripts. [For example, the readme shows how to use it to Annotate with GitHub actions](https://github.com/lawrencegripper/gripdev-json-schema-validator?tab=readme-ov-file#overview).

Enjoy!

Big shout out to [James King (Newtonsoft)](https://www.newtonsoft.com/jsonschema) for the heavy lifting here using the [AGPL version of his library](https://www.nuget.org/packages/Newtonsoft.Json.Schema/4.0.1/License).