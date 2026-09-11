# Gas Station

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Tushuntirish:
Siz 3-indeksdagi (4 litrli) shaxobchadan boshlaysiz.
Boringizdagi gaz = 4. 4-shaxobchaga borasiz: xarajat 1 litr. Qoldi 3, u yerdan +5 = 8.
Keyin 0-shaxobchaga: xarajat 2. Qoldi 6, +1 = 7.
Keyin 1-shaxobchaga: xarajat 3. Qoldi 4, +2 = 6.
Keyin 2-shaxobchaga: xarajat 4. Qoldi 2, +3 = 5.
Keyin yana 3-shaxobchaga: xarajat 5. Qoldi 0. Sikl tugadi.
```

## TASK

Aylanma marshrut bo'ylab `n` ta yoqilg'i shaxobchasi (gas station) joylashgan bo'lib, `i`-chi shaxobchada `gas[i]` litr benzin bor.

Sizning cheksiz yonilg'i bakiga ega bo'lgan avtomobilingiz bor va `i`-chi shaxobchadan `(i + 1)`-chi shaxobchaga borish uchun `cost[i]` litr benzin ketadi. Siz marshrutni ixtiyoriy bir shaxobchadan bo'sh bak bilan boshlaysiz.

Agar siz marshrut bo'ylab bir marta soat mili yo'nalishida aylanib chiqa olsangiz, boshlang'ich shaxobcha indeksini qaytaring. Agar buning iloji bo'lmasa, `-1` qaytaring. Agar yechim mavjud bo'lsa, u yagona (unique) bo'lishi kafolatlanadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(gas []int, cost []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir shaxobchani boshlang'ich nuqta sifatida sinab ko'rish shart emas — gaz va xarajat farqini kuzatib boruvchi greedy (ochko'z) yondashuvni o'ylab ko'ring.
2. Har bir i uchun diff = gas[i]-cost[i] ni umumiy yig'indiga qo'shib boring; agar joriy bakdagi yig'indi manfiy bo'lib qolsa, demak boshlanish nuqtasi o'sha yergacha bo'lgan hech qaysi shaxobcha bo'la olmaydi — boshlanishni i+1 ga ko'chirib bakni nolga qaytaring; oxirida agar umumiy yig'indi manfiy bo'lsa -1, aks holda topilgan boshlanish indeksi javob bo'ladi.
