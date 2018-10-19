# Go-JsonError

Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec.

[Json API Spec](https://jsonapi.org/format/#errors)

## Usuage

In order to use jsonError you will need to initliaze the ErrorJson struct. Then when you would like to add an error you must pass in a ErrorComp struct to the AddError().

This follows the jsonapi spec where the error response must return an array.

Here is a basic example

```
	var err ErrorJSON

	errorComposition := ErrorComp{
		Detail: "this is a error message",
		Code:   "This is the code",
		Source: Source{
			Pointer: "/unit/tests",
		},
		Title:  "Title Test",
		Status: 200,
	}

	err.AddError(errorComposition)

	err.Error()
	err.ByteError()
```



## Running the tests

This package is just using the standard test packge included with go. You can run the test cases with

```
go test ./...
```


## Contributing

Please feel free to submit any pull requests.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/ddymko/go-jsonerror/tags).

## Authors

* **David Dymko**


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

