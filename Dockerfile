FROM centos/mysql-80-centos7

USER root

# mysql container vars
ENV MYSQL_USER=du
ENV MYSQL_PASSWORD=pass
ENV MYSQL_DATABASE=testdb
ENV MYSQL_ROOT_PASSWORD=123456
# golang vars
ENV GOROOT=/usr/local/go
ENV CGO_ENABLED=1
# simeple-tree test vars
ENV BASE_URL=http://localhost
ENV LISTEN=:80
ENV GIN_MODE=test
ENV DB_DRIVER=mysql
ENV DB_USER=du
ENV DB_PASSWORD=pass
ENV DB_NAME=testdb
ENV DB_HOST=127.0.0.1:3306
ENV TEST_SERVER=http://localhost
ENV TEST_DSN=du:pass@tcp(localhost:3306)/testdb

ADD https://go.dev/dl/go1.23.2.linux-amd64.tar.gz /root/go.tar.gz
ADD http://mirrors.aliyun.com/repo/Centos-7.repo /etc/yum.repos.d/CentOS-Base.repo

RUN cd /root
RUN /bin/yum -q -y clean all
RUN /bin/yum --disablerepo=* --enablerepo=base -q -y makecache fast
RUN /bin/yum --disablerepo=* --enablerepo=base -q -y install gcc
RUN /bin/yum -q -y clean all
RUN /bin/tar -C /usr/local -xzf /root/go.tar.gz
RUN curl -sSf https://atlasgo.sh | sh -s -- -y

RUN mkdir -p /root/app
RUN groupadd app
RUN useradd -md /home/app -g app -s /bin/bash app
RUN chown -R app:app /home/app
RUN chmod 755 /home /home/app /root/app

RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> /root/.cshrc
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/app/.cshrc

USER mysql
