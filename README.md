# Traefik Ondemand Plugin

Traefik plugin to scale down to zero containers on docker swarm.

## Description

A container may be a simple nginx server serving static pages, when no one is using it, it still consume CPU and memory.

With this plugin you can scale down to zero when there is no request for the service.
It will scale back to 1 when there is a user requesting the service.

## Configuration

- `serviceUrl` the traefik-ondemand-service url
- `dockerServiceName` the service to sclae on demand name (docker service ls)
- `timeoutSeconds` timeout in seconds for the service to be scaled down to zero after the last request

See `config.yml` and docker-compose.yml for full configuration.

## Demo

The service whoami is scaled to 0. We configured a timeout of 10 seconds.

![Demo](./img/demo.gif)

## Run the demo

- `docker swarm init`
- `export TRAEFIK_PILOT_TOKEN=your_traefik_pilot_token`
- `docker stack deploy -c docker-compose.yml TRAEFIK_HACKATHON`

## Limitations

### Cannot use service labels

You cannot set the labels for a service inside the service definition.

Otherwise when scaling to 0 the specification would not be found because there is no more task running. So you have to write it under the dynamic configuration file.

### The need of "traefik-ondemand-service"

This is a small project developped to interact freely with the docker deamon and manage an independant lifecycle.

*We may try to update this plugin to embed the scaling behavior in a future.*

-> The source is available at https://github.com/acouvreur/traefik-ondemand-service

## Authors

[Alexis Couvreur](https://www.linkedin.com/in/alexis-couvreur/) (left) and [Alexandre Hiltcher](https://www.linkedin.com/in/alexandre-hiltcher/) (right)

![Alexandre and Alexis](./img/gophers-traefik.png)