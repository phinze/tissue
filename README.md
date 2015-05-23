# Tissue

A Github Issues workflow.

## Background

This is a tool I'm developing for myself to support my personal issues workflow.

## Usage

Add this to shell config:

```sh
which tissue > /dev/null && source <(tissue -sh)
```

Then `tj` ("tissue jump") will accept GitHub URLs, make dirs, drop some
metadata into them, and then `cd` you there.

## Behavior

### Issue URLs: Stage for Repro

 * Make a path with a descriptive title.
 * Drops `issue.markdown` with issue body contents.
 * Drops `issue.url` with issue URL.
 * `cd` to directory.

### PR URLs: Stage for Build / Test / Review

 * Save any WIP present in the source dir.
 * Fetch latest master so diff will be current.
 * Fetch and check out PR branch with descriptive name.


## Config

### `TISSUE_ISSUES_ROOT`

Root of paths created for issue repro. Defaults to `~/issues`.

### `TISSUE_GITHUB_{USER,TOKEN}`

Set these to authenticate the GitHub API requests. Useful for accessing private
repos and/or higher API limits.

### `TISSUE_DEBUG`

Set this to non-empty to see some logs on stderr.
