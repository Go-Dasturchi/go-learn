# 02 — Custom Errors

## THEORY

Oddiy `errors.New("...")` — bu shunchaki matn. Lekin ba'zan xatolik haqida ko'proq **ma'lumot** kerak bo'ladi — masalan, aynan qaysi qiymat noto'g'ri bo'lganini dastur ichida (matnni "parslashga" urinmasdan) bilish. Buning uchun o'zingizning **maxsus xatolik turingizni** yaratish mumkin.

**Eslang: `error` — bu shunchaki interfeys** ("Errors" darsida ko'rgan edingiz):

```go
type error interface {
	Error() string
}
```

Demak, `Error() string` methodiga ega bo'lgan **har qanday** struct — `error` sifatida ishlatilishi mumkin ("Interfaces" darsida ko'rgan implicit implementatsiyani eslang — hech qanday maxsus e'lon shart emas):

```go
type ManfiySonError struct {
	Qiymat int
}

func (e *ManfiySonError) Error() string {
	return fmt.Sprintf("manfiy son qabul qilinmaydi: %d", e.Qiymat)
}
```

Endi bu turdagi xatolikni oddiy xatolik kabi qaytarish mumkin:

```go
func ildiz(x int) (float64, error) {
	if x < 0 {
		return 0, &ManfiySonError{Qiymat: x}
	}
	return math.Sqrt(float64(x)), nil
}
```

**Nega bu foydali — struct sifatida qo'shimcha ma'lumot tashish.** Oddiy `errors.New("manfiy son")` bilan solishtiring — u yerda faqat matn bor, **qaysi son** manfiy ekanligi haqidagi ma'lumot yo'qolgan (matn ichiga uni "yozib qo'yish" mumkin, lekin uni dastur ichida qayta o'qib olish — matnni parslash — noqulay va ishonchsiz). Custom error bilan esa, chaqiruvchi tomon xohlasa xatolikni **struct**ga aylantirib, aynan `Qiymat` maydonining o'ziga to'g'ridan-to'g'ri murojaat qila oladi:

```go
_, err := ildiz(-9)
var manfiyErr *ManfiySonError
if errors.As(err, &manfiyErr) {
	fmt.Println("Noto'g'ri son:", manfiyErr.Qiymat) // 9 ni emas, -9 ni ko'rsatadi
}
```

`errors.As` — standart kutubxonadagi funksiya, u berilgan xatolik aslida qaysi **aniq turga** tegishli ekanligini tekshiradi va, agar mos kelsa, uni o'sha turga "aylantirib" beradi. Bu — katta dasturlarda turli xil xatolik turlarini bir-biridan ajratib, har biriga mos munosabatda bo'lish uchun keng qo'llaniladi.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

type ManfiySonError struct {
	Qiymat int
}

func (e *ManfiySonError) Error() string {
	return fmt.Sprintf("manfiy son qabul qilinmaydi: %d", e.Qiymat)
}

func ildiz(x int) (float64, error) {
	if x < 0 {
		return 0, &ManfiySonError{Qiymat: x}
	}
	return math.Sqrt(float64(x)), nil
}

func main() {
	natija, err := ildiz(16)
	fmt.Println(natija, err)

	_, err = ildiz(-9)
	fmt.Println(err)

	var manfiyErr *ManfiySonError
	if errors.As(err, &manfiyErr) {
		fmt.Println("Noto'g'ri son edi:", manfiyErr.Qiymat)
	}
}
```

Natija:

```
4 <nil>
manfiy son qabul qilinmaydi: -9
Noto'g'ri son edi: -9
```

## TASK

`ManfiySonError` struct'i (`Qiymat int` maydoni bilan, `Error() string` methodi allaqachon yozilgan) va `ildiz(x int) (float64, error)` funksiyasi berilgan. Funksiyani shunday to'ldiringki:

1. Agar `x < 0` bo'lsa, `0` va `&ManfiySonError{Qiymat: x}` qaytarsin.
2. Aks holda, `math.Sqrt(float64(x))` va `nil` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `&ManfiySonError{Qiymat: x}` — struct'ga pointer yaratib, uni xatolik sifatida qaytarasiz (`Error()` methodi pointer receiver bilan yozilgan).
2. `math` paketini import qilishni unutmang — `math.Sqrt` `float64` qabul qiladi, shuning uchun `x`ni avval `float64(x)` orqali aylantiring.
