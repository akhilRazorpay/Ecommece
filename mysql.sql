
CREATE TABLE `product` (
	`id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`name` varchar(250) NOT NULL,
	`price` int(11) NOT NULL,
	`quantity` int(11) NOT NULL,
	`image` VARCHAR(250) NOT NULL,
	`description` VARCHAR(300) NOT NULL
  ) ENGINE=MyISAM DEFAULT CHARSET=latin1;

  --
  -- Dumping data for table `product`
  --

  INSERT INTO `product` (`id`, `name`, `price`, `quantity`, `image`, `description`) VALUES
  (1, 'Tv 1', 20000, 4, 'https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$', "TV1 DESCRIPTION"),
  (2, 'Tv 2', 25000, 6, "https://images.philips.com/is/image/PhilipsConsumer/58PUT6604_94-IMS-en_IN?wid=840&hei=720&$jpglarge$", "TV2 DESCRIPTION"),
  (3, 'Laptop 1', 50000, 11, "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", "LAPTOP1 DESCRIPTION"),
  (4, 'Laptop 2', 54100, 6, "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", "LAPTOP2 DESCRIPTION"),
  (5, 'Laptop 3', 60000, 11, "https://www.cnet.com/a/img/resize/4e82f3a17554a5aff8089194237de5a3acfce3b4/2022/04/27/b796792b-5b34-4405-83eb-efc66371ee06/samsung-galaxy-book-2-pro-360-01.jpg?auto=webp&fit=crop&height=630&width=1200", "LAPTOP3 DESCRIPTION"),
  (6, 'Mobile 1', 12000, 20, "https://cdn-www.mediatek.com/page/Mobile-2_2021-10-20-155734_vspa.png", "Mobile1 DESCRIPTION"),
  (7, 'Mobile 2', 21000, 25, "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", "Mobile2 DESCRIPTION");
