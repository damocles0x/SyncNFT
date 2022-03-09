# SyncNFT

#DDL
CREATE TABLE `contract` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`created_time` datetime DEFAULT '2000-01-01 00:00:00',
`updated_time` datetime DEFAULT '2000-01-01 00:00:00',
`tx_from_address` varchar(64) DEFAULT NULL,
`tx_hash` varchar(64) DEFAULT NULL,
`contract_address` varchar(64) DEFAULT NULL,
`contract_name` longtext,
`contract_symbol` longtext,
PRIMARY KEY (`id`),
KEY `contract_addr` (`contract_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 

CREATE TABLE `nft` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`created_time` datetime DEFAULT '2000-01-01 00:00:00',
`updated_time` datetime DEFAULT '2000-01-01 00:00:00',
`tx_hash` varchar(64) DEFAULT NULL,
`tx_from_address` varchar(64) DEFAULT NULL,
`tx_to_address` varchar(64) DEFAULT NULL,
`token_id` varchar(255) DEFAULT NULL,
`contract_id` varchar(255) DEFAULT NULL,
PRIMARY KEY (`id`),
KEY `tokenId` (`token_id`),
KEY `contract_id` (`contract_id`),
KEY `mergre` (`contract_id`,`token_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 