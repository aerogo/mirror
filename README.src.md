# {name}

{go:header}

Reflect utilities for Go data types.

## API

### GetField

```go
field, dataType, value, err := mirror.GetField(movie, "Title")
```

```go
field, dataType, value, err := mirror.GetField(movie, "Director.Name")
```

```go
field, dataType, value, err := mirror.GetField(movie, "Actors[0].Name")
```

Using simple queries:

```go
field, dataType, value, err := mirror.GetField(movie, `Actors[Name="Tom Cruise"].Name`)
```

You can not use the `.` character in query strings yet. This is a bug that needs to be fixed sometime.

{go:footer}
