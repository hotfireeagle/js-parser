### letter
    
由于javascript的变量必须以字母，_, $开头，即可选范围在[A-Z]&[a-z]&[_,$]之间。
     
       |------ _ --------|
       |------ $ --------|
       |------ A --------|
       |------ B --------|
       |     ......      |
       |------ Z --------|
-----> |                 | ----->
       |------ a --------|
       |------ b --------|
       |     ......      |
       |------ z --------|
      
     
### digit
     
自然在[0-9]之间
     
      | ------- 0 ------- |
      | ------- 1 ------- |
----> | ------- 2 ------- | ----->
      | ------- 3 ------- |
      |      ......       |
      | ------- 9 ------- |
     

### word
     
以letter开头，后可跟任意的letter、digit、empty。
     
注：这里暂不实现变量标识符的长度限制。
      
因此word的状态图如下：
     
                     |-------> letter ------> |
                     |                        |
------> letter ----->|------------------------|----------> output(word)
          ^          |                        |     |
          |          |-------> digit -------> |     |
          |                                         |
          |                                         |
          <-----------------------------------------|