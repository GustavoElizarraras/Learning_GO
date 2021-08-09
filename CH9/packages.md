
- Repository : Version control.
# Module
Root of a Go library or app stored in a repo. It consists of one or more packages, which give organization and structure.
Every module has a globally unique identifier, in Go this is the path to the repo.

## go.mod
We need this file to indicate is a module, it is located in the rrot directory. To create we use _go mod init MODULE\_PATH_. The file created includes the Go version and requirements for the module.

## Organizing the module
- When the module is small, keep it all together. 
- When the project grows, create a _cmd_ folder at the root of the module and within it, create a folder for each binary built from the module. 
- When the project contains many files (shell scripts, docker, etc), all the Go code or logic should be in a folder name _pkg_  (besides the main packages under cmd).
    
    Limit the dependencies between packages, a common patter is to organize the code by slices of functionality.

## Rename gracefully
If a rename of some exported identifiers is needed or a move between packages, don't remove the originals, provide an alternate name instead, like a function or method that calls the original. For a constant, declare a new one of the same type and value, but with different name (an alias), if a new method or change of fields are needed, the change must be done in the original; if you want an alias type from another module, you can't use an alias to refer to the unexported methods and fields of the original
# Builing Packages
## Import
Allows to access exported constants, variables, functions and types in another package; all of these are called identifiers. Go uses **capitalization** to a package level identifier is visible outside of the package.

It is a compile-time error to import a package and not use it, this ensures the binary only contains used code. Absolute import paths clarify what you are importing and make easier to refactor the code

The name of a package is determined by its package clause (_package main_) not its import _path_. As a general rule, the package name should match the directorys name.If you use the same package in two different files, the import must be in both.

## Naming
Meaningful names. A practical example is the following:

The `util` package has to functions: `ExtractNames` and `FormatNames`It is clearer this way `extract.Names` and `format.names`, the first word is the package and the next one is a function. 

### Overriding a package name
Let's take two random packages `crypto/rand` and `math/rand`, one of these must have an alternate name in the current file, inside the import: `"crand crypto/rand"`

## internal Package
When you want to share an identifier within packages in the module, use the `internal` package name. This makes visible that package with the parent directory and the rest of the directories in the same level.

## init function
**Avoid if possible**. This sets up a state in the code and it is possible to have many in package level. Some packages uses them to register the database driver and image formats. 

Any package-level variables configured via `init`, should be _effectively inmmutable_. But Go does not have a way to enforce this, so it is better to put that state of the program inside a struct that is initialized and returned by a function in the package.

## Circular dependencies
Go does not allow this, if package **A** imports package **B** directlyor indirectly, **B** can't import **A**. This is caused by splitting packages to thin, maybe a merge could be done or move only the dependant parts of the code into a single one.

