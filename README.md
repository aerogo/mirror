# mirror

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

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

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/aerogo/mirror?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/mirror
[report-image]: https://goreportcard.com/badge/github.com/aerogo/mirror
[report-url]: https://goreportcard.com/report/github.com/aerogo/mirror
[tests-image]: https://cloud.drone.io/api/badges/aerogo/mirror/status.svg
[tests-url]: https://cloud.drone.io/aerogo/mirror
[coverage-image]: https://codecov.io/gh/aerogo/mirror/graph/badge.svg
[coverage-url]: https://codecov.io/gh/aerogo/mirror
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
