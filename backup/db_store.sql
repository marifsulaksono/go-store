-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 28, 2023 at 10:43 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_store`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`id`, `user_id`, `product_id`, `qty`) VALUES
(1, 3, 4, 3),
(2, 3, 36, 5),
(3, 3, 19, 96),
(4, 3, 3, 1),
(5, 1, 11, 7),
(6, 1, 29, 15);

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(3) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`) VALUES
(1, 'Elektronik'),
(2, 'Fashion'),
(3, 'Computer & Gadget'),
(4, 'Kitchen Set');

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE `clients` (
  `id` varchar(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `contact` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `name`, `contact`, `created_at`, `updated_at`, `deleted_at`) VALUES
('C001', 'Aka', '085331828630', NULL, '2023-06-04 10:50:15.112', NULL),
('C003', 'Yunus', '08741212345', NULL, NULL, NULL),
('C004', 'Nita', '08958218463', NULL, NULL, NULL),
('C005', 'Imaa', '08533159630', '2023-06-04 10:12:11.733', '2023-06-04 10:12:11.733', NULL),
('C006', 'Samsul', '08512543585', NULL, '2023-06-04 10:00:07.315', NULL),
('C007', 'Iqbal', '08747212952', '2023-06-03 12:33:13.973', '2023-06-03 12:33:13.975', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `employees`
--

CREATE TABLE `employees` (
  `id` int(11) NOT NULL,
  `name` varchar(150) NOT NULL,
  `address` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `isActive` varchar(10) NOT NULL DEFAULT 'Aktif'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `employees`
--

INSERT INTO `employees` (`id`, `name`, `address`, `role`, `isActive`) VALUES
(1, 'Arif', 'Kraksaan', 'Engineer', 'Aktif'),
(2, 'Aka', 'Kraksaan', 'Tech Lead', 'Aktif'),
(3, 'Samsul', 'Gending', 'Engineer', 'Aktif'),
(4, 'Munir', 'Gading', 'Java Engineer', 'Aktif'),
(5, 'Dias', 'Kraksaan', 'Operator', 'Nonaktif'),
(6, 'Roni', 'Brumbungan lor', 'Operator DZ', 'Aktif'),
(7, 'Dhoni', 'Paiton', 'Account Officer', 'Aktif'),
(8, 'Benny', 'Kregenan', 'Senior Backend', 'Aktif'),
(9, 'Sofyan', 'Bulu', 'Teknisi', 'Aktif'),
(10, 'Agung', 'Rondokuning', 'Frontliner', 'Nonaktif'),
(11, 'Danny', 'Maron', 'Senior Frontend', 'Aktif'),
(12, 'Bagus', 'Jabungsisir', 'Mekanik', 'Nonaktif'),
(15, 'Firman', 'Pajarakan', 'Perawat', 'Aktif'),
(16, 'Indra', 'Maron', 'Mobile Engineer', 'Aktif'),
(17, 'Maul', 'Bekasi', 'Enterprenuer', 'Aktif');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `stock` int(8) NOT NULL,
  `price` int(11) NOT NULL,
  `sold` int(11) DEFAULT NULL,
  `desc` text NOT NULL,
  `status` enum('sale','soldout') NOT NULL DEFAULT 'sale',
  `category_id` int(3) NOT NULL,
  `store_id` int(11) NOT NULL,
  `delete_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `stock`, `price`, `sold`, `desc`, `status`, `category_id`, `store_id`, `delete_at`) VALUES
(1, 'Monitor LG', 20, 2500000, 4, 'Monitor Bertipe LED dengan ukuran 21\'inch\r\nHemat daya listrik dan sangat ergonomis', 'sale', 1, 18, NULL),
(2, 'Keyboard R-ONE', 30, 145000, NULL, '', 'soldout', 3, 4, '2023-07-17 11:38:54'),
(3, 'Speaker Dolby', 50, 1245000, 2, 'Rasakan sensasi getaran suara yang membuatmu menikmati dunia', 'sale', 1, 18, NULL),
(4, 'PS3', 12, 1000000, 3, '', 'sale', 1, 18, NULL),
(6, 'Proyektor Epson', 30, 750000, 8, '', 'sale', 1, 18, NULL),
(7, 'Mouse', 18, 75000, 7, '', 'sale', 3, 4, NULL),
(8, 'Kaos', 25, 75000, NULL, '', 'soldout', 2, 1, '2023-07-26 12:51:24'),
(10, 'Jaket Kulit', 0, 75000, NULL, 'Jaket dengan bahan kulit berkualitas dan harga terjangkau', 'soldout', 2, 1, NULL),
(11, 'Tupperware', 90, 99000, 9, '', 'sale', 4, 1, NULL),
(12, 'Tas Kaliber', 12, 410000, 3, '', 'sale', 2, 1, NULL),
(14, 'Kursi Gaming', 13, 1500000, 2, '', 'sale', 1, 18, NULL),
(15, 'JBL Headphone', 20, 180000, NULL, '', 'sale', 1, 18, NULL),
(17, 'HP Vivo', 10, 2500000, NULL, '', 'soldout', 3, 4, '2023-07-17 13:10:08'),
(19, 'Kemeja Batik', 96, 75000, 4, '', 'sale', 2, 1, NULL),
(21, 'Jaket Bomber', 22, 100000, 3, '', 'sale', 2, 1, NULL),
(22, 'HP Nokia', 39, 450000, 1, '', 'sale', 3, 4, NULL),
(24, 'CCTV Dahua', 60, 4800000, NULL, '', 'sale', 1, 18, NULL),
(28, 'Alamo Botol', 90, 3500, 10, '', 'sale', 2, 1, NULL),
(29, 'Aqua Botol', 90, 3000, 10, '', 'sale', 3, 1, NULL),
(31, 'Le Mineral Botol', 90, 4000, 10, '', 'sale', 2, 1, NULL),
(32, 'Jam Tangan Rolex', 47, 300000, 3, '', 'sale', 2, 1, NULL),
(34, 'HP Iphone', 10, 15000000, NULL, '', 'soldout', 3, 4, '2023-07-17 13:10:26'),
(35, 'CCTV Vision', 36, 6300000, 4, '', 'sale', 1, 18, NULL),
(36, 'Charger Vivo', 35, 550000, 5, '', 'sale', 1, 18, NULL),
(37, 'Tas Eiger', 21, 360000, 4, '', 'sale', 2, 1, NULL),
(38, 'Tas Consina', 20, 315000, NULL, '', 'sale', 2, 1, NULL),
(39, 'Tas Adidas', 30, 255000, 10, '', 'sale', 2, 1, NULL),
(42, 'Seragam PNS', 50, 390000, 0, 'Seragam PNS harga terjangkau dan nyaman dipakai', 'sale', 2, 4, NULL),
(43, 'Seragam Siswa', 150, 330000, 0, 'Seragam siswa-siswi dengan harga terjangkau dan nyaman dipakai', 'sale', 2, 4, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `shipping_addresses`
--

CREATE TABLE `shipping_addresses` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `recipient_name` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `phonenumber` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `shipping_addresses`
--

INSERT INTO `shipping_addresses` (`id`, `user_id`, `recipient_name`, `address`, `phonenumber`) VALUES
(1, 2, 'Taslima', 'Perum Bulu Land Blok E4 Kraksaan', '85331888118'),
(2, 1, 'Arif', 'Surabaya', '85331824328'),
(3, 3, 'Nita', 'Malang', '85321824328'),
(4, 1, 'Taslima', 'Kraksaan, Probolinggo', '85331532850'),
(5, 5, 'Dika', 'Kraksaan', '82321456239');

-- --------------------------------------------------------

--
-- Table structure for table `stores`
--

CREATE TABLE `stores` (
  `id` int(11) NOT NULL,
  `name_store` varchar(255) NOT NULL,
  `address` text NOT NULL,
  `email` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `status` enum('active','inactive') NOT NULL,
  `user_id` int(11) NOT NULL,
  `create_at` date NOT NULL,
  `delete_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stores`
--

INSERT INTO `stores` (`id`, `name_store`, `address`, `email`, `description`, `status`, `user_id`, `create_at`, `delete_at`) VALUES
(1, 'Beli Beli Store', 'Ds. Bulu, Kec. Kraksaan, Kab. Probolinggo', 'belibelistore@gmail.com', 'Toko serbaguna yang menyediakan berbagai macam produk', 'active', 2, '2023-07-29', NULL),
(4, 'Nitnot Computer', 'Ds. Pesangrahan, Kec. Jangkar, Kab. Situbondo', 'nitnot.comp@gmail.com', 'Toko kami menyediakan peralatan dan aksesoris komputer', 'active', 3, '2023-07-29', NULL),
(18, 'Marfs Store Official', 'Ds. Kandangjati Kulon, Kec. Kraksaan, Kab. Probolinggo', 'marfs.official@gmail.com', 'Toko resmi yang memberikan kepercayaan dan kepuasan pelanggan', 'active', 1, '2023-07-31', NULL),
(20, 'Gatherloop Course', 'Perum New Kraksaan Land G16, Kec. Kraksaan, Kab. Probolinggo', 'official@gatherloop.co', 'Komunitas yang menawarkan kursus online', 'active', 4, '2023-08-16', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `date` datetime NOT NULL,
  `total` int(11) NOT NULL,
  `status` enum('waiting','shipping','cancel','success') NOT NULL,
  `shipping_address_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `delete_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `date`, `total`, `status`, `shipping_address_id`, `user_id`, `delete_at`) VALUES
(1, '2023-07-26 14:19:24', 10300000, 'waiting', 2, 1, NULL),
(3, '2023-07-27 14:04:28', 6891000, 'waiting', 1, 2, NULL),
(4, '2023-07-28 07:57:37', 2490000, 'waiting', 1, 2, NULL),
(5, '2023-07-28 17:40:59', 105000, 'waiting', 1, 2, NULL),
(6, '2023-07-28 17:45:28', 450000, 'waiting', 1, 2, NULL),
(7, '2023-07-28 17:48:18', 900000, 'waiting', 1, 2, NULL),
(8, '2023-08-04 02:18:56', 525000, 'waiting', 3, 3, NULL),
(9, '2023-08-04 02:48:41', 28140000, 'waiting', 3, 3, NULL),
(11, '2023-08-16 04:26:57', 1800000, 'waiting', 2, 1, NULL),
(12, '2023-08-25 15:13:09', 2550000, 'waiting', 3, 3, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transaction_items`
--

CREATE TABLE `transaction_items` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `subtotal` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_items`
--

INSERT INTO `transaction_items` (`id`, `transaction_id`, `product_id`, `qty`, `price`, `subtotal`) VALUES
(17, 1, 1, 4, 2500000, 10000000),
(18, 1, 19, 4, 75000, 300000),
(19, 3, 6, 8, 750000, 6000000),
(20, 3, 11, 9, 99000, 891000),
(21, 4, 3, 2, 1245000, 2490000),
(22, 5, 28, 10, 3500, 35000),
(23, 5, 29, 10, 3000, 30000),
(24, 5, 31, 10, 4000, 40000),
(25, 6, 22, 1, 450000, 450000),
(26, 7, 32, 3, 300000, 900000),
(27, 8, 7, 7, 75000, 525000),
(28, 9, 37, 4, 360000, 1440000),
(29, 9, 35, 4, 6300000, 25200000),
(30, 9, 14, 1, 1500000, 1500000),
(31, 11, 21, 3, 100000, 300000),
(32, 11, 14, 1, 1500000, 1500000),
(33, 12, 39, 10, 255000, 2550000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `name` longtext NOT NULL,
  `username` longtext NOT NULL,
  `password` longtext NOT NULL,
  `email` varchar(255) NOT NULL,
  `phonenumber` varchar(30) NOT NULL,
  `role` varchar(200) NOT NULL,
  `create_at` date NOT NULL,
  `update_at` date DEFAULT NULL,
  `delete_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `username`, `password`, `email`, `phonenumber`, `role`, `create_at`, `update_at`, `delete_at`) VALUES
(1, 'Muhammad Arif Sulaksono', 'arif', '$2a$10$gsscMbaLD5j3DwvPF91tdO3wCKvtp/Zn.2ETJ3BHYo6/LpLkb4g/C', 'arif@gmail.com', '8123456789', 'Admin', '2023-08-01', NULL, NULL),
(2, 'Taslima', 'imaa', '$2a$10$KbpMRpri3W3HHqCHFGJl8OSEps9EqshG3DVLvAKoeAk0EZ8rbGGhu', 'imaa@gmail.com', '8123456788', 'User', '2023-08-01', NULL, NULL),
(3, 'Dewi Yunita Anjaswati', 'nita', '$2a$10$upvNFWqGwFVuUfiUmmUjnOfaSla5bEr7N1.tPNgQ38Aem/fJgrtpm', 'nita@gmail.com', '8123456787', 'User', '2023-08-01', NULL, NULL),
(4, 'M Nindra zaka', 'aka', '$2a$10$sxXXVMCHQd9EsOVz9zMPzO5S/6c4P4Z20TF6hYqH9aRc4OWfrw7uO', 'aka@gmail.com', '85468932461', 'buyer', '2023-08-15', '0000-00-00', NULL),
(5, 'Tri Andika Pranata Wijaya', 'dika', '$2a$10$WsEUKgnLP.EDgYfEXE4nvuYKZ64WWQqolExC7qEEDMm2Iiz4cb9am', 'dika@gmail.com', '85467567366', 'buyer', '2023-08-24', '0000-00-00', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `clients`
--
ALTER TABLE `clients`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_clients_deleted_at` (`deleted_at`);

--
-- Indexes for table `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idCategory` (`category_id`);

--
-- Indexes for table `shipping_addresses`
--
ALTER TABLE `shipping_addresses`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stores`
--
ALTER TABLE `stores`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction_items`
--
ALTER TABLE `transaction_items`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`) USING HASH;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `carts`
--
ALTER TABLE `carts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `employees`
--
ALTER TABLE `employees`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=44;

--
-- AUTO_INCREMENT for table `shipping_addresses`
--
ALTER TABLE `shipping_addresses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `stores`
--
ALTER TABLE `stores`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `transaction_items`
--
ALTER TABLE `transaction_items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_categoryitem` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
