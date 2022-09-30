FROM --platform=$BUILDPLATFORM node:16-alpine3.16 as builder
ARG BUILDPLATFORM

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

COPY . .

RUN npm run build

FROM nginx:stable-alpine

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/conf.d /etc/nginx/conf.d

COPY --from=builder /app/dist/ui /usr/share/nginx/html

EXPOSE 80
