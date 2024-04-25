#!/usr/bin/env bash
xcaddy build --with github.com/acouvreur/sablier/plugins/caddy=.

sudo ./caddy run --config Caddyfile