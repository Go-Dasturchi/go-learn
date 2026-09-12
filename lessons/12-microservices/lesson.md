# 12 — Microservices

## THEORY

Shu paytgacha ko'rgan darslarning aksariyati, bitta dastur (**monolit**) ichida ishlaydi: bitta binary, bitta baza, bitta deploy. Dastur kichik bo'lganda, bu — eng oddiy va to'g'ri yondashuv. Lekin dastur va jamoa kattalashgani sari, muammolar chiqadi: kichik o'zgarish uchun **butun** dasturni qayta joylashtirish kerak, bitta qismdagi xato **hammasini** yiqitishi mumkin, va turli jamoalar bir xil kod ustida ishlashga majbur bo'ladi.

**Mikroservis arxitekturasi** — bitta katta dastur o'rniga, har biri **bitta vazifaga** javobgar bo'lgan, mustaqil ishga tushiriladigan **kichik xizmatlar** to'plami quradi. Masalan:

```
┌────────────────┐     ┌────────────────┐     ┌────────────────┐
│ Foydalanuvchi   │     │ Buyurtma        │     │ Bildirishnoma   │
│ Xizmati         │◄────┤ Xizmati         │────►│ Xizmati         │
│ (HTTP, port     │     │ (HTTP, port     │     │ (Kafka          │
│  8081)          │     │  8082)          │     │  consumer)      │
└────────────────┘     └────────────────┘     └────────────────┘
```

Har bir xizmat — **o'z bazasiga**, **o'z Docker konteyneriga**, va odatda **o'z jamoasiga** ega bo'ladi. Xizmatlar bir-biri bilan, oldingi darslarda ko'rgan usullar orqali gaplashadi: **HTTP/REST** ("HTTP & Backend"), **gRPC** ("gRPC"), yoki **Kafka** orqali async xabar almashish ("Kafka").

**Nega bu qiyin — tarqoq tizim muammolari.** Monolitda, bir funksiyani chaqirish deyarli hech qachon muvaffaqiyatsiz bo'lmaydi. Lekin mikroservislarda, **tarmoq orqali** so'rov yuborilgani uchun, boshqa xizmat **ishlamay qolishi**, **sekin javob berishi**, yoki **umuman mavjud bo'lmasligi** mumkin. Shuning uchun, har bir tarmoq chaqiruvi, **xatolikni kutish** bilan yozilishi kerak:

```go
func BuyurtmaYaratish(foydalanuvchiXizmati FoydalanuvchiXizmatiKlienti, foydalanuvchiID int) error {
	foydalanuvchi, err := foydalanuvchiXizmati.Ol(foydalanuvchiID)
	if err != nil {
		return fmt.Errorf("foydalanuvchi xizmatiga ulanib bo'lmadi: %w", err)
	}
	// ... buyurtma yaratish mantig'i
	return nil
}
```

**Bu — "Clean Architecture" va "Dependency Injection" darslarining, tarmoq chegarasiga qo'llanilishi.** `BuyurtmaXizmati`, boshqa xizmatga **to'g'ridan-to'g'ri** emas, balki **interfeys** (`FoydalanuvchiXizmatiKlienti`) orqali murojaat qiladi — bu, haqiqiy tarmoq bo'lmagan holatda ham (masalan, testda) uni **soxta klient** bilan sinash imkonini beradi:

```go
type FoydalanuvchiXizmatiKlienti interface {
	Ol(id int) (Foydalanuvchi, error)
}
```

**Nega mikroservislarga darhol o'tish kerak emas.** Ko'plab muvaffaqiyatli kompaniyalar, dasturlarini **monolit** sifatida boshlaydi, va faqat jamoa/tizim kattalashgach, kerakli qismlarni **alohida xizmatlarga ajratadi**. Erta boshlangan, ortiqcha murakkab mikroservis arxitekturasi — ko'pincha, hal qilingandan ko'ra ko'proq muammo tug'diradi ("hurmatsiz murakkablik" — accidental complexity).

**Terminalda ko'rish — bir nechta xizmatni birga ishga tushirish** ("Docker Compose" darsini eslang):

```bash
docker compose up  # foydalanuvchi-xizmati, buyurtma-xizmati, va bildirishnoma-xizmati birga ishga tushadi
```

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

type Foydalanuvchi struct {
	ID  int
	Ism string
}

type FoydalanuvchiXizmatiKlienti interface {
	Ol(id int) (Foydalanuvchi, error)
}

type SoxtaKlient struct {
	malumotlar map[int]Foydalanuvchi
}

func (k *SoxtaKlient) Ol(id int) (Foydalanuvchi, error) {
	f, bor := k.malumotlar[id]
	if !bor {
		return Foydalanuvchi{}, errors.New("foydalanuvchi xizmati: topilmadi")
	}
	return f, nil
}

func BuyurtmaYaratish(klient FoydalanuvchiXizmatiKlienti, foydalanuvchiID int) (string, error) {
	f, err := klient.Ol(foydalanuvchiID)
	if err != nil {
		return "", fmt.Errorf("foydalanuvchi xizmatiga ulanib bo'lmadi: %w", err)
	}
	return fmt.Sprintf("Buyurtma %s uchun yaratildi", f.Ism), nil
}

func main() {
	klient := &SoxtaKlient{malumotlar: map[int]Foydalanuvchi{1: {ID: 1, Ism: "Ali"}}}

	natija, err := BuyurtmaYaratish(klient, 1)
	fmt.Println(natija, err)

	_, err = BuyurtmaYaratish(klient, 99)
	fmt.Println(err)
}
```

Natija:

```
Buyurtma Ali uchun yaratildi <nil>
foydalanuvchi xizmatiga ulanib bo'lmadi: foydalanuvchi xizmati: topilmadi
```

## TASK

`BuyurtmaYaratish(klient FoydalanuvchiXizmatiKlienti, foydalanuvchiID int) (string, error)` funksiyasini to'ldiring:

1. `klient.Ol(foydalanuvchiID)` ni chaqiring.
2. Agar xatolik bo'lsa, `fmt.Errorf("foydalanuvchi xizmatiga ulanib bo'lmadi: %w", err)` bilan **o'ralgan** xatolikni qaytaring (`%w` — "Custom Errors" darsidagi xatolikni o'rash naqshi).
3. Aks holda, `fmt.Sprintf("Buyurtma %s uchun yaratildi", f.Ism)` ni `nil` xatolik bilan qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `f, err := klient.Ol(foydalanuvchiID)`, keyin `if err != nil { return "", fmt.Errorf("foydalanuvchi xizmatiga ulanib bo'lmadi: %w", err) }`.
2. `return fmt.Sprintf("Buyurtma %s uchun yaratildi", f.Ism), nil`.
