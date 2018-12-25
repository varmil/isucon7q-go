#!/bin/bash
USERNAME=isucon
PASSWORD=isucon
DATABASE=isubata


query="select id,name from image order by id"
result=$(mysql -u${USERNAME} -p ${DATABASE} -BN -e "${query}")


IFS=$'\n'; for row in ${result}
do
    declare -a columns
    IFS=$'\t' read -ra columns <<< "${row}"
    echo "id         : [${columns[0]}]"
    echo "name       : [${columns[1]}]"
    mysql -uisucon -pisucon isubata -e "SELECT data FROM image WHERE id = ${columns[0]} INTO DUMPFILE '/tmp/img/${columns[1]}';"
done
IFS=$org_ifs
