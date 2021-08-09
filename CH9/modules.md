# Modules

## Third-party code

Go compiles all code for your application in a single binary, own code or from third parties (imports) and any imports that are not already in the _go.mod_ file, are automatically updated in the file, it now includes the paths that contain the package and the version of the module; meanwhile, the _go.sum_ file is updated with the package name, its version, a hash of the module and a hash of the _go.mod_ file

## Versioning

Use version control tools to do this, for each version a branch is created.

### Semantic
The version are condsidered like this by convention: **vmajor.minor.patch**, the increments works this way: is a major version when there is no more compability with the previous one, is a **minor** when a new, backward-compatible feature is added (patch is set to 0 again), and is a **patch** when a bug is fixed.

### Working with versions
Go picks the latest version of a dependency. `go list`show what versions of the module are avaible, the `-m` flag list the modules, `-version` changes _go list_ to report on the avaible versions for the specified module. To downgrade we use the `go get` command and at the end of the package add: **@vma.mi.pa**, example: _go get github.com/learning-go-book/simpletax@v1.0.0_. _go.mod_ and _go.sum_ are updated too, they does not cause problems and we want to clean these files, we use `go mod tidy`.

### Minimum version selection
We will obtain the lowest version of a dependency that is declared to work in all of the _go.mod_ files across all of the dependencies.The import compability rule says that all minor and patch vesions of a module must be backward compatible, if not is a bug. In the case there are incompatibilities across modules, we need to contact the authors to fix it.

### Update to compatible versions
 
- Use `go get -u=patch IMPORT_PATH` to update the patch version.
- To especify a minor version use `go get IMPORT_PATH@vma.mi.pa` and then `go get -u=patch IMPORT_PATH` to update the patch version in that minor version.
- Use `go get -u IMPORT_PATH` to the most recent version.

### Update to incompatible versions
You can update the same package but with different major versions, Go consider this possible beacause internally they are different because the are incompatible between them.

## Vendoring
`go vendor` create a directory called _vendor_ at the top level of your module that contains all of the dependencies. If new are added, you need to run `go vendor` again to update. The downside is that the size is increases dramatically.

## Publishing
There are a lot of places where go packages are stored.
- The Go team created **pkg.go.dev** that indexes open source Go projects
- Module proxy servers are used to make safer the imports
- To use a private repo, set `GOPRIVATE=*.example.com/repo`

