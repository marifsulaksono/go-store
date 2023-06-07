-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 30, 2023 at 07:16 AM
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
(4, 'Munir', 'Gading', 'Engineer', 'Aktif'),
(5, 'Dias', 'Kraksaan', 'Operator', 'Nonaktif'),
(6, 'Roni', 'Brumbungan lor', 'Operator DZ', 'Aktif'),
(7, 'Dhoni', 'Paiton', 'Account Officer', 'Aktif'),
(8, 'Benny', 'Kregenan', 'Senior Backend', 'Aktif'),
(9, 'Sofyan', 'Bulu', 'Teknisi', 'Aktif'),
(10, 'Agung', 'Rondokuning', 'Frontliner', 'Nonaktif'),
(11, 'Danny', 'Maron', 'Senior Frontend', 'Aktif'),
(12, 'Bagus', 'Jabungsisir', 'Mekanik', 'Nonaktif'),
(13, 'Zuhri', 'Paiton', 'Koki', 'Nonaktif'),
(14, 'Joko', 'Kraksaan', 'Programmer', 'Aktif');

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `stock` int(8) NOT NULL,
  `price` int(11) NOT NULL,
  `isSale` int(1) NOT NULL COMMENT '1 = sale, 2 = not sale'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`id`, `name`, `stock`, `price`, `isSale`) VALUES
(1, 'Monitor LG', 25, 2500000, 1),
(2, 'Keyboard R-ONE', 30, 145000, 1),
(3, 'Speaker Dolby', 20, 2145000, 1),
(4, 'PS4', 10, 1000000, 1),
(5, 'Proyektor Epson', 8, 750000, 2),
(6, 'Proyektor Epson', 8, 750000, 1),
(7, 'Mouse', 25, 75000, 1),
(8, 'Kaos', 25, 75000, 1),
(9, 'Kaos kaki', 25, 75000, 2),
(10, 'Jaket', 40, 95000, 1),
(11, 'Tupperware', 99, 99000, 1),
(12, 'Tas Kaliber', 15, 410000, 1),
(13, 'Tas Palazzo', 20, 110000, 1),
(14, 'Kursi Gaming', 10, 1000000, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `employees`
--
ALTER TABLE `employees`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
