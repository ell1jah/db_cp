\copy brand FROM 'mnt/brand.csv' DELIMITER ';';
\copy webUser FROM 'mnt/user.csv' DELIMITER ';';
\copy item FROM 'mnt/item.csv' DELIMITER ';';
\copy ordering FROM 'mnt/ordering.csv' WITH DELIMITER ';' NULL AS 'null' csv;
\copy orderItems FROM 'mnt/orderItems.csv' DELIMITER ';';
