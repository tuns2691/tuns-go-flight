--// Bảng lưu trữ thông tin chuyến bay
CREATE TABLE "flights" (
  "id" varchar PRIMARY KEY,	--ID
  "name" varchar(200) NOT NULL,	--Tên chuyến bay
  "from" varchar(10) NOT NULL,	--Từ địa điểm - cất cánh
  "to" varchar(10) NOT NULL,	--Đến địa điểm - hạ cánh
  "depart_date" timestamptz NOT NULL,	--Thời gian bay
  "status" varchar(10) NOT NULL,	--Trạng thái	(1: Hoạt động, 0: Không hoạt động)
  "available_slot" int NOT NULL,	-- Số ghế còn trống
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Ngày giờ tạo
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Ngày giờ cập nhật
);

--// Thông tin người đặt/quản lý người dùng
CREATE TABLE "customers" (
  "id" varchar PRIMARY KEY,	--ID
  "role" int,	--Role (0: Chưa đăng ký, 1: Đã đăng ký, 2: Admin)
  "name" varchar(200) NOT NULL,	--Họ & Tên
  "email" varchar(200) NOT NULL,	--Email
  "phone_number" varchar(20) NOT NULL,	--SĐT
  "date_of_bith" varchar(20) NOT NULL,	-- Ngày sinh
  "identity_card" varchar(20) NOT NULL,	--Căn cước công dân/Hộ chiếu/Chứng minh thư nhân dân
  "address" varchar(200) NOT NULL,	--Địa chỉ
  "membership_card"  varchar(20),	--Số thẻ hội viên
  "password" varchar(200),	--Password
  "status" int,	--Trạng thái (0: inactive, 1: Active)
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Ngày giờ tạo
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Ngày giờ cập nhật
);


--// Lưu thông tin booking
CREATE TABLE "bookings" (
  "id" varchar PRIMARY KEY,
  "customer_id" varchar NOT NULL,	--ID người đặt vé
  "flight_id" varchar NOT NULL,	--ID chuyến bay
  "code" varchar(20) NOT NULL,	--Mã đặt vé
  "booked_slot" int,	-- Số ghế booking
  "status" varchar(10) NOT NULL,	-- Trạng thái
  "booked_date" timestamp NOT NULL DEFAULT 'now()',	-- Thời gian đặt vé
  "created_at" timestamptz NOT NULL DEFAULT 'now()',	-- Ngày giờ tạo
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'	-- Ngày giờ cập nhật
);

ALTER TABLE "bookings" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bookings" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");
