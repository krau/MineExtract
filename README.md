# MineExtract

**MineExtract** is a simple Go tool for extracting Minecraft assets files based on a JSON index.

It recreates the original directory structure by reading resource metadata, such as file paths and hashes, from the index file and locating the files in the `objects` directory.

## Features

- Parses `index.json` to determine the original file names and paths for Minecraft resource files.
- Extracts files from their hashed locations in `objects` into a structured directory.
- Preserves the correct directory and file structure required by Minecraft.

## Getting Started

Download the latest release from the [releases page](https://github.com/krau/MineExtract/releases) and extract the archive.

### Usage

```sh
./mineextract <path-to-index.json> <path-to-objects-directory> <output-directory>
```
