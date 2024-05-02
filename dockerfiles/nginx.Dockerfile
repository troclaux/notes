FROM nginx:latest

RUN apt-get update && \
    apt-get upgrade -y

RUN apt-get install -y \
    neovim \
    git

WORKDIR /etc/nginx

CMD ["nginx", "-g", "daemon off;"]
