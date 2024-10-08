FROM kong:latest
USER root

COPY kong.yml /usr/local/kong/declarative/kong.yml

RUN chmod -R 777 /usr/local/kong

USER kong

CMD ["kong", "start", "--v", "--conf", "/usr/local/kong/declarative/kong.yml"]

