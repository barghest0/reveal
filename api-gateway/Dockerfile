FROM kong:latest
USER root

RUN echo "KONG_DATABASE=postgres" >> /usr/local/kong/.kong_env && \
    echo "KONG_PG_HOST=postgres" >> /usr/local/kong/.kong_env && \
    echo "KONG_PG_USER=postgres" >> /usr/local/kong/.kong_env && \
    echo "KONG_PG_PASSWORD=barghest" >> /usr/local/kong/.kong_env && \
    # echo "KONG_PG_DATABASE=users" >> /usr/local/kong/.kong_env && \
    echo "KONG_PRFIX=/usr/local/kong" >> /usr/local/kong/.kong_env 

COPY kong.yml /usr/local/kong/kong.conf

RUN chmod -R 777 /usr/local/kong

USER kong

WORKDIR /usr/local/kong

CMD ["kong", "start", "--v"]
