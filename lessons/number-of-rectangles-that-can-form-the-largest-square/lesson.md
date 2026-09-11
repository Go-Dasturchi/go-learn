# Number of Rectangles That Can Form The Largest Square

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: rectangles = [[5,8],[3,9],[5,12],[16,5]]
Output: 3
Tushuntirish: Maksimal kvadrat tomoni: min(5,8)=5, min(3,9)=3, min(5,12)=5, min(16,5)=5. Maks 5. Uchta to'rtburchak 5x5 kvadrat hosil qila oladi.
```

## TASK

Sizga `rectangles` massivi berilgan, bunda `rectangles[i] = [li, wi]` i-chi to'rtburchakning uzunligi va eni.

`i`-chi to'rtburchakdan eng katta bo'lishi mumkin bo'lgan kvadratni hosil qilishingiz mumkin, va bu kvadratning tomonining uzunligi `min(li, wi)` ga teng.

Barcha to'rtburchaklardan hosil bo'lishi mumkin bo'lgan eng katta kvadratning tomonini aniqlang, keyin shu uzunlikdagi kvadrat hosil qila oladigan to'rtburchaklar sonini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(rectangles [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
