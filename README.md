# Goplays

This repository is for housing the source code of Playground examples I share around.

Read individual project [go.mods](./go.mod) for additional information.

## Running the example in Playground

Visit https://go.dev/play/p/zd9_cMCH187

## Generating Playground examples

This is the command to produce playground example contain all the files within the package structure:

```sh
find . -name '*.go' -or -name 'go.mod' | while read -r FILE; do
  echo "-- ${FILE/.\//} --"
  cat "$FILE"
done
```
