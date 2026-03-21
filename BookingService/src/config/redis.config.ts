import IORedis from 'ioredis';
import Redlock from 'redlock';
import { serverConfig } from '.';

// Create a Redis client
//what is the purpose of this client?
// The Redis client is used to interact with the Redis server. It allows us to perform various operations such as setting and getting values, managing locks, and handling pub/sub messaging. In this context, it is likely used for implementing distributed locking mechanisms to ensure that only one instance of a process can access a critical section of code at a time, which is essential for maintaining data consistency in a distributed environment.
const redisClient = new IORedis(serverConfig.REDIS_SERRVER_URL);

// Create a Redlock instance
//what is the purpose of this redlock instance?
// The Redlock instance is used to implement distributed locking using Redis. It provides a mechanism to acquire and release locks across multiple instances of an application, ensuring that only one instance can access a critical section of code at a time. This is particularly useful in scenarios where multiple instances of an application are running and need to coordinate access to shared resources, such as a database or a file system, to
const redlock = new Redlock([redisClient],{
    retryCount: 10,
    retryDelay: 200, // time in ms
    retryJitter: 200, // time in ms
    driftFactor: 0.01 // time in ms
});

export { redisClient, redlock };