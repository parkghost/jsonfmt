jsonfmt
=======
*A dead simple json formatter*

### Usage
```text
Usage of jsonfmt:
  -c=false: Compact the json data
  -i="stdin": The input file
  -o="stdout": The output file
```

### Installation
```
go get github.com/parkghost/jsonfmt
```

### Example
```text
echo '{ "foo": "lorem", "bar": "ipsum" }' | jsonfmt
{
    "bar": "ipsum",
    "foo": "lorem"
}
```

License
---------------------

This project is licensed under the MIT license
