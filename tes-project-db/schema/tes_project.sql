/*
SQLyog Enterprise v12.5.1 (64 bit)
MySQL - 8.0.17 : Database - test-tunaiku
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`test-tunaiku` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `test-tunaiku`;

/*Table structure for table `tbl_installment` */

DROP TABLE IF EXISTS `tbl_installment`;

CREATE TABLE `tbl_installment` (
  `id_installment` int(11) NOT NULL AUTO_INCREMENT,
  `loan_code` varchar(50) DEFAULT NULL,
  `user_code` varchar(50) DEFAULT NULL,
  `capital` bigint(20) DEFAULT NULL,
  `interest` bigint(20) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  `plan` int(11) DEFAULT NULL,
  `due_date` datetime DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_installment`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tbl_installment` */

insert  into `tbl_installment`(`id_installment`,`loan_code`,`user_code`,`capital`,`interest`,`total`,`plan`,`due_date`,`status`) values 
(1,'LOAN08152019155825',NULL,1000000,201600,1201600,1,'2018-10-01 00:00:00',NULL),
(2,'LOAN08152019155825',NULL,1000000,201600,1201600,2,'2018-11-01 00:00:00',NULL),
(3,'LOAN08152019155825',NULL,1000000,201600,1201600,3,'2018-12-01 00:00:00',NULL),
(4,'LOAN08152019155825',NULL,1000000,201600,1201600,4,'2019-01-01 00:00:00',NULL),
(5,'LOAN08152019155825',NULL,1000000,201600,1201600,5,'2019-02-01 00:00:00',NULL),
(6,'LOAN08152019155825',NULL,1000000,201600,1201600,6,'2019-03-01 00:00:00',NULL),
(7,'LOAN08152019155825',NULL,1000000,201600,1201600,7,'2019-04-01 00:00:00',NULL),
(8,'LOAN08152019155825',NULL,1000000,201600,1201600,8,'2019-05-01 00:00:00',NULL),
(9,'LOAN08152019155825',NULL,1000000,201600,1201600,9,'2019-06-01 00:00:00',NULL),
(10,'LOAN08152019155825',NULL,1000000,201600,1201600,10,'2019-07-01 00:00:00',NULL),
(11,'LOAN08152019155825',NULL,1000000,201600,1201600,11,'2019-08-01 00:00:00',NULL),
(12,'LOAN08152019155825',NULL,1000000,201600,1201600,12,'2019-09-01 00:00:00',NULL);

/*Table structure for table `tbl_loan_interest` */

DROP TABLE IF EXISTS `tbl_loan_interest`;

CREATE TABLE `tbl_loan_interest` (
  `id_loan_interest` int(11) NOT NULL AUTO_INCREMENT,
  `tenor` int(20) DEFAULT NULL,
  `interest` float DEFAULT NULL,
  PRIMARY KEY (`id_loan_interest`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tbl_loan_interest` */

insert  into `tbl_loan_interest`(`id_loan_interest`,`tenor`,`interest`) values 
(1,12,1.68),
(2,18,1.68),
(3,24,1.59),
(4,30,1.59),
(5,36,1.59);

/*Table structure for table `tbl_request_loans` */

DROP TABLE IF EXISTS `tbl_request_loans`;

CREATE TABLE `tbl_request_loans` (
  `id_request_loan` int(11) NOT NULL AUTO_INCREMENT,
  `user_code` varchar(50) DEFAULT NULL,
  `loan_code` varchar(50) DEFAULT NULL,
  `jumlah_pinjaman` bigint(20) DEFAULT NULL,
  `lama_tenor` int(5) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `deleted` int(1) NOT NULL DEFAULT '0',
  `deleted_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_request_loan`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tbl_request_loans` */

insert  into `tbl_request_loans`(`id_request_loan`,`user_code`,`loan_code`,`jumlah_pinjaman`,`lama_tenor`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted`,`deleted_by`) values 
(1,'Wir20010925','LOAN08152019155809',2000000,12,'2019-08-15 15:58:09',NULL,NULL,NULL,0,NULL);

/*Table structure for table `tbl_user` */

DROP TABLE IF EXISTS `tbl_user`;

CREATE TABLE `tbl_user` (
  `id_user` int(20) NOT NULL AUTO_INCREMENT,
  `user_code` varchar(50) DEFAULT NULL,
  `no_ktp` varchar(20) DEFAULT NULL,
  `nama` varchar(225) DEFAULT NULL,
  `tanggal_lahir` varchar(50) DEFAULT NULL,
  `gender` varchar(50) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(50) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `deleted` int(11) NOT NULL DEFAULT '0',
  `deleted_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tbl_user` */

insert  into `tbl_user`(`id_user`,`user_code`,`no_ktp`,`nama`,`tanggal_lahir`,`gender`,`created_at`,`created_by`,`updated_at`,`updated_by`,`deleted`,`deleted_by`) values 
(1,'Wir20010925','3522582509010002','Wiro','2001-09-25','Male','2019-08-15 15:58:09',NULL,NULL,NULL,0,NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
