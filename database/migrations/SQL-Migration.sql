--
-- Table structure for table `companies`
--
DROP TABLE IF EXISTS `companies`;
CREATE TABLE `companies` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ชื่อ companies',
  `is_active` tinyint(1) NOT NULL COMMENT 'เปิดใช้โรงงาน',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_register` tinyint(1) DEFAULT NULL COMMENT 'แสดงใน select Register หรือไม่',
  `path_image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ลิ้งรูป'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `devices`
--
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `employee_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'เจ้าของอุปกร',
  `os` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ยี่ห้อ',
  `identifier` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'รหัส',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Token'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `employees`
--
DROP TABLE IF EXISTS `employees`;
CREATE TABLE `employees` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ชื่อจริง',
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'นามสกุล',
  `mobile` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'เบอร์โทร',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'รหัสผ่าน',
  `company_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'บริษัท',
  `prefix_name_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'คำนำหน้าชื่อ',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `easy_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ชื่อเล่น',
  `is_active` tinyint(1) DEFAULT NULL COMMENT 'สามารถเข้าสู่ระบบได้',
  `level_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'ระดับแต่ละคน'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `history_manage_products`
--
DROP TABLE IF EXISTS `history_manage_products`;
CREATE TABLE `history_manage_products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `employee_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'พนักงานที่ทำรายการเปลี่ยนแปลสินค้า',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `history_manage_products_to_products`
--
DROP TABLE IF EXISTS `history_manage_products_to_products`;
CREATE TABLE `history_manage_products_to_products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `product_id` bigint(20) UNSIGNED DEFAULT NULL,
  `history_manage_product_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `levels`
--
DROP TABLE IF EXISTS `levels`;
CREATE TABLE `levels` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ชื่อ',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `menus`
--
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ชื่อเมนู',
  `is_active` tinyint(1) NOT NULL COMMENT 'เปิดใช้ผู้พัฒนา',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `menu_employees`
--
DROP TABLE IF EXISTS `menu_employees`;
CREATE TABLE `menu_employees` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `employee_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'พนักงาน',
  `menu_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'พนักงาน',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



-- --------------------------------------------------------

--
-- Table structure for table `migrations`
--


-- --------------------------------------------------------

--
-- Table structure for table `prefix_names`
--
DROP TABLE IF EXISTS `prefix_names`;
CREATE TABLE `prefix_names` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'คำนำหน้าชื่อ',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ชื่อ',
  `price` double NOT NULL COMMENT 'ราคา',
  `image_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ที่อยู่รูป',
  `num_product` int(11) NOT NULL COMMENT 'จำนวนสินค้า',
  `detail` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ข้อมูล',
  `company_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'บริษัท',
  `adder_id` bigint(20) UNSIGNED DEFAULT NULL COMMENT 'พนักงานที่เพิ่มสินค้า',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`id`),
  ADD KEY `devices_employee_id_foreign` (`employee_id`);

--
-- Indexes for table `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `employees_mobile_unique` (`mobile`),
  ADD KEY `employees_company_id_foreign` (`company_id`),
  ADD KEY `employees_prefix_name_id_foreign` (`prefix_name_id`),
  ADD KEY `employees_level_id_foreign` (`level_id`);

--
-- Indexes for table `history_manage_products`
--
ALTER TABLE `history_manage_products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `history_manage_products_employee_id_foreign` (`employee_id`);

--
-- Indexes for table `history_manage_products_to_products`
--
ALTER TABLE `history_manage_products_to_products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `history_manage_products_to_products_product_id_foreign` (`product_id`),
  ADD KEY `hmp_id_foreign` (`history_manage_product_id`);

--
-- Indexes for table `levels`
--
ALTER TABLE `levels`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `menu_employees`
--
ALTER TABLE `menu_employees`
  ADD PRIMARY KEY (`id`),
  ADD KEY `menu_employees_employee_id_foreign` (`employee_id`),
  ADD KEY `menu_employees_menu_id_foreign` (`menu_id`);

--
-- Indexes for table `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `prefix_names`
--
ALTER TABLE `prefix_names`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `products_company_id_foreign` (`company_id`),
  ADD KEY `products_adder_id_foreign` (`adder_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `devices`
--
ALTER TABLE `devices`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=73;

--
-- AUTO_INCREMENT for table `employees`
--
ALTER TABLE `employees`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=84;

--
-- AUTO_INCREMENT for table `history_manage_products`
--
ALTER TABLE `history_manage_products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `history_manage_products_to_products`
--
ALTER TABLE `history_manage_products_to_products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `levels`
--
ALTER TABLE `levels`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `menus`
--
ALTER TABLE `menus`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `menu_employees`
--
ALTER TABLE `menu_employees`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `prefix_names`
--
ALTER TABLE `prefix_names`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `devices`
--
ALTER TABLE `devices`
  ADD CONSTRAINT `devices_employee_id_foreign` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `employees`
--
ALTER TABLE `employees`
  ADD CONSTRAINT `employees_company_id_foreign` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `employees_level_id_foreign` FOREIGN KEY (`level_id`) REFERENCES `levels` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `employees_prefix_name_id_foreign` FOREIGN KEY (`prefix_name_id`) REFERENCES `prefix_names` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `history_manage_products`
--
ALTER TABLE `history_manage_products`
  ADD CONSTRAINT `history_manage_products_employee_id_foreign` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `history_manage_products_to_products`
--
ALTER TABLE `history_manage_products_to_products`
  ADD CONSTRAINT `history_manage_products_to_products_product_id_foreign` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `hmp_id_foreign` FOREIGN KEY (`history_manage_product_id`) REFERENCES `history_manage_products` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `menu_employees`
--
ALTER TABLE `menu_employees`
  ADD CONSTRAINT `menu_employees_employee_id_foreign` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `menu_employees_menu_id_foreign` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_adder_id_foreign` FOREIGN KEY (`adder_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `products_company_id_foreign` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
