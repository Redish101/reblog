====
指南
====

reblog是一个使用go编写的轻量化的动态博客框架，本文将介绍如何快速搭建reblog服务端。

前置条件
--------

- 性能足够的大脑
- 一台能够运行docker的linux服务器（reblog生产构建的运行最低需要7MB的RAM，但为了在高负载情况下的正常运行，建议分配至少128MB的内存）
- curl, docker, docker-compose

.. attention:: 强烈建议在装载linux的86_64架构服务器上通过docker进行部署。因reblog使用部分linux独占特性，故在windows server上的部署可能会出现无法预知的问题。

在开始之前，请确保服务器已安装docker与docker-compose。

执行如下脚本:

.. code-block:: bash

    curl https://github.com/redish101/reblog/raw/main/scripts/quick_start.sh | sh

此时docker-compose会自动拉取镜像后运行，之后可通过以下命令运行reblog:

.. code-block:: bash

    cd ~/reblog
    docker-compose up -d

若需更新，可执行:

.. code-block:: bash

    cd ~/reblog
    docker-compose update
    docker-compose up -d
