# mirror

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