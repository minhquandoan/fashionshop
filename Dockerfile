FROM alpine

ENV APP_HOME=/bcon/fashionshop/
WORKDIR ${APP_HOME}

COPY ./build/target/fashionshop ${APP_HOME}
CMD [ "./fashionshop" ]