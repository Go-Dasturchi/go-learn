# Container with most water

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
```

## TASK

Uzunligi `n` bo'lgan massiv `height` berilgan. Undagi elementlar vertikal chiziqlar balandligini bildiradi. Qaysi ikkita chiziq eng ko'p suvni sig'dira olishini hisoblang va o'sha hajmni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(height []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha juftliklarni tekshirish shart emas — ikkita ko'rsatkichni massivning ikki chetidan boshlab, ularni asta-sekin ichkariga siljitish g'oyasini (two pointers) ko'rib chiqing.
2. l=0 va r=n-1 dan boshlang, har qadamda joriy maydonni (past chiziq balandligi × masofa) hisoblab eng yaxshisini yangilang, so'ng har doim ikkita chiziqdan pastrog'iga qarab tegishli ko'rsatkichni ichkariga suring — chunki balandroq chiziqni harakatlantirish natijani hech qachon yaxshilamaydi.
