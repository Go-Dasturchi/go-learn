# Number of Sub-arrays With Odd Sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: arr = [1,3,5]
Output: 4
Tushuntirish: [1], [3], [5] va [1,3,5] — barchasi toq yig'indili qism-massivlar.
```

## TASK

Sizga butun sonlardan iborat `arr` massivi berilgan. Yig'indisi **toq** (odd) bo'lgan barcha qism-massivlarning sonini toping.

Natijani `10^9 + 7` ga bo'lingandagi qoldig'ini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(arr []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir qism-massivni alohida tekshirish shart emas — prefiks yig'indilarning juft/toqligini kuzatib borsangiz, qism-massiv yig'indisining juft yoki toqligini tezda aniqlash mumkinligini o'ylab ko'ring.
2. Joriy prefiks yig'indisining juft yoki toq ekanligini hisoblab boring va shu vaqtgacha uchragan juft hamda toq prefikslar sonini alohida sanang; joriy prefiks juft bo'lsa, undan oldingi toq prefikslar soni javobga qo'shiladi (chunki ular orasidagi farq toq bo'ladi), joriy prefiks toq bo'lsa, undan oldingi juft prefikslar soni qo'shiladi — natijani modul bo'yicha yig'ib boring.
