# Managing Go Dependencies

1. Install with go get i.e.

`
go get package 
`

2. Sync go.mod with GoLAnd IDE using context menu

3. Sync go.mod with Bazel WORKSPACE 

`
bazel run //:gazelle -- update-repos -from_file=go.mod
`
4. Generate or update build files 

`
bazel run //:gazelle
`

Double check for breaking changes 

```
bazel build
bazel build  //ib:go_default_library
bazel run    //example:main
```

Each of the targets should run just fine after having added new dependencies. 