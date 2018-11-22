# Go Transformer 

### Experimental transformer for Go that is similar to fractal

### Features

* Hide attribute
* Change the key name
* Custom value
* Nested Struct

From

```json
{
    "address":{
       "city":"Singapore",
       "postal_code":"770124",
       "country":"Singapore"
    },
    "age":"24",
    "firstName":"Nyan",
    "lastName":"Win"
}
```

To

```json
{
    "address":{
       "City":"Singapore",
       "Code":"770124"
    },
    "age":"24",
    "firstName":"Nyan",
    "fullName":"Nyan Win",
    "lastName":"Win"
}
```
