# Count Operations to Obtain Zero

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: num1 = 2, num2 = 3
Output: 3
Tushuntirish:
1. num2 = 3, num1 = 2. num2 = 3 - 2 = 1.
2. num1 = 2, num2 = 1. num1 = 2 - 1 = 1.
3. num1 = 1, num2 = 1. num1 = 1 - 1 = 0.
Num1 0 bo'ldi, jarayon tugadi. Jami 3 ta amal.
```

## TASK

Sizga ikkita manfiy bo'lmagan butun son `num1` va `num2` berilgan.

Agar `num1 >= num2` bo'lsa, `num1` dan `num2` ni ayiring, aks holda `num2` dan `num1` ni ayiring.

Qachonki sonlardan biri `0` bo'lsa amallarni to'xtating. Shunda jami bajarilgan operatsiyalar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(num1 int, num2 int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masalada tasvirlangan jarayonni to'g'ridan-to'g'ri simulyatsiya qiling — bu Evklid algoritmiga o'xshaydi.
2. num1 va num2 ikkalasi ham nolga teng bo'lmaguncha, kattarog'idan kichikrog'ini ayirib, har safar hisoblagichni oshirib boring.
