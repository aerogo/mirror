# mirror

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://www.patreon.com/eduardurbach)

[godoc-image]: https://godoc.org/github.com/aerogo/mirror?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/mirror
[report-image]: https://goreportcard.com/badge/github.com/aerogo/mirror
[report-url]: https://goreportcard.com/report/github.com/aerogo/mirror
[tests-image]: https://cloud.drone.io/api/badges/aerogo/mirror/status.svg
[tests-url]: https://cloud.drone.io/aerogo/mirror
[coverage-image]: https://codecov.io/gh/aerogo/mirror/graph/badge.svg
[coverage-url]: https://codecov.io/gh/aerogo/mirror
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
