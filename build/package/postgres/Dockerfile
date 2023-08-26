FROM postgres:11.5-alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB clothshop

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
CMD ["postgres"]

# docker build --tag=clothshop:latest .
# docker run -d --rm -p 5432:5432 clothshop

# psql postgresql://postgres:postgres@localhost:5432/clothshop

# \COPY person FROM 'mnt/data/person.csv' DELIMITER ';';
# \copy bank FROM 'mnt/data/bank.csv' DELIMITER ';';
# \copy company FROM 'mnt/data/company.csv' DELIMITER ';';
# \copy card FROM 'mnt/data/card.csv' DELIMITER ';';
# \copy store FROM 'mnt/data/store.csv' DELIMITER ';';
# \copy transaction FROM 'mnt/data/transaction.csv' DELIMITER ';';

# -- psql -U postgres -d clothshop -f "mnt/scripts/copy.sql"
