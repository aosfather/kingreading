# Profile
## 范围
 * 定义推送的终端类型 （kindle ...）
 * 定义推送的频率 (每日、每周) ，最晚发送时间
 * 定义内容的主题 (小说名、资讯、博客主、类型杂集( IT相关、java、情感生活...)) 
 
 ## 格式定义
  * JSON格式
    * RemoteType string //终端类型
    * Caption string    //内容标题
    * Rate    RateType  //频率
    * LastTime string   //最晚发出时间
    * EmptyEnabled bool //空内容是否也发送