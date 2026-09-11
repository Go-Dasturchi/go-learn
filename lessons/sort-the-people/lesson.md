# Sort the People

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: names = ["Mary","John","Emma"], heights = [180,165,170]
Output: ["Mary","Emma","John"]
Tushuntirish: Mary eng uzun (180), keyin Emma (170), va nihoyat John (165).
```

## TASK

Sizga bir xil uzunlikdagi `names` va `heights` massivlari berilgan. Har bir insonning ismi va bo'yi mos ravishda ifodalangan. `heights` dagi barcha qiymatlar takrorlanmas deb hisoblang.

Siz odamlarning ismlarini ularning bo'ylariga mos ravishda **kamayish tartibida (eng uzundan eng kaltagacha)** qaytarishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(names []string, heights []int) []string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ismlarni to'g'ridan-to'g'ri saralash o'rniga, indekslarni bo'ylarga qarab saralashni o'ylang.
2. `0` dan `len(names)-1` gacha indekslardan iborat massiv yarating va uni `heights` qiymatlariga qarab kamayish tartibida saralang, so'ng saralangan indekslar tartibida `names` dan mos ismlarni yig'ib natija hosil qiling.
