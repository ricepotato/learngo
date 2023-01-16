# learn go

## run

```bash
go run main.go
```

### function

```go
// 함수정의. func 키워드로 시작
// 타입을 지정해줘야 함, 뒤에만 타입을 지정해줘도 됨
func multiply(a, b int) int {
	return a * b
}
```

```go
// 함수가 여러개의 값을 리턴할 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
```
