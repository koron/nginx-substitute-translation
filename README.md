# Translation by substitution on nginx

NGINX の `sub_filter` (文字列置換)と `refplace_filter` (正規表現置換) を用いたオンタイム翻訳の実証実験とパフォーマンス計測。

## Try on Linux

    ../root/nginx/sbin/nginx -p `pwd`/nginx-env -c conf/nginx.conf

### Results

Measure average of time per request.

    $ ab -n ${TIME} ${URL}

All data's unit is millisecond.

(`sub_filter` だけを使ったケースと、何も使ってないケースとのパフォーマンス差)

Path (URL)     |none  |subst (255)
---------------|-----:|-----------:
/              |0.055 |      0.086
/admin.html    |0.048 |      0.192
/settings.html |0.065 |      1.397

Using regexp filter.
(`replace_filter` だけを使って並列数を挙げた時の計測)

Path (URL)     |924       |500      |100
---------------|---------:|--------:|------:
/              |  439.353 |  50.381 | 2.295
/admin.html    | 2118.383 | 178.126 | 7.367
/settings.html |24201.220 |2127.501 |86.995

Using substitute and regexp filter (in this order).
(`sub_filter` のあとに `replace_filter` をかけ、並列数を100に固定して計測)

Path (URL)     |100    
---------------|------:
/              | 2.223 
/admin.html    | 7.419 
/settings.html |88.373 

Using regexp and substitute filter (in this order).
(`replace_filter` のあとに `sub_filter` をかけ、並列数を100に固定して計測)

Path (URL)     |100    
---------------|------:
/              | 2.245 
/admin.html    | 7.429 
/settings.html |88.628 

## Remarks

*   <https://openresty.org/download/agentzh-nginx-tutorials-en.html>
