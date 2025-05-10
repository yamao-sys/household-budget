FROM node:22-alpine AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY app /app/app
COPY components /app/components
COPY const /app/const
COPY features /app/features
COPY lib /app/lib
COPY public /app/public
COPY types /app/types
COPY hooks /app/hooks
COPY middlewares /app/middlewares
COPY contexts /app/contexts
COPY services /app/services
COPY package.json package.json
COPY pnpm-lock.yaml pnpm-lock.yaml
COPY react-router.config.ts react-router.config.ts
COPY tsconfig.json tsconfig.json
COPY vite.config.ts vite.config.ts

COPY .env.production .env.production

RUN pnpm i
RUN pnpm build

FROM nginx:1.28.0-alpine

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/build/client /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
