# Find Center of Star Graph

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: edges = [[1,2],[2,3],[4,2]]
Output: 2
Tushuntirish: Ko'rinib turibdiki, 2-tugun barcha boshqa tugunlarga (1, 3, 4) bog'langan. Demak, markaz - 2.
```

## TASK

Yulduzli graf (star graph) — bu `n` ta tugundan iborat yo'naltirilmagan graf bo'lib, markazdagi bitta tugun boshqa barcha tugunlarga to'g'ridan to'g'ri bog'langan.

Sizga ushbu graf qirralari `edges` ko'rinishida berilgan, qayerda `edges[i] = [ui, vi]` shu ikki tugun orasida qirra borligini bildiradi.

Yulduzli grafning markazini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(edges [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Markaziy tugun barcha qirralarda ishtirok etadi — shuning uchun uni topish uchun barcha qirralarni ko'rish shart emas, faqat dastlabki ikkita qirrani solishtirish kifoya.
2. Birinchi ikkita qirrani oling; ularda umumiy bo'lgan tugun (ya'ni ikkalasida ham uchraydigan tugun) — bu markaz, chunki markaz har bir qirrada qatnashadi.
