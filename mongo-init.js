// The code in the docker-entrypoint-init.d folder is only executed if the database has never been initialized before
conn = new Mongo()
db = conn.getDB("cookbook");
// The tricky part was to understand that *.js files were run unauthenticated
admin = db.getSiblingDB('admin')
admin.auth("root", "example")

db.createCollection("recipes", {capped: false})
