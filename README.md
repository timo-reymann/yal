yal - Yet Another Landingpage
===
[![LICENSE](https://img.shields.io/github/license/timo-reymann/yal)](https://github.com/timo-reymann/yal/blob/main/LICENSE)
[![DockerHub Pulls](https://img.shields.io/docker/pulls/timoreymann/yal)](https://hub.docker.com/r/timoreymann/yal)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/yal)](https://goreportcard.com/report/github.com/timo-reymann/yal)
[![codecov](https://codecov.io/gh/timo-reymann/yal/graph/badge.svg?token=rsQYV5lODS)](https://codecov.io/gh/timo-reymann/yal)
[![CircleCI](https://circleci.com/gh/timo-reymann/yal.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/yal)
[![GitHub Release](https://img.shields.io/github/v/tag/timo-reymann/yal?label=version)](https://github.com/timo-reymann/yal/releases)
[![Renovate](https://img.shields.io/badge/renovate-enabled-green?logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAzNjkgMzY5Ij48Y2lyY2xlIGN4PSIxODkuOSIgY3k9IjE5MC4yIiByPSIxODQuNSIgZmlsbD0iI2ZmZTQyZSIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTUgLTYpIi8+PHBhdGggZmlsbD0iIzhiYjViNSIgZD0iTTI1MSAyNTZsLTM4LTM4YTE3IDE3IDAgMDEwLTI0bDU2LTU2YzItMiAyLTYgMC03bC0yMC0yMWE1IDUgMCAwMC03IDBsLTEzIDEyLTktOCAxMy0xM2ExNyAxNyAwIDAxMjQgMGwyMSAyMWM3IDcgNyAxNyAwIDI0bC01NiA1N2E1IDUgMCAwMDAgN2wzOCAzOHoiLz48cGF0aCBmaWxsPSIjZDk1NjEyIiBkPSJNMzAwIDI4OGwtOCA4Yy00IDQtMTEgNC0xNiAwbC00Ni00NmMtNS01LTUtMTIgMC0xNmw4LThjNC00IDExLTQgMTUgMGw0NyA0N2M0IDQgNCAxMSAwIDE1eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik04MSAxODVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzI1YzRjMyIgZD0iTTIyMCAxMDBsMjMgMjNjNCA0IDQgMTEgMCAxNkwxNDIgMjQwYy00IDQtMTEgNC0xNSAwbC0yNC0yNGMtNC00LTQtMTEgMC0xNWwxMDEtMTAxYzUtNSAxMi01IDE2IDB6Ii8+PHBhdGggZmlsbD0iIzFkZGVkZCIgZD0iTTk5IDE2N2wxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMDBhZmIzIiBkPSJNMjMwIDExMGwxMyAxM2M0IDQgNCAxMSAwIDE2TDE0MiAyNDBjLTQgNC0xMSA0LTE1IDBsLTEzLTEzYzQgNCAxMSA0IDE1IDBsMTAxLTEwMWM1LTUgNS0xMSAwLTE2eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xMTYgMTQ5bDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxZGRlZGQiIGQ9Ik0xMzQgMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxYmNmY2UiIGQ9Ik0xNTIgMTEzbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xNzAgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTYzIDE2N2wxOC0xOCAxOCAxOC0xOCAxOHpNOTggMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMzNGVkZWIiIGQ9Ik0xMzQgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTE1MyA3OGwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMzRlZGViIiBkPSJNODAgMTEzbDE4LTE3IDE4IDE3LTE4IDE4ek0xMzUgNjBsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzk4ZWRlYiIgZD0iTTI3IDEzMWwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjYjUzZTAyIiBkPSJNMjg1IDI1OGw3IDdjNCA0IDQgMTEgMCAxNWwtOCA4Yy00IDQtMTEgNC0xNiAwbC02LTdjNCA1IDExIDUgMTUgMGw4LTdjNC01IDQtMTIgMC0xNnoiLz48cGF0aCBmaWxsPSIjOThlZGViIiBkPSJNODEgNzhsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzAwYTNhMiIgZD0iTTIzNSAxMTVsOCA4YzQgNCA0IDExIDAgMTZMMTQyIDI0MGMtNCA0LTExIDQtMTUgMGwtOS05YzUgNSAxMiA1IDE2IDBsMTAxLTEwMWM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiMzOWQ5ZDgiIGQ9Ik0yMjggMTA4bC04LThjLTQtNS0xMS01LTE2IDBMMTAzIDIwMWMtNCA0LTQgMTEgMCAxNWw4IDhjLTQtNC00LTExIDAtMTVsMTAxLTEwMWM1LTQgMTItNCAxNiAweiIvPjxwYXRoIGZpbGw9IiNhMzM5MDQiIGQ9Ik0yOTEgMjY0bDggOGM0IDQgNCAxMSAwIDE2bC04IDdjLTQgNS0xMSA1LTE1IDBsLTktOGM1IDUgMTIgNSAxNiAwbDgtOGM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiNlYjZlMmQiIGQ9Ik0yNjAgMjMzbC00LTRjLTYtNi0xNy02LTIzIDAtNyA3LTcgMTcgMCAyNGw0IDRjLTQtNS00LTExIDAtMTZsOC04YzQtNCAxMS00IDE1IDB6Ii8+PHBhdGggZmlsbD0iIzEzYWNiZCIgZD0iTTEzNCAyNDhjLTQgMC04LTItMTEtNWwtMjMtMjNhMTYgMTYgMCAwMTAtMjNMMjAxIDk2YTE2IDE2IDAgMDEyMiAwbDI0IDI0YzYgNiA2IDE2IDAgMjJMMTQ2IDI0M2MtMyAzLTcgNS0xMiA1em03OC0xNDdsLTQgMi0xMDEgMTAxYTYgNiAwIDAwMCA5bDIzIDIzYTYgNiAwIDAwOSAwbDEwMS0xMDFhNiA2IDAgMDAwLTlsLTI0LTIzLTQtMnoiLz48cGF0aCBmaWxsPSIjYmY0NDA0IiBkPSJNMjg0IDMwNGMtNCAwLTgtMS0xMS00bC00Ny00N2MtNi02LTYtMTYgMC0yMmw4LThjNi02IDE2LTYgMjIgMGw0NyA0NmM2IDcgNiAxNyAwIDIzbC04IDhjLTMgMy03IDQtMTEgNHptLTM5LTc2Yy0xIDAtMyAwLTQgMmwtOCA3Yy0yIDMtMiA3IDAgOWw0NyA0N2E2IDYgMCAwMDkgMGw3LThjMy0yIDMtNiAwLTlsLTQ2LTQ2Yy0yLTItMy0yLTUtMnoiLz48L3N2Zz4=)](https://renovatebot.com)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_yal&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=timo-reymann_yal)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_yal&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=timo-reymann_yal)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_yal&metric=bugs)](https://sonarcloud.io/summary/new_code?id=timo-reymann_yal)

<p align="center">
	<img width="800" src="https://raw.githubusercontent.com/timo-reymann/yal/main/.github/images/example.png">
    <br />
    A simple link hub, to display and search links. Allows easy branding, runs with the least privileges and is simple to use.
</p>

## Demo

[Click here](https://timo-reymann.github.io/yal/)

## Features

- statically generated site
- single static-compiled go binary
- runs as non-root by default
- integrate any search engine
- simple and intuitive design
- integrated search
- inlines external images on start up
- dependency free
- fully accessible for blind people & screen reader users

## Requirements

- any container platform or supported base system

## Installation

### Run server as container

The container generates the static HTML page on startup, keeps it in memory and serves it using a go webserver.

You can simply run it using e.g. docker with docker-compose:

```yaml
version: "3.5"
services:
  yal:
    image: timoreymann/yal:latest # check for version to use if you would like to pin it
    restart: always
    ports:
      - <public-port>:2024
    volumes:
      - ./config:/app/config
      - ./icons:/app/icons # optional in case you want to use local icons
      - ./images:/app/images # contains favicon etc.
    environment:
      # Port to listen
      YAL_PORT: 2024
      # page title for html
      YAL_PAGE_TITLE: My link hub
      # Path to config files
      YAL_CONFIG_FOLDER: /app/config
      # Path to images, see below for available ones
      YAL_IMAGES_FOLDER: /app/images
      # Omit file extension as it will be picked up by name automatically
      # each of the entries is searched like ${YAL_IMAGES_FOLDER}/<icon>{png,jpg,jpeg,svg}
      # if that does not succeed, an attempt is made to load the path as is and if
      # there is no such file it tries to load it as an URL.
      YAL_MASCOT: mascot # the mascot to display on the left
      YAL_LOGO: logo # logo to display on the right
      YAL_BACKGROUND: background # background image for page
      YAL_FAVICON: favicon # favicon to serve
```

### Generate static HTML file

Specifying all env vars manually (if you want to customize them) and keeping the directory structure you can also just
generate a static HTML page.

1. Download the [latest release](https://github.com/timo-reymann/yal/releases/latest) for your platform
2. `./yal --render --output file.html`
3. Serve `file.html` using any static file server

### Use with your CI provider

Instead of the regular image use the CI version: `timoreymann/yal:ci` or for a versioned
tag `timoreymann/yal:{version}-ci`.

#### CircleCI

```yaml
version: 2.1

# Define the jobs we want to run for this project
jobs:
  build-page:
    docker:
      - image: timoreymann/yal:ci
    steps:
      # Checkout repo with yal config and assets
      - checkout
      - run:
          name: Generate page with yal
          command: yal
      # Upload the artifact html somewhere
      - store_artifacts:
          path: templated.html

workflows:
  build_link_hub:
    jobs:
      - build-page
```

#### Gitlab CI

```yaml
stages:
  - build

build-page:
  stage: build
  image:
    name: timoreymann/yal:ci
    entrypoint: [""]
  script:
    # Build page with yal config and assets from project
    - yal
  artifacts:
    paths:
      - templated.html
```

## Configuration

THe container comes with some demo data by default, while the CLI will fail with an error when you attempt to render the
page without the corresponding files present.

### Configuration

| Environment variable  | Flag                | Default                                     | Description                                                                                                                                                                                                      |
|:----------------------|:--------------------|:--------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| YAL_BACKGROUND        | --background        | `background`                                | Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as a background image                                                                                    |
| YAL_BACKGROUND_FILTER | --background-filter | `blur(5px) brightness(0.9)`                 | CSS Filter to apply to the background image. See [MDN docs](https://developer.mozilla.org/en-US/docs/Web/CSS/filter-function) for more information and examples for the filter CSS function for more information |
| YAL_CONFIG_FOLDER     | --config-folder     | `config`                                    | Relative or absolute path where the configuration files reside                                                                                                                                                   |
| YAL_FAVICON           | --favicon           | `favicon`                                   | Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as favicon for the page                                                                                  |
| YAL_IMAGES_FOLDER     | --images-folder     | `images`                                    | Relative or absolute path where the images reside                                                                                                                                                                |
| YAL_LOGO              | --logo              | `logo`                                      | Basename of a file  without extension (searched in images-folder) or an HTTP url of the image to be used as a logo on the right                                                                                  |
| YAL_MASCOT            | --mascot            | `mascot`                                    | Basename of a file without extension (searched in images-folder) or an HTTP url of the image to be used as a logo on the left                                                                                    |
| YAL_OUTPUT            | --output            | `templated.html`                            | File to render to if -render is specified, use - to render to stdout                                                                                                                                             |
| YAL_PAGE_TITLE        | --page-title        | `LinkHub - The place where it just clicks.` | Title of the HTML page generated                                                                                                                                                                                 |
| YAL_PORT              | --port              | `2024`                                      | The HTTP port of the server when run with serve (default)                                                                                                                                                        |
| YAL_RENDER            | --render            | `false`                                     | Render to output and exit                                                                                                                                                                                        |
| YAL_SERVE             | --serve             | `false`                                     | Render and Serve on HTTP                                                                                                                                                                                         |
| YAL_TEMPLATE_FILE     | --template-file     | `builtin`                                   | Template file to Render, builtin uses the bundled one with yal                                                                                                                                                   |

### Files

Besides the env vars, there are two config files to maintain:

#### searchEngines.json

Configures the search engines for search box to display as last elements

```json
[
  {
    "title": "Name",
    "urlPrefix": "https://my.search?text=<here search term will be appended>"
  }
]
```

#### items.json

Configures the links to display.

```json
[
  {
    "title": "<Title of the section>",
    "entries": [
      {
        "text": "<Display text for the link>",
        "link": "<link>",
        "description": "<short description for search and hover tooltip>",
        "icon": "<url or local path, can be relative; needs to be accessible by container and will be inlined on start up>"
      }
    ]
  }
]
```

## Motivation

There are a tons of landing pages out there, each has a unique set of features.

Some simply provide a lot of stuff that are not necessary for a simple link hub. Others look too cluttered or are not
intuitive to use.

This project aims to fill the gap and provide a link hub with search that is easy to brand and use. Nothing more or
less.

## Contributing

I love your input! I want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the configuration
- Submitting a fix
- Proposing new features
- Becoming a maintainer

To get started please read the [Contribution Guidelines](./CONTRIBUTING.md).

## Development

### Requirements

- [GNU make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/get-docker/)
- [pre-commit](https://pre-commit.com/)
- Go 1.21

### Test

```shell
make test
```

### Build

```shell
make build
```

### Alternatives

- [Heimdall](https://github.com/linuxserver/Heimdall) - in case you need more powerful features such as widgets etc.
- [jump](https://github.com/daledavies/jump) - if you want more whitespace and hidden functionality by default
- [homepage](https://github.com/gethomepage/homepage) - if you want to utilize all the space completely and need widgets
- [homer](https://github.com/bastienwirtz/homer) - for a bit more bloated UI
