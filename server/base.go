package server

const DEMO_QUEUE_NAME = "demoQueue"

//fanout 交换机name
const FANOUT_EXCHANGE_NAME = "amq.fanout"

//fanout routing_key
const FANOUT_ROUTING_KEY1 = "fanoutQueue1"
const FANOUT_ROUTING_KEY2 = "fanoutQueue2"

//topic 交换机name
const TOPIC_EXCHANGE_NAME = "amq.topics"

//topic
//#匹配一个或多个词
//*仅匹配一个词
const TOPIC1 = "china.people.life"
const TOPIC2 = "china.other"
