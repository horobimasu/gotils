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
1. `ExpandPath(path)` expands the env vars in `path` wrapped in `%%` and `~` to `%USERPROFILE%` if it is the first character

# exec
1. `RunCommand(program, args)` runs `program` with `args`
2. `StartCommand(program, args)` returns the cmd struct executing `program` with `args`
3. `GetCommandOutput(program, args)` runs `program` with `args` and gets the trimmed output

# crypt
1. `EncryptAesGcm(data, key)` encrypts `data` with `key` using AES-GCM
2. `DecryptAesGcm(data, key)` decrypts `data` with `key` using AES-GCM
