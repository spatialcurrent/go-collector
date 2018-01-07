# go-collector

# Description

**go-collector** is a simple library for collecting a list of files from a given path in a filesystem.

# Installation

```
go get github.com/spatialcurrent/go-collector
```

# Usage

**Import**

```
import (
  "github.com/spatialcurrent/go-collector/collector"
)
```

**CollectFilepaths**

```
func CollectFilepaths(basepath string, extensions []string, recursive bool, paths []string) ([]string, error) {
  ...
}
```

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-collector/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
