/**
 * mongo init script
 * @see https://www.mongodb.com/docs/manual/reference/method/
 */

console.log("mongo-init.js script executed");

const mongoDatabaseName = process.env.MONGO_DATABASE_NAME;
const mongoCustomUsername = process.env.MONGO_CUSTOM_USERNAME;
const mongoCustomPassword = process.env.MONGO_CUSTOM_PASSWORD;

const projectDb = db = db.getSiblingDB(mongoDatabaseName);

projectDb.createCollection('project');
projectDb.createCollection('');

// db will get created together with user
db.createUser({
    user: mongoCustomUsername,
    pwd: mongoCustomPassword,
    roles: [{
        role: 'userAdmin', // userAdminAnyDatabase // userAdmin // readWrite
        db: mongoDatabaseName
    }]
});