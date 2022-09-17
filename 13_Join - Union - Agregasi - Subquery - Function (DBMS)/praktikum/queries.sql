-- Insert - Select - Update - Delete

-- 1. Perintah Insert
-- a. Insert 5 operators ada table operators.
INSERT INTO operators (name) VALUES
("XL"),
("Telkomsel"),
("Axis"),
("Indosat"),
("3 (Three)");

-- b. Insert 3 product type
INSERT INTO product_types (name) VALUES
("Pulsa"),
("Paket Data"),
("Kartu Perdana");

-- c. Insert 2 product dengan product type id = 1, dan operators id = 3
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
(1, 3, "P-1", "Pulsa Axis 5.000", 100),
(1, 3, "P-2", "Pulsa Axis 10.000", 100);

-- d. Insert 3 product dengan product type id = 2, dan operators id = 1
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
(2, 1, "PD-1", "Paket Data XL 2GB + 1GB", 100),
(2, 1, "PD-2", "Paket Data XL 4GB + 1.5GB", 100),
(2, 1, "PD-3", "Paket Data XL 8GB + 3GB", 100);

-- e. Insert 3 product dengan product type id = 3, dan operators id = 4
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
(3, 4, "kp-1", "Kartu Perdana Indosat Bonus Gratis Telepon 60 Menit", 100),
(3, 4, "kp-2", "Kartu Perdana Indosat Bonus Paket Data 4GB", 100),
(3, 4, "kp-3", "Kartu Perdana Indosat Bonus Paket Data 6GB", 100);

-- f. Insert product description pada setiap product
INSERT INTO product_descriptions (product_id, description) VALUES
(1, 'Pulsa Axis senilai 5.000 10% lebih murah khusus pengguna baru'),
(2, 'Pulsa Axis senilai 10.000 10% lebih murah khusus pengguna baru'),
(3, 'Paket data XL beli 2GB bonus 1GB khusus area pulau Jawa dan Bali'),
(4, 'Paket data XL beli 4GB bonus 1.5GB khusus area pulau Jawa dan Bali'),
(5, 'Paket data XL beli 8GB bonus 3GB khusus area pulau Jawa dan Bali'),
(6, 'Kartu perdana Indosat bonus gratis telepon 60 menit ke semua operator'),
(7, 'Kartu perdana Indosat bonus paket data 4GB khusus pelanggan area Jawa dan Bali'),
(8, 'Kartu perdana Indosat bonus paket data 6GB khusus pelanggan area Jawa dan Bali');

-- g. Insert 3 payment methods
INSERT INTO payment_methods (name, status) VALUES
("Gopay", 100),
("Dana", 100),
("ShopeePay", 100);

-- h. Insert 5 user pada tabel users
INSERT INTO users (name, status, dob, gender) VALUES
('Arvin Paundra', 100, '2002-03-26', 'M'),
('John Cena', 100, '2004-09-08', 'M'),
('Marry Jane', 100, '2002-05-17', 'F'),
('Yayan Ruhian', 100, '1970-01-01', 'M'),
('Diana Thamrin', 100, '2008-02-28', 'F');

-- i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(1, 1, 'success', 1, '7000'),
(1, 2, 'pending', 1, '25000'),
(1, 3, 'success', 1, '48000'),
(2, 1, 'success', 1, '7000'),
(2, 2, 'success', 1, '38000'),
(2, 3, 'pending', 1, '12000'),
(3, 1, 'failed', 1, '56000'),
(3, 2, 'success', 1, '15000'),
(3, 3, 'pending', 1, '20000'),
(4, 1, 'success', 1, '38000'),
(4, 2, 'success', 1, '12000'),
(4, 3, 'success', 1, '7000'),
(5, 1, 'pending', 1, '56000'),
(5, 2, 'pending', 1, '38000'),
(5, 3, 'failed', 0, '48000');

-- j. Insert 3 product di masing-masing transaksi
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
(1, 1, 'success', 1, '48000'),
(2, 5, 'pending', 1, '25000'),
(3, 7, 'success', 1, '48000'),
(4, 1, 'success', 1, '7000'),
(5, 6, 'success', 1, '38000'),
(6, 2, 'pending', 1, '12000'),
(7, 8, 'failed', 1, '56000'),
(8, 3, 'success', 1, '15000'),
(9, 4, 'pending', 1, '20000'),
(10, 6, 'success', 1, '38000'),
(12, 1, 'success', 1, '7000'),
(13, 8, 'pending', 1, '56000'),
(14, 6, 'pending', 1, '38000'),
(15, 7, 'failed', 1, '48000');

-- 2. Perintah Select
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki atau M
SELECT name FROM users WHERE gender = "M";

-- b. Tampilkan product dengan id = 3
SELECT * FROM products WHERE id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung huruf 'a'
SELECT * FROM users WHERE created_at BETWEEN "2022-09-08 00:00:00" AND "2022-09-15 23:59:59" AND name LIKE "%a%";

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan
SELECT COUNT(*) AS total FROM users WHERE gender = "F";

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY name ASC;

-- f. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;

-- 3. Perintah Update
-- a. Ubah data product id 1 dengan nama 'product dummy'
UPDATE products SET name = "product dummy" WHERE id = 1;

-- b. Update qty = 3 pada transaction detail dengan product id = 1
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- Perintah Delete
-- a. Delete pada data table product dengan id = 1
DELETE FROM products WHERE id = 1;

-- b. Delete pada tabel product dengan product type id = 1
DELETE FROM products WHERE product_type_id = 1;

-- Join - Union - Subquery - Function

-- 1. Gabungkan data transaksi dari user id 1 dan user 2
SELECT * FROM transactions WHERE user_id = 1
UNION
SELECT * FROM transactions WHERE user_id = 2;

-- 2. Tampilkan jumlah harga transaksi user id 1
SELECT SUM(total_price) AS jumlah_harga FROM transactions WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2
SELECT COUNT(*) AS total_transaksi FROM transaction_details td
INNER JOIN products p ON td.product_id = p.id
WHERE p.product_type_id = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
SELECT p.* pt.name FROM products p
LEFT JOIN product_types pt ON p.product_type_id = pt.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user
SELECT t.*, p.name AS product_name, u.name AS user_name FROM transactions t
LEFT JOIN transaction_details td ON t.id = td.transaction_id
RIGHT JOIN products p ON td.product_id = p.id
RIGHT JOIN users u ON t.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud
-- Fungsi trigger
DELIMITER $$
CREATE TRIGGER delete_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE v_transaction_id INT;
SET v_transaction_id = OLD.id;
DELETE FROM transaction_details WHERE transaction_id = v_transaction_id;
END; $$

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus
DELIMITER $$
CREATE TRIGGER update_total_qty_transactions
BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN
DECLARE v_transaction_id, v_qty INT;
SET v_transaction_id = OLD.transaction_id;
SET v_qty = OLD.qty;
UPDATE transactions SET total_qty = total_qty - v_qty WHERE id = v_transaction_id;
END; $$

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan subquery
SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details);