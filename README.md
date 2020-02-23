![Travis (.org)](https://img.shields.io/travis/gastrodon/booru?label=%E2%80%8B&logo=travis)
![Codecov](https://img.shields.io/codecov/c/github/gastrodon/booru?label=%E2%80%8B&logo=codecov)

# Booru

A library for interacting with \*booru sites using go

### Running tests

Before running tests, an account that may log in to danbooru is needed.
Their username should be set to the environment variable `BOORU_LOGIN`,
and their api token to `BOORU_KEY`

### Searching

Function level docs will not cover search parameters that many search types share.
Instead they will be documented here. More information about them can be found by reading the
[danbooru api docs](https://danbooru.donmai.us/wiki_pages/help:api)

```
page:    search page offset, starts at 1
         Can also use b<id> or a<id> to search
         before and after post ids, respectively

limit:   maximum number of results per page

random:  Return results in random order?
```

References to these will still be in docstrings, but not with full documentation
