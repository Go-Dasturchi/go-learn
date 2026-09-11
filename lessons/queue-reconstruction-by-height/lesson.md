# Queue Reconstruction by Height

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
```

## TASK

Sizga `people` massivi berilgan, bunda `people[i] = [hi, ki]` degani `i`-chi shaxsning balandligi `hi` va uning oldida kamida `ki` ta baland bo'lgan (yoki teng) shaxs turishi lozimligini bildiradi.

Navbatni qayta tartibla va natijani qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(people [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Balandligi kattaroq odamlarni avval joylashtirsangiz, ular pastroq bo'ylagilarning "oldida kamida k ta baland odam" shartiga ta'sir qilmasligini o'ylab ko'ring — bu greedy + qadamma-qadam qo'yish (insertion) usuliga olib keladi.
2. Odamlarni balandlik bo'yicha kamayish tartibida, balandlik teng bo'lganda esa `k` bo'yicha o'sish tartibida saralang; so'ng natija ro'yxatini bo'shdan boshlab, har bir odamni to'g'ridan-to'g'ri o'zining `k`-indeksiga (pozitsiyasiga) kiritib boring — chunki undan balandroq yoki teng bo'lganlar allaqachon joylashtirilgan.
