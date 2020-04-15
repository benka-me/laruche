#Manager

This package provide function for managing dependencies of bees and hives.
It's still on progress but work fine.

###Add dependencies to the bee
```go
func BeeAddDependencies(bee *laruche.Bee, request laruche.Namespaces) error
```

###Add dependencies to the hive
```go 
func HiveAddDependencies(hive *laruche.Hive, request laruche.Namespaces) error
```



test files are located on: pkg/test/manager_test.go