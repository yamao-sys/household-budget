#!/bin/bash

pnpm i
pnpm playwright install
pnpm playwright install-deps
pnpm gen:api-spec:watch &
pnpm dev
