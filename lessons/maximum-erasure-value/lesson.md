# Maximum Erasure Value

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [4,2,4,5,6]
Output: 17
Tushuntirish: Eng yuqori ball [2,4,5,6] yig'indisiga teng: 2+4+5+6=17.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Siz massivdan takrorlanmaydigan elementlar bilan iborat **uzluksiz** qism-massivni o'chirish orqali ball yig'aysiz. Ball sifatida o'chirilgan elementlar **yig'indisi** hisoblanadi.

Bir ishda yig'ish mumkin bo'lgan maksimal ballni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. "Takrorlanmaydigan uzluksiz qism-massiv" degan shart o'zgaruvchan uzunlikdagi oyna (sliding window) qo'llashga ishora qiladi — oyna ichida qaysi sonlar borligini kuzatib boring.
2. Oyna ichidagi sonlarni saqlaydigan set (map) va joriy yig'indini yuriting; oynani o'ngga kengaytirganda agar shu son oynada allaqachon bo'lsa, chap chetdan sonlarni chiqarib (yig'indidan ayirib, setdan o'chirib) takrorlanish yo'qolguncha torayting, so'ng yangi sonni qo'shib joriy yig'indini eng yaxshi natija bilan solishtiring.
