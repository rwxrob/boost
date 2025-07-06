FROM ubuntu

LABEL maintainer="rob@rwx.gg"
LABEL io.k8s.description="Linux Beginner Boost (made with ❤️ by rwxrob)"
LABEL org.opencontainers.image.source="https://github.com/rwxrob/boost"

ENV DEBIAN_FRONTEND=noninteractive

COPY . /

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked apt-get update -q
RUN apt-get install -yq --no-install-recommends git

USER ubuntu
