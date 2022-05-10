CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
`customer_id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`date_of_birth` date NOT NULL,
`city` varchar(100) NOT NULL, 
`zipcode` varchar(10) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES
	(2000, 'Stephen', '1978-12-15', 'Delhi', '110075', 1),
    (2001, 'Adrian', '1988-12-15', 'Newburgh', '12550', 1),
    (2002, 'Haley', '1998-12-15', 'Englewood', '07631', 1),
    (2003, 'Borris', '1975-12-15', 'Manchester, NH', '03102', 0),
    (2004, 'Natalie', '1974-07-18', 'Clarkson, MI', '48348', 1),
    (2005, 'Osiris', '1984-12-15', 'Hyattsville, MD', '20782', 0);
    
    
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
`account_id` int(11) NOT NULL AUTO_INCREMENT,
`customer_id` int(11) NOT NULL,
`opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
`account_type` varchar(10) NOT NULL,
`pin` varchar(10) NOT NULL, 
`status` tinyint(4) NOT NULL DEFAULT '1', 
PRIMARY KEY (`account_id`),
KEY `accounts_FK' (`customer_id`),

