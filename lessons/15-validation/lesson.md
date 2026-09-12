# 15 — Validation

## THEORY

Bank kassiri sizdan pul yechish so'rovini olganda, avval "hisobingizda yetarli mablag' bormi", "so'ralgan summa musbatmi" kabi narsalarni tekshiradi — noto'g'ri so'rovni **amalga oshirishga urinishdan oldin** rad etadi. **Validation (validatsiya)** — dasturlashda aynan shu: kiritilgan ma'lumotni **qayta ishlashdan oldin**, u "to'g'ri shaklda"ligini tekshirish.

**Nega validatsiya qilinmasa, xavfli.** Agar handler foydalanuvchidan kelgan ma'lumotni (masalan, "JSON Request Body" darsida ko'rgan `Decode`dan keyingi struct'ni) **tekshirmasdan** to'g'ridan-to'g'ri ishlatsa — bo'sh ism, manfiy yosh, yoki boshqa mantiqsiz qiymatlar dasturning qolgan qismiga "sizib kirib", kutilmagan xatolarga sabab bo'lishi mumkin.

**Validatsiya funksiyasi — bir nechta shartni tekshirib, tavsiflovchi xatolik qaytarish:**

```go
func foydalanuvchiniValidatsiyaQil(ism string, yosh int) error {
	if ism == "" {
		return errors.New("ism bo'sh bo'lishi mumkin emas")
	}
	if yosh < 0 || yosh > 150 {
		return errors.New("yosh 0 dan 150 gacha bo'lishi kerak")
	}
	return nil
}
```

Bu — "Error Handling" darsida ko'rgan naqshning davomi: har bir shart mos kelmasa, **aniq, tushunarli** xatolik matni bilan qaytariladi — bu, mijozga (yoki handler'ni chaqirgan boshqa kodga) **nima noto'g'ri ekanligini** to'g'ridan-to'g'ri bildiradi.

**API'da qanday ishlatiladi — validatsiya, ma'lumotni qayta ishlashdan OLDIN:**

```go
func royxatdanOtishHandler(w http.ResponseWriter, r *http.Request) {
	var kirish Foydalanuvchi
	json.NewDecoder(r.Body).Decode(&kirish)

	if err := foydalanuvchiniValidatsiyaQil(kirish.Ism, kirish.Yosh); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400 — mijoz xatosi
		return
	}

	// ... endi ma'lumot ishonchli, qayta ishlashda davom etish mumkin ...
}
```

**Nega validatsiya xatoligi har doim `400 Bad Request` bilan qaytariladi.** "HTTP Status Codes" darsida ko'rgan kod toifalarini eslang: validatsiya muvaffaqiyatsizligi — bu **serverning emas, mijozning** xatosi (u noto'g'ri ma'lumot yubordi), shuning uchun `4xx` toifasidagi kod, aniqrog'i `400`, mos keladi — `500` (server xatosi) emas.

**Bir nechta shartni tekshirish tartibi.** Odatda, birinchi topilgan xatolik **darhol** qaytariladi (yuqoridagi misoldagi kabi) — bu, foydalanuvchiga har safar faqat **bitta** muammoni ko'rsatib, uni tuzatgach, keyingisini ko'rsatishni anglatadi. Ba'zi tizimlar esa **barcha** xatoliklarni bir vaqtda yig'ib, ro'yxat qilib qaytaradi — bu murakkabroq, lekin foydalanuvchiga bir martada barcha muammolarni ko'rsatadi.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

func foydalanuvchiniValidatsiyaQil(ism string, yosh int) error {
	if ism == "" {
		return errors.New("ism bo'sh bo'lishi mumkin emas")
	}
	if yosh < 0 || yosh > 150 {
		return errors.New("yosh 0 dan 150 gacha bo'lishi kerak")
	}
	return nil
}

func main() {
	fmt.Println(foydalanuvchiniValidatsiyaQil("Ali", 25))
	fmt.Println(foydalanuvchiniValidatsiyaQil("", 25))
	fmt.Println(foydalanuvchiniValidatsiyaQil("Ali", 200))
}
```

Natija:

```
<nil>
ism bo'sh bo'lishi mumkin emas
yosh 0 dan 150 gacha bo'lishi kerak
```

## TASK

`foydalanuvchiniValidatsiyaQil(ism string, yosh int) error` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Agar `ism` bo'sh bo'lsa, `errors.New("ism bo'sh bo'lishi mumkin emas")` qaytarsin.
2. Aks holda, agar `yosh` `0`dan kichik yoki `150`dan katta bo'lsa, `errors.New("yosh 0 dan 150 gacha bo'lishi kerak")` qaytarsin.
3. Aks holda, `nil` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if ism == "" { return errors.New("ism bo'sh bo'lishi mumkin emas") }` bilan boshlang.
2. `if yosh < 0 || yosh > 150 { return errors.New("yosh 0 dan 150 gacha bo'lishi kerak") }`, so'ng `return nil`.
