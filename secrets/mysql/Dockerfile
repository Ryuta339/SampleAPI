#FROM mysql:5.6.37
FROM --platform=linux/x86_64 mysql:8.0
COPY 1_init.sh /docker-entrypoint-initdb.d/
COPY custom.cnf /etc/mysql/conf.d/custom.cnf
RUN chmod 644 /etc/mysql/conf.d/custom.cnf
