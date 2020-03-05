

  逻辑分页


  数据库分页包括：物理分页、逻辑分页

    1.物理分页：是指在操作数据库时,在数据库中进行分页查询 例如使用 MySQL 的关键字 limit、SQLServer 的关键字 top 就可进行物理分页

    2.逻辑分页：逻辑分页分为前端逻辑分页、后端逻辑分页

        A.前端的逻辑分页：是指一次性返回给前端所有数据、让前端去选择性遍历部分数据的行为来达到分页显示效果.
                        这只是一种障眼法而已、只能用于少量数据的情况下，一旦数据变多、大量的数据在传输过程中、前端加载过程中都会变得很慢，而是实时性很差.

        B.后端的逻辑分页:是指在数据库查询到所有的数据、装在到容器中,根据页码和每页显示条数从容器中取出结果,返回给前端
                        其实，这也是一种障眼法、伪分页、前端每次请求的时候还是从数据库中查询所有的数据装在到容器中然后从容器中取出。
                        在与数据库交互次数变多、数据量变大的时候数据库压力大外,但是相比较前端逻辑分页来说，具有实时性.