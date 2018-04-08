package collector

import (
  "path"
  "strings"
  "os"
)

import (
  "github.com/mitchellh/go-homedir"
  "github.com/pkg/errors"
)

func CollectFilepaths(basepath string, extensions []string, recursive bool, paths []string) ([]string, error) {

  basepath_expanded, err := homedir.Expand(basepath)
  if err != nil {
    return nil, errors.New("Error: Could not expand home directory for path " + basepath + ".")
  } else {
    basepath = basepath_expanded
  }

  fi, err := os.Stat(basepath_expanded)
  if err != nil {
    return nil, errors.New("Error: Could not stat directory at " + basepath + ".")
  }
  switch mode := fi.Mode(); {
  case mode.IsRegular():
    return nil, errors.New("Error: File at path " + basepath + ".  Need directory.")
  }

  cwd, err := os.Open(basepath)
  if err != nil {
    return nil, errors.New("Error: Could not open directory at " + basepath + ".")
  }

  files_all, err := cwd.Readdir(0)
  if err != nil {
    return nil, errors.New("Error: Could not read directory at " + basepath + ".")
  }

	for _ , f := range files_all {
    if f.IsDir() {
      if recursive {
        paths, err = CollectFilepaths(basepath+"/"+f.Name(), extensions, recursive, paths)
        if err != nil {
          return nil, err
        }
      }
    } else {
      filename := path.Base(f.Name())
      valid := false
      for _ , ext := range extensions {
        valid = strings.HasSuffix(filename, "."+ext)
        if valid {
          break
        }
      }

      if valid {
        if basepath == "" || basepath == "/" {
          paths = append(paths, f.Name())
        } else if strings.HasSuffix(basepath, "/") {
          paths = append(paths, basepath+f.Name())
        } else {
          paths = append(paths, basepath+"/"+f.Name())
        }
      }
    }
  }

  return paths, err
}
