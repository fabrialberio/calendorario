FROM buildpack-deps:latest
RUN wget -O tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.16/tailwindcss-linux-x64
RUN chmod u+x tailwindcss
ENTRYPOINT [ "/tailwindcss" ]
