FROM node:22-alpine AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY . .

COPY .env.production .env.production

RUN pnpm i
RUN pnpm build

FROM nginx:1.28.0-alpine

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/build/client /usr/share/nginx/html
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
