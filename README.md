# 새김 언어 (Saegim Language)
'밑바닥부터 인터프리터 만들기 in go' 책을 통해 인터프리터 동작 원리를 학습한 후, 한글 프로그래밍 언어 '새김'을 제작했습니다.

## 지원 방언
- `.sg` - 표준어
- `.hbg` - 충청도 방언 버전

## 설치
```bash
git clone https://github.com/Jihyun3478/saegim-lang
cd saegim-lang
go build -o saegim
```

## 사용법
```bash
./saegim examples/hello.sg

./saegim examples/hello.hbg
```

## 예제
### 표준어 (.sg)
```
변수 이름 = "새김";
출력("안녕하세요, " + 이름);
```

### 충청도 (.hbg)
```
변수 이름 = "새김";
출력해유("안녕하세유, " + 이름);
```
