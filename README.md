Issue with sysl parser:
```
Server:
    !type LoginData:
        @description = "blah nah brah"
        password <: sequence of string:
            @json_tag = "pwd"
```

we expect the field to have only one attribute "json_tag"
but it has two: "description" and "json_tag":

`map[description:s:"blah nah brah" json_tag:s:"pwd"]`

This can be seen in [main.go]()
