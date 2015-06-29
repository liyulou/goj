# GO Json scanner

*goj* is a small low-level JSON scanning library.  It is representation-free, providing
no in-memory representation - that's your job.

*goj* may be useful to you if the following are true:

1. you need fast json parsing
2. you do not need a streaming parser (the distinct JSON documents you are parsing
   are delimited in some fasion)
3. you either want to extract a subset of JSON documents, or have your own data
   representation in memory, or wish to transform JSON into a different format.

## Example usage

installation:
```
go get github.com/lloyd/goj
```

A program to extract and print the top level `.name` property from a file containing
newline delimited json documents:
```

```

## Performance

Using the same JSON sample data as `encoding/json`, *goj* parses about 3x
faster than go's reflection based json parsing - which on my i7 desktop is
about 300mb/s vs 100mb/s.  See `test/bench_test.go` for the source.

Comparing against [jq](http://stedolan.github.io/jq/) is (a tiny and awesome tool
written in C that extracts nested values from json data), *goj* is more than 4x faster.

```
$ time go run example/main.go < 1.2gb_sample_data.json > /dev/null
real         0m5.712s
user         0m5.437s
sys          0m0.273s
$ time jq .name < 1.2gb_sample_data.json  > /dev/null
real         0m25.198s
user         0m25.017s
sys          0m0.180s
```
Comapred against [yajl](https://github.com/lloyd/yajl) (a fast streaming json parser
written in C) in a fair fight, *goj* is about 37% slower.

```
$ time json_verify -s < 1.2gb_sample_data.json
JSON is valid

real 0m3.560s
user 0m3.420s
sys  0m0.137s
$ time go run cmd/prof/main.go < 1.2gb_sample_data.json

real    0m4.897s
user    0m4.717s
sys     0m0.183s
```
