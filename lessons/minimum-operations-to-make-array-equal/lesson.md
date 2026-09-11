# Minimum Operations to Make Array Equal

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: n = 3
Output: 2
Tushuntirish: arr = [1, 3, 5]. Maqsad hamma element 3 bo'lsin. 1 ni 2 ga oshir, 5 ni 2 ga kamayt -> [3,3,3]. 2 operatsiya.
```

## TASK

`arr` massivi berilgan bo'lib, bunda `arr[i] = (2 * i) + 1` formula asosida belgilanadi (`0`-indekslanadi). Massiv uzunligi `n`.

Bir operatsiyada siz massivning ixtiyoriy bir elementini `x` ga oshirish va boshqa bir elementini `x` ga kamaytirishingiz mumkin.

Massivning barcha elementlarini teng qilish uchun nechta operatsiya kerakligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
