FROM postgres:11.5-alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB clothshop

COPY build/package/postgres/init.sql /docker-entrypoint-initdb.d/
COPY build/package/postgres/copy.sql /mnt/copy.sql
COPY build/package/postgres/brand.csv /mnt/brand.csv
COPY build/package/postgres/item.csv /mnt/item.csv
COPY build/package/postgres/ordering.csv /mnt/ordering.csv
COPY build/package/postgres/user.csv /mnt/user.csv
COPY build/package/postgres/orderItems.csv /mnt/orderItems.csv

EXPOSE 5432
CMD ["postgres"]

# docker build --tag=clothshop:latest . && docker run -d --rm -p 5432:5432 clothshop
# docker run -d --rm -p 5432:5432 clothshop

# psql postgresql://postgres:postgres@localhost:5432/clothshop

# -- psql -U postgres -d clothshop -f "mnt/copy.sql"
