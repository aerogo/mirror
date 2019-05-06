# {name}

Reflect utilities for Go data types.

## API

### GetProperty

```go
field, dataType, value, err := mirror.GetProperty(movie, "Title")
```

```go
field, dataType, value, err := mirror.GetProperty(movie, "Director.Name")
```

```go
field, dataType, value, err := mirror.GetProperty(movie, "Actors[0].Name")
```

Using simple queries:

```go
field, dataType, value, err := mirror.GetProperty(movie, `Actors[Name="Tom Cruise"].Name`)
```

You can not use the `.` character in query strings yet. This is a bug that needs to be fixed sometime.