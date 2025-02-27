# Timeseries
A timeseries is a series of data points in ascending time order,
where "point" is a combination of a numeric value and unix timestamp.

Whereas our users send us "messy" timeseries data (duplicate timestamps, time between points not consistent, etc),
our query responses should be "clean": no duplicate timestamps, consistent timestamp deltas, etc.

So we need a function that converts messy timeseries data to clean timeseries data. It should also only return data in the range queried for.
It looks like:

```
func Fix(in []Point, start, end, interval uint32) []Point
```

* The input and output timeseries data is implemented as a slice of points, where each point is an instance of the included Point struct.
* The interval is a number of seconds such as 10, 60, etc.
* The start/end are unix timestamps.

## Input Data Characteristics
1) The points are always in increasing timestamp order.
2) The deltas between the timestamps are not always consistent.
3) There might be gaps (missing points).
4) It may include points that are outside the time range we want to select.

## Return Data Characteristics
1) Points must have timestamps in increments of the provided interval (in seconds).
2) Points must have timestamps that are multiples of the interval.
   For example for interval=60, some valid timestamps are 60, 120, 180, etc.
   If the input timestamp is not a multiple of the interval, it should be adjusted to the next one that is.
3) If multiple points correspond to the same timestamp, the first one wins.
4) Points must have a timestamp >= `start` and < `end`.
5) Points without values should use the value `math.NaN()`.

See the included code for scaffolding code and unit tests.
