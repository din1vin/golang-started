# 4 ways define variable in Golang

## 1. no init value defined
```go
var number int64
fmt.Println("number = ", number)

--console
number =  0
```

## 2. type auto predicate
```go
var s = "hello"
fmt.Printf("s=%s, type of s is %T\n",s,s)

--console 
s=hello, type of s is string
```

## 3.  full define (type and initValue)
```go
var s string = "hello"
fmt.Println("s=",s)

--console
s=hello
```

## 4. popular define (without var)
```go
 e :=  100
 fmt.Println("e = ",e)
```

### Caution
**1. if defined global variable, only 1,2,3 worked**

**2. golang support multi variable defined in one line**