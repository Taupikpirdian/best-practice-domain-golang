-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Oct 02, 2022 at 12:24 AM
-- Server version: 5.7.33
-- PHP Version: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `salt_academy_exam_1`
--

-- --------------------------------------------------------

--
-- Table structure for table `article`
--

CREATE TABLE `article` (
  `id` int(11) NOT NULL,
  `code_article` varchar(100) DEFAULT NULL,
  `title_original` varchar(100) DEFAULT NULL,
  `text_original` text,
  `date` timestamp NULL DEFAULT NULL,
  `banner` varchar(100) DEFAULT NULL,
  `author` varchar(100) DEFAULT NULL,
  `thumbs` varchar(100) DEFAULT NULL,
  `is_highlight` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `article`
--

INSERT INTO `article` (`id`, `code_article`, `title_original`, `text_original`, `date`, `banner`, `author`, `thumbs`, `is_highlight`) VALUES
(12, 'XXXX-22029568', 'What is Lorem Ipsum?', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.', '2022-10-02 00:24:21', 'example.jpg', 'Taupik Pirdian', 'thumbs-example.jpg', 0);

-- --------------------------------------------------------

--
-- Table structure for table `translation`
--

CREATE TABLE `translation` (
  `id` int(11) NOT NULL,
  `code_article` varchar(100) DEFAULT NULL,
  `code_language` varchar(2) DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `text` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `translation`
--

INSERT INTO `translation` (`id`, `code_article`, `code_language`, `title`, `text`) VALUES
(6, 'XXXX-22029568', 'in', 'Apa itu Lorem Ipsum?', 'Lorem Ipsum hanyalah teks tiruan dari industri percetakan dan penyusunan huruf. Lorem Ipsum telah menjadi teks dummy standar industri sejak tahun 1500-an, ketika seorang pencetak yang tidak dikenal mengambil sekumpulan tipe dan mengacaknya untuk membuat buku spesimen tipe. Ini telah bertahan tidak hanya lima abad, tetapi juga lompatan ke pengaturan huruf elektronik, pada dasarnya tetap tidak berubah. Itu dipopulerkan pada 1960-an dengan merilis lembar Letraset yang berisi bagian-bagian Lorem Ipsum, dan baru-baru ini dengan perangkat lunak penerbitan desktop seperti Aldus PageMaker termasuk versi Lorem Ipsum.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `translation`
--
ALTER TABLE `translation`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `article`
--
ALTER TABLE `article`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `translation`
--
ALTER TABLE `translation`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
