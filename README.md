# 🕒 Delayed HTTP Job Queue (Go)

**Delayed HTTP Job Queue**, belirlediğiniz bir adrese, belirlediğiniz zamanda HTTP isteği gönderen basit ve genişletilebilir bir zamanlı kuyruk sistemidir.

## 🚀 Ne İşe Yarar?

Bu servis sayesinde:

- Belirli bir **URL'ye**
- Belirli bir **zaman sonra**
- Belirli bir **HTTP metodu**, **header** ve **body** ile

otomatik olarak HTTP isteği gönderebilirsiniz.

## 🔧 Özellikler

- Gecikmeli HTTP istek planlama (`POST /jobs`)
- API Token doğrulama
- Job durumlarını takip edebilme (`pending`, `done`, `failed`, `cancelled`)
- Basit retry (yeniden deneme) desteği
- SQLite veya PostgreSQL desteği (esnek yapı)

## 📝 Yapılacaklar

- [X] İstenilen tam tarih ve saat getirilmesi
- [X] Postman örnek dosyası paylaşılması
- [X] Validation sisteminin geliştirilmesi
- [X] Tüm modüllere validation eklenmesi
- [X] Klasör yapısının best practice'e uygun hale getirilmesi
- [X] Middleware Hazırlanması tokensız giriş yapılmaması
- [X] Basit bir UI hazırlanması
- [ ] İşlerin token'a göre kaydedilmesi
- [ ] İşlerin token'a göre listelenmesi

## 📦 Kurulum


### Backend Kurulumu
```bash
git clone https://github.com/beytullahk54/quego.git
cd quego/backend
go run cmd/app/main.go
```

### Frontend Kurulumu
```bash
cd quego/frontend
yarn install
yarn dev
```
