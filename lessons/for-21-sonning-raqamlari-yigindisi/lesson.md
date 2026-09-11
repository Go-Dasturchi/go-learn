# Sonning raqamlari yig'indisi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, For bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: N=3, nums=[123, 456, 789]
Output: 6, 15, 24
```

## TASK

N berilgan. N ta son kiriting va har birining raqamlari yig'indisini chiqaring.

Quyidagi funksiyani to'ldiring:

```go
func digitSum(N int, nums []int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: tashqi - sonlar uchun, ichki - raqamlar uchun
2. Raqamlarni ajratish: n % 10, n // 10
