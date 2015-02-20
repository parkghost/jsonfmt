jsonfmt
=======
*A dead simple json formatter*

### Usage
```text
Usage of jsonfmt:
  -c=false: Compact json data to output
```

### Installation
```
go get github.com/parkghost/jsonfmt
```

### Example
```text
$ curl -s -o - http://api.openweathermap.org/data/2.5/weather?q=Taiwan,taipei | jsonfmt
{
    "base": "cmc stations",
    "clouds": {
        "all": 20
    },
    "cod": 200,
    "coord": {
        "lat": 25.04,
        "lon": 121.56
    },
    "dt": 1.424368508e+09,
    "id": 1.67572e+06,
    "main": {
        "grnd_level": 1026.75,
        "humidity": 100,
        "pressure": 1026.75,
        "sea_level": 1035.47,
        "temp": 290.292,
        "temp_max": 290.292,
        "temp_min": 290.292
    },
    "name": "Taipei City",
    "sys": {
        "country": "Taiwan",
        "message": 0.0169,
        "sunrise": 1.4242983e+09,
        "sunset": 1.424339408e+09
    },
    "weather": [
        {
            "description": "few clouds",
            "icon": "02n",
            "id": 801,
            "main": "Clouds"
        }
    ],
    "wind": {
        "deg": 118.503,
        "speed": 7.16
    }
}
```

License
---------------------

This project is licensed under the MIT license
