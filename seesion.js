const session = db.getMongo().startSession();
session.startTransaction();
const sessionColl = session.getDatabase("test").getCollection("uc");
sessionColl.updateOne({name: "eric"}, {$set: {"value": 123}})
sessionColl.commitTransaction()