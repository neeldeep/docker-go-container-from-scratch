FROM scratch

EXPOSE 9190

COPY hello /
COPY template template

CMD ["/hello"]