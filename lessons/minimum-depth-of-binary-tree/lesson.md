# Minimum Depth of Binary Tree

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: depth_count = 3 (daraxt: [3,9,20,null,null,15,7])
Output: 2
Tushuntirish: [3 -> 9] minimal chuqurlik 2 (3 dan 9 gacha).
```

_Soddalashtirish: Biz bu yerda massiv berisak ham juda murakkab bo'ladi, shuning uchun biz n-size daraxt uchun minimal chuqurlikni matematik formula bilan hisoblaymiz._

## TASK

Ikkilik daraxtning ildizidan (root) eng yaqin barg tugunga (leaf node) qadar bo'lgan minimal chuqurlikni toping.

_Eslatma: Bu yerda daraxt massiv sifatida ifodalanadi (level-order traversal), bunda `null` bo'lmagan tugunlar ketma-ket keladi. Siz ushbu massiv yordamida chuqurlikni hisoblashingiz mumkin._

Siz minimal chuqurlikni hisoblashingiz kerak, bunda:

- Ildiz tugunning chuqurligi 1 ga teng.
- Har bir bolaning chuqurligi ota-onasiga nisbatan 1 ga katta.
- Barg tugun — bu bolasi bo'lmagan tugun.

Quyidagi funksiyani to'ldiring:

```go
func solve(leftChildren []int, rightChildren []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
