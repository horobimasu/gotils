# go-utils
random utilities i find myself using across my golang projects

# install:
```bash
go get github.com/horobimasu/gotils
```

# file
1. `CopyFile(srcPath, dstPath)` copies and file from `srcPath` to `dstPath`
2. `ReplaceFileBytesExact(filePath, toReplace, replaceWith)` replaces `toReplace` with `replaceWith`, they have to be exactly the same length or the function will return an error
3. `ReplaceFileBytesFromLeading(filePath, leadingBytes, replaceWith)` replaces all bytes starting from the first `leadingBytes` to the length of `replaceWith`

# parse
1. `ExpandPath(path)` expands env vars wrapped in `%%` and `~` to `%USERPROFILE%`
