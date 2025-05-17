# cacherefresh

This is a simple library for small cache which is refreshed concurrently.

`go get github.com/yonesko/cacherefresh@latest`

```go
c, err := cacherefresh.New(func() (string, error) {
	        //load from db and return the whole set
			return "Hello", nil
		}, time.Hour)

if err != nil {
    log.Fatal(err)
}

data, err := c.Get()
if err != nil {
    t.Fatal(err)
}

fmt.Println(data)

```