# golang_blog_demo_noORM

###简单文章CUDR
|路径 | 方法|用途|
|----|----|----|
|"/"           | get |首页|
|"/postarticle" |post |创建新博客文章|
|"/articles"    |get |请求所有文章列表|
|"/article/:id" |get |请求单独一页文章|
|"/comment/:id" |get |取得文章评论|
|"/comment/:id" |post |对当前id文章评论|