# fusion

A tool for pulling together multiple jsons through a CLI

## Initial Version

Roughly a day spread out across a few days, this is just the basic version.  It accepts a path which defaults to the current directory if none is specified

```
fusion ./my-json-files
```

# Tools used

* VSCode + Go Extension
* Go 1.24.4
* WSL2
* gitignore.io
* Github
* Gemini (my go is rusty and AI is amazingly good at reminding me and getting me back on the right track)

# Ideas for improvements

Just using this as a dumping ground for ideas I had during development/while thinking about the project

## CLI Improvements

Things I would improve in no particular order

* Making the json dynamic
* Filter e.g. --filter "vendor=boards r us,wifi=true" split on `,` then again on the `=` and get by key and filter by value during the `sortAndGatherMetrics`
* Output flag: --output-file=output.json
* Pretty Print Json
* Output Type
* Take/Limit/etc
* Recursive: -r
* Avoiding reading all the boards into memory

## Webapp Improvements

If I get there!

* Paging
* Basic Cache, if files haven't changed then use cache instead of having to rebuild each time
