// config/redis.config.ts

import Redis from "ioredis";
import { serverConfig } from "./index";

const redisConfig = {
    port: serverConfig.REDIS_PORT,
    host: serverConfig.REDIS_HOST,
    maxRetriesPerRequest: null,
};

let connection: Redis | null = null;

export function getRedisConnectionObject() {
    if (!connection) {
        connection = new Redis(redisConfig);

        connection.on("connect", () => {
            console.log(`Connected to Redis at ${redisConfig.host}:${redisConfig.port}`);
        });

        connection.on("error", (err) => {
            console.error(`Failed to connect to Redis: ${err}`);
        });
    }
    return connection;
}