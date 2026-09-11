# Number of Steps to Reduce a Number to Zero

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: num = 14
Output: 6
Tushuntirish:
1-qadam: 14 juft; 2 ga bo'lamiz va 7 hosil bo'ladi.
2-qadam: 7 toq; 1 ayiramiz va 6 hosil bo'ladi.
3-qadam: 6 juft; 2 ga bo'lamiz va 3 hosil bo'ladi.
4-qadam: 3 toq; 1 ayiramiz va 2 hosil bo'ladi.
5-qadam: 2 juft; 2 ga bo'lamiz va 1 hosil bo'ladi.
6-qadam: 1 toq; 1 ayiramiz va 0 hosil bo'ladi.
Jami 6 qadam.
```

## TASK

Sizga manfiy bo'lmagan butun `num` soni berilgan. Uni nolga (0 ga) kamaytirish uchun qilingan qadamlar sonini qaytaring.

Har bir qadamda, agar joriy son juft bo'lsa, uni 2 ga bo'lishingiz kerak, aks holda undan 1 ni ayirishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(num int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
