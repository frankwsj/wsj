include:
  - nginx.install


nginx_conf:
  file.managed:
    - name: /usr/local/nginx/conf/nginx.conf
    - source: salt://nginx/files/nginx.conf
    - template: jinja
    - defaults:
        nginx_user: nginxnginx
        num_cpus: undefined

nginx_service:        ##Nginx服务管理
  file.managed:
    - name: /etc/init.d/nginx
    - user: root
    - mode: 755
    - source: salt://nginx/files/nginx
  cmd.run:        ## 设置Nginx服务开机启动
    - names:
       - /sbin/chkconfig --add nginx
       - /sbin/chkconfig --level 35 nginx on
    - unless: /sbin/chkconfig --list nginx
  service.running:        ##Nginx启动状态
    - watch:
      - file: /usr/local/nginx/conf/nginx.conf
    - name: nginx
    - enable: True
    - reload: True



