# Max Sum of Rectangle No Larger Than K

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **50 daqiqa**

## EXAMPLE

```
Input: matrix = [[1,0,1],[0,-2,3]], k = 2
Output: 2
Tushuntirish: Eng katta yig'indiga ega bo'lgan to'rtburchak [[0, 1], [-2, 3]] dir va uning yig'indisi 0 + 1 - 2 + 3 = 2 ga teng. 2 <= 2.
```

## TASK

Sizga `m x n` o'lchamli butun sonlardan iborat `matrix` va butun son `k` berilgan.

Matritsa ichidan yig'indisi ko'pi bilan `k` bo'lgan (ya'ni `<=` `k`) va yig'indisi bo'lishi mumkin bo'lgan eng katta to'rtburchakni toping.
Natijada ana shu maksimal yig'indini qaytaring.

Shart shuki, har doim bunday yig'indiga ega bo'lgan kamida bitta to'rtburchak mavjud bo'ladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(matrix [][]int, k int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
