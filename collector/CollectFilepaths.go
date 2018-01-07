package collector

import (
  "errors"
  "log"
  "path"
  "strings"
  "os"
)

import (
  "github.com/ttacon/chalk"
  "github.com/mitchellh/go-homedir"
)

func CollectFilepaths(basepath string, extensions []string, recursive bool, paths []string) ([]string, error) {

  log.Println(chalk.Green, "Collecting filepaths from ", basepath, ".", chalk.Reset)

  basepath_expanded, err := homedir.Expand(basepath)
  if err != nil {
    return nil, errors.New("Error: Could not expand home directory for path " + basepath + ".")
  } else {
    basepath = basepath_expanded
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
        } else {
          paths = append(paths, basepath+"/"+f.Name())
        }
      }
    }
  }

  return paths, err
}