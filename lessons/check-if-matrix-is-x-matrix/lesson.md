# Check if Matrix Is X-Matrix

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: grid = [[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
Output: true
Tushuntirish: Diagonallar: 2,3,2,2 va 1,1,5,4. Barchasi > 0. Boshqa barcha elementlar 0.
```

## TASK

Sizga `n x n` o'lchamli butun sonlardan iborat kvadratik `grid` matritsasi berilgan.

Matritsa **X-matritsa** hisoblanishi uchun quyidagi ikkala shart bajarilishi kerak:

1. Matritsaning har ikkala diagonalidagi barcha elementlar **nol bo'lmasligi** kerak.
2. Diagonallardan tashqaridagi barcha elementlar **nol bo'lishi** kerak.

Agar `grid` X-matritsa bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(grid [][]int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
