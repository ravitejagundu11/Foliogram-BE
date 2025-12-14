-- MySQL dump 10.13  Distrib 8.0.44, for Linux (aarch64)
--
-- Host: localhost    Database: foliogram_db
-- ------------------------------------------------------
-- Server version	8.0.44

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `analytics_views`
--

DROP TABLE IF EXISTS `analytics_views`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `analytics_views` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `viewer_ip` varchar(45) DEFAULT NULL,
  `session_id` varchar(255) DEFAULT NULL,
  `user_agent` varchar(500) DEFAULT NULL,
  `referrer` varchar(500) DEFAULT NULL,
  `viewed_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_analytics_views_portfolio_id` (`portfolio_id`),
  KEY `idx_analytics_views_viewed_at` (`viewed_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `analytics_views`
--

LOCK TABLES `analytics_views` WRITE;
/*!40000 ALTER TABLE `analytics_views` DISABLE KEYS */;
INSERT INTO `analytics_views` VALUES ('0069a19a-c2b4-40f8-8e90-8b831af186b3','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('00987d55-f2de-4ee9-9f21-0562cd98e7d2','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('02728807-397d-4478-870d-590060e80450','54ee2f47-ab6c-4f4a-be77-036f02a73743','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('05d9736a-7b00-4845-bfcb-355807b0b28c','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('0b099060-f7ab-47ae-bfff-fd6f391a5099','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('0f18a2e1-6605-49bc-b529-fabbea957cda','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('0f64988b-7419-47fa-8238-29e3254080fe','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('11c294aa-ed7c-4181-95a1-b69eaac8a6e4','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('13a160dc-9163-47a6-a8a7-48aca593ccef','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('1c6d8de3-512a-475d-85b3-1d19caa25a5f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('1f4df5df-fe89-4e10-b908-ac957a8e325f','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('20c0dc56-96b7-422e-b0e4-7e1026281f4a','641cd400-5781-4254-b75e-c3f8d142f5af','::1','curl/8.7.1',NULL,NULL,NULL,NULL),('242acbd3-186d-404e-8e39-4b5f5045bf66','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('27fca792-e70f-4260-8bdf-f90cf756a243','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('297fff25-6645-41a1-b827-ea3c9dd9aee6','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('30408c57-0ccd-4ea9-9241-6624946512c2','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('32b04f44-153d-44ac-b7d0-77e7f9f3d673','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('3a93921b-bbed-43b5-b432-b0ea386b37d8','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('410d7f29-c04c-4573-873b-9cfd35bec18c','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('4330158b-67c0-4396-a09f-e4681e09427d','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('46fae40f-4cc1-4e6d-a36a-1c2c986fa02a','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('4ba1cddb-2950-44f6-8899-ef2f8772403d','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('4cb8aaa5-6cfd-4554-95bd-2c6e5dbe97c2','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('558d25c9-71c3-49de-82f1-3bedf7bfbfa5','641cd400-5781-4254-b75e-c3f8d142f5af','::1','curl/8.7.1',NULL,NULL,NULL,NULL),('5ecc14b4-75e3-4071-a0c3-e1444d92354f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('6018015e-cc43-480f-9f14-26b57f47b38c','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('634512d9-9573-4865-8797-1244c575d2c4','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('65434da8-94d7-40b6-8758-10e44ba2bf1c','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('690ede1f-3b0f-4eae-9b7c-41ced2aa325f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('6d779d51-e792-44f4-bb32-c0a7dfea9e57','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('7b8cffe2-f660-4a37-9b9d-14636783989e','54ee2f47-ab6c-4f4a-be77-036f02a73743','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('7d2946e5-8a2d-467c-a694-e38d7b45435d','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('7dcba9a2-be57-43b8-a4ce-7927fe7aa82c','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('822a406f-2220-45c5-9fa3-f7afb3da269f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('833e953e-e66d-4c43-a75b-2c6d26cd69e8','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('84fc24ce-0ef7-4ff5-a8d8-f7df1426886e','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('860be284-333b-416e-b237-88b1d4ade296','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('88c0741d-6891-4abf-8306-1cb5ba7eb7d4','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('89287e1a-1c99-4c56-b662-2888c08491cf','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('899c4603-a3ed-4b0e-bd17-8aec96cb3826','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('8d8ba905-1831-43b3-a6ad-43d513382f4f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('8e139f8b-4a7a-4d12-ac34-916eaf64d1e6','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('8f829afc-f806-4093-b1c9-1df7fd4b239f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('92588b1a-20c7-497d-9d8d-f5e471faa771','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('94034358-8f3b-439d-a89d-5953ffe6d2ef','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('981752db-c900-4f2c-a973-17f9526c0edf','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('999152de-b7f6-4e61-b1fd-d51c8c881541','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('9ad5725e-e8a8-45ce-b032-385aba9ea0f8','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('9bec3d61-295a-4094-9217-a2832a0e7138','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('9c445fbd-de79-4d33-b54f-7ffee9c0c243','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('a19e01d3-7cd5-4b6c-92a1-31ea1ab5fdd8','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('a4f86910-6c59-4444-8212-798525a48be1','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('a73d9c51-ffde-4f56-babe-dca03d85e59a','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('a822e39e-5411-4b38-a2b2-e8d0597f7f5c','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('afd2470c-10ee-49cb-a01f-125a9c3a15a2','641cd400-5781-4254-b75e-c3f8d142f5af','::1','curl/8.7.1',NULL,NULL,NULL,NULL),('b4c77fbe-c8f1-4984-977b-f7902ea9870f','54ee2f47-ab6c-4f4a-be77-036f02a73743','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('b64491a6-32e0-45dd-b200-84ae3e623cc7','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('ba05f52d-ba08-40ed-9c3e-1457d86630b9','54ee2f47-ab6c-4f4a-be77-036f02a73743','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('bbbc447e-6280-4eb4-b784-604ea96e7e0f','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('bbe6f13e-8c49-45b6-a7af-6b8860cd20e5','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('bfb491c3-a67e-4b5c-8c07-5cd12a1d122c','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('c1c60627-5d50-4565-a817-43a683bfdb23','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('c8d44eae-8b80-49ae-8a57-dbb183a7cccd','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('cc0c0392-8a5e-444f-a83e-6f34b121ca94','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('cd7f2f85-62e5-4cae-989b-d2a3a21ee82e','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('d440f795-8b4f-47e5-88f3-23bbf76eb34b','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('d4f00aee-209a-4ca7-8fb4-c88543180425','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('d56ee221-3e90-43d9-afb8-f9961fd27ea3','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('d5cffabe-b782-44e3-87e0-ed07e1a51c64','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('d8fff3ef-41d4-4663-9e7d-7abf57441806','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('da6a00cd-83e7-488d-831a-82ac35a3ae10','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('e36e130b-d1f1-4951-be00-a6aaa795ec98','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('e63ca3fb-f134-4373-9544-c269eb6a727b','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('ec059910-024e-4750-8e0e-819f0b3d2ed0','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('ecf0dba8-3d24-4eec-bd73-cccb2fe3efdc','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('ed6a2bf3-b359-458b-bd18-dff95a30bc82','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('f27c32dd-75c0-4b33-afdb-0204d7d4d57e','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('f341cf15-5edb-49ea-a202-e96528773063','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('fb135052-ee62-4fed-9a29-cc7a2e63b909','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL),('fb77a592-1171-4ffc-bd77-72a482518ca6','641cd400-5781-4254-b75e-c3f8d142f5af','::1','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36',NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `analytics_views` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contacts`
--

DROP TABLE IF EXISTS `contacts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contacts` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `address` text,
  `city` varchar(100) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `map_url` varchar(500) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_portfolio_contact` (`portfolio_id`),
  KEY `idx_contacts_deleted_at` (`deleted_at`),
  KEY `idx_contacts_portfolio_id` (`portfolio_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contacts`
--

LOCK TABLES `contacts` WRITE;
/*!40000 ALTER TABLE `contacts` DISABLE KEYS */;
INSERT INTO `contacts` VALUES ('46ec63be-0afc-4631-9a96-91ffbe7cb027','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','raviteja@gmail.com','12345678','','','','','2025-12-10 09:30:58.299','2025-12-10 09:33:26.682',NULL),('5283bb2d-d419-4566-b98b-2ff1e55fcf32','641cd400-5781-4254-b75e-c3f8d142f5af','pebaruhu@mailinator.com','+1 (658) 113-1592','Iste est ipsum ass','','','','2025-12-10 09:05:31.185','2025-12-10 09:12:54.939',NULL);
/*!40000 ALTER TABLE `contacts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `education`
--

DROP TABLE IF EXISTS `education`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `education` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `institution` varchar(255) NOT NULL,
  `degree` varchar(255) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL,
  `start_date` varchar(50) DEFAULT NULL,
  `end_date` varchar(50) DEFAULT NULL,
  `grade` varchar(50) DEFAULT NULL,
  `description` text,
  `sort_order` int DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_education_portfolio_id` (`portfolio_id`),
  KEY `idx_education_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `education`
--

LOCK TABLES `education` WRITE;
/*!40000 ALTER TABLE `education` DISABLE KEYS */;
INSERT INTO `education` VALUES ('63dbd6bd-7a45-4c73-a51c-9b110e31aed4','641cd400-5781-4254-b75e-c3f8d142f5af','Totam totam quis com','Voluptas eos esse ','Corrupti cumque id ','2022-08-23','2011-06-24','','',0,'2025-12-10 09:05:31.167','2025-12-10 09:12:54.915',NULL),('6c660c45-6a35-4996-9e1a-701a8146bfad','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','Purdue University, Fort Wayne      		','Master\'s','Computer Science','2025-01-08','2027-01-05','','',0,'2025-12-10 09:30:58.291','2025-12-10 09:33:26.667',NULL);
/*!40000 ALTER TABLE `education` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `experiences`
--

DROP TABLE IF EXISTS `experiences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `experiences` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `company` varchar(255) NOT NULL,
  `position` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `start_date` varchar(50) DEFAULT NULL,
  `end_date` varchar(50) DEFAULT NULL,
  `is_current` tinyint(1) DEFAULT '0',
  `description` text,
  `sort_order` int DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_experiences_portfolio_id` (`portfolio_id`),
  KEY `idx_experiences_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `experiences`
--

LOCK TABLES `experiences` WRITE;
/*!40000 ALTER TABLE `experiences` DISABLE KEYS */;
INSERT INTO `experiences` VALUES ('2f38ea55-a6cb-4ea4-9609-4056f61a5b29','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','Sureify','SOFTWARE ENGINEER III','Full Time','2021-01-01','2025-01-01',0,'',0,'2025-12-10 09:30:58.295','2025-12-10 09:33:26.675',NULL),('5119a39e-dc60-4764-bebc-93be62099570','641cd400-5781-4254-b75e-c3f8d142f5af','Omnis explicabo Vol','Qui quo quo sint a n','Internship','1987-05-28','2016-09-24',0,'',0,'2025-12-10 09:05:31.173','2025-12-10 09:12:54.923',NULL);
/*!40000 ALTER TABLE `experiences` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `portfolios`
--

DROP TABLE IF EXISTS `portfolios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `portfolios` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `template_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `description` text,
  `is_published` tinyint(1) DEFAULT '0',
  `custom_domain` varchar(255) DEFAULT NULL,
  `config` json DEFAULT NULL,
  `meta_title` varchar(255) DEFAULT NULL,
  `meta_description` text,
  `meta_keywords` varchar(500) DEFAULT NULL,
  `view_count` bigint DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_portfolios_slug` (`slug`),
  KEY `idx_portfolios_user_id` (`user_id`),
  KEY `idx_portfolios_template_id` (`template_id`),
  KEY `idx_portfolios_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `portfolios`
--

LOCK TABLES `portfolios` WRITE;
/*!40000 ALTER TABLE `portfolios` DISABLE KEYS */;
INSERT INTO `portfolios` VALUES ('641cd400-5781-4254-b75e-c3f8d142f5af','78540f78-2154-40ad-b24b-fd83a9e727ec','48877b9f-d581-11f0-b19b-6ada76144272','Perferendis explicab','perferendis-explicab','Beatae adipisicing e',1,'','{\"bodyFont\": \"Open Sans\", \"headline\": \"Inventore qui dolor \", \"headingFont\": \"Inter\", \"layoutStyle\": \"{\\\"headerStyle\\\":\\\"centered\\\",\\\"spacing\\\":\\\"comfortable\\\",\\\"cardStyle\\\":\\\"rounded\\\"}\", \"socialLinks\": {\"github\": \"https://www.kemo.org.au\", \"twitter\": \"https://www.wijyguneposukyl.net\", \"website\": \"https://www.pupigoge.org\", \"linkedin\": \"https://www.qabihy.info\"}, \"primaryColor\": \"#3B82F6\", \"profilePicture\": \"http://localhost:8080/uploads/images/ecdf8e5e-92d2-48e0-a6e2-527a208096ce_1765371829.png\", \"secondaryColor\": \"#10B981\"}','','','',0,'2025-12-10 07:33:06.785','2025-12-10 09:12:54.852',NULL),('7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','78540f78-2154-40ad-b24b-fd83a9e727ec','48877cdb-d581-11f0-b19b-6ada76144272','Raviteja Gundu','raviteja-gundu','Computer Science Master’s student at Purdue University, Fort Wayne, and a software engineer with 4+ years of experience in building scalable web applications and RESTful APIs. Expertise in developing high-performance backend microservices using Go and AWS architecture. Proficient in Test-Driven Development, with a strong focus on efficiency and code quality.',1,'','{\"bodyFont\": \"Inter\", \"headline\": \"Master\'s Student at Purdue University\", \"headingFont\": \"Inter\", \"layoutStyle\": \"{\\\"headerStyle\\\":\\\"centered\\\",\\\"spacing\\\":\\\"comfortable\\\",\\\"cardStyle\\\":\\\"rounded\\\"}\", \"socialLinks\": {\"github\": \"https://github.com/ravitejagundu11\", \"twitter\": \"\", \"website\": \"\", \"linkedin\": \"https://www.linkedin.com/in/ravitejagundu/\"}, \"primaryColor\": \"#667eea\", \"profilePicture\": \"http://localhost:8080/uploads/images/6acbff03-035f-401a-a3e6-7d6ae1b421f9_1765377058.jpg\", \"secondaryColor\": \"#764ba2\"}','','','',0,'2025-12-10 09:30:58.247','2025-12-10 09:33:26.613',NULL);
/*!40000 ALTER TABLE `portfolios` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `projects` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text,
  `image_url` varchar(500) DEFAULT NULL,
  `gallery_urls` json DEFAULT NULL,
  `demo_url` varchar(500) DEFAULT NULL,
  `github_url` varchar(500) DEFAULT NULL,
  `tags` json DEFAULT NULL,
  `technologies` json DEFAULT NULL,
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  `is_featured` tinyint(1) DEFAULT '0',
  `sort_order` bigint DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_projects_portfolio_id` (`portfolio_id`),
  KEY `idx_projects_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES ('04818ad1-1f11-4990-87dc-a19f5382dec4','54ee2f47-ab6c-4f4a-be77-036f02a73743','Voluptate illo exerc','Repellendus Invento','','[]','https://www.gyweqenijyqova.org.au','https://www.nuhydijun.info','[]','[]',NULL,NULL,1,0,'2025-12-10 07:14:14.202','2025-12-10 07:14:14.202',NULL),('148629a8-7992-4955-8088-448b39406524','641cd400-5781-4254-b75e-c3f8d142f5af','Ipsum itaque alias ','Quia quasi animi la','','[]','https://www.sytibyhy.ca','https://www.hokivegapyn.net','[]','[]',NULL,NULL,0,0,'2025-12-10 07:43:23.435','2025-12-10 09:12:54.873',NULL),('93f1a511-c08f-4b10-a7ce-a68b58c25cbf','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','EASY SIGNATURE SERVICE','Developed the Easy Signature backend microservice with the Bee-go framework, enabling efficient integration with DocuSign and Adobe Sign via REST APIs for a streamlined e-signature experience.\nStreamlined document processing with webhook architecture, enhancing efficiency and reducing manual intervention, improving response times by 7ms.\nDeveloped a cost-effective serverless solution with AWS Lambda and Python to dynamically add form fields to documents, filling a gap in Go\'s PDF processing capabilities.\nOptimized service performance by caching documents using Redis and Memcached, reducing response times by 5ms.\n','','[]','','','[]','[]',NULL,NULL,0,0,'2025-12-10 09:30:58.271','2025-12-10 09:33:26.629',NULL),('b0ebc1aa-fe59-4ec3-8544-d34d72ab59ee','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','ADMIN DASHBOARD','Developed a React.js and TypeScript frontend dashboard using Material UI for managing backend service configurations, enabling efficient CRUD operations via API interactions.\nCreated reusable npm packages to enhance integration efficiency and consistency across client projects.\nLeveraged React states, contexts, and hooks to manage data flow between the UI and backend APIs effectively.\n','','[]','','','[]','[]',NULL,NULL,1,0,'2025-12-10 09:30:58.271','2025-12-10 09:33:26.629',NULL),('c3172177-ef72-4b98-9969-2ddc46ed8692','3d224f12-5987-433c-8e3e-27a35af021da','Updated Test Project','Updated description','https://images.unsplash.com/photo-1498050108023-c5249f4df085','[]','https://updated.com','https://github.com/updated','[]','[\"React\", \"TypeScript\", \"Node.js\"]',NULL,NULL,0,1,'2025-12-10 06:50:20.073','2025-12-10 06:50:33.352',NULL),('caf59aca-6f35-4d45-be8d-38a00862deba','6b8b248c-15ce-4003-b8e6-c62cb3f07f82','Updated Project Title','Updated description here','','[]','https://updated.example.com','https://github.com/updated','[]','[\"React\", \"TypeScript\"]',NULL,NULL,1,1,'2025-12-10 06:55:18.973','2025-12-10 06:55:29.585',NULL),('f7d32a52-405f-4ae2-8f0a-4198ed20dade','54ee2f47-ab6c-4f4a-be77-036f02a73743','Voluptate illo exerc','Repellendus Invento','','[]','https://www.gyweqenijyqova.org.au','https://www.nuhydijun.info','[]','[]',NULL,NULL,1,0,'2025-12-10 07:00:00.276','2025-12-10 07:00:00.276',NULL),('fd8bc64f-999a-47a4-afee-87a4aa35e459','641cd400-5781-4254-b75e-c3f8d142f5af','Dolorum aut laborum','Dolores magna et cul','','[]','https://www.cycabamububyn.net','https://www.razacyhis.com.au','[]','[]',NULL,NULL,1,0,'2025-12-10 07:33:06.802','2025-12-10 09:12:54.873',NULL);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `publications`
--

DROP TABLE IF EXISTS `publications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `publications` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `title` varchar(500) NOT NULL,
  `authors` text,
  `journal` varchar(255) DEFAULT NULL,
  `conference` varchar(255) DEFAULT NULL,
  `year` varchar(10) DEFAULT NULL,
  `doi` varchar(255) DEFAULT NULL,
  `url` varchar(500) DEFAULT NULL,
  `abstract` text,
  `type` varchar(50) DEFAULT NULL,
  `sort_order` int DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_publications_portfolio_id` (`portfolio_id`),
  KEY `idx_publications_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `publications`
--

LOCK TABLES `publications` WRITE;
/*!40000 ALTER TABLE `publications` DISABLE KEYS */;
INSERT INTO `publications` VALUES ('9f22b806-1c1e-44dc-8d89-5736d0785dbc','641cd400-5781-4254-b75e-c3f8d142f5af','Iure unde et magna s','Culpa magna est qu','Culpa magna est qu','','1997','','','Animi amet quod fa','',0,'2025-12-10 09:05:31.178','2025-12-10 09:12:54.930',NULL);
/*!40000 ALTER TABLE `publications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `skills`
--

DROP TABLE IF EXISTS `skills`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `skills` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `category` varchar(100) DEFAULT NULL,
  `level` varchar(50) DEFAULT NULL,
  `percentage` bigint DEFAULT NULL,
  `icon_url` varchar(500) DEFAULT NULL,
  `sort_order` bigint DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_skills_portfolio_id` (`portfolio_id`),
  KEY `idx_skills_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `skills`
--

LOCK TABLES `skills` WRITE;
/*!40000 ALTER TABLE `skills` DISABLE KEYS */;
INSERT INTO `skills` VALUES ('2e17e69e-7d19-43c3-beb2-75e403b001b9','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','React','Frontend','',0,'',0,'2025-12-10 09:30:58.282','2025-12-10 09:33:26.648',NULL),('2ecb792e-24d0-4f0b-b217-101920a77744','54ee2f47-ab6c-4f4a-be77-036f02a73743','Esse molestias et do','Other','',0,'',0,'2025-12-10 07:13:23.024','2025-12-10 07:13:23.024',NULL),('7b8547f6-9a55-4b4a-b56e-1311838dd66c','641cd400-5781-4254-b75e-c3f8d142f5af','Magna aut assumenda ','Tools','',0,'',0,'2025-12-10 07:33:06.842','2025-12-10 09:12:54.892',NULL),('b2516e76-d70b-4b85-8c8a-bcdbb5f49c68','54ee2f47-ab6c-4f4a-be77-036f02a73743','Esse molestias et do','Other','',0,'',0,'2025-12-10 07:13:57.870','2025-12-10 07:13:57.870',NULL),('bed83fbb-6e7a-4cfc-9f38-6174ddb9de7e','641cd400-5781-4254-b75e-c3f8d142f5af','Tempora non natus qu','Backend','',0,'',0,'2025-12-10 07:43:23.444','2025-12-10 09:12:54.892',NULL),('d3bf3bd2-0437-4f02-a7d6-047037b6f091','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','GO','Backend','',0,'',0,'2025-12-10 09:30:58.282','2025-12-10 09:33:26.648',NULL),('dea892f5-7679-4ed3-8cf2-7412b287c3d1','54ee2f47-ab6c-4f4a-be77-036f02a73743','Esse molestias et do','Other','',0,'',0,'2025-12-10 07:14:14.209','2025-12-10 07:14:14.209',NULL),('f49caa50-b186-440c-8179-903ddc782ac5','641cd400-5781-4254-b75e-c3f8d142f5af','Est ducimus eius it','Tools','',0,'',0,'2025-12-10 07:43:23.444','2025-12-10 09:12:54.892',NULL);
/*!40000 ALTER TABLE `skills` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `templates`
--

DROP TABLE IF EXISTS `templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `templates` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `category` varchar(50) NOT NULL,
  `preview_url` varchar(500) DEFAULT NULL,
  `thumbnail_url` varchar(500) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `is_premium` tinyint(1) DEFAULT '0',
  `features` json DEFAULT NULL,
  `default_config` json DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `templates`
--

LOCK TABLES `templates` WRITE;
/*!40000 ALTER TABLE `templates` DISABLE KEYS */;
INSERT INTO `templates` VALUES ('4886bde8-d581-11f0-b19b-6ada76144272','Modern Developer','Clean and modern portfolio template perfect for developers and engineers','developer','https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1498050108023-c5249f4df085?w=400&h=300&fit=crop',1,0,'[\"Responsive Design\", \"Dark Mode\", \"Project Showcase\", \"Skills Section\", \"Contact Form\"]','{\"layout\": \"modern\", \"fontFamily\": \"Inter\", \"primaryColor\": \"#3B82F6\", \"secondaryColor\": \"#10B981\"}',NULL,NULL,NULL),('4886c2f7-d581-11f0-b19b-6ada76144272','Minimal Tech','Minimalist template focused on code and projects','developer','https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1461749280684-dccba630e2f6?w=400&h=300&fit=crop',1,0,'[\"Clean Layout\", \"GitHub Integration\", \"Code Snippets\", \"Timeline\", \"Blog\"]','{\"layout\": \"minimal\", \"fontFamily\": \"Roboto Mono\", \"primaryColor\": \"#000000\", \"secondaryColor\": \"#FFFFFF\"}',NULL,NULL,NULL),('4886c49d-d581-11f0-b19b-6ada76144272','Full Stack Pro','Comprehensive template for full-stack developers','developer','https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1517694712202-14dd9538aa97?w=400&h=300&fit=crop',1,1,'[\"Multi-page\", \"Tech Stack Showcase\", \"Testimonials\", \"Case Studies\", \"Analytics Dashboard\"]','{\"layout\": \"professional\", \"fontFamily\": \"Poppins\", \"primaryColor\": \"#8B5CF6\", \"secondaryColor\": \"#EC4899\"}',NULL,NULL,NULL),('4886fc45-d581-11f0-b19b-6ada76144272','Creative Studio','Bold and creative template for designers and artists','designer','https://images.unsplash.com/photo-1561070791-2526d30994b5?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1561070791-2526d30994b5?w=400&h=300&fit=crop',1,0,'[\"Gallery Grid\", \"Animation Effects\", \"Color Palette\", \"Process Timeline\", \"Client Logos\"]','{\"layout\": \"creative\", \"fontFamily\": \"Playfair Display\", \"primaryColor\": \"#F59E0B\", \"secondaryColor\": \"#EF4444\"}',NULL,NULL,NULL),('4886fd74-d581-11f0-b19b-6ada76144272','UI/UX Portfolio','Showcase your design process and case studies','designer','https://images.unsplash.com/photo-1559028012-481c04fa702d?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1559028012-481c04fa702d?w=400&h=300&fit=crop',1,1,'[\"Case Study Pages\", \"Figma Integration\", \"Before/After\", \"Design System\", \"Testimonials\"]','{\"layout\": \"showcase\", \"fontFamily\": \"Inter\", \"primaryColor\": \"#6366F1\", \"secondaryColor\": \"#8B5CF6\"}',NULL,NULL,NULL),('4886fdca-d581-11f0-b19b-6ada76144272','Artistic Vision','Elegant template for visual artists and illustrators','designer','https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1513364776144-60967b0f800f?w=400&h=300&fit=crop',1,0,'[\"Full-screen Gallery\", \"Lightbox\", \"Categories\", \"About Artist\", \"Shop Integration\"]','{\"layout\": \"gallery\", \"fontFamily\": \"Lora\", \"primaryColor\": \"#EC4899\", \"secondaryColor\": \"#F472B6\"}',NULL,NULL,NULL),('4887403b-d581-11f0-b19b-6ada76144272','Photo Gallery','Image-focused template for photographers','photographer','https://images.unsplash.com/photo-1452587925148-ce544e77e70d?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1452587925148-ce544e77e70d?w=400&h=300&fit=crop',1,0,'[\"Masonry Grid\", \"Full-screen Slideshow\", \"EXIF Data\", \"Collections\", \"Booking\"]','{\"layout\": \"gallery\", \"fontFamily\": \"Montserrat\", \"primaryColor\": \"#0F172A\", \"secondaryColor\": \"#64748B\"}',NULL,NULL,NULL),('48874193-d581-11f0-b19b-6ada76144272','Wedding Photography','Romantic template for wedding photographers','photographer','https://images.unsplash.com/photo-1519741497674-611481863552?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1519741497674-611481863552?w=400&h=300&fit=crop',1,1,'[\"Elegant Galleries\", \"Client Portal\", \"Package Pricing\", \"Testimonials\", \"Contact Form\"]','{\"layout\": \"elegant\", \"fontFamily\": \"Cormorant Garamond\", \"primaryColor\": \"#FBBF24\", \"secondaryColor\": \"#FCA5A5\"}',NULL,NULL,NULL),('48877b9f-d581-11f0-b19b-6ada76144272','Professional Business','Corporate template for consultants and agencies','business','https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?w=400&h=300&fit=crop',1,0,'[\"Services Section\", \"Team Members\", \"Client Logos\", \"Statistics\", \"Newsletter\"]','{\"layout\": \"corporate\", \"fontFamily\": \"Open Sans\", \"primaryColor\": \"#1E40AF\", \"secondaryColor\": \"#3B82F6\"}',NULL,NULL,NULL),('48877cdb-d581-11f0-b19b-6ada76144272','Entrepreneur','Dynamic template for entrepreneurs and startups','business','https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=400&h=300&fit=crop',1,1,'[\"Product Showcase\", \"Metrics Dashboard\", \"Press Mentions\", \"Video Integration\", \"CTA Sections\"]','{\"layout\": \"modern\", \"fontFamily\": \"Space Grotesk\", \"primaryColor\": \"#059669\", \"secondaryColor\": \"#10B981\"}',NULL,NULL,NULL),('48879cb5-d581-11f0-b19b-6ada76144272','Content Creator','Vibrant template for YouTubers and influencers','creator','https://images.unsplash.com/photo-1611162617474-5b21e879e113?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1611162617474-5b21e879e113?w=400&h=300&fit=crop',1,0,'[\"Video Embeds\", \"Social Media Links\", \"Latest Content\", \"Subscriber Count\", \"Merch Shop\"]','{\"layout\": \"dynamic\", \"fontFamily\": \"Nunito\", \"primaryColor\": \"#DC2626\", \"secondaryColor\": \"#F59E0B\"}',NULL,NULL,NULL),('48879e1f-d581-11f0-b19b-6ada76144272','Blogger','Clean template for writers and bloggers','creator','https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=800&h=600&fit=crop','https://images.unsplash.com/photo-1499750310107-5fef28a66643?w=400&h=300&fit=crop',1,0,'[\"Blog Posts\", \"Categories\", \"Search\", \"Newsletter\", \"Social Sharing\"]','{\"layout\": \"blog\", \"fontFamily\": \"Merriweather\", \"primaryColor\": \"#4338CA\", \"secondaryColor\": \"#818CF8\"}',NULL,NULL,NULL);
/*!40000 ALTER TABLE `templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testimonials`
--

DROP TABLE IF EXISTS `testimonials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `testimonials` (
  `id` char(36) NOT NULL,
  `portfolio_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `position` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `content` text NOT NULL,
  `rating` bigint DEFAULT NULL,
  `avatar_url` varchar(500) DEFAULT NULL,
  `linked_in_url` varchar(500) DEFAULT NULL,
  `is_visible` tinyint(1) DEFAULT '1',
  `sort_order` bigint DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_testimonials_portfolio_id` (`portfolio_id`),
  KEY `idx_testimonials_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testimonials`
--

LOCK TABLES `testimonials` WRITE;
/*!40000 ALTER TABLE `testimonials` DISABLE KEYS */;
INSERT INTO `testimonials` VALUES ('1653f1ee-06ad-4449-bd5e-4a33a58df386','641cd400-5781-4254-b75e-c3f8d142f5af','Nostrum maxime ab si','','Est consequuntur fug','Quae ullamco quis ut',5,'','',0,0,'2025-12-10 07:43:23.458','2025-12-10 09:12:54.905',NULL),('899e502c-2109-4da4-9983-89f9d9ac011d','7a2d81bb-bfb8-41b9-83a2-fbb1fe735636','Dolor et sit explic','','Voluptatem magnam vo','Consequatur Sit tem',5,'','',0,0,'2025-12-10 09:30:58.287','2025-12-10 09:33:26.658',NULL);
/*!40000 ALTER TABLE `testimonials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `mobile` varchar(20) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `profile_picture` text,
  `role` varchar(20) NOT NULL DEFAULT 'user',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('78540f78-2154-40ad-b24b-fd83a9e727ec','raviteja','ravi@gmail.com','$2a$10$siNXGYSbExxPbBQYBwDf.ufJ9Hj6GkBIIaktulxVAdRZZZJYKruze','','','','user','2025-12-04 15:53:34.774','2025-12-09 23:07:46.203'),('9f09707c-df23-4c10-b531-b09d8b6f2fd1','Akeem Mcintosh','duwes@mailinator.com','$2a$10$Loy6oliq0lUQPLSNY2X2quCyZr6ShPG1U7uyg/O4gavTria/oiKE2','','','','user',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-12-10 14:37:29
