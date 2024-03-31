#!/bin/bash

FILENAME=$1

# ${VAR%.* }含义：从$VAR中删除位于 % 右侧的通配符左右匹配的字符串，
# 通配符从右向左进行匹配。现在给变量 name 赋值，name=text.gif，
# 那么通配符从右向左就会匹配到 .gif，所有从 $VAR 中删除匹配结果
# %为非贪婪匹配，%%为贪婪匹配
BASENAME=${FILENAME%.*}

# 从左向右匹配：# 和 ## ,两个##表示贪婪匹配
EXTNAME=${FILENAME##*.}

if [ "$EXTNAME" == "sh" ] || [ "$EXTNAME" == "SH" ]
then
    echo "#!/bin/bash" > $FILENAME
    echo "" >> $FILENAME
    echo "# function: " >> $FILENAME
    echo "# author  : liangjisheng" >> $FILENAME
    echo "# date    : `date +"%Y/%m/%d %H:%M:%S"`" >> $FILENAME
    echo "# version : 1.0" >> $FILENAME
    echo "" >> $FILENAME
    echo "" >> $FILENAME
    echo "" >> $FILENAME
    echo "exit 0" >> $FILENAME
else [ "$EXTNAME" == "c" ] || [ "$EXTNAME" == "cc" ] || [ "$EXTNAME" == "cpp" ] || [ "$EXTNAME" == "h" ] || [ "$EXTNAME" == "hpp" ]
    echo "" > $FILENAME
    echo "// function: " >> $FILENAME
    echo "// author  : liangjisheng" >> $FILENAME
    echo "// date    : `date +"%Y/%m/%d %H:%M:%S"`" >> $FILENAME
    echo "// version : 1.0" >> $FILENAME
    echo "" >> $FILENAME
    echo "" >> $FILENAME
fi

exit 0
