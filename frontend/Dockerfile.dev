FROM node:22-bookworm-slim

WORKDIR /app

COPY . .

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN npm install -g pnpm \
    && SHELL=bash pnpm setup \
    && source /root/.bashrc \
    && pnpm install turbo --global
