# password


Hey there to run download the code and use the follow command:

go run main.go

### use the follow route:

POST http://llocalhost:1323/verify

entry example:
{
    "password": "Te111S",
    "rules": [
        {"rule": "minSize", "value":5},
        {"rule": "minUppercase", "value":2},
        {"rule": "minLowercase", "value":1},
        {"rule": "noRepeted", "value":0},
        {"rule": "minSpecialChars", "value":0 },
        {"rule": "minDigit", "value":3}
    ]
}

into rules put the follow rule that you want, and the value of option reference
of rule:

minSize,
minUppercase,
minLowercase,
noRepeted,
minSpecialChars,
minDigit

### test

cd src/service 

use command: go test
