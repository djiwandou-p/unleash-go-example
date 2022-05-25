# unleash-go-example
contoh sederhana menggunakan feature flag/toggle/switch dari [Unleash](https://www.getunleash.io/) dengan Go SDK

**Unleash**

adalah sistem yang digunakan untuk menyalakan-mematikan sebuah fitur/modul tanpa membutuhkan instalasi-deployment ulang ke server, menurut saya tujuan utamanya penggunaannya agar tim bisa lebih sering "rilis" dalam skala kecil di production server 

## Hasil akhir
https://user-images.githubusercontent.com/74530990/170183380-c8cda092-a6c8-4074-bba1-644a1062d443.mp4

### Langkah-langkah menggunakan Unleash Go SDK 
1. Cek koneksi ke Unleash demo server dan credentials, dengan cara run command berikut di terminal:
   
   ~~~
   curl https://app.unleash-hosted.com/demo/api/client/features \ -H "Authorization: 56907a2fa53c1d16101d509a10b78e36190b0f918d9f122d";
   ~~~


<img width="682" alt="Screen Shot 2022-05-25 at 06 27 34" src="https://user-images.githubusercontent.com/74530990/170263265-aeb1c73e-f34f-49f8-9f3b-012d588db46e.png">

   
2. Buat 1 toggle/feature flag di demo server Unleash https://app.unleash-hosted.com/demo/ , caranya:
* Setelah login -> Pilih `New feature toggle` 
![Screen Shot 2022-05-25 at 08 35 01](https://user-images.githubusercontent.com/74530990/170263957-a93d6b64-3ab4-4a82-85c3-48e35727dcfd.png)
* Lengkapi isian didalamnya (usahakan `Name` feature toggle nya unik ya) -> Klik `Save`
![Screen Shot 2022-05-25 at 08 42 58](https://user-images.githubusercontent.com/74530990/170264312-178e38a9-95d7-4c3f-b60d-660d1d2604c4.png)

3. Jalankan repo ini -> fork(opsional) -> clone dan run command berikut di terminal: 
~~~
go run main/main.go
~~~

4. Ubah status toggle/feature flag di dashboard demo server Unleash dan lihat `Hasil Akhir` diatas 👆


### Referensi
* https://docs.getunleash.io/user_guide/quickstart
* https://docs.gitlab.com/ee/operations/feature_flags.html
