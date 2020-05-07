CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `users` (`id`, `name`, `created`) VALUES (1, 'Olga', '1978-01-06 15:27:33');
INSERT INTO `users` (`id`, `name`, `created`) VALUES (2, 'Scot', '2002-07-19 04:24:42');
