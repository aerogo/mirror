# mirror
Reflect utilities for Go data types.

## API

### GetProperty

```go
dataType, value, err := mirror.GetProperty(movie, "Title")
```

```go
dataType, value, err := mirror.GetProperty(movie, "Director.Name")
```

```go
dataType, value, err := mirror.GetProperty(movie, "Actors[0].Name")
```