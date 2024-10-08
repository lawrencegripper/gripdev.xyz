---
author: gripdev
draft: true
category:
  - programming
  - docker
  - docker-compose
date: "2024-10-08T21:02:00+00:00"
title: "Advanced Docker: Mount inline config files from docker-compose file"
url: /2024/10/08/advanced-docker-techniques
---

So you have some config that you want to provide to a service started from `docker compose`. 

Did you know you can keep things simple by inlining that config in the compose file, rather 
than messing with mounts?

## Inline content in `docker-compose.yaml` 

In our case we want a `json` file with 👇 mounted at `/etc/consul/client.json` 

```json
{
  "node_name": "consul-client",
  "data_dir": "/consul/data",
  "log_level": "ERROR",
  "server": false,
  "retry_interval": "10s",
  "retry_join":[
      "consul-server"
  ]
}
```

You can create a `config` section and underneath provide the content of each named config. 

Create an entry for the file and use the [`|` sign on the `content`](https://yaml.org/spec/1.2-old/spec.html#id2760844) field to inline. 

This accepts a multiline string and preserves new lines `/n` in the output. 

Then we can add a `configs` section under the service where we want to mount the file, specifying the source and `target` as the path on disk where the file should be placed in the container.

```yaml
    configs:
      - source: consul_client.json
        target: /etc/consul/client.json
```

All together this looks like 👇

```yaml
configs:
  consul_client.json:
    content: |
      {
          "node_name": "consul-client",
          "data_dir": "/consul/data",
          "log_level": "ERROR",
          "server": false,
          "retry_interval": "10s",
          "retry_join":[
              "consul-server"
          ]
      }
services:
  windmill-host:
    build:
      context: .
      dockerfile: ./mock-services/some-host/Dockerfile
    container_name: windmill-host
    ...
    configs:
      - source: consul_client.json
        target: /etc/consul/client.json

```

## Further reading

Here are the Docker docks which go into more detail:
- [Config section in Docker Compose](https://docs.docker.com/reference/compose-file/configs/)