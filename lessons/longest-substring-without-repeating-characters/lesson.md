# Longest substring without repeating characters

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: s = "abcabcbb"
Output: 3
Tushuntirish: Javob "abc", uning uzunligi 3.
```

## TASK

`s` satri berilgan. Hech qanday belgi takrorlanmaydigan eng uzun qism-satrning uzunligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir qism-satrni alohida tekshirish shart emas — o'zgaruvchan uzunlikdagi oyna (sliding window) va har bir belgining oxirgi ko'rilgan pozitsiyasini saqlashni o'ylab ko'ring.
2. Har bir belgining oxirgi uchragan indeksini map da saqlang; oynani o'ngga kengaytirib boring, agar joriy belgi oyna ichida oldin uchragan bo'lsa, oynaning chap chegarasini o'sha belgidan keyingi pozitsiyaga sakratib o'tkazing va har qadamda joriy oyna uzunligi bilan eng yaxshi natijani solishtiring.
