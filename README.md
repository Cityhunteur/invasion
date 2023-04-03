# invasion

Invasion is a program that simulates an alien invasion. It uses a map indicating
a list of cities and the roads connecting them.

The aliens can land randomly in any city, and can wander around randomly
following available roads to another city. Each should discover at least
10,000 cities unless it gets killed.

In case two aliens meet, they fight each other and ends up destroying each other
and the city as well.

The program ends if all the aliens are destroyed or all of the remaining
ones have discovered at least 10,000 cities.

## Requirements

* Go >= 1.20

## Usage
```shell
go run main.go --aliens 10 --map testdata/example.map
```
where
    `aliens` specify the number of aliens
    `map` file containing map of the world 

## Build

```shell
make build
```

## Test

```shell
make test
```

## Lint

```shell
make lint
```

## Run

```shell
make run
```

## map

Following is an example of a map

```text
Foo north=Bar west=Baz south=Quux
Bar south=Foo west=Corge
Corge east=Bar
```

## Assumptions

* City and alien names are unique
