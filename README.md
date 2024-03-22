# ImpactMapper

Golang program that makes it easy to build [impact maps](https://www.impactmapping.org/drawing.html) with graphviz

Also uses [toml](https://github.com/toml-lang/toml) files parsed by a [toml parser](https://github.com/BurntSushi/toml) to actually build the structures used.

## Modes

The mapper supports two modes, horizontal and vertical. The horizontal mode matches what the impact maps site while the vertical looks more like a tree.

## Development

This project utilizes the new [Go modules](https://github.com/golang/go/wiki/Modules) feature introduced in Go 1.11 and as a result at least this version is required for development. You just need to clone this repo and develop as usual.