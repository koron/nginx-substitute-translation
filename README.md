# Translation by substitution on nginx

## Try on Linux

    ../root/nginx/sbin/nginx -p `pwd`/nginx-env -c conf/nginx.conf

### Results

Measure average of time per request.

    $ ab -n ${TIME} ${URL}

All data's unit is millisecond.

               |none  |subst (255)
---------------|-----:|-----------:
/              |0.055 |      0.086
/admin.html    |0.048 |      0.192
/settings.html |0.065 |      1.397

Using regexp filter.

Regex entries  |924       |500      |100
---------------|---------:|--------:|------:
/              |  439.353 |  50.381 | 2.295
/admin.html    | 2118.383 | 178.126 | 7.367
/settings.html |24201.220 |2127.501 |86.995

Using substitute and regexp filter (in this order).

               |100    
---------------|------:
/              | 2.223 
/admin.html    | 7.419 
/settings.html |88.373 

Using regexp and substitute filter (in this order).

               |100    
---------------|------:
/              | 2.245 
/admin.html    | 7.429 
/settings.html |88.628 

## Remarks

*   <https://openresty.org/download/agentzh-nginx-tutorials-en.html>
