# Noti Bok Duai (โนติ บอกด้วย)

โปรแกรมแจ้งเตือนโพสประกาศเช่าคอนโด จากเว็บ https://www.renthub.in.th/

## คุณสมบัติ
* แจ้งเตือนผ่าน Line Notify
* รันที่ไหนก็ได้
	* เขียนด้วย Go และ Compile เป็น exec file
	* Database เก็บเป็นไฟล์
	
## การติดตั้ง
ไฟล์ `config.yaml` แก้ `your line notify token` เป็น `line token`
```yaml
token: "your line notify token"
```
ไฟล์ `config.yaml` แก้ `link` เป็น url ที่คุณต้องการ
```yaml
renthub:
  link: [
    "https://www.renthub.in.th/search/condo_filter?search_price=monthly&condo_monthly_price%5Bprice_range_2%5D=1&temp%5Bzone_id%5D=115&locale=th"
  ]
```

## วิธีการได้ link
![Image of renthub](https://i.ibb.co/74h5HmT/Screenshot-from-2019-06-30-20-58-42.png)

## การใช้งาน

ถ้าต้องการรันเทส
```bash
$ ./ndb.run skip
```

ถ้าต้องการใช้รัน Notify
```bash
$ ./ndb.run noskip
```

ใส่ crontab เพื่อให้โปรแกรมรันทุกๆ 10 นาที
```bash
*/10 * * * * /home/kreang/nbd.run onskip
```

## การแจ้งเตือน
![Image of line](https://i.ibb.co/8BXyKV4/DF7-A59-A3-0-F59-4-D5-E-9910-92-D0-DD719947.jpg)